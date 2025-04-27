package main

import (
	"fmt"

	"github.com/thenaveensharma/students-api/internal/config"
)

func main() {
	//load config
	cfg := config.MustLoad()
	fmt.Println(cfg)
}
