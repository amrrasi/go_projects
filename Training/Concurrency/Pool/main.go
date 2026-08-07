package main

import (
	"fmt"
	"sync"
)

type DBConnection struct {
	Host     string
	DbName   string
	User     string
	Password string
}

var connectionPool sync.Pool = sync.Pool{

	New: func() interface{} {
		return &DBConnection{
			Host:     "localhost",
			DbName:   "Test",
			User:     "Amir",
			Password: "Root",
		}
	},
}

func main() {

	connection := connectionPool.Get().(*DBConnection)
	fmt.Printf("%v\n", connection)

	connectionPool.Put(connection)

}
