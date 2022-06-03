package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sula0/ibanator/v2/pkg/iban"
	"net/http"
)

func validateIBAN(c *gin.Context) {
	rawData, err := c.GetRawData()
	if err != nil {
		panic(err)
	}

	valid := iban.ValidateIBAN(string(rawData))

	c.IndentedJSON(http.StatusOK, valid)
}

func main() {
	r := gin.Default()

	r.POST("/", validateIBAN)

	r.Run()
}
