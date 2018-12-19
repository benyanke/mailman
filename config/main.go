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

// Loads all the configuration values into the struct from the various locations
func (c Configuration) LoadConfig() {

	// Load the base directory for the configuration
	// c.configDir = getConfigDir()

	// TODO: Add a method here to use a flag or env at runtime to specify a custom dir

	// Set the name of the configuration file
	var configFile string = "config.yaml"

	// Set config type
	// viper.SetConfigType(configType)
	// viper.SetConfigName(configName) // name of config file (without extension)

	// This replaces the two commands above
	viper.SetConfigFile(configFile)

	// System-wide config
	viper.AddConfigPath("/etc/mailman")
	// Look for config in the working directory
	viper.AddConfigPath(".")
	// Look in dotfile directory
	viper.AddConfigPath("$HOME/.mailman")

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {

		// Handle specific errors before falling back to general error

		// Check if error is "config not found"
		if strings.HasPrefix(err.Error(), "Config File \"" + configFile + "\" Not Found") {
			panic(fmt.Errorf("Configuration file could not be found. Can not continue"))

			// TODO: add example0-config writing with viper.WriteConfigAs() here later

		// Throw generic error
		} else {
			panic(err)
		}

	}


	c.loaded = true

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
