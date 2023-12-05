package shapes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Nickadimas79/calculate-lib-example/pkg/geometry/circle"
	"github.com/Nickadimas79/calculate-lib-example/pkg/shapes"
)

type Rectangle struct {
	Length float32 `json:"Length"`
	Width  float32 `json:"width"`
	Height float32 `json:"height"`
}

func GetRectangleArea(c *gin.Context) {
	re := &Rectangle{}

	err := c.BindJSON(re)
	if err != nil {
		fmt.Println("error with JSON binding:", err)
	}

	c.JSON(http.StatusOK, gin.H{"area": shapes.GetRectangleArea(re.Length, re.Width)})
}

func GetRectangleVolume(c *gin.Context) {
	re := &Rectangle{}

	err := c.BindJSON(re)
	if err != nil {
		fmt.Println("error with JSON binding:", err)
	}

	c.JSON(http.StatusOK, gin.H{"volume": shapes.GetRectangleVolume(re.Length, re.Width, re.Height)})
}

// GetCircleArea HTTP implementation for getting area of a circle.
func GetCircleArea(c *gin.Context) {
	cr := &circle.Circle{}

	err := c.BindJSON(cr)
	if err != nil {
		fmt.Println("error with JSON binding:", err)
	}

	c.JSON(http.StatusOK, gin.H{"area": cr.Area()})

}
