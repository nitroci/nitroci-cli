/*
Copyright 2021 The NitroCI Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
)

func globalConfig() string {
	nitrociConfg := os.Getenv("NITROCICONFIG")
	if len(nitrociConfg) == 0 {
		nitrociConfg = "$HOME/.nitroci/config"
	}
	return nitrociConfg
}

func FindGlobalconfig() (file string) {
	return globalConfig() + ".ini"
}

func EnsureConfiguration() {
	nitrociConfg := globalConfig()
	configHome := filepath.Dir(nitrociConfg)
	configName := filepath.Base(nitrociConfg)
	configType := "ini"
	configPath := filepath.Join(configHome, configName)
	viper.AddConfigPath(configHome)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	if _, err := os.Stat(configHome); os.IsNotExist(err) {
		os.MkdirAll(configHome, 0700)
	}
	_, err := os.Stat(configPath)
	if !os.IsExist(err) {
		if err := viper.SafeWriteConfig(); err != nil {
		}
	}
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
}

func GetGlobalConfigString(profile string, key string) string {
	EnsureConfiguration()
	value := viper.Get(fmt.Sprintf("%v.%v", profile, key))
	if value != nil {
		return fmt.Sprintf("%v", value)
	}
	return ""
}

func SetGlobalConfigString(profile string, key string, value string) {
	EnsureConfiguration()
	viper.Set(fmt.Sprintf("%v.%v", profile, key), value)
	viper.WriteConfig()
}

func PromptGlobalConfigKeyAndSave(profile string, label string, secret bool, key string, save bool) (bool, string) {
	prompt := promptui.Prompt{
		Label:       label,
		HideEntered: true,
	}
	if secret == true {
		prompt.Mask = '*'
	}
	value, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return false, ""
	}
	if save == true {
		SetGlobalConfigString(profile, key, value)
	}
	return true, value
}

func PromptGlobalConfigKey(profile string, label string, secret bool) (bool, string) {
	return PromptGlobalConfigKeyAndSave(profile, label, secret, "", false)
}
