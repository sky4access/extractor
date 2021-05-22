package cmd

import (
	"github.com/sky4access/extractor/internal/server"
	"github.com/spf13/cobra"
)



func NewCmdService(service *server.Service) *cobra.Command {
	command := &cobra.Command{
		Use: "service",
		Short: "run extractor service",
		Long: "run  extractor service",
		RunE: func(cmd *cobra.Command, args[]string) error {
			return service.Run()
		},
	 }

	return command
}
