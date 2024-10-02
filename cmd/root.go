/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/justinschuster/worc/internal/addon"
	"github.com/spf13/cobra"
)

var (
	defaultPath = "/home/justin/Games/battlenet/drive_c/Program Files (x86)/World of Warcraft/"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "worc",
	Short: "CLI World of Warcraft addon manager",
	Long: `worc: a CLI World of Warcraft addon manager.
	Code found at: https://github.com/justinschuster/worc`,
	Run: func(cmd *cobra.Command, args []string) {
		addonPath := filepath.Join(defaultPath, "Interface", "Addons")
		if !addon.CheckAddonPath(addonPath) {
			if err := addon.CreateAddonPath(); err != nil {
				fmt.Printf("Failed to create addon path: %v\n", err)
				os.Exit(1)
			}
		}

		if !addon.CheckGameVersionRetail() {
			fmt.Println("Retail version of World of Warcraft not found.")
			os.Exit(1)
		}

		addons, err := addon.LoadAddons()
		if err != nil {
			fmt.Printf("Failed to load addons: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("\nListing Addons:")
		fmt.Println()
		for _, addonValue := range addons {
			fmt.Println("Name: ", addonValue.Name)
			fmt.Println("Version: ", addonValue.Version)
			fmt.Println("Addon Path: ", addonValue.AddonPath)
			fmt.Println()
		}
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.worc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
