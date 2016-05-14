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
		name:      name,
		modifiers: []string{"NOT NULL"},
	}
}

// SQL generates the SQL needed to create the column.
func (col *LongBlobColumn) SQL() string {
	modifiers := strings.Join(col.modifiers, " ")
	return fmt.Sprintf("`%s` LONGBLOB %s", col.name, modifiers)
}

// Nullable allows the column to contain NULLs.
func (col *LongBlobColumn) Nullable() *LongBlobColumn {
	for i, m := range col.modifiers {
		if m == "NOT NULL" {
			col.modifiers = append(col.modifiers[:i], col.modifiers[i+1:]...)
			break
		}
	}

	return col
}
