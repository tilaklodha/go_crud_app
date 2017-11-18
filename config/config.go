package config

type Config struct {
	dbConfig *dbConfig
}

var appConfig *Config

func Load() {
	appConfig = &Config{
		dbConfig: newDBConfig(),
	}
}

func DBConfig() *dbConfig {
	return appConfig.dbConfig
}
