package config

import (
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

//Load is a function that set configuration
func Load(configStruct interface{}, configPath string) error {
	dir, file := filepath.Split(configPath)
	ext := path.Ext(file)
	fileExt := strings.TrimPrefix(ext, ".")
	fileName := strings.TrimSuffix(file, ext)

	viper.SetConfigName(fileName)
	viper.SetConfigType(fileExt)
	viper.AddConfigPath(dir)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(configStruct)
	if err != nil {
		return err
	}

	return nil
}
