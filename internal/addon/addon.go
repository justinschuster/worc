package addon

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var (
	defaultPath = "/home/justin/Games/battlenet/drive_c/Program Files (x86)/World of Warcraft/"
)

type Addon struct {
	AddonPath string
	Name      string
	Version   string
}

func CheckAddonPath(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	fmt.Println("Game directory found at:", path)
	return info.IsDir()
}

func CheckGameVersionRetail() bool {
	filePath := filepath.Join(defaultPath, "_retail_", "Wow.exe")
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	fmt.Println("Retail version found at:", filePath)
	return true
}

func CreateAddonPath() error {
	interfacePath := filepath.Join(defaultPath, "_retail_", "Interface")
	addonPath := filepath.Join(interfacePath, "Addons")
	
	if err := os.MkdirAll(addonPath, 0750); err != nil {
		return fmt.Errorf("failed to create addon path: %w", err)
	}
	return nil
}

func LoadAddons() ([]Addon, error) {
	var addons []Addon
	addonPath := filepath.Join(defaultPath, "_retail_", "Interface", "Addons")
	entries, err := os.ReadDir(addonPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read addon directory: %w", err)
	}
	
	for _, entry := range entries {
		if entry.IsDir() {
			fullPath := filepath.Join(addonPath, entry.Name())
			versionNumber := FindVersion(fullPath)
			addonName := FindName(fullPath)
			addon := Addon{
				AddonPath: fullPath,
				Version:   versionNumber,
				Name:      addonName,
			}
			addons = append(addons, addon)
		}
	}
	return addons, nil
}

func FindVersion(path string) string {
	return findInFile(path, `^## Version: (.+)`)
}

func FindName(path string) string {
	return findInFile(path, `^## Title: (.+)`)
}

func findInFile(path, pattern string) string {
	file, err := os.Open(filepath.Join(path, filepath.Base(path)+".toc"))
	if err != nil {
		fmt.Println("Error opening file:", err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	regex := regexp.MustCompile(pattern)
	
	for scanner.Scan() {
		matches := regex.FindStringSubmatch(scanner.Text())
		if len(matches) > 1 {
			return matches[1]
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
	return ""
}
