package schema

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/altipla-consulting/database"
	"github.com/altipla-consulting/schema/column"
	"github.com/juju/errors"
	"golang.org/x/net/context"
)

// Table stores information about a table
type Table struct {
	name    string
	columns []column.Column
}

// NewTable initializes the info about a new table.
func NewTable(name string) *Table {
	return &Table{name: name}
}

// Create sends the commands to create a new table in the DB.
func (t *Table) Create(ctx context.Context) error {
	conn := database.FromContext(ctx)

	lines := make([]string, len(t.columns))
	for i, col := range t.columns {
		lines[i] = col.CreateSQL()
	}

	content := strings.Join(lines, ",")
	statement := fmt.Sprintf("CREATE TABLE `%s` (%s) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci", t.name, content)
	if _, err := conn.DB.Exec(statement); err != nil {
		return errors.Trace(err)
	}

	return nil
}

// Drop sends the commands to drop the table in the DB.
func (t *Table) Drop(ctx context.Context) error {
	conn := database.FromContext(ctx)

	statement := fmt.Sprintf("DROP TABLE `%s`", t.name)
	if _, err := conn.DB.Exec(statement); err != nil {
		return errors.Trace(err)
	}

	return nil
}

// Exists checks if the table is already present in the DB.
func (t *Table) Exists(ctx context.Context) (bool, error) {
	conn := database.FromContext(ctx)

	var name string
	err := conn.DB.QueryRow(fmt.Sprintf("SHOW TABLES LIKE '%s';", t.name)).Scan(&name)
	switch {
	case err == sql.ErrNoRows:
		return false, nil

	case err != nil:
		return false, errors.Trace(err)

	default:
		return true, nil
	}
}

// CreateIfNotExists creates the table if it isn't already present.
func (t *Table) CreateIfNotExists(ctx context.Context) error {
	exists, err := t.Exists(ctx)
	switch {
	case err != nil:
		return errors.Trace(err)

	case !exists:
		if err := t.Create(ctx); err != nil {
			return errors.Trace(err)
		}
	}

	return nil
}

// AddColumn add a new column to the list that will be added when the table is created.
func (t *Table) AddColumn(column column.Column) {
	t.columns = append(t.columns, column)
}

// RenameColumn changes the name of a column. It needs the current type of the column
// and is not recommended to change the type manually with that string (though it's possible).
func (t *Table) RenameColumn(ctx context.Context, oldName, name, columnType string) error {
	conn := database.FromContext(ctx)

	statement := fmt.Sprintf("ALTER TABLE `%s` CHANGE `%s` `%s` %s", t.name, oldName, name, columnType)
	if _, err := conn.DB.Exec(statement); err != nil {
		return errors.Trace(err)
	}

	return nil
}

// CreateColumn directly creates a new column in an already present table.
func (t *Table) CreateColumn(ctx context.Context, col column.Column) error {
	conn := database.FromContext(ctx)

	statement := fmt.Sprintf("ALTER TABLE `%s` ADD %s", t.name, col.CreateSQL())
	if _, err := conn.DB.Exec(statement); err != nil {
		return errors.Trace(err)
	}

	return nil
}

// DropColumn removes a column from the table.
func (t *Table) DropColumn(ctx context.Context, name string) error {
	conn := database.FromContext(ctx)

	statement := fmt.Sprintf("ALTER TABLE `%s` DROP COLUMN %s", t.name, name)
	if _, err := conn.DB.Exec(statement); err != nil {
		return errors.Trace(err)
	}

	return nil
}

// DropPrimaryKey removes the primary key from a table (not the column, only the index).
func (t *Table) DropPrimaryKey(ctx context.Context) error {
	conn := database.FromContext(ctx)

	statement := fmt.Sprintf("ALTER TABLE `%s` DROP PRIMARY KEY", t.name)
	if _, err := conn.DB.Exec(statement); err != nil {
		return errors.Trace(err)
	}

	return nil
}

// AssignPrimaryKey sets the new primary key of the table (it should have been dropped
// before, or not exist in the first place)
func (t *Table) AssignPrimaryKey(ctx context.Context, name string) error {
	conn := database.FromContext(ctx)

	statement := fmt.Sprintf("ALTER TABLE `%s` ADD PRIMARY KEY (`%s`)", t.name, name)
	if _, err := conn.DB.Exec(statement); err != nil {
		return errors.Trace(err)
	}

	return nil
}

// AddUnique adds a new unique index to a column.
func (t *Table) AddUnique(ctx context.Context, column string) error {
	conn := database.FromContext(ctx)

	statement := fmt.Sprintf("ALTER TABLE `%s` ADD UNIQUE INDEX (`%s`)", t.name, column)
	if _, err := conn.DB.Exec(statement); err != nil {
		return errors.Trace(err)
	}

	return nil
}

// DropUnique removes the unique index of a column.
func (t *Table) DropUnique(ctx context.Context, column string) error {
	conn := database.FromContext(ctx)

	statement := fmt.Sprintf("ALTER TABLE `%s` DROP INDEX `%s`", t.name, column)
	if _, err := conn.DB.Exec(statement); err != nil {
		return errors.Trace(err)
	}

	return nil
}

// Rename changes the name of the table to a new one.
func (t *Table) Rename(ctx context.Context, name string) error {
	conn := database.FromContext(ctx)

	statement := fmt.Sprintf("RENAME TABLE `%s` TO `%s`", t.name, name)
	if _, err := conn.DB.Exec(statement); err != nil {
		return errors.Trace(err)
	}

	return nil
}
