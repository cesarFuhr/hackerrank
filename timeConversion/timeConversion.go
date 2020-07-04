package timeconversion

import (
	"fmt"
	"strconv"
	"strings"
)

func TimeConversion(s string) string {
	splited := strings.Split(s, ":")
	rawhh, _ := strconv.Atoi(splited[0])
	t := splited[2][2:]
	if t == "PM" {
		rawhh += 12
		if rawhh == 24 {
			rawhh = 12
		}
	} else {
		if rawhh == 12 {
			rawhh = 0
		}
	}
	hh := fmt.Sprintf("%02d", rawhh)
	return hh + ":" + splited[1] + ":" + strings.TrimSuffix(splited[2], t)
}
