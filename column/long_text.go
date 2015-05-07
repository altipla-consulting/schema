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
		name: name,
	}
}

// CreateSQL generates the SQL needed to create the column.
func (col *LongTextColumn) CreateSQL() string {
	modifiers := strings.Join(col.modifiers, " ")
	return fmt.Sprintf("`%s` LONGTEXT %s", col.name, modifiers)
}

// NotNull flags the column so it can't contain NULLs.
func (col *LongTextColumn) NotNull() *LongTextColumn {
	col.modifiers = append(col.modifiers, "NOT NULL")
	return col
}
