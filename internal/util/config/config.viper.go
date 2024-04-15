package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type config struct {
	viper.Viper
}

func NewConfig() *config {
	// Initialize a new viper instance
	v := viper.New()
	// Load global envs
	v.AutomaticEnv()
	return &config{Viper: *v}
}

// Load config file
func (r *config) LoadConfigFile(path, configType, configFile string) error {
	if _, err := os.Stat(fmt.Sprintf("%s/%s", path, configFile)); err == nil {
		r.AddConfigPath(path)
		r.SetConfigType(configType)
		r.SetConfigName(configFile)
		err := r.ReadInConfig()
		if err != nil {
			return err
		}
	}
	return nil
}
