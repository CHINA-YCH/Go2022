package config

/*
编写配置结构体

在项目根目录下新建文件夹 config，用于存放所有配置对应的结构体
定义 Configuration 结构体，其 App 属性对应 config2.yaml 中的 app
*/

type Configuration struct {
	App      App      `mapstructure:"app" json:"app" yaml:"app"`
	Log      Log      `mapstructure:"log" json:"log" yaml:"log"`
	Database Database `mapstructure:"database" json:"database" yaml:"database"`
	Jwt      Jwt      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis    Redis    `mapstructure:"redis" json:"redis" yaml:"redis"`
}
