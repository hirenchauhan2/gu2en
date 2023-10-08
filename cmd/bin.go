package main

import (
	"fmt"
	"os"

	"github.com/hirenchauhan2/gu2en"
)

func main() {
	arg_count := len(os.Args)
	if arg_count < 2 {
		fmt.Println("Missing the argument for gujarati text")
		fmt.Println(`
Demo:
$> ./bin "આ ગુજરાતી ભાષા જેવી સરળતા બીજે ક્યાક જો જોવા મળે તો તમે નવી શોધ કરી છે, એમ માનજો."`)
		os.Exit(1)
	}

	text := os.Args[1]
	fmt.Println(gu2en.Transliterate(text))
}
