package config

type Config struct {
	ENV                     string
	MONGO_CONNECTION_STRING string
	REDIS_HOST              string
	REDIS_PASSWORD          string
}

func LoadConfig() Config {

}
