package main

import "mh-api/api/middleware"

func main() {
	// s, err := controller.NewServer()
	// if err != nil {
	// 	panic(err)
	// }
	// s.Run()

	middleware.New()
}