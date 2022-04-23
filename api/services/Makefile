run-single-test:
	ENV_FILE=../.env go test -v -count=1 -run $(test)

run-all-tests:
	ENV_FILE=../.env go test ./... -v -parallel=1 -count=1