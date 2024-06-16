package config

type EnvConfig struct {
	Application ApplicationConfig `mapstructure:"app"`
	Database    DatabaseConfig    `mapstructure:"database"`
}
type (
	ApplicationConfig struct {
		Name string `mapstructure:"name"`
		Env  string `mapstructure:"env"`
		Port int    `mapstructure:"port"`
	}
	DatabaseConfig struct {
		Username     string `mapstrucutre:"username"`
		Password     string `mapstructure:"password"`
		Hostname     string `mapstrucutre:"hostname"`
		Port         int    `mapstructure:"port"`
		Name         string `mapstructure:"name"`
		Driver       string `mapstructure:"driver"`
		Charset      string `mapstructure:"charset"`
		Timezone     string `mapstructure:"timezone"`
		ReadTimeout  int    `mapstructure:"read_timeout"`
		WriteTimeout int    `mapstructure:"write_timeout"`
		MaxIdleConn  int    `mapstructure:"max_idle_conn"`
		MaxOpenConn  int    `mapstructure:"max_open_conn"`
	}
)
