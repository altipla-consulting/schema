package column

import (
	"fmt"
	"strings"
)

// StringColumn models a column with a string type.
type StringColumn struct {
	name      string
	length    int
	modifiers []string
}

// String creates a new string column.
func String(name string, length int) *StringColumn {
	return &StringColumn{
		name:      name,
		length:    length,
		modifiers: []string{"NOT NULL"},
	}
}

// SQL generates the SQL needed to create the column.
func (col *StringColumn) SQL() string {
	modifiers := strings.Join(col.modifiers, " ")
	return fmt.Sprintf("`%s` VARCHAR(%d) %s", col.name, col.length, modifiers)
}

// Nullable allows the column to contain NULLs.
func (col *StringColumn) Nullable() *StringColumn {
	for i, m := range col.modifiers {
		if m == "NOT NULL" {
			col.modifiers = append(col.modifiers[:i], col.modifiers[i+1:]...)
			break
		}
	}

	return col
}

// PrimaryKey flags the column as the primary key of the table.
func (col *StringColumn) PrimaryKey() *StringColumn {
	col.modifiers = append(col.modifiers, "PRIMARY KEY")
	return col
}

// After sets the column next to the one we're inserting as a reference point
// for the update/create column.
func (col *StringColumn) After(name string) *StringColumn {
	col.modifiers = append(col.modifiers, fmt.Sprintf("AFTER `%s`", name))
	return col
}
