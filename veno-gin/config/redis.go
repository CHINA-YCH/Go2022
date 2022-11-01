package config

/*
 * @Author: ych
 * @Description: ...
 * @File: redis
 * @Version: ...
 * @Date: 2022-11-01 16:24:17
 */

type Redis struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
