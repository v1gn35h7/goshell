package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewVersionCommand() *cobra.Command {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "goshellctl version",
		Long:  "Prints Goshell version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("GshellCtl ", viper.GetString("goshellctl.version"))
		},
	}
	return versionCmd
}
