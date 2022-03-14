package models

import (
	"api/constants"
	"api/env"
	"api/store"
	"database/sql"
	"github.com/beego/beego/v2/core/logs"
	"github.com/joho/godotenv"
	"os"
	"os/exec"
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

func clearTestDb() {
	// use credentials without db in order to create db
	db, err := sql.Open(env.GetSQLDriver(), env.GetSQLServerDsn())
	if err != nil {
		logs.Error(err)
	}
	ctx := store.GetDefaultDBContext()
	conn, err := db.Conn(ctx)
	if err != nil {
		logs.Error(err)
	}
	defer conn.Close()
	_, err = conn.ExecContext(ctx, "DROP DATABASE "+env.GetMySQLTestDb())
	if err != nil {
		logs.Error(err)
	}
	_, err = conn.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+env.GetMySQLTestDb())
	if err != nil {
		logs.Error(err)
	}
}

func beforeEachTest() {
	if !testAppInited {
		initConfig()
		setTestAppEnv()
		store.InitTestORM()
		initFixtures()
		testAppInited = true
	}
	clearTestDb()
	restoreFromDump()
}

func setTestAppEnv() {
	err := os.Setenv(constants.AppEnvEnvVar, constants.TestingAppEnv)
	if err != nil {
		logs.Error(err)
	}
}

func restoreFromDump() {
	os.Chdir(env.GetAppRoot())
	cmd := exec.Command("make", constants.RestoreTestDbDump)
	_, err := cmd.Output()
	if err != nil {
		logs.Error(err)
	}
}

func getDbDumpFile() string {
	return path.Join(env.GetAppRoot(), constants.TmpFolder, constants.TestDbDumpFile)
}
