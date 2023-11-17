package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"urbathon-backend-2023/pkg/projectpath"
)

var config *viper.Viper

// Init Инициализация конфигурационного файла из файла
func Init(env string) {
	var err error
	config = viper.New()
	config.SetConfigType("json")
	config.SetConfigName(env)
	config.AddConfigPath("configs/")
	config.AddConfigPath("/configs/")
	config.AddConfigPath("/app/configs/")
	config.AddConfigPath(fmt.Sprintf("%s/configs/", projectpath.Root))
	err = config.ReadInConfig()
	if err != nil {
		log.Println("error on parsing configuration file", err)
	}

	initEnvVars(config)
	initFlags()
}

func initEnvVars(c *viper.Viper) {
	setConfigVarFromEnv(c, "database.address", "POSTGRES_HOST")
	setConfigVarFromEnv(c, "database.user", "POSTGRES_USER")
	setConfigVarFromEnv(c, "database.password", "POSTGRES_PASSWORD")
	setConfigVarFromEnv(c, "database.dbname", "POSTGRES_DB")
	setConfigVarFromEnv(c, "database.port", "POSTGRES_PORT")
}

func GetConfig() *viper.Viper {
	return config
}

// initFlags Чтение флагов, переданных при запуске исполняемого файла
func initFlags() {
	dbAddr := flag.String("dbAddr", "", "database.address")
	dbUser := flag.String("dbUser", "", "database.user")
	dbPass := flag.String("dbPass", "", "database.password")
	dbName := flag.String("dbName", "", "database.dbname")
	dbPort := flag.String("dbPort", "", "database.port")
	srvAddr := flag.String("srvAddr", "localhost", "server.address")
	srvPort := flag.String("srvPort", "", "server.port")

	flag.Parse()

	setConfigVarFromCLI(config, "database.address", *dbAddr)
	setConfigVarFromCLI(config, "database.user", *dbUser)
	setConfigVarFromCLI(config, "database.password", *dbPass)
	setConfigVarFromCLI(config, "database.dbname", *dbName)
	setConfigVarFromCLI(config, "database.port", *dbPort)
	setConfigVarFromCLI(config, "server.address", *srvAddr)
	setConfigVarFromCLI(config, "server.port", *srvPort)
}

func setConfigVarFromEnv(c *viper.Viper, key string, envKey string) {
	if os.Getenv(envKey) != "" {
		c.Set(key, os.Getenv(envKey))
	}
}

func setConfigVarFromCLI(c *viper.Viper, key string, cliKey string) {
	if cliKey != "" {
		c.Set(key, cliKey)
	}
}
