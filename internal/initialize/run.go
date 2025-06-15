package initialize

import (
	"fmt"

	"github.com/buiminhhoat/go-ecommerce-backend-api/global"
)

func Run() {
	LoadConfig()
	fmt.Println("Loading configuration mysql: ", global.Config.Mysql.Username)
	InitLogger()
	InitMySql()
	InitRedis()

	r := InitRouter()

	r.Run(":8888")
}
