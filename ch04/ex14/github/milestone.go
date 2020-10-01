package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ListMilestones(owner, repo string) ([]Milestone, error) {
	u := fmt.Sprintf("%s/repos/%s/%s/milestones", Endpoint, owner, repo)
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list milestones failed: %s", resp.Status)
	}

	var result []Milestone
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
