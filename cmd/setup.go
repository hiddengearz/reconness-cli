package cmd

import (
	"github.com/hiddengearz/reconness-cli/internal"
	"github.com/hiddengearz/reconness-cli/sdk/client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	SetupCmd.Flags().StringVarP(&Username, "username", "u", "default", "username (required)")
	SetupCmd.Flags().StringVarP(&Password, "password", "p", "default", "password (required)")
	SetupCmd.Flags().StringVarP(&internal.EncKey, "encryption key", "k", "", "(opional) Pass key used to encrypt data stored on disk. Key can be any length")
	SetupCmd.Flags().StringVarP(&internal.Hostname, "hostname", "h", "reconness.com", "(opional) URL for the reconness server") //really should be mandatory for now

	SetupCmd.MarkFlagRequired("username")
	SetupCmd.MarkFlagRequired("password")
	SetupCmd.MarkFlagRequired("hostname")

	RootCmd.AddCommand(SetupCmd)

}

var SetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "s",
	Long:  `Setup the reconness CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("SetupCmd executed")
		//-u for username
		// -p for password
		// -e for encryption (optional)
		//-h for host (optional, default to cloud)
		if internal.Hostname != "" {
			internal.ClientCfg = client.DefaultTransportConfig().WithHost("internal.Hostname")
		} else {
			internal.ClientCfg = client.DefaultTransportConfig()
		}
		internal.Client = client.NewHTTPClientWithConfig(nil, internal.ClientCfg)

		//Login with credentials, save JWT
		//If login successful and encryption is enabled encrypt and store credentials in badger

		/*
			_, _, err := auth.Login(&Username, &Password)
			if err != nil {
				log.Error(err)
			}
		*/
	},
}
