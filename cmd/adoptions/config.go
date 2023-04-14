package main

import (
	"time"
	"errors"
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

const (
	ConfigPathDev string = "/home/mr-papi/SoftwareCode/Projects/PetPark/"
	ConfigPathSystem string = "/etc/"
	ConfigPathUser   string = "/home/mr-papi/.config/"
	ConfigFileName string = "petpark.conf"
	ConfigFileType string = "toml"
)

type configuration struct {
	LogFile    string `mapstructure:"LogFile"`
	DbDriver   string `mapstructure:"DbDriver"`
	DbProtocol string `mapstructure:"DbProtocol"`
	DbHost     string `mapstructure:"DbHost"`
	DbPort     int32  `mapstructure:"DbPort"`
	DbUser     string `mapstructure:"DbUser"`
	DbPassword string `mapstructure:"DbPassword"`
	DbName     string `mapstructure:"DbName"`
	ServerHost string `mapstructure:"ServerHost"`
	ServerPort string `mapstructure:"ServerPort"`
	BaseTimeout time.Duration `mapstructure:"BaseTimeout"`
}

var (
	Config configuration
	LogFile *os.File
	Logger  zerolog.Logger
)

func InitConfig() (err error) {
	viper.SetConfigName(ConfigFileName)
	viper.SetConfigType(ConfigFileType)
	viper.AddConfigPath(ConfigPathDev)
	viper.AddConfigPath(ConfigPathSystem)
	viper.AddConfigPath(ConfigPathUser)

	err = viper.ReadInConfig()
	if err != nil {
		panic(errors.New("failed to load config file"))
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(errors.New("failed to save config"))
	}

	return nil
}

func InitLogger() (err error) {
	LogFile, err = os.OpenFile(Config.LogFile, os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(errors.New("failed to open logfile"))
	}

	fileWriter := zerolog.New(LogFile).With().Logger()
	Logger = zerolog.New(fileWriter).With().Timestamp().Logger()
	return nil
}
