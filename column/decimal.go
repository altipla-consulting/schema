package column

import (
	"fmt"
	"strings"
)

// DecimalColumn models a column with a decimal type.
type DecimalColumn struct {
	name                     string
	length, decimalPositions int
	modifiers                []string
}

// Decimal creates a new decimal column.
func Decimal(name string, length, decimalPositions int) *DecimalColumn {
	return &DecimalColumn{
		name:             name,
		length:           length,
		decimalPositions: decimalPositions,
		modifiers:        []string{"NOT NULL"},
	}
}

// SQL generates the SQL needed to create the column.
func (col *DecimalColumn) SQL() string {
	modifiers := strings.Join(col.modifiers, " ")
	return fmt.Sprintf("`%s` DECIMAL(%d, %d) %s", col.name, col.length, col.decimalPositions, modifiers)
}

// Nullable allows the column to contain NULLs.
func (col *DecimalColumn) Nullable() *DecimalColumn {
	for i, m := range col.modifiers {
		if m == "NOT NULL" {
			col.modifiers = append(col.modifiers[:i], col.modifiers[i+1:]...)
			break
		}
	}

	return col
}
