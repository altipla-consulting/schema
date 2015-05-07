package column

import (
	"fmt"
	"strings"
)

// BooleanColumn models a column with a boolean type.
type BooleanColumn struct {
	name      string
	modifiers []string
}

// Boolean creates a new boolean column.
func Boolean(name string) *BooleanColumn {
	return &BooleanColumn{
		name: name,
	}
}

// CreateSQL generates the SQL needed to create the column.
func (col *BooleanColumn) CreateSQL() string {
	modifiers := strings.Join(col.modifiers, " ")
	return fmt.Sprintf("`%s` TINYINT(1) %s", col.name, modifiers)
}

// NotNull flags the column so it can't contain NULLs.
func (col *BooleanColumn) NotNull() *BooleanColumn {
	col.modifiers = append(col.modifiers, "NOT NULL")
	return col
}
