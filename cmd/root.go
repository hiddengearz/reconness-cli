package cmd

import (
	"os"

	"github.com/hiddengearz/reconness-cli/internal"
	"github.com/hiddengearz/reconness-cli/internal/db"
	l "github.com/hiddengearz/reconness-cli/internal/logger"
	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string
	//userLicense      string
	Username string
	Password string
	//Email            string
	//FirstName        string
	//LastName         string
	//getAllProjects   bool
	//getAllSubDomains bool
	//projectName      string
	//subdomainName    string

	RootCmd = &cobra.Command{
		Use:   "Laelap",
		Short: "Agent for the [redacted] project",
		Long:  `[redacted]`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			f, err := os.OpenFile("log.info", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				log.Fatal(err)

			}
			if internal.Debug && internal.SilentMode {
				log.Error("Please choose Debug mode or silent mode. Not both.")

			} else if internal.Debug {
				l.InitDetailedLogger(f)
				log.Debug("Debug mode enabled")

			} else if internal.SilentMode {
				log.SetLevel(log.PanicLevel)

			}

		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			db.Cleanup()
			db.Close()
		},
	}
)

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(db.InitDB)

	RootCmd.PersistentFlags().BoolVarP(&internal.Debug, "debug", "D", false, "Enable Debug mode")
	RootCmd.PersistentFlags().BoolVarP(&internal.SilentMode, "silent", "S", false, "Enable Silent mode")
	RootCmd.PersistentFlags().BoolVarP(&internal.OfflineMode, "offline", "O", false, "Enable Offline mode")
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "Config", "C", "", "Specify a config File to be used")

	//internal.CoreClientCfg = cclient.DefaultTransportConfig().WithHost("localhost:8081")
	//internal.CoreClient = cclient.NewHTTPClientWithConfig(nil, internal.CoreClientCfg)

	//internal.UserClientCfg = cclient.DefaultTransportConfig().WithHost("localhost:8082")
	//internal.UserClient = cclient.NewHTTPClientWithConfig(nil, internal.UserClientCfg)

}

//Reads the config & reads variables saved in it.
func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	if cfgFile != "" { // Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else { //default is to check the home dir ~/reconness/rcli.yaml

		viper.AddConfigPath(home + "/reconness/")
		viper.SetConfigName("rcli.yaml")
		viper.SetConfigType("yaml")
	}

	//viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil { //read the config
		log.Debug("Using config file:", viper.ConfigFileUsed())
	} else { //otherwise create a config file in the default location
		log.Error(err)

		path := home + "/reconness/"

		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			os.Mkdir(path, 0700)
		}

		log.Info("creating config file at ", path+"rcli.yaml")

		err = viper.SafeWriteConfigAs(path + "rcli.yaml")
		if err != nil {
			log.Fatal(err)
		}

		initConfig()

	}

}
