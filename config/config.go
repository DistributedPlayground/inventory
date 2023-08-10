package config

type vars struct {
	ENV            string `required:"true"`
	PORT           string `required:"true"`
	REDIS_HOST     string `required:"true"`
	REDIS_PORT     string `required:"true"`
	REDIS_PASSWORD string `required:"true"`
}

var Var vars
