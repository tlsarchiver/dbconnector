package dbconnector

import (
	"log"
)

func checkErr(err error) {
	if err != nil {
		log.Println(err)
		panic("Aborting.")
	}
}
