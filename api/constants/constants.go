package constants

import "time"

const (
	//common
	SqlDsnFormat   = `%s:%s@tcp(%s:%s)/%s`
	TmpFolder      = "tmp"
	TestDbDumpFile = "test_db_dump.sql"
	SecretEnvVar   = "SECRET"

	// app config env vars
	JWTExpireDurationDays = "JWT_EXPIRE_DURATION_DAYS"

	// DEPRECATED! use default db & user for testing purposes
	// test db & test user
	TestingDB                   = "testing" // reference docker/db/docker-entrypoint-initdb.d/test.db-sql
	MysqlTestUserEnvVar         = "MYSQL_TEST_USER"
	MysqlTestUserPasswordEnvVar = "MYSQL_TEST_USER_PASSWORD"
	MysqlTestDBEnvVar           = "MYSQL_TEST_DATABASE"

	// env vars
	MysqlUserEnvVar     = "MYSQL_USER"
	MysqlDBEnvVar       = "MYSQL_DATABASE"
	MysqlPasswordEnvVar = "MYSQL_PASSWORD"

	AppRootEnvVar     = "APP_ROOT"
	ProjectRootEnvVar = "PROJECT_ROOT"

	SqlDriverEnvVar          = "SQL_DRIVER"
	MysqlPortEnvVar          = "MYSQL_PORT"
	MysqlHostEnvVar          = "MYSQL_HOST"
	DockerMysqlServiceEnvVar = "DOCKER_MYSQL_SERVICE"
	DockerizedDBEnvVar       = "DOCKERIZED_DB"

	// app env
	EnvFileVar   = "ENV_FILE"
	AppEnvEnvVar = "APP_ENV"

	DevelopmentAppEnv = "development"
	ProductionAppEnv  = "production"
	TestingAppEnv     = "testing"

	// makefile commands
	RestoreTestDbDump = "restore-test-db-dump"

	// db
	DefaultDBAlias = "default"

	// db tables
	MigrationDBTable = "migration"
	ProductDBTable   = "product"
	CustomerDBTable  = "customer"
	OrderDBTable     = "order"
	TaxDBTable       = "tax"
	JWTInfoDBTable   = "jwt_info"
	UserDBTable      = "user"
	CategoryDBTable  = "category"
	DiscountDBTable  = "discount"

	OrderTaxDBTable             = "order_tax"
	OrderProductDBTable         = "order_product"
	OrderDiscountDBTable        = "order_discount"
	OrderProductTaxDBTable      = "order_product_tax"
	OrderProductDiscountDBTable = "order_product_discount"

	Product2CategoryDBTable = "product_2_category"

	// Server
	DefaultWriteTimout  = 60 * time.Second
	DefaultReadTimeout  = 60 * time.Second
	DefaultStoreTimeout = 60 * time.Second

	// model names
	MigrationModel = "Migration"
	ProductModel   = "Product"
	UserModel      = "User"
	JWTInfoModel   = "JWTInfo"
	CategoryModel  = "Category"
	DiscountModel  = "Discount"

	//tax types
	TaxCartType     = 1
	TaxCategoryType = 2
	TaxProductType  = 3

	//discount types
	DiscountCartType     = 1
	DiscountCategoryType = 2
	DiscountProductType  = 3
)
