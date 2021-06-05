package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/sky4access/extractor/cmd"
	"os"
)

var commit = "unknown-commit"

func main() {

	contextLogger := log.WithFields(log.Fields{
		"commitHas": commit,
		"dateTime":  "",
	})

	command := cmd.NewCommand(contextLogger)

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stdout, "%v\n", err)
	}

}
