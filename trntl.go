package main

import (
	"fmt"
	"log"

	"github.com/tarantool/go-tarantool"
)

type TarantoolModel struct {
	conn  *tarantool.Connection
	space string
	index string
	id    string
}

func (m TarantoolModel) Select(c interface{}) (*tarantool.Response, error) {
	resp, err := m.conn.Select(m.space, m.index, 0, 1, tarantool.IterGt, c)
	if err != nil {
		log.Fatalf("Failed to select: %s", err)
	}

	if resp.Code != tarantool.OkCode {
		return resp, fmt.Errorf("Select failed: %s", resp.Error)
	}

	return resp, nil
}

var KeyValueModel TarantoolModel = TarantoolModel{
	space: "tester",
	index: "scanner",
	id:    "id",
}

func Connect() (*tarantool.Connection, error) {
	conn, err := tarantool.Connect("127.0.0.1:3301", tarantool.Opts{
		User: "admin",
		Pass: "pass",
	})

	if err != nil {
		return nil, err
	}

	return conn, nil
}
