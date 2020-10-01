package omdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func SearchMovie(token, title string) (*Response, error) {
	u := fmt.Sprintf("%s/?apikey=%s&t=%s", EndPoint, token, url.QueryEscape(title))
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search movie failed: %s", resp.Status)
	}

	var res Response
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}
