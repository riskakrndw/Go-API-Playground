package config

const DB_CONNECTION_STRING = "root:welcome12345@tcp(localhost:3306)/rest_schema?charset=utf8&parseTime=True&loc=Local"
const TEST_DB_CONNECTION_STRING = "root:welcome12345@tcp(localhost:3306)/test_schema?charset=utf8&parseTime=True&loc=Local"

/*
go test -v -coverpkg=./... -coverprofile=profile.cov ./...
go tool cover -func profile.cov
*/
