// Package clickhouse_go_rows_utils implement utils for working with rows at clickhouse-go.
package clickhouse_go_rows_utils

//go:generate go run go.uber.org/mock/mockgen -destination=./mock/rows_mock.go -package mock github.com/ClickHouse/clickhouse-go/v2/lib/driver Rows

import "github.com/ClickHouse/clickhouse-go/v2/lib/driver"

// CollectableRow is the subset of Rows methods that a RowToFunc is allowed to call.
type CollectableRow interface {
	Scan(dest ...any) error
	ScanStruct(dest any) error
	ColumnTypes() []driver.ColumnType
	Totals(dest ...any) error
	Columns() []string
}

// RowToFunc is a function that scans or otherwise converts row to a T.
type RowToFunc[T any] func(row CollectableRow) (T, error)

// CollectRows iterates through rows, calling fn for each row, and collecting the results into a slice of T.
func CollectRows[T any](rows driver.Rows, fn RowToFunc[T]) ([]T, error) {
	return AppendRows([]T{}, rows, fn)
}

// AppendRows iterates through rows, calling fn for each row, and appending the results into a slice of T.
func AppendRows[T any, S ~[]T](slice S, rows driver.Rows, fn RowToFunc[T]) (S, error) {
	err := ForEachRow(rows, func(row CollectableRow) error {
		value, err := fn(row)
		if err != nil {
			return err
		}

		slice = append(slice, value)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return slice, nil
}

// ForEachRowFunc calls at each row.
type ForEachRowFunc func(row CollectableRow) error

// ForEachRow iterates through rows, calling fn for each row.
func ForEachRow(rows driver.Rows, fn ForEachRowFunc) error {
	defer rows.Close() //nolint:errcheck

	for rows.Next() {
		if err := fn(rows); err != nil {
			return err
		}
	}

	return rows.Err()
}
