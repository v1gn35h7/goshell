package goshellctl

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/v1gn35h7/goshell/internal/datastore/cassdb"
	"github.com/v1gn35h7/goshell/pkg/cmd/cli"
	"github.com/v1gn35h7/goshell/pkg/constants"
)

var (
	configPath string
)

func NewCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "goshellctl",
		Short: "goshellctl service",
		Long:  "Goshell service starts goshell command line services",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			printLogo()

			// Init read config
			fmt.Println("Config path set to: ", configPath)
			readConfig(configPath)

			// Init database
			cassdb.SetUpSession()
		},
	}

	// Bind cli flags
	rootCmd.PersistentFlags().StringVar(&configPath, "configPath", "", "config file path")

	// Add sub commands
	rootCmd.AddCommand(cli.NewVersionCommand())
	rootCmd.AddCommand(cli.NewStartCommand())
	rootCmd.AddCommand(cli.NewStopCommand())
	rootCmd.AddCommand(cli.NewDebugCommand())

	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))

	return rootCmd
}

func printLogo() {
	fmt.Println("#########################################################")
	fmt.Println("                                                         ")
	fmt.Println("                                                         ")
	fmt.Println("        \"\"\"\"\"\"\"                                   ")
	fmt.Println("       \"\"       \"\"                                  ")
	fmt.Println("      \"\"                                              ")
	fmt.Println("      \"\"                                               ")
	fmt.Println("      \"\"                \"\"\"\"\"\"                   ")
	fmt.Println("      \"\"                \"        \"                   ")
	fmt.Println("      \"\"         \"\"   \"        \"                   ")
	fmt.Println("       \"\"        \"\"   \"        \"                   ")
	fmt.Println("         \"\"      \"\"   \"        \"                   ")
	fmt.Println("           \"\"\"\"\"\"   \"\"\"\"\"\"                  ")
	fmt.Println("                                                         ")
	fmt.Println("                                                         ")
	fmt.Println("#########################################################")
}

func readConfig(configPath string) {
	// Read config
	//logger.Info("Reading config from file", "confi_path", configPath)
	viper.SetConfigName(constants.ConfigName) // name of config file (without extension)
	viper.SetConfigType(constants.ConfigType) // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(configPath)           // path to look for the config file in
	viper.AddConfigPath(".")                  // optionally look for config in the working directory
	err := viper.ReadInConfig()               // Find and read the config file

	if err != nil {
		// Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
