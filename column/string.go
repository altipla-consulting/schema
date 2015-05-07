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
		name:   name,
		length: length,
	}
}

// CreateSQL generates the SQL needed to create the column.
func (col *StringColumn) CreateSQL() string {
	modifiers := strings.Join(col.modifiers, " ")
	return fmt.Sprintf("`%s` VARCHAR(%d) %s", col.name, col.length, modifiers)
}

// NotNull flags the column so it can't contain NULLs.
func (col *StringColumn) NotNull() *StringColumn {
	col.modifiers = append(col.modifiers, "NOT NULL")
	return col
}

// PrimaryKey flags the column as the primary key of the table.
func (col *StringColumn) PrimaryKey() *StringColumn {
	col.modifiers = append(col.modifiers, "PRIMARY KEY")
	return col
}
