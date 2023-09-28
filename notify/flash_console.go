package notify

import (
	"syscall"
	"unsafe"
)

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

// Checks for error before getting the current window handle
// Does error checking based on primary return value as per instructions
func getWindowHandle() (uintptr, error) {
	if err := getConsoleWindow.Find(); err != nil {
		return 0, err
	}
	handle, _, err := getConsoleWindow.Call()
	if handle == 0 {
		return 0, err
	}
	return handle, nil
}

// Checks for error before flashing the console window
// Does error checking based on primary return value as per instructions
func flashConsoleWindow(info flashInfo) error {
	if err := flashWindowEx.Find(); err != nil {
		return err
	}
	r1, _, err := flashWindowEx.Call(uintptr(unsafe.Pointer(&info)))
	if r1 == 0 {
		return err
	}
	return nil
}
