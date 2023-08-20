package cli

import (
	"github.com/spf13/cobra"
	"github.com/v1gn35h7/goshell/internal/datastore/cassdb"
	"github.com/v1gn35h7/goshell/pkg/logging"
	"github.com/v1gn35h7/goshell/pkg/service"
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

			// Init database
			cassdb.SetUpSession()

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
