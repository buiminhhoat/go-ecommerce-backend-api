package main

import (
	"github.com/buiminhhoat/go-ecommerce-backend-api/internal/router"
)

func main() {
	r := router.NewRouter()

	r.Run(":8888") // listen and serve on 0.0.0.0:8888 (for windows "localhost:8888")
}
