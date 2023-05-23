package main

import "mh-api/api/handler"

func main() {
	s := handler.NewServer()
	s.Run()
}