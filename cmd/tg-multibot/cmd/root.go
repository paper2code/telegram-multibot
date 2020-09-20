package cmd

import (
	"fmt"
	"os"

	"github.com/k0kubun/pp"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var (
	log = logrus.New()
)

var options struct {
	verbose         bool
	debug           bool
	generateDoc     bool
	generateDocPath string
	dbDriver        string
	dbHost          string
	dbUser          string
	dbName          string
	dbPass          string
	dbPort          int
}

// RootCmd is the root command for limo
var RootCmd = &cobra.Command{
	Use:   "tg-multibot",
	Short: "A multibot aggregating knowledge from everywhere.",
	Long:  `A multibot aggregating knowledge from multiple sources and models.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	if options.generateDoc {
		err := doc.GenMarkdownTree(RootCmd, "./docs")
		if err != nil {
			log.Fatal(err)
		}
		// todo: generate changelog and todo from code
	}
}

func init() {
	flags := RootCmd.PersistentFlags()
	flags.StringVarP(&options.dbDriver, "db-driver", "", "sqlite3", "database driver")
	flags.StringVarP(&options.dbHost, "db-host", "", os.Getenv("P2TG_MYSQL_HOST"), "database host.")
	flags.StringVarP(&options.dbUser, "db-user", "", os.Getenv("P2TG_MYSQL_USER"), "database username.")
	flags.StringVarP(&options.dbPass, "db-pass", "", os.Getenv("P2TG_MYSQL_PASSWORD"), "database password.")
	flags.StringVarP(&options.dbName, "db-name", "", os.Getenv("P2TG_MYSQL_DATABASE"), "database name.")
	flags.IntVarP(&options.dbPort, "db-port", "", 3306, "database port.")
	flags.BoolVarP(&options.generateDoc, "doc-gen", "g", false, "generate commands/sub-commands documentation.")
	flags.StringVarP(&options.generateDocPath, "doc-path", "", "./docs", "destination directory for documentation.")
	flags.BoolVarP(&options.debug, "debug", "d", false, "debug mode.")
	flags.BoolVarP(&options.verbose, "verbose", "v", false, "verbose output.")
}

// fail fast on initialization
func must(i interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return i
}
