package main

import (
	"fmt"
	"io/ioutil"
	"log"

	be "src/backend"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var fileName string
	fmt.Scanf("%s", &fileName)
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	fmt.Println(be.CheckDNAInput(string(body)))
}
