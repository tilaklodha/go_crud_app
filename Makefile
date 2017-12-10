all: build

APP = go_crud_app
DB_USER = "postgres"
APP_EXECUTABLE = "./out/$(APP)"
DB_NAME = "users"
TEST_DB_NAME = "go_crud_app_test"

build-deps:
	glide install

compile:
	go build -o $(APP_EXECUTABLE)

build: build-deps compile
