package configurator

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/spf13/viper"
	"log"
)

type ServiceConfig struct {
	Domain           string
	Port             int
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
}

type MailConfig struct {
	SMTPServerDomain string
	SMTPServerPort int
	Mail string
	Password string
}

type TarantoolConfig struct {
	Domain string
	Port int
	User string
	Password string
}

type Config struct {
	Auth           ServiceConfig
	Profile        ServiceConfig
	App            ServiceConfig
	DatabaseDomain string
	DatabasePort   int
	Tarantool TarantoolConfig
	Mail 		   MailConfig
}

func Run(configPath string) (*Config, error) {
	viper.SetDefault("Auth.Domain", "auth")
	viper.SetDefault("Profile.Domain", "profile")
	viper.SetDefault("App.Domain", "app")

	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Unable to read config file: %s", err)
		return nil, models.ErrFooIncorrectPath
	}

	config := new(Config)
	config.Auth.Domain = viper.GetString("Auth.Domain")
	config.Auth.Port = viper.GetInt("Auth.Port")
	config.Auth.DatabaseName = viper.GetString("Auth.Database.Name")
	config.Auth.DatabaseUser = viper.GetString("Auth.Database.User")
	config.Auth.DatabasePassword = viper.GetString("Auth.Database.Password")
	config.Profile.Domain = viper.GetString("Profile.Domain")
	config.Profile.Port = viper.GetInt("Profile.Port")
	config.Profile.DatabaseName = viper.GetString("Profile.Database.Name")
	config.Profile.DatabaseUser = viper.GetString("Profile.Database.User")
	config.Profile.DatabasePassword = viper.GetString("Profile.Database.Password")
	config.App.Domain = viper.GetString("App.Domain")
	config.App.Port = viper.GetInt("App.Port")
	config.App.DatabaseName = viper.GetString("App.Database.Name")
	config.App.DatabaseUser = viper.GetString("App.Database.User")
	config.App.DatabasePassword = viper.GetString("App.Database.Password")
	config.DatabaseDomain = viper.GetString("Database.Domain")
	config.DatabasePort = viper.GetInt("Database.Port")
	config.Mail.Mail = viper.GetString("Mail.Mail")
	config.Mail.Password = viper.GetString("Mail.Password")
	config.Mail.SMTPServerDomain = viper.GetString("Mail.Domain")
	config.Mail.SMTPServerPort = viper.GetInt("Mail.Port")

	config.Tarantool.Domain = viper.GetString("Tarantool.Domain")
	config.Tarantool.Port = viper.GetInt("Tarantool.Port")
	config.Tarantool.User = viper.GetString("Tarantool.User")
	config.Tarantool.Password = viper.GetString("Tarantool.Password")


	return config, nil
}
