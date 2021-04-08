package config

import (
	"encoding/json"
	"fmt"
	"hermes-service/model"
	"os"
)

func BuildConfig(env string) model.Config {
	var c model.Config
	err := readEnvJSON(env, &c)
	if err != nil {
		panic(fmt.Sprintf("can't read config: %v", err))
	}

	return c
}

// readEnvJSON unmarshalls the content of "env_{env} file in out struct pointer"
func readEnvJSON(env string, out interface{}) error {
	file, err := os.Open(fmt.Sprintf("env_%s.json", env))
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(out)
	if err != nil {
		file.Close()
		return err
	}

	return file.Close()
}
