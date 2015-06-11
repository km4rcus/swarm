package utils

import (
		"strings"
		"strconv"
)

// convert a string list as "0-2,7" in a slice of strings ["0" "1" "2" "7"]
func StringListSplit(in string) []string {
	out := []string{}
	for _,v := range strings.Split(in, ",") {
		if s := strings.Split(v, "-"); len(s) == 2 {
			begin,_ := strconv.ParseInt(s[0], 10, 64);
			end,_ := strconv.ParseInt(s[1], 10, 64)
			for i := begin; i <= end; i++ {
				out = append(out, strconv.Itoa(int(i)))
			}
		} else {
			out = append(out, v)
		}
	}
	return out
}