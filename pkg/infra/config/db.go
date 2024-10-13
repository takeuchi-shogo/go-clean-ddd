package config

type DBConfig struct {
	Database struct {
		Host     string `mapstructure:"host"`
		UserName string `mapstructure:"user_name"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"database_name"`
	}
}
