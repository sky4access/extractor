package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/sky4access/extractor/internal/server"
	"github.com/spf13/cobra"
)

const ()

func NewCommand(logger *logrus.Entry) *cobra.Command {
	cmds := &cobra.Command{
		Use:   "bext",
		Short: "Bible extractor",
		Long:  "Extract Bible verses",
		Run:   rootCmd,
	}

	service := server.NewService(logger)
	cmds.AddCommand(NewCmdService(service))
	return cmds
}

func rootCmd(cmd *cobra.Command, args []string) {
	cmd.Help()
}
