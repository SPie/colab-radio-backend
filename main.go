package main

import (
    "github.com/spie/colab-radio-backend/router"
)

func main() {
    router := router.SetUp()
    
    err := router.Run()
    if err != nil {
	panic(err)
    }
}
