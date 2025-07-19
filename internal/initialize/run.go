package initialize

import (
	"fmt"

	"github.com/buiminhhoat/go-ecommerce-backend-api/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Run() *gin.Engine {
	LoadConfig()
	fmt.Println("Loading configuration mysql: ", global.Config.Mysql.Username)
	InitLogger()
	global.Logger.Info("Config log ok!", zap.String("ok", "success"))
	InitMySql()
	InitMySqlC()
	InitServiceInterface()
	// InitRedis()
	InitRedisSentinel()
	InitKafka()

	r := InitRouter()

	return r
}
