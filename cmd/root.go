/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

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

func checkAddonPath(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	fmt.Println("Game directory found at: ", path)
	return info.IsDir()
}

func checkGameVersionRetail() bool {
	file_path := default_path + "_retail_/Wow.exe"
	info, err := os.Stat(file_path)
	if os.IsNotExist(err) {
		return false
	}
	fmt.Println("Retail version found at: ", file_path)
	return info.IsDir()
}

func createAddonPath() {
	interface_path := default_path + "_retail_/Interface"
	err := os.Mkdir(interface_path, 0750)
	if err != nil && !os.IsExist(err) {
		fmt.Println(err)
		os.Exit(1)
	}

	addon_path := interface_path + "/Addons/"
	err = os.Mkdir(addon_path, 0750)
	if err != nil && !os.IsExist(err) {
		fmt.Println(err)
		os.Exit(1)
	}
}

func loadAddons() ([]string, error) {
	var dirs []string
	addon_path := default_path + "_retail_/Interface/Addons/"
	entries, err := os.ReadDir(addon_path)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, filepath.Join(addon_path, entry.Name()))
		}
	}
	return dirs, nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	addon_path := default_path + "Interface/Addons/"
	if err != nil {	
		if !checkAddonPath(addon_path) {
			createAddonPath()
		}
		checkGameVersionRetail()
		addons, err := loadAddons()
		if err != nil {
			fmt.Println(err)	
		} else {
			fmt.Println("\nListing Addon paths:")
			for _, value := range addons {
				fmt.Println(value)
			}
		}
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


