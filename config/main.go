// Find config dir: if env or flag override not set, use default
// Set default config dir in viper

package config

import (
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"strings"
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

	// Set the name of the configuration file
	var configName string = "config"
	var configType string = "yaml"

	// Set config type
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	// Setup the struct
	var configuration Configuration

	// Set the default values to later be overridden
	setConfigDefaults()

	// Paths to check for config files
	viper.AddConfigPath("/etc/mailman")
	viper.AddConfigPath("$HOME/.mailman")
	// Implement XDG config path here (XDG_CONFIG_HOME, if not set, ~/.config/appname)
	// TODO: Add a method here to use a flag or env at runtime to specify a custom config dir or config file

	// Find and read the config file
	errRead := viper.ReadInConfig()

	if errRead != nil {

		// Handle specific errors before falling back to general error

		// Check if error is "config not found"
		if strings.HasPrefix(errRead.Error(), "Config File \""+configName+"\" Not Found") {
			panic(fmt.Errorf("Configuration file (~/.mailman/" + configName + ".yml) could not be found - can not continue.\n\nCreate an empty file to continue with development.\n\nTODO: Add piece which could create config with --firstrun later"))

			// TODO: add example0-config writing with viper.WriteConfigAs() here later

		} else {
			// Throw generic error if a more specific one not found
			panic(errRead)
		}

	}

	// Marshall the config into the structs
	errMarsh := viper.Unmarshal(&configuration)
	if errMarsh != nil {
//		log.Fatalf("unable to decode into struct, %v", errMarsh)
	}



//        log.Printf("port is %d", configuration.ImapServer.Port)

	// TODO: remove later
	viper.Debug()

}

// Handles all the default configuration, allowing them to be easily
// set in one place
func setConfigDefaults() {
	// TODO: improve seperation of imap module into it's own seperate backend piece
	// Default imap timeout
	viper.SetDefault("ImapServer.Port", "993")

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
