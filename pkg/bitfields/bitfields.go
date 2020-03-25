package bitfields

import "fmt"

const (
	INFO   = 0x01 // 00000001
	NOTICE = 0x02 // 00000010
	WARN   = 0x03 // 00000011
	ERROR  = 0x04 // 00000100
	DEBUG  = 0x05 // 00000101
	MASK   = 0x07 // 00000111
)

func Log(msg string, options uint8) {

	const (
		InfoColor    = "\033[1;34m%s\033[0m"
		NoticeColor  = "\033[1;36m%s\033[0m"
		WarningColor = "\033[1;33m%s\033[0m"
		ErrorColor   = "\033[1;31m%s\033[0m"
		DebugColor   = "\033[0;36m%s\033[0m"
	)

	switch options & MASK {
	case INFO:
		fmt.Printf(InfoColor, msg)
	case NOTICE:
		fmt.Printf(NoticeColor, msg)
	case WARN:
		fmt.Printf(WarningColor, msg)
	case ERROR:
		fmt.Printf(ErrorColor, msg)
	case DEBUG:
		fmt.Printf(DebugColor, msg)
	}
	fmt.Println("")
}
