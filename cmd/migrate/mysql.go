package migrate

import (
	"fmt"
	"tennet/gethired/infrastructure/repository/mysql"

	"github.com/spf13/cobra"
)

var (
	Postgres bool
)

// PostgresCmd represents the mysql command
var PostgresCmd = &cobra.Command{
	Use:   "mysql",
	Short: "Migrate PostgreSQL database",
	Long:  `The mysql command is used to migrate the PostgreSQL database to its latest schema version`,
	Run: func(cmd *cobra.Command, args []string) {
		if Postgres {
			err := mysql.MigrateMysql(PostgresDB)
			if err != nil {
				_ = fmt.Errorf("fatal error in migrating mysql: %s", err)
				panic(err)
			}
			return
		}

		cmd.Help()
	},
}

func init() {
	// migrating flag
	PostgresCmd.PersistentFlags().BoolVarP(&Postgres, "migrate", "m", false, "perform database migration")
}
