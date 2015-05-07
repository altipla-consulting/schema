// Package schema helps to migrate the schema of a MySQL database.
// 
// The context all functions receive has to be initialized before using the package
// https://github.com/altipla-consulting/database. Example:
// 
//     ctx := context.NewBackground()
//     var err error
//     ctx, err = database.WithContext(ctx, "username", "password", "unix(/var/lib/mysql/mysql.sock)", "database")
//     if err != nil {
//       log.Fatal(err)
//     }
//     
//     // ... use ctx to call the functions in this package
//
package schema
