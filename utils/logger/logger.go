package logger

import "log"

func INFO(message string, data interface{}) {
	log.Printf(message, data)
}

func PANIC(message string, err error) {
	if err != nil {
		log.Panic(message, err)
	}
}
