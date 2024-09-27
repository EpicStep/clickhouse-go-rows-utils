# clickhouse-go-rows-utils [![Go Reference](https://pkg.go.dev/badge/github.com/EpicStep/clickhouse-go-rows-utils.svg)](https://pkg.go.dev/github.com/EpicStep/clickhouse-go-rows-utils) [![Lint](https://github.com/EpicStep/clickhouse-go-rows-utils/actions/workflows/lint.yml/badge.svg)](https://github.com/EpicStep/clickhouse-go-rows-utils/actions/workflows/lint.yml) [![Test](https://github.com/EpicStep/clickhouse-go-rows-utils/actions/workflows/test.yml/badge.svg)](https://github.com/EpicStep/clickhouse-go-rows-utils/actions/workflows/test.yml)

## Install

```shell
go get github.com/EpicStep/clickhouse-go-rows-utils
```

## Documentation

Documentation for all functions available at [pkg.go.dev](https://pkg.go.dev/github.com/EpicStep/clickhouse-go-rows-utils)

## Usage example

```go
package main

import (
	"context"
	
	rowsutils "github.com/EpicStep/clickhouse-go-rows-utils"
)

func main() {
	// connect to database.
	rows, err := db.Query(context.Background(), "SELECT 1")
	if err != nil {
		panic(err)
	}
	
	err = rowsutils.ForEachRow(rows, func(row rowsutils.CollectableRow) error {
		// scan your rows or do other work.
		return nil
	})
	if err != nil {
		panic(err)
	}
}
```
