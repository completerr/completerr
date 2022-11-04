package config

import (
	"completerr/tasks"
	"completerr/utils"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"strings"
)

var logger = utils.GetLogger()

func initConfig() {
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigName("config")            // name of config file (without extension)
	viper.SetConfigType("yaml")              // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/config/")          // path to look for the config file in
	viper.AddConfigPath("/etc/completerr/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.completerr") // call multiple times to add many search paths
	viper.AddConfigPath(".")                 // optionally look for config in the working directory

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		logger.Info(fmt.Sprintf("Config file changed: %s", e.Name))
		tasks.RestartScheduler()
	})
	viper.WatchConfig()
}
func init() {
	initConfig()
}
