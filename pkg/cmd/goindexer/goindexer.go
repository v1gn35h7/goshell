package goindexer

import (
	"fmt"
	_ "net/http/pprof"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/v1gn35h7/goshell/pkg/logging"
	"github.com/v1gn35h7/goshell/pkg/service"
)

var (
	configPath string
)

func NewCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "goIndexer service",
		Short: "goIndexer service",
		Long:  "goIndexer service started",
		Run: func(cmd *cobra.Command, args []string) {

			//Init
			printLogo()

			//Bootstrap server
			bootStrapServer()
		},
	}

	// Bind cli flags
	rootCmd.PersistentFlags().StringVar(&configPath, "configPath", "", "config file path")

	return rootCmd
}

func printLogo() {
	fmt.Println("#########################################################")
	fmt.Println("                                                         ")
	fmt.Println("                                                         ")
	fmt.Println("        \"\"\"\"\"\"\"                   \"\"\"\"\"\"\"\"\"\"\"\"                ")
	fmt.Println("       \"\"       \"\"                             \"\"    							")
	fmt.Println("      \"\"                                         \"\"    ")
	fmt.Println("      \"\"                                         \"\"     ")
	fmt.Println("      \"\"                \"\"\"\"\"\"             \"\"     ")
	fmt.Println("      \"\"                \"        \"             \"\"     ")
	fmt.Println("      \"\"         \"\"   \"        \"             \"\"     ")
	fmt.Println("       \"\"        \"\"   \"        \"             \"\"     ")
	fmt.Println("         \"\"      \"\"   \"        \"             \"\"     ")
	fmt.Println("           \"\"\"\"\"\"   \"\"\"\"\"\"    \"\"\"\"\"\"\" \"\"\"\"    ")
	fmt.Println("                                                         ")
	fmt.Println("                                                         ")
	fmt.Println("#########################################################")
}

func bootStrapServer() {
	//Logger
	logger := logging.Logger()

	viper.SetConfigName("app")      // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(configPath) // path to look for the config file in
	viper.AddConfigPath(".")        // optionally look for config in the working directory
	err := viper.ReadInConfig()     // Find and read the config file

	if err != nil {
		// Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	fmt.Println(viper.AllSettings())

	service.IndexerService(logger)
}
