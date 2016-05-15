package column

import (
	"fmt"
	"strings"
)

// DateTimeColumn models a column with a datetime type.
type DateTimeColumn struct {
	name      string
	modifiers []string
}

// DateTime creates a new datetime column.
func DateTime(name string) *DateTimeColumn {
	return &DateTimeColumn{
		name:      name,
		modifiers: []string{"NOT NULL"},
	}
}

// SQL generates the SQL needed to create the column.
func (col *DateTimeColumn) SQL() string {
	modifiers := strings.Join(col.modifiers, " ")
	return fmt.Sprintf("`%s` DATETIME %s", col.name, modifiers)
}

// Nullable allows the column to contain NULLs.
func (col *DateTimeColumn) Nullable() *DateTimeColumn {
	for i, m := range col.modifiers {
		if m == "NOT NULL" {
			col.modifiers = append(col.modifiers[:i], col.modifiers[i+1:]...)
			break
		}
	}

	return col
}

// DefaultCurrent sets the default value of the column to the current datetime.
func (col *DateTimeColumn) DefaultCurrent() *DateTimeColumn {
	col.modifiers = append(col.modifiers, "DEFAULT CURRENT_TIMESTAMP")
	return col
}
