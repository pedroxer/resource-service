package config

type Config struct{
	Postgres Postgres `json:"postgres"`
	Port int `json:"port"`
}

type Postgres struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Db       string `json:"db"`
	User     string `env:"PG_USER,notEmpty"`
	SSLMode  string `json:"sslmode"`
	Password string `env:"PG_PASSWORD,notEmpty"`

}