module github.com/SENERGY-Platform/analytics-parser

go 1.24.0

toolchain go1.24.6

//replace github.com/SENERGY-Platform/analytics-fog-lib => ../analytics-fog/analytics-fog-lib

require (
	github.com/SENERGY-Platform/analytics-fog-lib v1.1.26
	github.com/SENERGY-Platform/go-service-base/struct-logger v0.4.1
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/gorilla/mux v1.8.1
	github.com/joho/godotenv v1.5.1
	github.com/parnurzeal/gorequest v0.3.0
	github.com/pkg/errors v0.9.1
	github.com/rs/cors v1.11.1
)

require (
	github.com/elazarl/goproxy v0.0.0-20210110162100-a92cc753f88e // indirect
	github.com/moul/http2curl v1.0.0 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	golang.org/x/net v0.43.0 // indirect
)
