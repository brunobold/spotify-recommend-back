package spotify

import (
	"io"
	"net/http"
)

func GetProfile(token string) (string, error) {
    req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)
    if err != nil {
        return "", err
    }
    req.Header.Set("Authorization", "Bearer "+token)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    responseData, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    return string(responseData), nil
}