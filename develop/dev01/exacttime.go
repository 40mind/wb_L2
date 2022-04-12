package exacttime

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func ShowTime() {
	exactTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(exactTime)
}
