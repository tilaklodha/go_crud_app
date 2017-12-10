all: build testdb.create test

APP = go_crud_app
DB_USER = "postgres"
APP_EXECUTABLE = "./out/$(APP)"
DB_NAME = "users"
TEST_DB_NAME = "go_crud_app_test"
UNIT_TEST_PACKAGES=$(shell glide novendor | grep -v "featuretests")

build-deps:
	glide install

compile:
	go build -o $(APP_EXECUTABLE)

build: build-deps compile

testdb.create: testdb.drop
	createdb -O $(DB_USER) -Eutf8 $(TEST_DB_NAME)

testdb.drop:
	dropdb --if-exists -U $(DB_USER) $(TEST_DB_NAME)

test:
	ENVIRONMENT=test go test $(UNIT_TEST_PACKAGES) -p=1
