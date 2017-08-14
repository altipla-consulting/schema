package schema

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/juju/errors"
)

// Connection to the DB.
type Connection struct {
	db           *sql.DB
	tablesEngine string
}

// NewConnection stores a connection to a DB to apply schema changes to it.
func NewConnection(db *sql.DB) *Connection {
	return &Connection{
		db:           db,
		tablesEngine: "InnoDB",
	}
}

// NewTestConnection stores a connection to a DB to apply schema changes to it. All
// tables created with this connection will reside in memory for tests.
func NewTestConnection(db *sql.DB) *Connection {
	return &Connection{
		db:           db,
		tablesEngine: "MEMORY",
	}
}

// Column it's the common interface between all type of columns.
type Column interface {

	// SQL generates the SQL needed to create the column.
	SQL() string
}

// CreateTable creates a new table.
func (conn *Connection) CreateTable(name string, columns []Column) error {
	lines := make([]string, len(columns))
	for i, col := range columns {
		lines[i] = col.SQL()
	}

	stmt := fmt.Sprintf("CREATE TABLE `%s` (%s) ENGINE=%s DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin", name, strings.Join(lines, ","), conn.tablesEngine)
	if _, err := conn.db.Exec(stmt); err != nil {
		return errors.Trace(err)
	}

	return nil
}

// DropTable drops a table.
func (conn *Connection) DropTable(name string) error {
	stmt := fmt.Sprintf("DROP TABLE `%s`", name)
	if _, err := conn.db.Exec(stmt); err != nil {
		return errors.Trace(err)
	}

	return nil
}

// TableExists checks if the table already exists in the DB.
func (conn *Connection) TableExists(name string) (bool, error) {
	var n string
	err := conn.db.QueryRow(fmt.Sprintf("SHOW TABLES LIKE '%s';", name)).Scan(&n)
	switch {
	case err == sql.ErrNoRows:
		return false, nil

	case err != nil:
		return false, errors.Trace(err)

	default:
		return true, nil
	}
}

// CreateTableIfNotExists creates the table if it is not already present.
func (conn *Connection) CreateTableIfNotExists(name string, columns []Column) error {
	exists, err := conn.TableExists(name)
	switch {
	case err != nil:
		return errors.Trace(err)

	case !exists:
		return errors.Trace(conn.CreateTable(name, columns))
	}

	return nil
}

// RenameColumn changes the name of a column. It needs the current type of the column.
// It is not recommended to change the type manually with that string (though it's possible).
func (conn *Connection) RenameColumn(tableName, oldColumnName, columnName, columnType string) error {
	stmt := fmt.Sprintf("ALTER TABLE `%s` CHANGE `%s` `%s` %s", tableName, oldColumnName, columnName, columnType)
	if _, err := conn.db.Exec(stmt); err != nil {
		return errors.Trace(err)
	}

	return nil
}

// AddColumn creates a new column in a table that already exists.
func (conn *Connection) AddColumn(tableName string, col Column) error {
	stmt := fmt.Sprintf("ALTER TABLE `%s` ADD %s", tableName, col.SQL())
	if _, err := conn.db.Exec(stmt); err != nil {
		return errors.Trace(err)
	}

	return nil
}

// AlterColumn changes the properties of a column. It needs the current type of the column.
func (conn *Connection) AlterColumn(tableName, columnName, columnType string) error {
	stmt := fmt.Sprintf("ALTER TABLE `%s` MODIFY `%s` %s", tableName, columnName, columnType)
	if _, err := conn.db.Exec(stmt); err != nil {
		return errors.Trace(err)
	}

	return nil
}

// DropColumn removes a column from a table.
func (conn *Connection) DropColumn(tableName, columnName string) error {
	stmt := fmt.Sprintf("ALTER TABLE `%s` DROP COLUMN %s", tableName, columnName)
	if _, err := conn.db.Exec(stmt); err != nil {
		return errors.Trace(err)
	}

	return nil
}

// DropPrimaryKey removes the primary key from a table (not the column, only the index).
func (conn *Connection) DropPrimaryKey(name string) error {
	stmt := fmt.Sprintf("ALTER TABLE `%s` DROP PRIMARY KEY", name)
	if _, err := conn.db.Exec(stmt); err != nil {
		return errors.Trace(err)
	}

	return nil
}

// AssignPrimaryKey sets the new primary key of the table. It should have been dropped
// before, or not exist previously.
func (conn *Connection) AssignPrimaryKey(tableName string, columnNames []string) error {
	stmt := fmt.Sprintf("ALTER TABLE `%s` ADD PRIMARY KEY (%s)", tableName, quoteCols(columnNames))
	if _, err := conn.db.Exec(stmt); err != nil {
		return errors.Trace(err)
	}

	return nil
}

// AddUnique adds a new unique index to a column.
func (conn *Connection) AddUnique(tableName, indexName string, columnNames []string) error {
	stmt := fmt.Sprintf("ALTER TABLE `%s` ADD UNIQUE INDEX %s(%s)", tableName, indexName, quoteCols(columnNames))
	if _, err := conn.db.Exec(stmt); err != nil {
		return errors.Trace(err)
	}

	return nil
}

// DropUnique removes the unique index of a column.
func (conn *Connection) DropUnique(tableName, indexName string) error {
	stmt := fmt.Sprintf("ALTER TABLE `%s` DROP INDEX %s", tableName, indexName)
	if _, err := conn.db.Exec(stmt); err != nil {
		return errors.Trace(err)
	}

	return nil
}

// RenameTable changes the name of the table to a new one.
func (conn *Connection) RenameTable(oldName, name string) error {
	stmt := fmt.Sprintf("RENAME TABLE `%s` TO `%s`", oldName, name)
	if _, err := conn.db.Exec(stmt); err != nil {
		return errors.Trace(err)
	}

	return nil
}

func quoteCols(cols []string) string {
	result := []string{}
	for _, col := range cols {
		result = append(result, fmt.Sprintf("`%s`", col))
	}
	return strings.Join(result, ",")
}
