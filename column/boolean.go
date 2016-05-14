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
		name:      name,
		modifiers: []string{"NOT NULL"},
	}
}

// SQL generates the SQL needed to create the column.
func (col *BooleanColumn) SQL() string {
	modifiers := strings.Join(col.modifiers, " ")
	return fmt.Sprintf("`%s` TINYINT(1) %s", col.name, modifiers)
}

// Nullable allows the column to contain NULLs.
func (col *BooleanColumn) Nullable() *BooleanColumn {
	for i, m := range col.modifiers {
		if m == "NOT NULL" {
			col.modifiers = append(col.modifiers[:i], col.modifiers[i+1:]...)
			break
		}
	}

	return col
}
