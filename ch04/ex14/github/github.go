package github

const Endpoint = "https://api.github.com"

type Issue struct {
	Number  int
	HTMLURL string `json:"html_url"`
	Title   string
	State   string
	User    *User
}

type User struct {
	HTMLURL string `json:"html_url"`
	Login   string
}

type Milestone struct {
	Number  int
	HTMLURL string `json:"html_url"`
	Title   string
	State   string
	Creator *User
}
