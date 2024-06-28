package core

import (
	"encoding/json"
	"os"
	"time"

	"github.com/backend/middleware/config"
	charmLog "github.com/charmbracelet/log"
	"github.com/google/uuid"
)

var (
	connectionCredentials = make(map[string]time.Time)
	log                   = charmLog.NewWithOptions(os.Stderr, charmLog.Options{
		ReportCaller:    false,
		ReportTimestamp: true,
		TimeFormat:      time.DateTime,
		Prefix:          "Hex Credentials 👾 ",
	})
)

func JobCredentials() {
	go func() {
		for {
			time.Sleep(1 * time.Second)
			log.Info("server started", "connection", GetServicesConnections())
			for credential := range connectionCredentials {
				ValidateCredentials(credential)
			}
		}
	}()
}

func NewCredentials() string {
	credential := uuid.New().String()
	connectionCredentials[credential] = time.Now()
	log.Info("new credential generated", "credential", credential)
	return credential
}

func ValidateCredentials(credential string) bool {
	foundCredential, ok := connectionCredentials[credential]
	if !ok {
		log.Debug("credential not found", "credential", credential)
		return false
	}
	if time.Since(foundCredential) > config.CredentialLifetime*time.Second {
		log.Debug("credential expired", "credential", credential)
		delete(connectionCredentials, credential)
		log.Info("credential deleted", "credential", credential)
		return false
	}
	log.Debug("credential valid", "credential", credential)
	return true
}

func GetServicesConnections() string {
	info, err := json.MarshalIndent(&connectionCredentials, "", "  ")
	if err != nil {
		log.Error("error marshalling connections", "error", err)
		return ""
	}
	return string(info)
}
