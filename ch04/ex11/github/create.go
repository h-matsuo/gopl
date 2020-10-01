package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type createBody struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func CreateIssue(user, token, owner, repo, title, body string) (*Issue, error) {
	u := fmt.Sprintf("%s/repos/%s/%s/issues", EndPoint, owner, repo)
	jsonBody, _ := json.Marshal(&createBody{title, body})
	req, err := http.NewRequest("POST", u, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	req.SetBasicAuth(user, token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("create issue failed: %s", resp.Status)
	}

	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil
}
