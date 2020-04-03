// Code generated by yo. DO NOT EDIT.
// Package models contains the types.
package models

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"
	"google.golang.org/grpc/codes"
)

// Mytable represents a row from 'mytable'.
type Mytable struct {
	A spanner.NullInt64 `spanner:"a" json:"a"` // a
	B spanner.NullInt64 `spanner:"b" json:"b"` // b
}

func MytablePrimaryKeys() []string {
	return []string{
		"a",
	}
}

func MytableColumns() []string {
	return []string{
		"a",
		"b",
	}
}

func (m *Mytable) columnsToPtrs(cols []string, customPtrs map[string]interface{}) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		if val, ok := customPtrs[col]; ok {
			ret = append(ret, val)
			continue
		}

		switch col {
		case "a":
			ret = append(ret, &m.A)
		case "b":
			ret = append(ret, &m.B)
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}
	return ret, nil
}

func (m *Mytable) columnsToValues(cols []string) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		switch col {
		case "a":
			ret = append(ret, m.A)
		case "b":
			ret = append(ret, m.B)
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}

	return ret, nil
}

// newMytable_Decoder returns a decoder which reads a row from *spanner.Row
// into Mytable. The decoder is not goroutine-safe. Don't use it concurrently.
func newMytable_Decoder(cols []string) func(*spanner.Row) (*Mytable, error) {
	customPtrs := map[string]interface{}{}

	return func(row *spanner.Row) (*Mytable, error) {
		var m Mytable
		ptrs, err := m.columnsToPtrs(cols, customPtrs)
		if err != nil {
			return nil, err
		}

		if err := row.Columns(ptrs...); err != nil {
			return nil, err
		}

		return &m, nil
	}
}

// Insert returns a Mutation to insert a row into a table. If the row already
// exists, the write or transaction fails.
func (m *Mytable) Insert(ctx context.Context) *spanner.Mutation {
	return spanner.Insert("mytable", MytableColumns(), []interface{}{
		m.A, m.B,
	})
}

// Update returns a Mutation to update a row in a table. If the row does not
// already exist, the write or transaction fails.
func (m *Mytable) Update(ctx context.Context) *spanner.Mutation {
	return spanner.Update("mytable", MytableColumns(), []interface{}{
		m.A, m.B,
	})
}

// InsertOrUpdate returns a Mutation to insert a row into a table. If the row
// already exists, it updates it instead. Any column values not explicitly
// written are preserved.
func (m *Mytable) InsertOrUpdate(ctx context.Context) *spanner.Mutation {
	return spanner.InsertOrUpdate("mytable", MytableColumns(), []interface{}{
		m.A, m.B,
	})
}

// UpdateColumns returns a Mutation to update specified columns of a row in a table.
func (m *Mytable) UpdateColumns(ctx context.Context, cols ...string) (*spanner.Mutation, error) {
	// add primary keys to columns to update by primary keys
	colsWithPKeys := append(cols, MytablePrimaryKeys()...)

	values, err := m.columnsToValues(colsWithPKeys)
	if err != nil {
		return nil, newErrorWithCode(codes.InvalidArgument, "Mytable.UpdateColumns", "mytable", err)
	}

	return spanner.Update("mytable", colsWithPKeys, values), nil
}

// FindMytable gets a Mytable by primary key
func FindMytable(ctx context.Context, db YORODB, a spanner.NullInt64) (*Mytable, error) {
	key := spanner.Key{a}
	row, err := db.ReadRow(ctx, "mytable", key, MytableColumns())
	if err != nil {
		return nil, newError("FindMytable", "mytable", err)
	}

	decoder := newMytable_Decoder(MytableColumns())
	m, err := decoder(row)
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "FindMytable", "mytable", err)
	}

	return m, nil
}

// Delete deletes the Mytable from the database.
func (m *Mytable) Delete(ctx context.Context) *spanner.Mutation {
	values, _ := m.columnsToValues(MytablePrimaryKeys())
	return spanner.Delete("mytable", spanner.Key(values))
}
