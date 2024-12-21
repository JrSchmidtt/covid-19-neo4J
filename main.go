package main

import (
	server "github.com/JrSchmidtt/covid-19-neo4J/src"
)

func main() {
	err := server.Run(); if err != nil {
		panic("Server failed to start" + err.Error())
	}
}