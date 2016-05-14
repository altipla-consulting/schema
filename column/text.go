package column

import (
	"fmt"
	"strings"
)

// TextColumn models a column with a text type.
type TextColumn struct {
	name      string
	modifiers []string
}

// Text creates a new text column.
func Text(name string) *TextColumn {
	return &TextColumn{
		name:      name,
		modifiers: []string{"NOT NULL"},
	}
}

// SQL generates the SQL needed to create the column.
func (col *TextColumn) SQL() string {
	modifiers := strings.Join(col.modifiers, " ")
	return fmt.Sprintf("`%s` TEXT %s", col.name, modifiers)
}

// Nullable allows the column to contain NULLs.
func (col *TextColumn) Nullable() *TextColumn {
	for i, m := range col.modifiers {
		if m == "NOT NULL" {
			col.modifiers = append(col.modifiers[:i], col.modifiers[i+1:]...)
			break
		}
	}

	return col
}
