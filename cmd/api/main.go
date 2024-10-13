package main

import (
	"fmt"
	"log"

	"github.com/takeuchi-shogo/go-clean-ddd/internal/api/infrastructure"
	"github.com/takeuchi-shogo/go-clean-ddd/pkg/infra/config"
	"github.com/takeuchi-shogo/go-clean-ddd/pkg/infra/database"
)

func main() {
	c := infrastructure.NewConfig("LOCAL")
	fmt.Printf("%+v", c)

	r := infrastructure.NewRouter(c)

	cfg := config.New("LOCAL")
	fmt.Printf("%+v\n", cfg)

	_, err := database.NewDB(*cfg)
	if err != nil {
		log.Println(err)
	}

	r.Run()
}
