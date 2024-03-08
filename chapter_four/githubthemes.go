package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

var (
	baseURL    = "https://api.github.com/repos/"
	httpClient = http.DefaultClient
	username   = "your_username"
	token      = "your_token"
)

type Issue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func createIssue(owner, repo string, issue Issue) error {
    // Launch preferred text editor for entering issue body
    editor := os.Getenv("EDITOR")
    if editor == "" {
        editor = "vim" // default to vim if EDITOR is not set
    }
    tempFile := "issue_body.txt"
    err := ioutil.WriteFile(tempFile, []byte(issue.Body), 0644)
    if err != nil {
        return err
    }

    cmd := exec.Command(editor, tempFile)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        return err
    }

    // Read issue body from temporary file
    updatedBody, err := ioutil.ReadFile(tempFile)
    if err != nil {
        return err
    }

    issue.Body = string(updatedBody)

    // Create issue
    url := fmt.Sprintf("%s%s/%s/issues", baseURL, owner, repo)
    payload, err := json.Marshal(issue)
    if err != nil {
        return err
    }
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
    if err != nil {
        return err
    }
    req.SetBasicAuth(username, token)
    req.Header.Set("Content-Type", "application/json")

    resp, err := httpClient.Do(req)
    if err != nil {
        return err
   }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusCreated {
        return fmt.Errorf("failed to create issue: %s", resp.Status)
    }
    fmt.Println("Issue created successfully")
    return nil
}

func readIssue(owner, repo string, issueNumber int) (Issue, error) {
	url := fmt.Sprintf("%s%s/%s/issues/%d", baseURL, owner, repo, issueNumber)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Issue{}, err
	}
	req.SetBasicAuth(username, token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return Issue{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Issue{}, fmt.Errorf("failed to read issue: %s", resp.Status)
	}

	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return Issue{}, err
	}

	return issue, nil
}

func updateIssue(owner, repo string, issueNumber int) error {
	// Read existing issue body
	issue, err := readIssue(owner, repo, issueNumber)
	if err != nil {
		return err
	}

	// Launch preferred text editor for updating issue body
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim" // default to vim if EDITOR is not set
	}
	tempFile := "issue_body.txt"
	err = ioutil.WriteFile(tempFile, []byte(issue.Body), 0644)
	if err != nil {
		return err
	}

	cmd := exec.Command(editor, tempFile)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	// Read updated issue body from temporary file
	updatedBody, err := ioutil.ReadFile(tempFile)
	if err != nil {
		return err
	}

	// Update issue with new body
	updatedIssue := Issue{
		Title: issue.Title,
		Body:  string(updatedBody),
	}
	err = createIssue(owner, repo, updatedIssue)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	owner := "owner" // Replace with your GitHub username or organization
	repo := "repo"   // Replace with your GitHub repository name

	fmt.Println("1. Create Issue")
	fmt.Println("2. Read Issue")
	fmt.Println("3. Update Issue")
	fmt.Println("Enter your choice:")
	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		var title, body string
		fmt.Println("Enter issue title:")
		fmt.Scanln(&title)
		fmt.Println("Enter issue body:")
		issue := Issue{Title: title, Body: body}
		err := createIssue(owner, repo, issue)
		if err != nil {
			fmt.Println("Error:", err)
		}
	case 2:
		var issueNumber int
		fmt.Println("Enter issue number:")
		fmt.Scanln(&issueNumber)
		issue, err := readIssue(owner, repo, issueNumber)
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Println("Issue:", issue)
	case 3:
		var issueNumber int
		fmt.Println("Enter issue number:")
		fmt.Scanln(&issueNumber)
		err := updateIssue(owner, repo, issueNumber)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Issue updated successfully")
		}
	default:
		fmt.Println("Invalid choice")
	}
}
