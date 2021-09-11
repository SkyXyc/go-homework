package main

import (
	"calderxu_workshop1_msg_notification/logic"
	"log"
	"os"
)

func main() {
	// 在最外层只会接收到nil或是wrap后的err，在log中记录详细堆栈
	if err := logic.MockLogic(); err != nil {
		log.Printf("FATAL: %+v", err)
		os.Exit(1)
	}
}
