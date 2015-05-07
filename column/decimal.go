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
	}
}

// CreateSQL generates the SQL needed to create the column.
func (col *DecimalColumn) CreateSQL() string {
	modifiers := strings.Join(col.modifiers, " ")
	return fmt.Sprintf("`%s` DECIMAL(%d, %d) %s", col.name, col.length, col.decimalPositions, modifiers)
}

// NotNull flags the column so it can't contain NULLs.
func (col *DecimalColumn) NotNull() *DecimalColumn {
	col.modifiers = append(col.modifiers, "NOT NULL")
	return col
}
