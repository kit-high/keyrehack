package windows

// https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-kbdllhookstruct
type KBDLLHOOKSTRUCT struct {
	VkCode      VKCode  // Virtual key code
	ScanCode    uint32  // Hardware scan code
	Flags       uint32  // Flags related to the key press
	Time        uint32  // Timestamp for this message
	DwExtraInfo uintptr // Extra info
}

// https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-msg
type MSG struct {
	Hwnd    uintptr
	Message uint32
	WParam  uint32
	LParam  uint32
	Time    uint32
	POINT
	// lPrivate uint32 this is in reference
}

// https://learn.microsoft.com/en-us/windows/win32/api/windef/ns-windef-point
type POINT struct {
	X int32
	Y int32
}

type INPUT struct {
	Type uint32     // Type of the input event
	Ki   KEYBDINPUT // Keyboard input
	Mi   struct{}   // Mouse input (not used here)
	Hi   struct{}   // Hardware input (not used here)
}

type KEYBDINPUT struct {
	Vk          uint16  // Virtual key code
	Scan        uint16  // Hardware scan code
	DwFlags     DwFlags // Key event flags
	Time        uint32  // Timestamp for this message
	DwExtraInfo uintptr // Extra info
}

// https://learn.microsoft.com/ja-jp/windows/win32/api/winuser/ns-winuser-keybdinput
type DwFlags uint32
