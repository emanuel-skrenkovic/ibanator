package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sula0/ibanator/v2/pkg/iban"
	"net/http"
)

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func validateIBAN(c *gin.Context) {
	rawData, err := c.GetRawData()

	if err != nil || len(rawData) == 0 {
		errorMessage := "failed to parse input"
		c.JSON(http.StatusBadRequest, Response{Error: errorMessage})
		return
	}

	valid := iban.ValidateIBAN(string(rawData))

	c.JSON(http.StatusOK, Response{Data: valid})
}

func main() {
	r := gin.Default()
	r.POST("/iban/validate", validateIBAN)
	r.Run()
}
