package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/marcomilon/ezphp/engine/php"
)

func StartTerminal(ioCom php.IOCom) {
Terminal:
	for {
		select {
		case outmsg := <-ioCom.Outmsg:
			fmt.Print(outmsg)
		case errMsg := <-ioCom.Errmsg:
			fmt.Print(errMsg)
		case confirmMsg := <-ioCom.Confirm:
			Confirm(confirmMsg, ioCom.Confirm)
		case <-ioCom.Done:
			byebye()
			break Terminal
		}
	}

	os.Exit(0)
}

func byebye() {
	fmt.Print("\nPress 'Enter' to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func Confirm(question string, result chan string) {

	var confirmation string

	fmt.Printf("%s [y/N] ", question)
	fmt.Scanln(&confirmation)

	confirmation = strings.TrimSpace(confirmation)
	confirmation = strings.ToLower(confirmation)

	if confirmation == "y" {
		result <- "Yes"
		return
	}

	result <- "No"
}
