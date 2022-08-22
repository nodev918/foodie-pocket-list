package utils

import (
	"fmt"
	"log"
)

type H = func()

func Typeof(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case H:
		return "func()"
	default:
		return "unknown"
	}

}

func Log(target interface{}) {

	if Typeof(target) == "string" {
		fmt.Printf("%s(%T)\n", target, target)
	} else {
		fmt.Printf("type is: %T\n", target)
	}
}

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}
