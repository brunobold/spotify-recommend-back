package spotify

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Track struct {
    Seed string `json:"uri"`
    // Add other fields as needed
}

type ListeningDataResponse struct {
    Items []Track `json:"items"`
}

func GetListeningData(token string) ([]Track, error) {
    req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me/top/tracks?limit=30&time_range=short_term", nil)
    if err != nil {
        return nil, err
    }
    req.Header.Set("Authorization", "Bearer "+token)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    responseData, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var listeningData ListeningDataResponse
    err = json.Unmarshal(responseData, &listeningData)
    if err != nil {
        return nil, err
    }

    // Select 5 random songs from the fetched data
    randomTracks := selectRandomTracks(listeningData.Items, 5)

    // Trim "spotify:track:" from the Seed string of the selected songs
    for i := range randomTracks {
        randomTracks[i].Seed = strings.TrimPrefix(randomTracks[i].Seed, "spotify:track:")
    }

    return randomTracks, nil
}

func selectRandomTracks(tracks []Track, count int) []Track {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    r.Shuffle(len(tracks), func(i, j int) {
        tracks[i], tracks[j] = tracks[j], tracks[i]
    })

    if count > len(tracks) {
        count = len(tracks)
    }

    return tracks[:count]
}