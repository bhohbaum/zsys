package main

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/ubuntu/zsys/cmd/zsys/client"
	"github.com/ubuntu/zsys/cmd/zsys/daemon"
)

func main() {
	var rootCmd *cobra.Command
	var errFunc func() error

	if filepath.Base(os.Args[0]) == "zsysd" {
		rootCmd = daemon.Cmd()
		errFunc = daemon.Error
	} else {
		rootCmd = client.Cmd()
		errFunc = client.Error
	}

	if err := rootCmd.Execute(); err != nil {
		// This is a usage Error (we don't use postfix E commands other than usage)
		// Usage error should be the same format than other errors
		log.SetFormatter(&log.TextFormatter{
			DisableLevelTruncation: true,
			DisableTimestamp:       true,
		})
		log.Error(err)
		os.Exit(2)
	}
	err := errFunc()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
