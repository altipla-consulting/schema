package column

// AutoIncrement creates a new integer column that autoincrements and it's
// the primary key of the table.
func AutoIncrement() *IntegerColumn {
	return Integer("id", 11).AutoIncrement().PrimaryKey()
}
