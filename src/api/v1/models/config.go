package models

type ServerConfig struct {
	Env           string `env:"ENV"`
	Name          string `env:"NAME_SERVER"`
	Port          string `env:"PORT_SERVER,required"`
	Host          string `env:"HOST_SERVER,required"`
	DBConfig      DBConfig
	ElasticConfig ElasticConfig
	Redis         Redis
}

type DBConfig struct {
	Name     string `env:"NAME_POSTGRES,required"`
	Host     string `env:"HOST_POSTGRES,required"`
	Port     string `env:"PORT_POSTGRES,required"`
	User     string `env:"USER_POSTGRES"`
	Password string `env:"PASS_POSTGRES"`
}

type ElasticConfig struct {
	Host     string `env:"HOST_ELASTICSEARCH,required"`
	Port     string `env:"PORT_ELASTICSEARCH,required"`
	User     string `env:"USER_ELASTICSEARCH"`
	Password string `env:"PASS_ELASTICSEARCH"`
	Index    string `env:"INDEX_ELASTICSEARCH,required"`
}

type Redis struct {
	Host     string `env:"REDIS_HOST"`
	Port     string `env:"REDIS_PORT"`
	Password string `env:"REDIS_PASSWORD"`
}
