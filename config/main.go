// Find config dir: if env or flag override not set, use default
// Set default config dir in viper

package config

import (
	"flag"
	"github.com/spf13/pflag"
	"fmt"
	"github.com/spf13/viper"
)

func main() {
       fmt.Println("Config directory is " + GetConfigDir())
}


// Returns the configuration directory after checking
// the flags to see if overridden
// TODO: make this idempotent, or rework how config is handled
func GetConfigDir() string {

	flag.String("configdir", "~/.mailman/", "Override configuration directory")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)


	// TODO: Add logging here so that in verbose mode, it is
        // noted whether the default or override is used

	var cfgDir string = viper.GetString("configdir")

        return cfgDir

}
