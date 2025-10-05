package goindexer

import (
	"fmt"
	_ "net/http/pprof"

	"github.com/mbndr/figlet4go"
	"github.com/spf13/cobra"
)

var (
	configPath string
	ascii      = figlet4go.NewAsciiRender()
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
	// Adding the colors to RenderOptions
	options := figlet4go.NewRenderOptions()
	options.FontColor = []figlet4go.Color{
		// Colors can be given by default ansi color codes...
		figlet4go.ColorMagenta,
	}
	options.FontName = "larry3d"

	renderStr, _ := ascii.RenderOpts("GoIndexer", options)

	fmt.Print(renderStr)
}
