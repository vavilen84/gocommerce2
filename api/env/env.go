package env

import (
	"api/constants"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"os"
	"strconv"
)

func GetDbDsn(dbname string) string {
	return fmt.Sprintf(
		constants.SqlDsnFormat,
		os.Getenv(constants.MysqlUserEnvVar),
		os.Getenv(constants.MysqlPasswordEnvVar),
		os.Getenv(constants.MysqlHostEnvVar),
		os.Getenv(constants.MysqlPortEnvVar),
		dbname,
	)
}

func GetSQLServerDsn() string {
	return fmt.Sprintf(
		constants.SqlDsnFormat,
		os.Getenv(constants.MysqlUserEnvVar),
		os.Getenv(constants.MysqlPasswordEnvVar),
		os.Getenv(constants.MysqlHostEnvVar),
		os.Getenv(constants.MysqlPortEnvVar),
		"",
	)
}

func GetMySQLUser() string {
	return os.Getenv(constants.MysqlUserEnvVar)
}

func GetMySQLTestUser() string {
	return os.Getenv(constants.MysqlTestUserEnvVar)
}

func GetMySQLUserPass() string {
	return os.Getenv(constants.MysqlPasswordEnvVar)
}

func getEnvVarIntValue(v string) int {
	resultString := os.Getenv(v)
	resultInt, err := strconv.Atoi(resultString)
	if err != nil {
		logs.Error(err)
	}
	return resultInt
}

func getEnvVarBoolValue(v int) bool {
	if v == 1 {
		return true
	}
	return false
}

func GetDockerMysqlService() string {
	return os.Getenv(constants.DockerMysqlServiceEnvVar)
}

func GetDockerizedDB() bool {
	return getEnvVarBoolValue(getEnvVarIntValue(constants.DockerizedDBEnvVar))
}

func GetMySQLTestUserPass() string {
	return os.Getenv(constants.MysqlTestUserPasswordEnvVar)
}

func GetSQLDriver() string {
	return os.Getenv(constants.SqlDriverEnvVar)
}

func GetMySQLDb() string {
	return os.Getenv(constants.MysqlDBEnvVar)
}

func GetMySQLTestDb() string {
	return os.Getenv(constants.MysqlTestDBEnvVar)
}

func GetAppRoot() string {
	return os.Getenv(constants.AppRootEnvVar)
}

func GetSecret() string {
	return os.Getenv(constants.SecretEnvVar)
}

func GetJWTExpireDurationDays() int {
	return getEnvVarIntValue(constants.JWTExpireDurationDays)
}
