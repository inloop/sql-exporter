package cmd

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql" // for mysql support
	"github.com/joho/sqltocsv"
	_ "github.com/lib/pq" // for postgres support
	"github.com/urfave/cli"
)

// Dump ...
func Dump() cli.Command {
	return cli.Command{
		Name: "dump",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "u,db-url",
				Usage:  "database connection string ({type}://{username}:{password}@{host}/{db})",
				EnvVar: "DATABASE_URL",
			},
			cli.StringFlag{
				Name:   "q,query",
				Usage:  "SQL query as source for csv data",
				EnvVar: "QUERY",
			},
			cli.StringFlag{
				Name:   "o,output",
				Usage:  "output file name",
				EnvVar: "OUTPUT",
				Value:  "output.csv",
			},
		},
		Action: func(c *cli.Context) error {
			connStr := c.String("u")
			query := c.String("q")
			output := c.String("o")

			if err := dumpSQL(connStr, query, output); err != nil {
				return cli.NewExitError(err, 1)
			}
			return nil
		},
	}
}

func dumpSQL(connStr, query, filename string) error {

	if connStr == "" {
		return fmt.Errorf("db-url attribute not provided [DATABASE_URL]")
	}
	if query == "" {
		return fmt.Errorf("query attribute not provided [QUERY]")
	}

	connType := strings.Split(connStr, ":")[0]
	db, err := sql.Open(connType, connStr)
	if err != nil {
		return cli.NewExitError(fmt.Sprintf("Error opening testdb %v", err), 1)
	}

	rows, err := db.Query(query)
	if err != nil {
		return err
	}

	fmt.Printf("Exporting to file %s", filename)

	return sqltocsv.WriteFile(filename, rows)
}