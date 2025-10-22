package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	searchURL = "https://itunes.apple.com/search"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

type Podcast struct {
	ArtistName     string `json:"artistName"`
	CollectionName string `json:"collectionName"`
	FeedURL        string `json:"feedUrl"`
	TrackCount     int    `json:"trackCount"`
}

type SearchResponse struct {
	ResultCount int       `json:"resultCount"`
	Results     []Podcast `json:"results"`
}

func (c *Client) SearchPodcasts(term string, limit int) (*SearchResponse, error) {
	if term == "" {
		return nil, fmt.Errorf("search term cannot be empty")
	}

	u, err := url.Parse(searchURL)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("term", term)
	q.Set("media", "podcast")
	q.Set("entity", "podcast")
	if limit > 0 {
		q.Set("limit", fmt.Sprintf("%d", limit))
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var searchResponse SearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&searchResponse); err != nil {
		return nil, err
	}

	return &searchResponse, nil
}
