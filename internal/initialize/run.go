package initialize

import (
	"fmt"

	"github.com/buiminhhoat/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	fmt.Println("Loading configuration mysql: ", global.Config.Mysql.Username)
	InitLogger()
	global.Logger.Info("Config log ok!", zap.String("ok", "success"))
	InitMySql()
	InitRedis()

	r := InitRouter()

	r.Run(":8888")
}
