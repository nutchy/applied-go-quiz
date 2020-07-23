package app

import "github.com/spf13/viper"

type Configs struct {
	v *viper.Viper

	// Config from flag
	State string
	Port  string

	// Config from yml file
	// Tag `mapstructure` used for map config file to string (like json tag)
	EnablePing bool `mapstructure: "enable_ping`
	// Mongo      Mongo `mapstructure: "mongo"`
}
