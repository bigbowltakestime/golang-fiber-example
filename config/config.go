package config

import (
	"github.com/spf13/viper"
)

// Config는 애플리케이션 설정을 담는 구조체입니다.
type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Database struct {
		ConnectionString string `mapstructure:"connection_string"`
	} `mapstructure:"database"`
	// 다른 설정들을 필요에 따라 추가할 수 있습니다.
}

// LoadConfig는 설정을 로드하는 함수입니다.
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config") // 설정 파일 이름 (예: config.yaml)
	viper.AddConfigPath(".")      // 설정 파일이 위치한 디렉토리

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
