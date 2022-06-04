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

type IBANValidationResponse struct {
	Valid         bool   `json:"valid"`
	InvalidReason string `json:"invalidReason,omitempty"`
}

func validateIBAN(c *gin.Context) {
	rawData, err := c.GetRawData()
	if err != nil || len(rawData) == 0 {
		c.JSON(http.StatusBadRequest, Response{Error: "failed to parse input"})
		return
	}

	// TODO: not sure if error is the correct value to use
	// as a "bad" response.
	valid, err := iban.ValidateIBAN(string(rawData))
	validationResponse := IBANValidationResponse{Valid: valid}
	if err != nil {
		validationResponse.InvalidReason = err.Error()
	}

	c.JSON(http.StatusOK, Response{Data: validationResponse})
}

func main() {
	r := gin.Default()
	r.POST("/iban/validate", validateIBAN)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
