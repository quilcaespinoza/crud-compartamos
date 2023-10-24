package config

type DatabaseConfig struct {
	Username     string
	Password     string
	Host         string
	Port         string
	DatabaseName string
}

func LoadDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Username:     "compartamos",
		Password:     "pwdcompartamos",
		Host:         "127.0.0.1",
		Port:         "3306",
		DatabaseName: "compartamos",
	}
}
