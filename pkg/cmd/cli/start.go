package cli

import (
	"github.com/goShell/pkg/logging"
	"github.com/goShell/pkg/service"
	"github.com/spf13/cobra"
)

func NewStartCommand() *cobra.Command {
	var startCmd = &cobra.Command{
		Use:   "start",
		Short: "goshellctl start",
		Long:  "Starts Goshell cli process",
		Run: func(cmd *cobra.Command, args []string) {
			// Set-up logger
			logger := logging.Logger()
			logger.Info("Logger initated...")

			goshellCtlSrvc := service.NewService(logger)
			started, err := goshellCtlSrvc.StartService()

			if err != nil {
				logger.Error(err, "Starting goshellstl service failed")
			}

			if started {
				logger.Info("Started service...")
			}
		},
	}
	return startCmd
}
