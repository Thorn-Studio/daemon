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
	"github.com/VaultedUI/daemon/system"
	"github.com/mitchellh/colorstring"
	"github.com/spf13/cobra"
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
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.daemon.yaml)")
	rootCmd.AddCommand(versionCmd)
}

func rootCmdRun(cmd *cobra.Command, _ []string) {
	printLogo()
	slog.Info("Lauching daemon...")
	_, err := client.Docker()
	if err != nil {
		slog.Error(err.Error())
	}
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
