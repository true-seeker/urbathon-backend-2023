package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"log"
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
	initFlags()
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

	if *dbAddr != "" {
		config.Set("database.address", *dbAddr)
	}
	if *dbUser != "" {
		config.Set("database.user", *dbUser)
	}
	if *dbPass != "" {
		config.Set("database.password", *dbPass)
	}
	if *dbName != "" {
		config.Set("database.dbname", *dbName)
	}
	if *dbPort != "" {
		config.Set("database.port", *dbPort)
	}
	if *srvAddr != "" {
		config.Set("server.address", *srvAddr)
	}
	if *srvPort != "" {
		config.Set("server.port", *srvPort)
	}
}
