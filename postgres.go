package athena

import (
	"crypto/tls"
	"fmt"
	"strings"

	"github.com/go-pg/pg/v9"
)

// PostgresEnv is the definition of environment variables needed
// to setup a postgres connection
type PostgresEnv struct {
	PostgresHost       string `required:"true" split_words:"true"`
	PostgresPort       int    `required:"true" split_words:"true"`
	PostgresSecureMode bool   `required:"true" split_words:"true"`
	PostgresUser       string `required:"true" split_words:"true"`
	PostgresPassword   string `required:"true" split_words:"true"`
	PostgresDatabase   string `required:"true" split_words:"true"`
}

// NewDB creates a connection to a postgres DB and ensures the connection is live.
func NewDB(env PostgresEnv) (*pg.DB, error) {
	opts := &pg.Options{
		Addr:     fmt.Sprintf("%s:%d", env.PostgresHost, env.PostgresPort),
		User:     env.PostgresUser,
		Password: env.PostgresPassword,
		Database: env.PostgresDatabase,
	}

	if env.PostgresSecureMode {
		opts.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}

	db := pg.Connect(opts)
	_, err := db.Exec("select 1")

	return db, err
}

// CleanUpTables removes all the rows in the passed tables. It is
// useful for cleaning up the DB for tests.
func CleanUpTables(db *pg.DB, tables ...string) error {
	query := fmt.Sprintf("truncate %s cascade", strings.Join(tables, ","))
	if _, err := db.Exec(query); err != nil {
		return err
	}

	return nil
}
