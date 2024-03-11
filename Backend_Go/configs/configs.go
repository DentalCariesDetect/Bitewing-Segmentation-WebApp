package configs

import (
	"os"
	"strconv"
)

type (
	Config struct {
		App App
		Db  Db
	}

	App struct {
		Port int
	}

	Db struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		SSLMode  string
		TimeZone string
	}

	Jwt struct {
		SecretKey string
	}
)

// The `init` function in Go initializes Viper to read configuration settings from a .env file in
// KEY=VALUE format and automatically load environment variables.
func init() {
	// viper.SetConfigFile(".env")
	// viper.SetConfigType("env") // If your .env file has KEY=VALUE format
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	log.Fatalf("Error reading .env file, %s", err)
	// }
	// viper.AutomaticEnv()
}

func GetConfig() Config {
	// Deploy
	echo_port, _ := strconv.Atoi(os.Getenv("ECHO_PORT"))
	postgres_port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	return Config{
		App: App{
			Port: echo_port,
		},
		Db: Db{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     postgres_port,
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			DBName:   os.Getenv("POSTGRES_DB"),
			SSLMode:  os.Getenv("POSTGRES_SSL_MODE"),
			TimeZone: os.Getenv("POSTGRES_TIMEZONE"),
		},
	}
	// return Config{
	// 	App: App{
	// 		Port: viper.GetInt("ECHO_PORT"),
	// 	},
	// 	Db: Db{
	// 		Host:     viper.GetString("POSTGRES_HOST"),
	// 		Port:     viper.GetInt("POSTGRES_PORT"),
	// 		User:     viper.GetString("POSTGRES_USER"),
	// 		Password: viper.GetString("POSTGRES_PASSWORD"),
	// 		DBName:   viper.GetString("POSTGRES_DB"),
	// 		SSLMode:  viper.GetString("POSTGRES_SSL_MODE"),
	// 		TimeZone: viper.GetString("POSTGRES_TIMEZONE"),
	// 	},
	// }
}

func GetJwtConfig() Jwt {
	// Deploy
	return Jwt{
		SecretKey: os.Getenv("JWT_SECRET"),
	}
	//	return Jwt{
	//		SecretKey: viper.GetString("JWT_SECRET"),
	//	}
}
