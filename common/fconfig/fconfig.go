package fconfig

import (
	"github.com/spf13/viper"
)

func GetNodes() map[string]string {

	nodes := make(map[string]string)

	//set config file to read
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return nodes
	}

	nodes = viper.GetStringMapString("nodes")

	return nodes
}

func GetVaultIndex() map[string]string {

	vaultIndex := make(map[string]string)

	//set config file to read
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return vaultIndex
	}

	vaultIndex = viper.GetStringMapString("vaultIndex")

	return vaultIndex
}

func GetLogfile() map[string]string {

	logFile := make(map[string]string)

	//set config file to read
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return logFile
	}

	logFile = viper.GetStringMapString("logfile")

	return logFile
}

func GetPlugin() map[string]string {

	plugin := make(map[string]string)

	//set config file to read
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return plugin
	}

	plugin = viper.GetStringMapString("plugin")

	return plugin
}
