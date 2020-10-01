package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/h-matsuo/gopl/ch04/ex11/github"
)

var (
	cmdCreate = flag.Bool("create", false, "Create an issue")
	cmdGet    = flag.Bool("get", false, "Get an issue")
	cmdUpdate = flag.Bool("update", false, "Update an issue")
	cmdClose  = flag.Bool("close", false, "Close an issue")

	user  = flag.String("user", "", "GitHub username")
	token = flag.String("token", "", "GitHub personal access token")

	owner  = flag.String("owner", "", "Repository owner")
	repo   = flag.String("repo", "", "Repository name")
	number = flag.Int("number", -1, "Issue number")
	title  = flag.String("title", "", "Issue title")
	body   = flag.String("body", "", "Issue body")
)

func main() {

	flag.Parse()

	if *owner == "" || *repo == "" {
		fmt.Fprintln(os.Stderr, "`-owner` and `-repo` are required.")
		os.Exit(1)
	}

	switch {
	case *cmdCreate:
		if *user == "" || *token == "" || *title == "" {
			fmt.Fprintln(os.Stderr, "`-user`, `-token` and `-title` are required.")
			os.Exit(1)
		}
		createIssue(*user, *token, *owner, *repo, *title, *body)
	case *cmdGet:
		if *number < 1 {
			fmt.Fprintln(os.Stderr, "`-number` is required.")
			os.Exit(1)
		}
		getIssue(*user, *token, *owner, *repo, *number)
	case *cmdUpdate:
		if *user == "" || *token == "" || *number < 1 || *title == "" {
			fmt.Fprintln(os.Stderr, "`-user`, `-token`, `-number` and `-title` are required.")
			os.Exit(1)
		}
		updateIssue(*user, *token, *owner, *repo, *number, *title, *body)
	case *cmdClose:
		if *user == "" || *token == "" || *number < 1 {
			fmt.Fprintln(os.Stderr, "`-user`, `-token` and `-number` are required.")
			os.Exit(1)
		}
		closeIssue(*user, *token, *owner, *repo, *number)
	default:
		fmt.Fprintln(os.Stderr, "Specify one of the subcommands: `-create`, `-get`, `-update`, `-close`")
		os.Exit(1)
	}
}

func createIssue(user, token, owner, repo, title, body string) {
	if body == "" {
		body = editWithEditor("<!-- Edit markdown body -->\n")
	}
	issue, err := github.CreateIssue(user, token, owner, repo, title, body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	printIssue(issue)
}

func getIssue(user, token, owner, repo string, number int) {
	issue, err := github.GetIssue(user, token, owner, repo, number)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	printIssue(issue)
}

func updateIssue(user, token, owner, repo string, number int, title, body string) {
	if body == "" {
		issue, err := github.GetIssue(user, token, owner, repo, number)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		body = editWithEditor(issue.Body)
	}
	issue, err := github.UpdateIssue(user, token, owner, repo, number, title, body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	printIssue(issue)
}

func closeIssue(user, token, owner, repo string, number int) {
	issue, err := github.CloseIssue(user, token, owner, repo, number)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	printIssue(issue)
}

// editWithEditor edits the string `before`
// with the editor specified by an environment variable `EDITOR`.
// Returns the edited string.
func editWithEditor(before string) string {
	tmpFile, err := ioutil.TempFile("", "")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer os.Remove(tmpFile.Name())
	fmt.Fprint(tmpFile, before)

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "nano"
	}
	cmd := exec.Command("sh", "-c", editor+" "+tmpFile.Name())
	fmt.Printf("%v\n", cmd.Args)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	buf, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return string(buf)
}

func printIssue(issue *github.Issue) {
	fmt.Printf("Issue #%d, State: %s\n", issue.Number, issue.State)
	fmt.Printf("Created At: %s\n", issue.CreatedAt)
	fmt.Printf("Updated At: %s\n", issue.UpdatedAt)
	fmt.Println()
	fmt.Printf("Title:\n%s\n", issue.Title)
	fmt.Println()
	fmt.Printf("Body:\n%s\n", issue.Body)
}
