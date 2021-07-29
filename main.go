package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tarantool/go-tarantool"
)

func main() {
	conn, err := Connect()
	if err != nil {
		log.Fatalf("Connection refused %s", err)
	}
	defer conn.Close()

	spaceName := "tester"
	indexName := "scanner"
	idFn := conn.Schema.Spaces[spaceName].Fields["id"].Id

	var tuplesPerRequest uint32 = 10
	cursor := []interface{}{}

	for {
		resp, err := conn.Select(spaceName, indexName, 0, tuplesPerRequest, tarantool.IterGt, cursor)
		if err != nil {
			log.Fatalf("Failed to select: %s", err)
		}

		if resp.Code != tarantool.OkCode {
			log.Fatalf("Select failed: %s", resp.Error)
		}

		if len(resp.Data) == 0 {
			fmt.Println(resp)
			break
		}

		fmt.Println("Iteration")

		tuples := resp.Tuples()
		for _, tuple := range tuples {
			fmt.Printf("\t%v\n", tuple)
		}

		lastTuple := tuples[len(tuples)-1]
		cursor = []interface{}{lastTuple[idFn]}
	}

	router := gin.Default()
	router.GET("/kv/:id", getKv)
	router.DELETE("/kv/:id", deleteKv)
	router.POST("/kv", createKv)
	router.PUT("/kv", updateKv)

	router.Run("localhost:8080")
}
