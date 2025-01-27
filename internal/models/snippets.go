package models

import (
	"database/sql"
	"errors"
	"time"
)

// Define a Snippet type to hold the data for an individual snippet
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// Define a SnippetModel type which wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

// Inserts a new snippet into the database
func (m *SnippetModel) Insert(
	title string,
	content string,
	expires int,
) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
            VALUES(?, ?, DATETIME('now'), DATETIME('now', ?||' days'))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Returns a specific snippet based on id
func (m *SnippetModel) Get(id int) (Snippet, error) {
	// Use strftime to format the datetime as a string in a format Go can parse
	stmt := `SELECT id, title, content, 
		strftime('%Y-%m-%d %H:%M:%S', created) as created,
		strftime('%Y-%m-%d %H:%M:%S', expires) as expires
		FROM snippets 
		WHERE expires > datetime('now') AND id = ?`

	// Temporary variables to hold the datetime strings
	var createdStr, expiresStr string
	var s Snippet

	err := m.DB.QueryRow(stmt, id).Scan(
		&s.ID,
		&s.Title,
		&s.Content,
		&createdStr,
		&expiresStr,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		}
		return Snippet{}, err
	}

	// Parse the datetime strings into time.Time
	s.Created, err = time.Parse("2006-01-02 15:04:05", createdStr)
	if err != nil {
		return Snippet{}, err
	}

	s.Expires, err = time.Parse("2006-01-02 15:04:05", expiresStr)
	if err != nil {
		return Snippet{}, err
	}

	return s, nil
}

// Returns the 10 most recently created snippets
func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
