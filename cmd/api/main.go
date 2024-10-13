package main

import (
	"fmt"

	"github.com/takeuchi-shogo/go-clean-ddd/internal/api/infrastructure"
)

func main() {
	c := infrastructure.NewConfig("LOCAL")
	fmt.Printf("%+v", c)

	r := infrastructure.NewRouter(c)

	r.Run()
}
