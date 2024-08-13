/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
  "fmt"
	"os"
  "time"
  "log/slog"
	"github.com/spf13/cobra"
  "github.com/mitchellh/colorstring"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vault",
	Short: "Runs the Vault daemon, needed for the backend to function.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: rootCmdRun,
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.daemon.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func rootCmdRun(cmd *cobra.Command, _ []string) {
  printLogo()
  slog.Info("Hi")  
}

func printLogo() {
  fmt.Printf(colorstring.Color(`[bold][yellow]Version 0.0.1 by[reset] [blue][bold]RyzechDev[reset][blue][bold]
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
in all copies or substantial portions of the Software.%s`), time.Now().Year(), "\n\n")
}


