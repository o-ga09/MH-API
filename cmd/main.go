package main

import "mh-api/api/controller"

func main() {
	s, err := controller.NewServer()
	if err != nil {
		panic(err)
	}
	s.Run()
}