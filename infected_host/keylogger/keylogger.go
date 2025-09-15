package keylogger

import "os"

func Check_unix() int {
	if os.Getegid() != -1 {
		return 0
	}
	return 1
}
