// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"time"
)

const dateTimeFormatWithSecond = "2006/01/02 15:04:05"

func main() {
	JST := time.FixedZone("JST", 9*3600)
	ut := time.Now().Unix()
	str := time.Unix(ut, 0).In(JST).Format(dateTimeFormatWithSecond)
	fmt.Println(str)
}
