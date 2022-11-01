package config

/*
 * @Author: ych
 * @Description: ...
 * @File: config
 * @Version: ...
 * @Date: 2022-10-31 15:20:50
 */

type Configuration struct {
	App      App      `mapstructure:"app" json:"app" yaml:"app"`
	Log      Log      `mapstructure:"log" json:"log" yaml:"log"`
	Database Database `mapstructure:"database" json:"database" yaml:"database"`
	Redis    Redis    `mapstructure:"redis" json:"redis" yaml:"redis"`
	Storage  Storage  `mapstructure:"storage" json:"storage" yaml:"storage"`
}
