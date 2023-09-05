package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"model_generate/server"
)

var ormName string
var dns string

func init() {
	flag.StringVar(&ormName, "orm", "", "choose orm")
	flag.StringVar(&dns, "dns", "", fmt.Sprintf("datasource example: %s", "root:password@tcp(127.0.0.1:3306)/testdb?charset=utf8"))
	flag.Parse()
	if dns == "" {
		panic(fmt.Errorf("datasource is empty"))
	}

	server.Init(dns)
}

func main() {
	server.GenStruct(ormName)
}
