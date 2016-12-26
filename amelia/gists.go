package amelia

import "time"

// User represents a GitHub user.
type User struct {
	Login *string `json:"login,omitempty"`
	Name  *string `json:"name,omitempty"`
}

// GistFilename represents filename on a gist.
type GistFilename string

// GistFile represents a file on a gist.
type GistFile struct {
	Filename *string `json:"filename,omitempty"`
	Content  *string `json:"content,omitempty"`
}

// Gist represents a GitHub gist.
type Gist struct {
	Description *string                   `json:"description,omitempty"`
	Public      *bool                     `json:"public,omitempty"`
	Owner       *User                     `json:"owner,omitempty"`
	Files       map[GistFilename]GistFile `json:"files,omitempty"`
	HTMLURL     *string                   `json:"html_url,omitempty"`
	CreatedAt   *time.Time                `json:"created_at,omitempty"`
}
