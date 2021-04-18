package middleware

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

var (
	cfg Config
)

// Config ...
type Config struct {
	Kafka struct {
		Topic         string `yaml:"kafka.topic"`
		EnableKafka   bool   `yaml:"kafka.enable"`
		BrokerAddress string `yaml:"kafka.brokerAddress"`
	}
	Keys struct {
		GoogleCivic string `yaml:"keys.google-civic"`
		Propublica  string `yaml:"keys.propublica"`
	}
}

func init() {
	readYAML(&cfg)
	readKeys(&cfg)
}

func readYAML(cfg *Config) {
	f, _ := os.Open("config/config.yml")

	defer f.Close()
	decoder := yaml.NewDecoder(f)
	decoder.Decode(cfg)
	fmt.Println(cfg.Kafka.EnableKafka)
}

func readEnv(cfg *Config) {
	var fileName string
	fileName = "env/" + os.Getenv("ENV") + ".env"
	fmt.Println(fileName)
	godotenv.Load(fileName)
	envconfig.Process("", cfg)
}

func readKeys(cfg *Config) {
	f, _ := os.Open("keys/keys.yml")
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	decoder.Decode(cfg)
	fmt.Println("Printing Keys")
	fmt.Println(cfg.Keys)
}
