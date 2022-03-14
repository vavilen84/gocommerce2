package env

import (
	"api/constants"
	"fmt"
	"os"
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
