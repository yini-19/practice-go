package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Date      string   `json:"dates"`
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type LocationsIndex struct {
	Index []Location `json:"index"`
}

type DatesIndex struct {
	Index []Date `json:"index"`
}

type RelationsIndex struct {
	Index []Relation `json:"index"`
}

func fetchArtist(url string) ([]Artist, error) {

	body, err := fetchJSON(url)
	if err != nil {
		return nil, err
	}

	var artists []Artist
	if err := json.Unmarshal(body, &artists); err != nil {
		return nil, fmt.Errorf("unmarshalling artists: %w", err)
	}
	return artists, nil
}

func fetchLocations(url string) (LocationsIndex, error) {
	body, err := fetchJSON(url)
	if err != nil {
		return LocationsIndex{}, err
	}
	var location LocationsIndex
	if err := json.Unmarshal(body, &location); err != nil {
		return location, fmt.Errorf("unmarshalling locations: %w", err)
	}
	return location, nil
}

func fetchDates(url string) (DatesIndex, error) {
	body, err := fetchJSON(url)
	if err != nil {
		return DatesIndex{}, err
	}
	var date DatesIndex
	if err := json.Unmarshal(body, &date); err != nil {
		return date, fmt.Errorf("unmarshalling dates: %w", err)
	}
	return date, nil
}

func fetchRelations(url string) (RelationsIndex, error) {
	body, err := fetchJSON(url)
	if err != nil {
		return RelationsIndex{}, err
	}
	var relation RelationsIndex
	if err := json.Unmarshal(body, &relation); err != nil {
		return relation, fmt.Errorf("unmarshalling relation: %w", err)
	}
	return relation, nil
}

/*
---------------------------------
Helper
---------------------------------
*/
func fetchJSON(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetching url: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	return body, nil
}

//
