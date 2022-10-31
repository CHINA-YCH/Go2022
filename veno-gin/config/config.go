package config

/*
 * @Author: ych
 * @Description: ...
 * @File: config
 * @Version: ...
 * @Date: 2022-10-31 15:20:50
 */

type Configuration struct {
	App App `mapstructure:"app" json:"app" yaml:"app"`
	Log Log `mapstructure:"log" json:"log" yaml:"log"`
}
