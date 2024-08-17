/*
Copyright © 2024 RyzechDev <ryzechdev@ryzech.net>
*/
package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/VaultedUI/daemon/client"
	"github.com/VaultedUI/daemon/config"
	"github.com/VaultedUI/daemon/system"
	"github.com/mitchellh/colorstring"
	"github.com/spf13/cobra"
)

var (
	configPath  = config.ConfigPath
	debug       = false
	ignoreDebug = false
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vault",
	Short: "Runs the Vault daemon, needed for the backend to function.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: rootCmdRun,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Gives the current version of the application.",
	Run: func(cmd *cobra.Command, _ []string) {
		fmt.Printf("Vault %s\nCopyright © 2024 - %d RyzechDev\n", system.Version, time.Now().Year())
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&configPath, "config", config.ConfigPath, "config file path")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "sets the debug mode for vault, used for diagnostics")

	rootCmd.Flags().BoolVar(&ignoreDebug, "ignore-debug", false, "ignore the warning when debug is turned on")

	rootCmd.AddCommand(versionCmd)
}

func rootCmdRun(cmd *cobra.Command, _ []string) {
	if debug == true {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	printLogo()
	slog.Info("Lauching daemon...")
	if debug == true && !ignoreDebug == true {
		slog.Warn("Debug mode is enabled! If this is intended, you may ignore this message.")
	}
	slog.Debug("Creating docker client.")
	_, err := client.Docker()
	if err != nil {
		slog.Error(err.Error())
		slog.Debug("Failed to create client, exiting now.")
		os.Exit(1)
	}
	slog.Debug("Created docker client successfully.")
}

func printLogo() {
	fmt.Printf(colorstring.Color(`[bold][yellow]Version %s by[reset] [blue][bold]RyzechDev[reset][blue][bold]
 ___      ___ ________  ___  ___  ___   _________
|\  \    /  /|\   __  \|\  \|\  \|\  \ |\___   ___\
\ \  \  /  / | \  \|\  \ \  \\\  \ \  \\|___ \  \_|
 \ \  \/  / / \ \   __  \ \  \\\  \ \  \    \ \  \
  \ \    / /   \ \  \ \  \ \  \\\  \ \  \____\ \  \
   \ \__/ /     \ \__\ \__\ \_______\ \_______\ \__\
    \|__|/       \|__|\|__|\|_______|\|_______|\|__|[reset]

Copyright © 2024 - %d RyzechDev

This software is made available under the terms of the MIT license.
The above copyright notice and this permission notice shall be included
in all copies or substantial portions of the Software.%s`), system.Version, time.Now().Year(), "\n\n")
}
