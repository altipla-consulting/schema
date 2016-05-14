package column

// Column it's the common interface between all type of columns.
type Column interface {

	// SQL generates the SQL needed to create the column.
	SQL() string
}
