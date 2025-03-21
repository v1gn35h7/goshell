package goindexer

import (
	"fmt"
	_ "net/http/pprof"

	"github.com/spf13/cobra"
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
		},
	}

	// Bind cli flags
	rootCmd.PersistentFlags().StringVar(&configPath, "configPath", "", "config file path")

	// Add cli
	rootCmd.AddCommand(NewStartCommand())
	rootCmd.AddCommand(NewSeedCommand())

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
