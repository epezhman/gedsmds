package config

import (
	"github.com/IBM/gedsMDS/internal/logger"
	"github.com/IBM/gedsMDS/internal/profiling"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

var Config *Configuration

type Configuration struct {
	UUID                string `mapstructure:"UUID"`
	TargetSystem        string `mapstructure:"TARGET_SYSTEM"`
	MDSServerPort       string `mapstructure:"TRANSACTION_SERVER_PORT"`
	ProfilingServerPort string `mapstructure:"PROFILING_SERVER_PORT"`
	ProfilingEnabled    string `mapstructure:"PROFILING_ENABLED"`
}

func init() {
	var err error
	Config, err = LoadConfig()
	if err != nil {
		logger.FatalLogger.Fatalln(err)
	}
}

func LoadConfig() (*Configuration, error) {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		return &Configuration{}, err
	}
	appUUID := viper.GetString("UUID")
	if len(appUUID) == 0 {
		viper.Set("UUID", uuid.NewString())
		err = viper.WriteConfig()
		if err != nil {
			return &Configuration{}, err
		}
	}
	config := &Configuration{}
	err = viper.Unmarshal(config)
	if err != nil {
		return &Configuration{}, err
	}
	if config.ProfilingEnabled == "enable_server" {
		profiling.StartProfilingServer(config.ProfilingServerPort)
	} else if config.ProfilingEnabled == "cpu" {
		profiling.StartCPUProfiling()
	} else if config.ProfilingEnabled == "memory" {
		profiling.StartMemoryProfiling()
	} else if config.ProfilingEnabled == "bandwidth" {
		profiling.StartBandWithProfiling()
	}
	return config, nil
}
