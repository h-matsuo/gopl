package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/h-matsuo/gopl/ch04/ex14/github"
)

type metadata struct {
	Owner         string
	Repo          string
	Issues        []github.Issue
	Milestones    []github.Milestone
	Collaborators []github.User
}

var metainfo = template.Must(template.New("metainfo").Parse(`
<h1>github.com/{{.Owner}}/{{.Repo}} meta data</h1>

<h2>Issues</h2>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Issues}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>

<h2>Milestones</h2>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>Creator</th>
  <th>Title</th>
</tr>
{{range .Milestones}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.Creator.HTMLURL}}'>{{.Creator.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>

<h2>Collaborators</h2>
<ul>
{{range .Collaborators}}
<li><a href='{{.HTMLURL}}'>{{.Login}}</a></li>
{{end}}
</ul>
`))

type config struct {
	user  string
	token string
	owner string
	repo  string
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	config, err := parseQuery(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err.Error())
		return
	}
	generatePage(w, config)
}

func parseQuery(r *http.Request) (*config, error) {
	var c config
	// user := r.URL.Query().Get("user")
	// if user == "" {
	// 	return nil, fmt.Errorf(`Query "user" required.`)
	// }
	// c.user = user
	// token := r.URL.Query().Get("token")
	// if token == "" {
	// 	return nil, fmt.Errorf(`Query "token" required.`)
	// }
	// c.token = token
	owner := r.URL.Query().Get("owner")
	if owner == "" {
		return nil, fmt.Errorf(`Query "owner" required.`)
	}
	c.owner = owner
	repo := r.URL.Query().Get("repo")
	if repo == "" {
		return nil, fmt.Errorf(`Query "repo" required.`)
	}
	c.repo = repo
	return &c, nil
}

func generatePage(w http.ResponseWriter, c *config) {
	issues, err := github.ListIssues(c.owner, c.repo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err.Error())
		return
	}
	milestones, err := github.ListMilestones(c.owner, c.repo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err.Error())
		return
	}
	collaborators, err := github.ListCollaborators(c.owner, c.repo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err.Error())
		return
	}

	if err := metainfo.Execute(w, &metadata{c.owner, c.repo, issues, milestones, collaborators}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
		return
	}
}
