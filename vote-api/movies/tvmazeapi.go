package movies

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var TvMazeApiBaseUrl = "https://api.tvmaze.com"

const (
	TvMazeShowsByIdUri = "/lookup/shows"
)

type TvMazeMovieDto struct {
	ID             int      `json:"id"`
	URL            string   `json:"url"`
	Name           string   `json:"name"`
	Type           string   `json:"type"`
	Language       string   `json:"language"`
	Genres         []string `json:"genres"`
	Status         string   `json:"status"`
	Runtime        any      `json:"runtime"`
	AverageRuntime int      `json:"averageRuntime"`
	Premiered      string   `json:"premiered"`
	Ended          string   `json:"ended"`
	OfficialSite   string   `json:"officialSite"`
	Schedule       struct {
		Time string   `json:"time"`
		Days []string `json:"days"`
	} `json:"schedule"`
	Rating struct {
		Average float64 `json:"average"`
	} `json:"rating"`
	Weight     int `json:"weight"`
	Network    any `json:"network"`
	WebChannel struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		Country      any    `json:"country"`
		OfficialSite string `json:"officialSite"`
	} `json:"webChannel"`
	DvdCountry any `json:"dvdCountry"`
	Externals  struct {
		Tvrage  any    `json:"tvrage"`
		Thetvdb int    `json:"thetvdb"`
		Imdb    string `json:"imdb"`
	} `json:"externals"`
	Image struct {
		Medium   string `json:"medium"`
		Original string `json:"original"`
	} `json:"image"`
	Summary string `json:"summary"`
	Updated int    `json:"updated"`
	Links   struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Previousepisode struct {
			Href string `json:"href"`
			Name string `json:"name"`
		} `json:"previousepisode"`
	} `json:"_links"`
}

func GetMovieByImdbId(imdbId string) (Movie, error) {
	baseUrlString, err := url.JoinPath(TvMazeApiBaseUrl, TvMazeShowsByIdUri)
	if err != nil {
		return Movie{}, fmt.Errorf("communication failed with tvmaze API: %w", err)
	}
	baseUrl, err := url.Parse(baseUrlString)
	if err != nil {
		return Movie{}, fmt.Errorf("communication failed with tvmaze API: %w", err)
	}

	query := baseUrl.Query()
	query.Set("imdb", imdbId)
	baseUrl.RawQuery = query.Encode()

	response, err := http.Get(baseUrl.String())
	if err != nil {
		return Movie{}, fmt.Errorf("communication failed with tvmaze API: %w", err)
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return Movie{}, fmt.Errorf("communication failed with tvmaze API: %w", err)
	}

	movieDto := TvMazeMovieDto{}

	err = json.NewDecoder(response.Body).Decode(&movieDto)
	if err != nil {
		return Movie{}, fmt.Errorf("communication failed with tvmaze API: %w", err)
	}

	return movieDto.toDomain(), nil
}

func (dto TvMazeMovieDto) toDomain() Movie {
	return Movie{
		Title:         dto.Name,
		ImdbId:        dto.Externals.Imdb,
		CoverImageUrl: dto.Image.Medium,
		ReleaseDate:   dto.Premiered,
		Genres:        dto.Genres,
		Summary:       dto.Summary,
	}
}
