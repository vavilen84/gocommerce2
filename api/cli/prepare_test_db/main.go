package main

import (
	"api/constants"
	"api/env"
	"api/helpers"
	"api/store"
	"database/sql"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/joho/godotenv"
	"os"
	"path"
)

func initConfig() {
	envFile := os.Getenv(constants.EnvFileVar)
	err := godotenv.Load(envFile)
	if err != nil {
		logs.Error(err)
	}
}

func main() {
	initConfig()
	clearTestDb()
	runMigrations()
	createDbDump()
}

func runMigrations() {
	helpers.RunCmd(
		"bee",
		"migrate",
		"-driver="+env.GetSQLDriver(),
		"-conn="+env.GetDbDsn(env.GetMySQLTestDb()),
	)
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

func createDbDump() {
	os.Remove(getDbDumpFile())
	dockerizedDb := env.GetDockerizedDB()
	dbDockerService := env.GetDockerMysqlService()
	if dockerizedDb {
		format := "docker exec %s /usr/bin/mysqldump -u %s --password=%s %s > %s"
		cmd := fmt.Sprintf(
			format,
			dbDockerService,
			env.GetMySQLUser(),
			env.GetMySQLUserPass(),
			env.GetMySQLTestDb(),
			getDbDumpFile(),
		)
		helpers.RunCmd("/bin/sh", "-c", cmd)
	} else {
		helpers.RunCmd(
			"mysqldump",
			"-u"+env.GetMySQLUser(),
			"-p"+env.GetMySQLUserPass(),
			env.GetMySQLTestDb(),
			"--result-file="+getDbDumpFile(),
		)
	}
}

func getDbDumpFile() string {
	return path.Join(env.GetAppRoot(), constants.TmpFolder, constants.TestDbDumpFile)
}
