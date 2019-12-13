package mysql

import "database/sql"

import "github.com/calbim/snippetbox/pkg/models"

import "fmt"

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
	fmt.Println(id)
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

//Get a snippet from database based on its id
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

//Latest returns the 10 latest snippets from a database
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
