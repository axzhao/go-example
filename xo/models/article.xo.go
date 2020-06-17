// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Article represents a row from 'public.article'.
type Article struct {
	ID          int            `json:"id"`          // id
	Title       string         `json:"title"`       // title
	Createddate pq.NullTime    `json:"createddate"` // createddate
	Body        sql.NullString `json:"body"`        // body
	Userid      sql.NullInt64  `json:"userid"`      // userid

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Article exists in the database.
func (a *Article) Exists() bool {
	return a._exists
}

// Deleted provides information if the Article has been deleted from the database.
func (a *Article) Deleted() bool {
	return a._deleted
}

// Insert inserts the Article to the database.
func (a *Article) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if a._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.article (` +
		`title, createddate, body, userid` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, a.Title, a.Createddate, a.Body, a.Userid)
	err = db.QueryRow(sqlstr, a.Title, a.Createddate, a.Body, a.Userid).Scan(&a.ID)
	if err != nil {
		return err
	}

	// set existence
	a._exists = true

	return nil
}

// Update updates the Article in the database.
func (a *Article) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !a._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if a._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.article SET (` +
		`title, createddate, body, userid` +
		`) = ( ` +
		`$1, $2, $3, $4` +
		`) WHERE id = $5`

	// run query
	XOLog(sqlstr, a.Title, a.Createddate, a.Body, a.Userid, a.ID)
	_, err = db.Exec(sqlstr, a.Title, a.Createddate, a.Body, a.Userid, a.ID)
	return err
}

// Save saves the Article to the database.
func (a *Article) Save(db XODB) error {
	if a.Exists() {
		return a.Update(db)
	}

	return a.Insert(db)
}

// Upsert performs an upsert for Article.
//
// NOTE: PostgreSQL 9.5+ only
func (a *Article) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if a._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.article (` +
		`id, title, createddate, body, userid` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, title, createddate, body, userid` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.title, EXCLUDED.createddate, EXCLUDED.body, EXCLUDED.userid` +
		`)`

	// run query
	XOLog(sqlstr, a.ID, a.Title, a.Createddate, a.Body, a.Userid)
	_, err = db.Exec(sqlstr, a.ID, a.Title, a.Createddate, a.Body, a.Userid)
	if err != nil {
		return err
	}

	// set existence
	a._exists = true

	return nil
}

// Delete deletes the Article from the database.
func (a *Article) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !a._exists {
		return nil
	}

	// if deleted, bail
	if a._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.article WHERE id = $1`

	// run query
	XOLog(sqlstr, a.ID)
	_, err = db.Exec(sqlstr, a.ID)
	if err != nil {
		return err
	}

	// set deleted
	a._deleted = true

	return nil
}

// GetMostRecentArticle returns n most recent rows from 'public.article',
// ordered by "created_at" in descending order.
func GetMostRecentArticle(db XODB, n int) ([]*Article, error) {
	const sqlstr = `SELECT ` +
		`id, title, createddate, body, userid` +
		`FROM public.article ` +
		`ORDER BY created_at DESC LIMIT $1`

	q, err := db.Query(sqlstr, n)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	var res []*Article
	for q.Next() {
		a := Article{}

		// scan
		err = q.Scan(&a.ID, &a.Title, &a.Createddate, &a.Body, &a.Userid)
		if err != nil {
			return nil, err
		}

		res = append(res, &a)
	}

	return res, nil
}

// ArticleByID retrieves a row from 'public.article' as a Article.
//
// Generated from index 'article_pkey'.
func ArticleByID(db XODB, id int) (*Article, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, title, createddate, body, userid ` +
		`FROM public.article ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	a := Article{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&a.ID, &a.Title, &a.Createddate, &a.Body, &a.Userid)
	if err != nil {
		return nil, err
	}

	return &a, nil
}
