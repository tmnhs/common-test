module github.com/tmnhs/common-test

go 1.16

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.8.1
	github.com/pkg/errors v0.9.1
	github.com/tmnhs/common v1.0.1
	go.uber.org/atomic v1.9.0 // indirect
	golang.org/x/sync v0.1.0
	gorm.io/gorm v1.23.10
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
