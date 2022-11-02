package config

import (
	"github.com/spf13/viper"
)

type ConfigFile struct {
	ConfigFileName string // "config"
	ConfigFileType string // "yaml"
	ConfigFilePath string // "./conf"
}

// Find and read the config file
func readConfigFile(configFile *ConfigFile) error {

	//set config file to read
	viper.SetConfigName(configFile.ConfigFileName)
	viper.SetConfigType(configFile.ConfigFileType)
	viper.AddConfigPath(configFile.ConfigFilePath)

	err := viper.ReadInConfig()
	return err
}

func (configFile *ConfigFile) GetNodes() map[string]string {

	nodes := make(map[string]string)

	err := readConfigFile(configFile)
	if err != nil { // Handle errors reading the config file
		return nodes
	}

	nodes = viper.GetStringMapString("nodes")

	return nodes
}

func (configFile *ConfigFile) GetVaultIndex() map[string]string {

	vaultIndex := make(map[string]string)

	err := readConfigFile(configFile)
	if err != nil { // Handle errors reading the config file
		return vaultIndex
	}

	vaultIndex = viper.GetStringMapString("vaultIndex")

	return vaultIndex
}

func (configFile *ConfigFile) GetLogfile() map[string]string {

	logFile := make(map[string]string)

	err := readConfigFile(configFile)
	if err != nil { // Handle errors reading the config file
		return logFile
	}

	logFile = viper.GetStringMapString("logfile")

	return logFile
}

func (configFile *ConfigFile) GetPlugin() map[string]string {

	plugin := make(map[string]string)

	err := readConfigFile(configFile)
	if err != nil { // Handle errors reading the config file
		return plugin
	}

	plugin = viper.GetStringMapString("plugin")

	return plugin
}
