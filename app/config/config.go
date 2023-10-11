package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type AppConfig struct {
	SERVERPORT int
	DBPORT     int
	DBHOST     string
	DBUSER     string
	DBPASS     string
	DBNAME     string
}

func InitConfig() *AppConfig {
	var res = new(AppConfig)
	res = loadConfig()

	if res == nil {
		logrus.Fatal("Config : Cannot start program, failed to load configuration")
		return nil
	}

	return res
}

func loadConfig() *AppConfig {
	var res = new(AppConfig)

	err := godotenv.Load(".env")

	if err != nil {
		logrus.Error("Config : Cannot load config file,", err.Error())
	}

	// var isRead = false
	if val, found := os.LookupEnv("SERVERPORT"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config : invalid port value,", err.Error())
			return nil
		}
		// isRead = true
		res.SERVERPORT = port
	}

	if val, found := os.LookupEnv("DBPORT"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config : invalid db port value,", err.Error())
			return nil
		}
		// isRead = true
		res.DBPORT = port
	}

	if val, found := os.LookupEnv("DBHOST"); found {
		// isRead = true
		res.DBHOST = val
	}

	if val, found := os.LookupEnv("DBUSER"); found {
		// isRead = true
		res.DBUSER = val
	}

	if val, found := os.LookupEnv("DBPASS"); found {
		// isRead = true
		res.DBPASS = val
	}

	if val, found := os.LookupEnv("DBNAME"); found {
		// isRead = true
		res.DBNAME = val
	}

	return res

}