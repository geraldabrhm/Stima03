package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

func checkDNAInput(isiFile string) bool {
	var re = regexp.MustCompile(`^[AGCT]+$`)
	return re.MatchString(isiFile)
}

func main() {
	var fileName string
	fmt.Scanf("%s", &fileName)
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	fmt.Println(checkDNAInput(string(body)))
}
