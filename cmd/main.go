package main

import "github.com/MH-API/mh-api/api"

func main() {
	s := api.NewServer()
	s.Run()
}