package auth

import (
	"github.com/hiddengearz/reconness-cli/internal"
	"github.com/hiddengearz/reconness-cli/sdk/client/auth"
	log "github.com/sirupsen/logrus"
)

func Login(username *string, password *string) (err error) {
	log.Debug("Attempting to login user")

	params := auth.NewPostAPIAuthLoginParams()
	params.Body.UserName = username
	params.Body.Password = password

	data, err := internal.Client.Auth.PostAPIAuthLogin(params, nil)
	if err != nil {
		log.Debug("Failed to login with given credentials")
		return err
	}
	//swagger is missing the response, unable to get JWT

}
