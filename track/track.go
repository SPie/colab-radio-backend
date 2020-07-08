package track

import (
	"fmt"

	"github.com/zmb3/spotify"
)

// Track represents a spotify song
type Track struct {
	TrackID  string   `json:"trackId"`
	Name     string   `json:"name"`
	Endpoint string   `json:"endpoint"`
	Artists  []Artist `json:"artists"`
}

func newTrack(track spotify.FullTrack) Track {
	return Track{TrackID: track.ID.String(), Name: track.Name, Endpoint: track.Endpoint, Artists: parseArtists(track)}
}

func parseArtists(track spotify.FullTrack) []Artist {
	artists := make([]Artist, len(track.Artists))
	for i, artist := range track.Artists {
		artists[i] = newArtist(artist)
	}

	return artists
}

// Artist represents a spotify artist
type Artist struct {
	ArtistID string `json:"artistId"`
	Name     string `json:"name"`
	Endpoint string `json:"endpoint"`
}

func newArtist(artist spotify.SimpleArtist) Artist {
	return Artist{ArtistID: artist.ID.String(), Name: artist.Name, Endpoint: artist.Endpoint}
}

// Service handles business logic for tracks
type Service interface {
	SearchTrack(query string, client SearchClient) ([]Track, error)
}

type service struct{}

// SearchClient interface for spotify track search
type SearchClient interface {
	Search(query string, t spotify.SearchType) (*spotify.SearchResult, error)
}

// NewService initializes a TrackService instance
func NewService() Service {
	return service{}
}

// SearchTrack searches for a song in Spotify
func (service service) SearchTrack(query string, client SearchClient) ([]Track, error) {
	fmt.Printf("Query: %s\n", query)
	result, err := client.Search(query, spotify.SearchTypeTrack)
	if err != nil {
		return []Track{}, err
	}

	tracks := make([]Track, len(result.Tracks.Tracks))
	for i, track := range result.Tracks.Tracks {
		tracks[i] = newTrack(track)
	}

	return tracks, nil
}
