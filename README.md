# seed
Package for seeding data.

It's supporting RDMS such as MySQL, MS SQL, Postgre,...; more data source types will be added in the future.

## Installation
Install the package with:
```
go get github.com/nguyendangminh/seed
```

## Documentation

Read document at [godoc](http://godoc.org/github.com/nguyendangminh/seed)

## Example

Seeding data against MySQL database

```
package mytest

import (
	"testing"
    "database/sql"

	seedsql "github.com/nguyendangminh/seed/sql"
    _ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func InitMySQLDB() (*sql.DB, error) {
    return db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname?multiStatements=true")
}

func TestDatabaseIntegration(t *testing.T) {
	db, err := InitMySQLDB()
	defer db.Close()

	// Seeding
	err = seedsql.SeedByFile(db, "/path/to/seed-file.sql")
	assert.NoError(t, err)

	// Your other tests belows

    seedsql.CleanByFile(db, "/path/to/cleanup-file.sql")
}
```