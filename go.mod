module github.com/SENERGY-Platform/analytics-parser

go 1.21.3

//replace github.com/SENERGY-Platform/analytics-fog-lib => ../analytics-fog/analytics-fog-lib

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/gorilla/mux v1.8.0
	github.com/joho/godotenv v1.5.1
	github.com/parnurzeal/gorequest v0.2.16
	github.com/pkg/errors v0.9.1
	github.com/rs/cors v1.10.1
)

require (
	github.com/SENERGY-Platform/analytics-fog-lib v1.0.15 // indirect
	github.com/elazarl/goproxy v0.0.0-20210110162100-a92cc753f88e // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	golang.org/x/net v0.15.0 // indirect
	moul.io/http2curl v1.0.0 // indirect
)
