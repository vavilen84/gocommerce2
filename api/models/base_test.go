package models

import (
	"api/constants"
	"api/env"
	"api/helpers"
	"api/store"
	"database/sql"
	"github.com/beego/beego/v2/core/logs"
	"os"
	"path"
	"sync"
	"testing"
)

func TestMain(m *testing.M) {
	var once sync.Once
	onceCall := func() {
		beforeAllTestRun()
	}
	once.Do(onceCall)
	code := m.Run()
	os.Exit(code)
}

func beforeEachTest() {
	restoreFromDump()
}

func beforeAllTestRun() {
	setTestAppEnv()
	store.InitTestORM()
	prepareTestDB()
}

func setTestAppEnv() {
	err := os.Setenv(constants.AppEnvEnvVar, constants.TestingAppEnv)
	if err != nil {
		logs.Error(err)
	}
}

func runMigrations() {
	helpers.RunCmd(
		"bee",
		"migrate",
		"-driver="+env.GetSQLDriver(),
		"-conn="+env.GetDbDsn(env.GetMySQLTestDb()),
	)
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

func createDbDump() {
	os.Remove(getDbDumpFile())
	helpers.RunCmd(
		"mysqldump",
		"-u"+env.GetMySQLUser(),
		"-p"+env.GetMySQLUserPass(),
		env.GetMySQLTestDb(),
		"--result-file="+getDbDumpFile(),
	)
}

/**
 * ! IMPORTANT - dont use for production DB !
 */
func prepareTestDB() {
	clearTestDb()
	runMigrations()
	createDbDump()
}

/**
 * ! IMPORTANT - dont use for production DB !
 */
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
