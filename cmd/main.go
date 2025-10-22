package main

import (
	"fmt"
	"log"

	"PodTUI/internal/client"
)

func main() {

	c := client.NewClient()

	searchTerm := "Thmaniyah"
	searchResponse, err := c.SearchPodcasts(searchTerm, 5)
	if err != nil {
		log.Fatalf("Error searching for podcasts: %v", err)
	}

	fmt.Printf("Found %d results for '%s':\n", searchResponse.ResultCount, searchTerm)
	for _, podcast := range searchResponse.Results {
		fmt.Printf("  - %s by %s (%d tracks)\n", podcast.CollectionName, podcast.ArtistName, podcast.TrackCount)
		fmt.Printf("    Feed URL: %s\n", podcast.FeedURL)
	}
}
