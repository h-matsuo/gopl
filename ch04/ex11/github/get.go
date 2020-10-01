package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetIssue(user, token, owner, repo string, number int) (*Issue, error) {
	u := fmt.Sprintf("%s/repos/%s/%s/issues/%d", EndPoint, owner, repo, number)
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	if user != "" && token != "" {
		req.SetBasicAuth(user, token)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get issue failed: %s", resp.Status)
	}

	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil
}
