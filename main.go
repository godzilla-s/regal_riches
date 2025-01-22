package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/godzilla-s/regal-riches/pkg/service"
)

func main() {
	engine := gin.Default()
	svc := service.NewService(&service.Config{})

	route := engine.Group("/regal_riches/v1")
	route.POST("/login", svc.Login)
	route.POST("/registry", svc.Registry)
	route.POST("/recieve", svc.ReciveRR)
	route.POST("/pay", svc.PayRR)
	route.GET("/balance/rr", svc.GetRRBalance)
	route.POST("/deposit", svc.DepositTON)
	route.POST("/withdraw", svc.WithdrawProposal)
	route.GET("/balance/ton", svc.GetTonBalance)
	log.Panic(engine.Run(":8082"))
}
