package notify

import "syscall"

type flashInfo struct {
	cbSize    uint32         // Size of structure
	hwnd      syscall.Handle // Window handle
	dwFlags   uint32         // Type of flashing
	uCount    uint32         // Number of flashes
	dwTimeout uint32         // Duration of each flash
}

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	flashWindowEx    = user32.NewProc("FlashWindowEx")
	kernel32         = syscall.NewLazyDLL("kernel32.dll")
	getConsoleWindow = kernel32.NewProc("GetConsoleWindow")
)
