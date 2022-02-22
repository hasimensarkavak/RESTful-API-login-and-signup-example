package helpers

import (
	"log"
)

func IsEmpty(data string) bool {
	if len(data) == 0 {
		return true
	} else {
		return false
	}
}

func ErrorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
