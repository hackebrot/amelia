package amelia

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const gistURL = "https://api.github.com/gists"

// GitHubClient holds user auth info and a http client
type GitHubClient struct {
	Username string
	Token    string
	client   http.Client
}

func (c *GitHubClient) post(URL string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", URL, body)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Username, c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateGist using the GitHub API
func (c *GitHubClient) CreateGist(g *Gist) (*Gist, error) {

	j, err := json.Marshal(g)
	if err != nil {
		return nil, fmt.Errorf("unable to serialize gist: %v", err)
	}

	resp, err := c.post(gistURL, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("POST %v request unsuccessful %v %v", gistURL, resp.Status, string(j))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read body %v", err)
	}

	var r Gist
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, fmt.Errorf("unable to deserialize gist %v", err)
	}
	return &r, nil
}
