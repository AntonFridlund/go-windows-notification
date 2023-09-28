package notify

import "syscall"

type flashInfo struct {
	cbSize    uint32         // Size of structure
	hwnd      syscall.Handle // Window handle
	dwFlags   uint32         // Type of flashing
	uCount    uint32         // Number of flashes
	dwTimeout uint32         // Duration of each flash
}
