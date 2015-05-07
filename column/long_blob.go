package column

import (
	"fmt"
	"strings"
)

// LongBlobColumn models a column with a large binary type.
type LongBlobColumn struct {
	name      string
	modifiers []string
}

// LongBlob creates a new large binary column.
func LongBlob(name string) *LongBlobColumn {
	return &LongBlobColumn{
		name: name,
	}
}

// CreateSQL generates the SQL needed to create the column.
func (col *LongBlobColumn) CreateSQL() string {
	modifiers := strings.Join(col.modifiers, " ")
	return fmt.Sprintf("`%s` LONGBLOB %s", col.name, modifiers)
}

// NotNull flags the column so it can't contain NULLs.
func (col *LongBlobColumn) NotNull() *LongBlobColumn {
	col.modifiers = append(col.modifiers, "NOT NULL")
	return col
}
