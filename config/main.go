// Find config dir: if env or flag override not set, use default
// Set default config dir in viper

package config

import (
	"strings"
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// This struct contains the config
/*
type Configuration struct {
	// Set to false until the initial load happens
	loaded bool

	// Configuration directory
	configDir string

	imap_host string
	imap_port string
	imap_user string
	imap_pass string
}
*/

// Loads all the configuration values into the struct from the various locations
func (c Configuration) LoadConfig() {

	// Load the base directory for the configuration
	// c.configDir = getConfigDir()

	// TODO: Add a method here to use a flag or env at runtime to specify a custom dir

	// Set the name of the configuration file
	var configName string = "config"
	var configType string = "yaml"

	// Set config type
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)


	// Paths to check for config files
	viper.AddConfigPath("/etc/mailman")
	viper.AddConfigPath("$HOME/.mailman")

	// TODO: remove this, perhaps
	// viper.AddConfigPath(".") // working directory

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {

		// Handle specific errors before falling back to general error

		// Check if error is "config not found"
		if strings.HasPrefix(err.Error(), "Config File \"" + configName + "\" Not Found") {
			panic(fmt.Errorf("Configuration file could not be found. Can not continue"))

			// TODO: add example0-config writing with viper.WriteConfigAs() here later

		// Throw generic error
		} else {
			panic(err)
		}

	}


	// c.loaded = true

}

// Config fetcher - gets the base configuration
func getConfigDir() string {

	flag.String("configdir", "~/.mailman/", "Override configuration directory")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)

	// TODO: Add logging here so that in verbose mode, it is
	// noted whether the default or override is used

	var cfgDir string = viper.GetString("configdir")

	return cfgDir

}


// good example of how to structure: https://scene-si.org/2017/04/20/managing-configuration-with-viper/
