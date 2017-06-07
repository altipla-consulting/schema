package column

import (
	"fmt"
	"strings"
)

// DateColumn models a column with a date type.
type DateColumn struct {
	name      string
	modifiers []string
}

// Date creates a new date column.
func Date(name string) *DateColumn {
	return &DateColumn{
		name:      name,
		modifiers: []string{"NOT NULL"},
	}
}

// SQL generates the SQL needed to create the column.
func (col *DateColumn) SQL() string {
	modifiers := strings.Join(col.modifiers, " ")
	return fmt.Sprintf("`%s` DATE %s", col.name, modifiers)
}

// Nullable allows the column to contain NULLs.
func (col *DateColumn) Nullable() *DateColumn {
	for i, m := range col.modifiers {
		if m == "NOT NULL" {
			col.modifiers = append(col.modifiers[:i], col.modifiers[i+1:]...)
			break
		}
	}

	return col
}
