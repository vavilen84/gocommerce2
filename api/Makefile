include .env
export $(shell sed 's/=.*//' .env)

run:
	bee run -gendoc=true -downdoc=true

# development commands
run-migrations:
	bee migrate -driver=mysql -conn="${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}"