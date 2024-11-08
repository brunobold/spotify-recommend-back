package spotify

type Token struct {
	AccessToken string `json:"access_token"`
}

func RefreshToken(token string) (string, error) {
	// data := url.Values{}
	// data.Set("grant_type", "refresh_token")
    // data.Set("refresh_token", token)
    // data.Set("client_id", clientID)
    // data.Set("client_secret", clientSecret)

	// res, err := http.Post("https://accounts.spotify.com/api/token", url.Values{""})
	// if err != nil {
	// 	return "", err
	// }
	return "", nil
}