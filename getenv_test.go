package getenv

import (
	"fmt"
	"testing"
)

var FlagEnvFile = "./.env.example"

func TestGetEnv(t *testing.T) {
	env := GetEnv{}
	env.SetFilePath(FlagEnvFile).Init()

	appName := env.GetVal("APP_NAME").String()
	fmt.Println("appName: ", appName)
}
