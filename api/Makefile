include .env
export $(shell sed 's/=.*//' .env)

run:
	bee run -gendoc=true -downdoc=true

# development commands
run-test-db-migrations:
	bee migrate -driver=mysql -conn="${MYSQL_TEST_USER}:${MYSQL_TEST_USER_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_TEST_DATABASE}"

create-test-db-dump:
	mysqldump -u${MYSQL_TEST_USER} -p${MYSQL_TEST_USER_PASSWORD} ${MYSQL_TEST_DATABASE} --result-file=tmp/test_db_dump.sql

prepare-test-db:
	ENV_FILE=.env go run cli/prepare_test_db/main.go