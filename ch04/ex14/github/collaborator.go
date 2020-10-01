package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ListCollaborators(owner, repo string) ([]User, error) {
	u := fmt.Sprintf("%s/repos/%s/%s/collaborators", Endpoint, owner, repo)
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	req.SetBasicAuth("h-matsuo", "cad04b9bdff7ed9ff47ad7f659823080e998a6d1")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list collaborators failed: %s", resp.Status)
	}

	var result []User
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
