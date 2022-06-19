package config

import (
	"github.com/OpenPeeDeeP/xdg"
	"os"
)

type AppConfig struct {
	Debug     bool `long:"debug" env:"DEBUG" default:"false"`
	ConfigDir string
}

func NewAppConfig(debugFlag bool) (*AppConfig, error) {
	configDir, err := loadOrCreateConfigDir()
	if err != nil {
		return nil, err
	}

	appConfig := &AppConfig{
		Debug:     debugFlag || os.Getenv("DEBUG") == "TRUE",
		ConfigDir: configDir,
	}

	return appConfig, nil
}

func loadOrCreateConfigDir() (string, error) {
	folder := configDir()

	err := os.MkdirAll(folder, 0o755)
	if err != nil {
		return "", err
	}

	return folder, nil
}

func configDirForVendor(vendor, project string) string {
	envConfigDir := os.Getenv("CONFIG_DIR")
	if envConfigDir != "" {
		return envConfigDir
	}

	configDirs := xdg.New(vendor, project)
	return configDirs.ConfigHome()
}

func configDir() string {
	return configDirForVendor("", "task-me")
}
