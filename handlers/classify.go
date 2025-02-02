package handlers

import (
	"net/http"
	"num_class_api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Number     int      `json:"number"`
	IsPrime    bool     `json:"is_prime"`
	IsPerfect  bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	DigitSum   int      `json:"digit_sum"`
	FunFact    string   `json:"fun_fact"`
}

func classifyNumber(c *gin.Context) {
	numberstr := c.Query("number")

	num, err := strconv.Atoi(numberstr)
	if err != nil {
		type Response struct {
			Number string `json:"number"`
			Error  bool   `json:"error"`
		}
		resp := Response{
			Number: "alphabet",
			Error:  true,
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	primeNum := utils.IsPrime(num)
	perfectNum := utils.IsPerfect(num)
	prop := utils.ClassifyProperties(num)
	digitSum := utils.CalculateDigitSum(num)
	funfact, _ := utils.GetFuncFact(num)

	resp := Response{
		Number:     num,
		IsPrime:    primeNum,
		IsPerfect:  perfectNum,
		Properties: prop,
		DigitSum:   digitSum,
		FunFact:    funfact,
	}

	c.JSON(http.StatusOK, resp)

}

func RegisterRoutes(router *gin.Engine) {
	router.GET("/api/classify-number", classifyNumber)
}
