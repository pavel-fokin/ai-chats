package sqlite

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTx_Tx(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)

	db, err := NewDB("file:memdb1?mode=memory&cache=shared")
	// db, err := NewDB("/tmp/test.db")
	assert.NoError(err)
	defer db.Close()

	_, err = db.Exec(
		`CREATE TABLE user (
			id TEXT PRIMARY KEY,
			username TEXT NOT NULL
		);

		CREATE TABLE chat (
			id TEXT PRIMARY KEY,
			user_id TEXT NOT NULL,
			FOREIGN KEY (user_id) REFERENCES user (id)
		);
		`,
	)
	assert.NoError(err)

	tx := NewTx(db)

	t.Run("success", func(t *testing.T) {
		err := tx.Tx(ctx, func(ctx context.Context) error {
			db := MaybeHaveTx(ctx)

			_, err := db.ExecContext(
				ctx, "INSERT INTO user (id, username) VALUES (?, ?)", "1", "username",
			)
			if err != nil {
				return err
			}

			_, err = db.Exec("INSERT INTO chat (id, user_id) VALUES (?, ?)", "1", "1")
			if err != nil {
				return err
			}

			return nil
		})
		assert.NoError(err)

		rows, err := db.Query("SELECT * FROM user WHERE id = 1")
		assert.NoError(err)
		defer rows.Close()

		var id, username string
		for rows.Next() {
			err = rows.Scan(&id, &username)
			assert.NoError(err)
		}
		assert.Equal("1", id)
	})

	t.Run("error", func(t *testing.T) {
		err := tx.Tx(ctx, func(ctx context.Context) error {
			db := MaybeHaveTx(ctx)

			_, err := db.Exec("INSERT INTO user (id, username) VALUES (?, ?)", "2", "username")
			if err != nil {
				return err
			}

			_, err = db.Exec("INSERT INTO non_existing_table (value) VALUES (?)", "this will fail")
			if err != nil {
				return err
			}

			return nil
		})
		assert.ErrorContains(err, "transaction failed: SQL logic error: no such table: non_existing_table")

		row := db.QueryRow("SELECT * FROM user WHERE id = 2")

		var id, username string
		err = row.Scan(&id, &username)
		assert.ErrorIs(err, sql.ErrNoRows)

	})
}
