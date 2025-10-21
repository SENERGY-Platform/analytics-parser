module github.com/SENERGY-Platform/analytics-parser

go 1.24.0

toolchain go1.24.6

//replace github.com/SENERGY-Platform/analytics-fog-lib => ../analytics-fog/analytics-fog-lib

require (
	github.com/SENERGY-Platform/analytics-flow-repo-v2/lib v0.0.0-20251021074433-b7305284d859
	github.com/SENERGY-Platform/analytics-fog-lib v1.1.26
	github.com/SENERGY-Platform/analytics-parser/lib v0.0.0-00010101000000-000000000000
	github.com/SENERGY-Platform/gin-middleware v0.12.0
	github.com/SENERGY-Platform/go-service-base/config-hdl v1.2.0
	github.com/SENERGY-Platform/go-service-base/srv-info-hdl v0.2.0
	github.com/SENERGY-Platform/go-service-base/struct-logger v0.4.1
	github.com/SENERGY-Platform/go-service-base/util v1.1.0
	github.com/SENERGY-Platform/service-commons v0.0.0-20250903071414-1b34f1965afa
	github.com/gin-contrib/cors v1.7.6
	github.com/gin-contrib/requestid v1.0.5
	github.com/gin-gonic/gin v1.11.0
	github.com/parnurzeal/gorequest v0.3.0
	github.com/pkg/errors v0.9.1
)

require (
	github.com/SENERGY-Platform/go-env-loader v0.5.3 // indirect
	github.com/bytedance/gopkg v0.1.3 // indirect
	github.com/bytedance/sonic v1.14.1 // indirect
	github.com/bytedance/sonic/loader v0.3.0 // indirect
	github.com/cloudwego/base64x v0.1.6 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/elazarl/goproxy v1.7.2 // indirect
	github.com/gabriel-vasile/mimetype v1.4.10 // indirect
	github.com/gin-contrib/sse v1.1.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.28.0 // indirect
	github.com/goccy/go-json v0.10.5 // indirect
	github.com/goccy/go-yaml v1.18.0 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.18.1 // indirect
	github.com/klauspost/cpuid/v2 v2.3.0 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.3-0.20250322232337-35a7c28c31ee // indirect
	github.com/moul/http2curl v1.0.0 // indirect
	github.com/pelletier/go-toml/v2 v2.2.4 // indirect
	github.com/pierrec/lz4/v4 v4.1.22 // indirect
	github.com/quic-go/qpack v0.5.1 // indirect
	github.com/quic-go/quic-go v0.55.0 // indirect
	github.com/segmentio/kafka-go v0.4.49 // indirect
	github.com/smartystreets/goconvey v1.8.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.3.0 // indirect
	go.mongodb.org/mongo-driver v1.17.4 // indirect
	go.uber.org/mock v0.6.0 // indirect
	golang.org/x/arch v0.22.0 // indirect
	golang.org/x/crypto v0.43.0 // indirect
	golang.org/x/exp v0.0.0-20251017212417-90e834f514db // indirect
	golang.org/x/mod v0.29.0 // indirect
	golang.org/x/net v0.46.0 // indirect
	golang.org/x/sync v0.17.0 // indirect
	golang.org/x/sys v0.37.0 // indirect
	golang.org/x/text v0.30.0 // indirect
	golang.org/x/tools v0.38.0 // indirect
	google.golang.org/protobuf v1.36.10 // indirect
)

replace github.com/SENERGY-Platform/analytics-parser/lib => ./lib
