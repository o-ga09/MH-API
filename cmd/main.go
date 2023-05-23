package main

import "mh-api/api"

func main() {
	s := api.NewServer()
	s.Run()
}