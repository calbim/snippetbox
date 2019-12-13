package mysql

import (
	"database/sql"

	"github.com/calbim/snippetbox/pkg/models"
)

//SnippetModel type wraps a sql database connection pool
type SnippetModel struct {
	DB *sql.DB
}

//Insert a new snippet into a database
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := ` INSERT INTO snippets (title, content, created, expires) VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY)) `
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

//Get a snippet from database based on its id
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := ` SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() AND id = ? `
	s := &models.Snippet{}
	err := m.DB.QueryRow(stmt, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		return nil, models.ErrRecordNotFound
	}
	return s, nil
}

//Latest returns the 10 latest snippets from a database
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	stmt := ` SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10 `
	rows, err := m.DB.Query(stmt)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	snippets := []*models.Snippet{}

	for rows.Next() {
		s := &models.Snippet{}
		rows.Scan(&s.ID, &s.Title, &s.Content, &s.Content, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}
	return snippets, nil
}
