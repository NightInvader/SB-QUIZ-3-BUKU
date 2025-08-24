package database

import (
	"database/sql"
	"embed"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

//go:embed sql_migrations/*.sql
var dbMigrations embed.FS

var Dbconnection *sql.DB

func DBMigrate(dbParam *sql.DB) {
	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "sq_migration",
	}

	n, errs := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
	if errs != nil {
		panic(errs)
	}

	Dbconnection = dbParam

	fmt.Println("Migration Success, applied", n, "migrations!")

}
