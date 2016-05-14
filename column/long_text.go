package column

import (
	"fmt"
	"strings"
)

// LongTextColumn models a column with a large ASCII text type.
type LongTextColumn struct {
	name      string
	modifiers []string
}

// LongText creates a new large ASCII text column.
func LongText(name string) *LongTextColumn {
	return &LongTextColumn{
		name:      name,
		modifiers: []string{"NOT NULL"},
	}
}

// SQL generates the SQL needed to create the column.
func (col *LongTextColumn) SQL() string {
	modifiers := strings.Join(col.modifiers, " ")
	return fmt.Sprintf("`%s` LONGTEXT %s", col.name, modifiers)
}

// Nullable allows the column to contain NULLs.
func (col *LongTextColumn) Nullable() *LongTextColumn {
	for i, m := range col.modifiers {
		if m == "NOT NULL" {
			col.modifiers = append(col.modifiers[:i], col.modifiers[i+1:]...)
			break
		}
	}

	return col
}
