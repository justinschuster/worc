/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	//"github.com/spf13/viper"
)

var (
	default_path = "/home/justin/Games/battlenet/drive_c/Program Files (x86)/World of Warcraft/"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "worc",
	Short: "CLI World of Warcraft addon manager",
	Long: `worc: a CLI World of Warcraft addon manager.
	Code found at: https://github.com/justinschuster/worc`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) { },
}

func CheckAddonPath() bool {
	info, err := os.Stat(default_path)
	if os.IsNotExist(err) {
		return false
	}
	fmt.Println("Game directory found at: ", default_path)
	return info.IsDir()
}

func CheckGameExists() bool {
	file_path := default_path + "_retail_/Wow.exe"
	info, err := os.Stat(file_path)
	if os.IsNotExist(err) {
		return false
	}
	fmt.Println("Retail version found at: ", file_path)
	return info.IsDir()
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {	
		CheckAddonPath()
		CheckGameExists()
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.worc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


