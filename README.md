# bigint

bigint is a wrapper around math/big package to let us use big.int type in postgresql.

This project is forked from https://github.com/d-fal/bigint.

## Example use with go-pg

**go-pg** is an amazing orm for gophers to utilize postgres. This package is used to help **go-pg** users implement **math/big** functionalities.

```go
package main

import (
	"github.com/iam047801/bigint"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func main() {
	type UserBalance struct {
		tableName struct{} `pg:"balances"`

		UserID uint64         `pg:",pk"`
		Value  *bigint.BigInt `pg:"type:numeric"`
	}

	db := pg.Connect(&pg.Options{
		Addr:     "127.0.0.1:5432",
		User:     "postgres",
		Password: "postgres",
		Database: "postgres",
	})

	err := db.Model((*UserBalance)(nil)).CreateTable(&orm.CreateTableOptions{
		Temp:          true,
		FKConstraints: true,
		IfNotExists:   true,
	})
	if err != nil {
		panic(err)
	}
}
```
