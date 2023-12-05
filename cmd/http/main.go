package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/Nickadimas79/calculate-lib-example/api/shapes"
)

const PORT = "8080"

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/area/rect", shapes.GetRectangleArea)
		v1.GET("/area/circle", shapes.GetCircleArea)
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/area/circle", shapes.GetCircleArea)
	}

	err := r.Run(":" + PORT)
	if err != nil {
		fmt.Println("error starting server")
	}

	fmt.Println("server running on port", PORT)
}
