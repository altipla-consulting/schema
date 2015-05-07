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
		name: name,
	}
}

// CreateSQL generates the SQL needed to create the column.
func (col *TextColumn) CreateSQL() string {
	modifiers := strings.Join(col.modifiers, " ")
	return fmt.Sprintf("`%s` TEXT %s", col.name, modifiers)
}

// NotNull flags the column so it can't contain NULLs.
func (col *TextColumn) NotNull() *TextColumn {
	col.modifiers = append(col.modifiers, "NOT NULL")
	return col
}
