package column

import (
	"fmt"
)

// TimestampColumn models a column with a timestamp type.
type TimestampColumn struct {
	name string
}

// Timestamp creates a new timestamp column.
func Timestamp(name string) *TimestampColumn {
	return &TimestampColumn{
		name: name,
	}
}

// SQL generates the SQL needed to create the column.
func (col *TimestampColumn) SQL() string {
	return fmt.Sprintf("`%s` TIMESTAMP NOT NULL DEFAULT '0000-00-00 00:00:00'", col.name)
}

// CreatedTimestamp creates a new column called "created_at" thought to store
// the creation timestamp of the tuple.
func CreatedTimestamp() *TimestampColumn {
	return Timestamp("created_at")
}

// UpdatedTimestamp creates a new column called "updated_at" thought to store
// the update timestamp of the tuple.
func UpdatedTimestamp() *TimestampColumn {
	return Timestamp("updated_at")
}
