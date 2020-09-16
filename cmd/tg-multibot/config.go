package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/paper2code/golang-telegram-multibot/v2/pkg/context"
)

// var options *context.Options

// LoadConfig function loads configuration file and set options
func LoadConfig(configName string) (options *context.Options, err error) {
	log.Warnf("Load configuration file...")

	viper.SetConfigName(configName)
	viper.AddConfigPath(".")
	viper.AddConfigPath("../..")
	viper.AddConfigPath("../../shared/config")
	viper.AddConfigPath("/etc")
	viper.AddConfigPath("/usr/local/etc")

	viper.SetDefault("main.plugin_dir", "plugins")

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	options = &context.Options{
		AppName:         configName,
		APIKey:          viper.GetString("main.api_key"),
		DBDriver:        viper.GetString("db.driver"),
		DBUser:          viper.GetString("db.user"),
		DBHost:          viper.GetString("db.host"),
		DBPort:          viper.GetString("db.port"),
		DBName:          viper.GetString("db.name"),
		DBPassword:      viper.GetString("db.password"),
		DBSecure:        viper.GetBool("db.ssl"),
		LogLevel:        viper.GetString("log.level"),
		Debug:           viper.GetBool("main.debug"),
		PluginDir:       viper.GetString("main.plugin_dir"),
		PluginsSettings: make(map[string]map[string]interface{}),
	}
	return
}

func loadPLuginConfig(pluginName string) {
	log.Debugf("Loading plugin \"%s\" settings from configuration file...", pluginName)
	sub := viper.Sub(pluginName)
	if sub == nil {
		log.Debugf("No settings for plugin \"%s\"", pluginName)
		return
	}
	tempMap := make(map[string]interface{})
	for _, key := range sub.AllKeys() {
		tempMap[key] = sub.Get(key)
	}
	options.PluginsSettings[pluginName] = tempMap
	log.Debugf("Settings for plugin \"%s\" loaded successful", pluginName)
}

// write config
