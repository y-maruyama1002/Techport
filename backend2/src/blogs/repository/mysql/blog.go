package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/y-maruyama1002/Techport/domain"
)

type mysqlBlogRepository struct {
	Conn *sql.DB
}

func NewMysqlBlogRepository(conn *sql.DB) domain.BlogRepository {
	return &mysqlBlogRepository{conn}
}

func (r *mysqlBlogRepository) fetch(query string, args ...interface{}) (result []domain.Blog, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := r.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	result = make([]domain.Blog, 0)
	for rows.Next() {
		t := domain.Blog{}
		err = rows.Scan(
			&t.ID, &t.Title, &t.Body, &t.CreatedAt, &t.UpdatedAt,
		)
		if err != nil {
			logrus.Error(err)
		}
		fmt.Println(result)
		result = append(result, t)
	}

	return result, nil
}

func (r *mysqlBlogRepository) GetById(id int64) (res domain.Blog, err error) {
	query := "SELECT id, title, body, created_at, updated_at FROM blogs WHERE id = ?"
	list, err := r.fetch(query, id)
	if err != nil {
		return
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}
	return
}
