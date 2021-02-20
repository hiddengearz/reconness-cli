package cmd

import (
	"os"

	"github.com/hiddengearz/reconness-cli/internal"
	l "github.com/hiddengearz/reconness-cli/internal/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile          string
	userLicense      string
	Username         string
	Password         string
	Email            string
	FirstName        string
	LastName         string
	getAllProjects   bool
	getAllSubDomains bool
	projectName      string
	subdomainName    string

	RootCmd = &cobra.Command{
		Use:   "Laelap",
		Short: "Agent for the [redacted] project",
		Long:  `[redacted]`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			f, err := os.OpenFile("log.info", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				l.Log.Fatal(err)
			}
			if internal.Debug && internal.SilentMode {
				l.Log.Error("Please choose Debug mode or silent mode. Not both.")
			} else if internal.Debug {

				l.InitDetailedLogger(f)
				l.Log.SetLevel(logrus.DebugLevel)

				l.Log.Debug("Debug mode enabled")
			} else if internal.SilentMode {
				l.Log.SetLevel(logrus.PanicLevel)
			}

		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			db.Cleanup()
			db.Close()
		},
	}
)

func Execute() {

}
