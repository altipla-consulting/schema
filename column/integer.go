package column

import (
	"fmt"
	"strings"
)

// IntegerColumn models a column with a integer type.
type IntegerColumn struct {
	name      string
	length    int
	modifiers []string
}

// Integer creates a new integer column.
func Integer(name string, length int) *IntegerColumn {
	return &IntegerColumn{
		name:   name,
		length: length,
	}
}

// CreateSQL generates the SQL needed to create the column.
func (col *IntegerColumn) CreateSQL() string {
	modifiers := strings.Join(col.modifiers, " ")
	return fmt.Sprintf("`%s` INT(%d) %s", col.name, col.length, modifiers)
}

// NotNull flags the column so it can't contain NULLs.
func (col *IntegerColumn) NotNull() *IntegerColumn {
	col.modifiers = append(col.modifiers, "NOT NULL")
	return col
}

// AutoIncrement flags the column with a default value that increments in each row inserted.
func (col *IntegerColumn) AutoIncrement() *IntegerColumn {
	col.modifiers = append(col.modifiers, "AUTO_INCREMENT")
	return col
}

// PrimaryKey flags the column as the primary key of the table.
func (col *IntegerColumn) PrimaryKey() *IntegerColumn {
	col.modifiers = append(col.modifiers, "PRIMARY KEY")
	return col
}

// After sets the column next to the one we're inserting as a reference point
// for the update/create column.
func (col *IntegerColumn) After(name string) *IntegerColumn {
	col.modifiers = append(col.modifiers, fmt.Sprintf("AFTER `%s`", name))
	return col
}

// First sets the column first in the table.
func (col *IntegerColumn) First() *IntegerColumn {
	col.modifiers = append(col.modifiers, "FIRST")
	return col
}
