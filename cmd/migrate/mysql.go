package migrate

import (
	"fmt"
	"tennet/gethired/infrastructure/repository/mysql"

	"github.com/spf13/cobra"
)

var (
	Mysql bool
)

// MysqlCmd represents the mysql command
var MysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "Migrate PostgreSQL database",
	Long:  `The mysql command is used to migrate the PostgreSQL database to its latest schema version`,
	Run: func(cmd *cobra.Command, args []string) {
		if Mysql {
			err := mysql.MigrateMysql(MysqlDB)
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
	MysqlCmd.PersistentFlags().BoolVarP(&Mysql, "migrate", "m", false, "perform database migration")
}
