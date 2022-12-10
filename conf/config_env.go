package conf

import "github.com/caarlos0/env/v6"

// AppConfig presents app conf
type AppConfig struct {
	Port string `env:"PORT" envDefault:"8001"`
	//DB CONFIG
	LogFormat string `env:"LOG_FORMAT" envDefault:"127.0.0.1"`
	DBHost    string `env:"DB_HOST" envDefault:"localhost"`
	DBPort    string `env:"DB_PORT" envDefault:"5432"`
	DBUser    string `env:"DB_USER" envDefault:"postgres"`
	DBPass    string `env:"DB_PASS" envDefault:"postgres"`
	DBName    string `env:"DB_NAME" envDefault:"postgres"`
	DBSchema  string `env:"DB_SCHEMA" envDefault:"public"`
	EnableDB  string `env:"ENABLE_DB" envDefault:"true"`
	// ENV
	EnvName        string `env:"ENV_NAME" envDefault:"dev"`
	MqttBrokerHost string `env:"MQTT_BROKER" envDefault:"broker.emqx.io"`
	MqttBrokerPort string `env:"MQTT_BROKER" envDefault:"1883"`
	MqttClientID   string `env:"MQTT_CLIENT_ID" envDefault:"go_mqtt_client"`
	MqttUsername   string `env:"MQTT_USERNAME" envDefault:"emqx"`
	MqttPassword   string `env:"MQTT_PASSWORD" envDefault:"public"`
	MqttTopic      string `env:"MQTT_TOPIC" envDefault:"topic/test"`
}

var config AppConfig

func SetEnv() {
	_ = env.Parse(&config)
}

func GetEnv() AppConfig {
	return config
}
