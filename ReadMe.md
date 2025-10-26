go mod init github.com/yourname/url-shortener
go get github.com/gin-gonic/gin gorm.io/gorm gorm.io/driver/postgres github.com/go-redis/redis/v8 github.com/rs/zerolog


go get gorm.io/driver/sqlite


go mod tidy
go run ./cmd/server
