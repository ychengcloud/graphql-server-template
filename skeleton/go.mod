module {{ .Extra.pkgpath }}

go 1.16

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/HdrHistogram/hdrhistogram-go v1.0.0 // indirect
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-contrib/pprof v1.3.0
	github.com/gin-contrib/zap v0.0.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.4.0
	github.com/google/uuid v1.0.0
	github.com/google/wire v0.4.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/ory/kratos-client-go v0.6.3-alpha.1
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.7.1
	github.com/smartystreets/assertions v1.0.0 // indirect
	github.com/spf13/viper v1.7.1
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.4.0+incompatible
	github.com/vektah/gqlparser/v2 v2.1.0
	github.com/vmihailenco/msgpack/v5 v5.3.1
	go.uber.org/zap v1.16.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v3 v3.0.0-20200605160147-a5ece683394c // indirect
	gorm.io/driver/mysql v1.0.3
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.20.7
)
