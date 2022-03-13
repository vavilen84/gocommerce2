package models

import (
	"api/constants"
	"api/env"
	"api/helpers"
	"api/store"
	"github.com/beego/beego/v2/core/logs"
	"github.com/joho/godotenv"
	"os"
	"path"
)

var testAppInited = false

func initConfig() {
	envFile := os.Getenv(constants.EnvFileVar)
	// means, we use IDE test debug run
	if envFile == "" {
		envFile = "../.env"
	}
	err := godotenv.Load(envFile)
	if err != nil {
		logs.Error(err)
	}
}

func beforeEachTest() {
	if !testAppInited {
		initConfig()
		setTestAppEnv()
		store.InitTestORM()
		testAppInited = true
	}
	restoreFromDump()
}

func setTestAppEnv() {
	err := os.Setenv(constants.AppEnvEnvVar, constants.TestingAppEnv)
	if err != nil {
		logs.Error(err)
	}
}

func restoreFromDump() {
	helpers.RunCmd(
		"mysql",
		"-u"+env.GetMySQLUser(),
		"-p"+env.GetMySQLUserPass(),
		env.GetMySQLTestDb(),
		"<",
		getDbDumpFile(),
	)
}

func getDbDumpFile() string {
	return path.Join(env.GetAppRoot(), constants.TmpFolder, constants.TestDbDumpFile)
}
