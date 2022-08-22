package start

import (
	"fmt"

	"github.com/uwindsorcss/site/pkg/config"
)

func Run() {
	_, err := config.Load("config.hcl")
	if err != nil {
		fmt.Printf("start.Run: %s\n", err)
		return
	}
	fmt.Printf("Hi\n")
}
