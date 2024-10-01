package addon

import {
	"fmt"
	"os"
	"path/filepath"
}

var (
	default_path = "/home/justin/Games/battlenet/drive_c/Program Files (x86)/World of Warcraft/"
)

func CheckAddonPath(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	fmt.Println("Game directory found at: ", path)
	return info.IsDir()
}

func CheckGameVersionRetail() bool {
	file_path := default_path + "_retail_/Wow.exe"
	info, err := os.Stat(file_path)
	if os.IsNotExist(err) {
		return false
	}
	fmt.Println("Retail version found at: ", file_path)
	return info.IsDir()
}

func CreateAddonPath() {
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

func LoadAddons() ([]string, error) {
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
