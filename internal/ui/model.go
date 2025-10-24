package ui

import (
	"PodTUI/internal/client"
	"PodTUI/internal/rss"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type viewState int

const (
	searchView viewState = iota
	podcastListView
	episodeListView
)

type Model struct {
	client      *client.Client
	state       viewState
	textInput   textinput.Model
	podcasts    []client.Podcast
	episodes    []rss.Episode
	selected    int
	err         error
	loading     bool
	searchTerm  string
	podcastName string
	width       int
	height      int
}

func NewModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Enter podcast name..."
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 50

	return Model{
		client:    client.NewClient(),
		state:     searchView,
		textInput: ti,
		selected:  0,
	}
}

type searchResultMsg struct {
	podcasts []client.Podcast
}

type episodesResultMsg struct {
	episodes []rss.Episode
}

type errMsg struct {
	err error
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func searchPodcasts(c *client.Client, term string) tea.Cmd {
	return func() tea.Msg {
		resp, err := c.SearchPodcasts(term, 10)
		if err != nil {
			return errMsg{err}
		}
		return searchResultMsg{podcasts: resp.Results}
	}
}

func fetchEpisodes(feedURL string) tea.Cmd {
	return func() tea.Msg {
		feed, err := rss.ParseFeed(feedURL)
		if err != nil {
			return errMsg{err}
		}
		return episodesResultMsg{episodes: feed.Channel.Items}
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			if m.state == searchView {
				return m, tea.Quit
			}
			// Go back on q
			if m.state == episodeListView {
				m.state = podcastListView
				m.selected = 0
				m.episodes = nil
				return m, nil
			}
			if m.state == podcastListView {
				m.state = searchView
				m.selected = 0
				m.podcasts = nil
				m.textInput.SetValue("")
				return m, nil
			}

		case "esc":
			if m.state == episodeListView {
				m.state = podcastListView
				m.selected = 0
				m.episodes = nil
				return m, nil
			}
			if m.state == podcastListView {
				m.state = searchView
				m.selected = 0
				m.podcasts = nil
				m.textInput.SetValue("")
				return m, nil
			}

		case "enter":
			if m.state == searchView && m.textInput.Value() != "" {
				m.searchTerm = m.textInput.Value()
				m.loading = true
				m.state = podcastListView
				m.selected = 0
				return m, searchPodcasts(m.client, m.searchTerm)
			}
			if m.state == podcastListView && len(m.podcasts) > 0 {
				podcast := m.podcasts[m.selected]
				m.podcastName = podcast.CollectionName
				m.loading = true
				m.state = episodeListView
				m.selected = 0
				return m, fetchEpisodes(podcast.FeedURL)
			}

		case "up", "k":
			if m.state != searchView && m.selected > 0 {
				m.selected--
			}

		case "down", "j":
			if m.state == podcastListView && m.selected < len(m.podcasts)-1 {
				m.selected++
			}
			if m.state == episodeListView && m.selected < len(m.episodes)-1 {
				m.selected++
			}
		}

	case searchResultMsg:
		m.loading = false
		m.podcasts = msg.podcasts
		m.err = nil

	case episodesResultMsg:
		m.loading = false
		m.episodes = msg.episodes
		m.err = nil

	case errMsg:
		m.loading = false
		m.err = msg.err
	}

	if m.state == searchView {
		m.textInput, cmd = m.textInput.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m Model) View() string {
	switch m.state {
	case searchView:
		return m.renderSearchView()
	case podcastListView:
		return m.renderPodcastList()
	case episodeListView:
		return m.renderEpisodeList()
	}
	return ""
}

func (m Model) renderSearchView() string {
	var b strings.Builder

	title := titleStyle.Render("🎙️  PODCAST BROWSER")
	subtitle := subtitleStyle.Render("Discover and explore your favorite podcasts")

	b.WriteString("\n")
	b.WriteString(title)
	b.WriteString("\n")
	b.WriteString(subtitle)
	b.WriteString("\n\n")

	prompt := inputPromptStyle.Render("Search: ")
	input := inputStyle.Render(m.textInput.View())
	b.WriteString(prompt)
	b.WriteString(" ")
	b.WriteString(input)
	b.WriteString("\n\n")

	help := helpStyle.Render(
		helpKeyStyle.Render("enter") + " search • " +
			helpKeyStyle.Render("ctrl+c") + " quit",
	)
	b.WriteString(help)

	return containerStyle.Render(b.String())
}

func (m Model) renderPodcastList() string {
	var b strings.Builder

	title := titleStyle.Render(fmt.Sprintf("🔍 Search Results: \"%s\"", m.searchTerm))
	b.WriteString("\n")
	b.WriteString(title)
	b.WriteString("\n")

	if m.loading {
		loading := loadingStyle.Render("⏳ Loading podcasts...")
		b.WriteString(loading)
		b.WriteString("\n")
		return containerStyle.Render(b.String())
	}

	if m.err != nil {
		errMsg := errorStyle.Render(fmt.Sprintf("❌ Error: %v", m.err))
		b.WriteString(errMsg)
		b.WriteString("\n")
		return containerStyle.Render(b.String())
	}

	if len(m.podcasts) == 0 {
		noResults := statusStyle.Render("No podcasts found. Try a different search term.")
		b.WriteString(noResults)
		b.WriteString("\n")
		return containerStyle.Render(b.String())
	}

	count := successStyle.Render(fmt.Sprintf("Found %d podcasts", len(m.podcasts)))
	b.WriteString(count)
	b.WriteString("\n")
	b.WriteString(dividerStyle.Render(strings.Repeat("─", 80)))
	b.WriteString("\n")

	for i, podcast := range m.podcasts {
		var line string
		num := itemNumberStyle.Render(fmt.Sprintf("%2d.", i+1))

		podcastInfo := fmt.Sprintf("%s\n    by %s • %d episodes",
			podcast.CollectionName,
			podcast.ArtistName,
			podcast.TrackCount,
		)

		if i == m.selected {
			line = selectedItemStyle.Render("▶ " + podcastInfo)
		} else {
			line = itemStyle.Render("  " + podcastInfo)
		}

		b.WriteString(num + " " + line)
		b.WriteString("\n")
	}

	b.WriteString(dividerStyle.Render(strings.Repeat("─", 80)))
	b.WriteString("\n")
	help := helpStyle.Render(
		helpKeyStyle.Render("↑/↓ j/k") + " navigate • " +
			helpKeyStyle.Render("enter") + " select • " +
			helpKeyStyle.Render("esc/q") + " back • " +
			helpKeyStyle.Render("ctrl+c") + " quit",
	)
	b.WriteString(help)

	return containerStyle.Render(b.String())
}

func (m Model) renderEpisodeList() string {
	var b strings.Builder

	title := titleStyle.Render(fmt.Sprintf("🎧 %s", m.podcastName))
	b.WriteString("\n")
	b.WriteString(title)
	b.WriteString("\n")

	if m.loading {
		loading := loadingStyle.Render("⏳ Loading episodes...")
		b.WriteString(loading)
		b.WriteString("\n")
		return containerStyle.Render(b.String())
	}

	if m.err != nil {
		errMsg := errorStyle.Render(fmt.Sprintf("❌ Error: %v", m.err))
		b.WriteString(errMsg)
		b.WriteString("\n")
		return containerStyle.Render(b.String())
	}

	if len(m.episodes) == 0 {
		noEpisodes := statusStyle.Render("No episodes found.")
		b.WriteString(noEpisodes)
		b.WriteString("\n")
		return containerStyle.Render(b.String())
	}

	count := successStyle.Render(fmt.Sprintf("📻 %d episodes available", len(m.episodes)))
	b.WriteString(count)
	b.WriteString("\n")
	b.WriteString(dividerStyle.Render(strings.Repeat("─", 80)))
	b.WriteString("\n")

	for i, episode := range m.episodes {
		var line strings.Builder
		num := itemNumberStyle.Render(fmt.Sprintf("%2d.", i+1))

		title := episode.Title
		if len(title) > 70 {
			title = title[:67] + "..."
		}

		meta := ""
		if episode.PubDate != "" {
			meta = episode.PubDate
		}
		if episode.Duration != "" {
			if meta != "" {
				meta += " • "
			}
			meta += episode.Duration
		}

		if i == m.selected {
			line.WriteString(selectedItemStyle.Render("▶ " + title))
			if meta != "" {
				line.WriteString("\n  " + selectedItemStyle.Render("  "+meta))
			}
		} else {
			line.WriteString(itemStyle.Render("  " + title))
			if meta != "" {
				line.WriteString("\n  " + episodeMetaStyle.Render("  "+meta))
			}
		}

		b.WriteString(num + " " + line.String())
		b.WriteString("\n")
	}

	b.WriteString(dividerStyle.Render(strings.Repeat("─", 80)))
	b.WriteString("\n")
	help := helpStyle.Render(
		helpKeyStyle.Render("↑/↓ j/k") + " navigate • " +
			helpKeyStyle.Render("esc/q") + " back • " +
			helpKeyStyle.Render("ctrl+c") + " quit",
	)
	b.WriteString(help)

	return containerStyle.Render(b.String())
}
