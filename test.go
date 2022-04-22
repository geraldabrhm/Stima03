// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"regexp"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func checkDNAInput(isiFile string) bool {
// 	var re = regexp.MustCompile(`^[AGCT]+$`)
// 	return re.MatchString(isiFile)
// }

// func main() {
// 	var fileName string
// 	fmt.Scanf("%s", &fileName)
// 	body, err := ioutil.ReadFile(fileName)
// 	if err != nil {
// 		log.Fatalf("unable to read file: %v", err)
// 	}
// 	fmt.Println(checkDNAInput(string(body)))

// 	db, err := sql.Open("mysql", "user:password@/dbname")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()