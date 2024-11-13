package utils

import (
	"os"
	"path"
)

func GetDbPath() string {
	home, _ := os.UserHomeDir()
	return path.Join(home, "."+APP_NAME, APP_NAME+".db")
}

func GetConfigPaths() ([]string, error) {
	configDir, _ := os.UserConfigDir()
	home, _ := os.UserHomeDir()

	possible := make([]string, 0)

	// get current directory
	possible = append(possible, path.Join(".", CONFIG_FILE))
	// in home directory
	possible = append(possible, path.Join(home, "."+CONFIG_FILE))
	// in config directory
	possible = append(possible, path.Join(configDir, APP_NAME, CONFIG_FILE))

	// filter out the files that do not exist
	exists := make([]string, 0)
	for _, p := range possible {
		if _, err := os.Stat(p); err == nil {
			exists = append(exists, p)
		}
	}

	return exists, nil
}

func GetConfigPath() (string, error) {
	paths, err := GetConfigPaths()
	if err != nil {
		return "", err
	}
	return paths[0], nil
}

func InitPaths() {
	// TODO: Handle the error
	configDir, _ := os.UserConfigDir()
	home, _ := os.UserHomeDir()

	os.MkdirAll(path.Join(configDir, APP_NAME), 0755)
	os.MkdirAll(path.Join(home, "."+APP_NAME), 0755)
}
