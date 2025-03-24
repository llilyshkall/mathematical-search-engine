package config

type DatabaseConfig struct {
	Host           string `yaml:"host"`
	Port           int    `yaml:"port"`
	UserName       string `yaml:"userName"`
	Password       string `yaml:"password"`
	DbName         string `yaml:"dbName"`
	MigrationsPath string `yaml:"migrationsPath" default:"migrations/"`
}

type Configuration struct {
	Database *DatabaseConfig `yaml:"database"`
}
