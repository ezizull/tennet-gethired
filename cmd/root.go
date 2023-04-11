package cmd

import (
	"os"
	"tennet/gethired/cmd/migrate"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "toolbox",
	Short: "A command line tool for executing custom commands based on your needs.",
	Long: `The toolbox command line service allows you 
        to perform a variety of extra commands tailored to 
        your needs. It offers a range of functionalities, 
        from simple to complex, that you can use to streamline 
        your workflow. Whether you want to automate repetitive 
        tasks or perform advanced operations, the toolbox is a 
        versatile and easy-to-use solution.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(mysql *gorm.DB) {
	// set migrating database
	migrate.SetMigrateDB(mysql)

	// mysql migrating flag
	rootCmd.AddCommand(migrate.MysqlCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
