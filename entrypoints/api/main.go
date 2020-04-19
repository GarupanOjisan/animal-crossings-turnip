package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oharai/animal-crossings-turnip/app/api/buyer"
	"github.com/oharai/animal-crossings-turnip/app/api/seller"
	"github.com/oharai/animal-crossings-turnip/config"
	"log"
)

func main() {
	r := gin.Default()

	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	registerRoutes(r, conf)

	log.Println("start listening on :8080")
	if err := r.Run(":8090"); err != nil {
		log.Fatal(err)
	}
}

func registerRoutes(r *gin.Engine, c *config.Config) {
	s := seller.Endpoint{Conf: c}
	sr := r.Group("/seller")
	sr.POST("", s.Register)
	sr.GET("", s.Get)

	b := buyer.Endpoint{Conf: c}
	br := r.Group("/buyer")
	br.POST("", b.Register)
	br.GET("", b.Get)
}
