package main

import (
	"fmt"
	"github.com/dikopylov/highload-architect/internal/handlers"
	"github.com/dikopylov/highload-architect/internal/infrastructure/config"
	"github.com/dikopylov/highload-architect/internal/model/auth"
	"github.com/dikopylov/highload-architect/internal/model/users"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func main() {
	err := loadConfig()
	if err != nil {
		log.Fatalln(fmt.Errorf("fatal error config file: %w", err))
	}

	dsn := fmt.Sprintf(
		"%s://%s:%s@%s:%d/%s?sslmode=disable",
		viper.GetString(config.DatabaseDriver),
		viper.GetString(config.DatabaseUser),
		viper.GetString(config.DatabasePassword),
		viper.GetString(config.DatabaseHost),
		viper.GetInt(config.DatabasePortInContainer),
		viper.GetString(config.DatabaseName),
	)

	db, err := sqlx.Connect(viper.GetString(config.DatabaseDriver), dsn)
	if err != nil {
		log.Fatalln(err)
	}

	userRepository := users.NewPgsqlRepository(*db)
	userService := users.NewService(userRepository, auth.NewInMemoryStorage())
	server := handlers.NewHTTPServer(userService)
	r := handlers.InitRouter(server)
	err = r.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Fatalln(fmt.Errorf("r.Run() failed: %w", err))
	}
}

func loadConfig() error {
	viper.SetConfigType("env") // REQUIRED if the config file does not have the extension in the name
	viper.SetConfigFile(".env")
	viper.AddConfigPath("./config") // call multiple times to add many search paths
	viper.AddConfigPath(".")        // optionally look for config in the working directory
	viper.AutomaticEnv()
	return viper.ReadInConfig() // Find and read the config file
}
