package keyboard

import (
	"keyrehack/internal/combinationReplacer"
	"keyrehack/internal/replacer"
	simplereplacer "keyrehack/internal/simpleReplacer"
	. "keyrehack/internal/windows"
	"log"
	"syscall"
	"unsafe"
)

const (
	// https://learn.microsoft.com/ja-jp/windows/win32/api/winuser/nf-winuser-setwindowshookexw
	WH_KEYBOARD_LL uintptr = 13
)

var (
	modUser32, _ = syscall.LoadDLL("user32.dll")

	procCallNextHookEx, _ = modUser32.FindProc("CallNextHookEx")
	// https://learn.microsoft.com/ja-jp/windows/win32/api/winuser/nf-winuser-setwindowshookexw
	procSetWindowsHookExW, _   = modUser32.FindProc("SetWindowsHookExW")
	procGetMessageW, _         = modUser32.FindProc("GetMessageW")
	procTranslateMessage, _    = modUser32.FindProc("TranslateMessage")
	procDispatchMessageW, _    = modUser32.FindProc("DispatchMessageW")
	procUnhookWindowsHookEx, _ = modUser32.FindProc("UnhookWindowsHookEx")

	keyboardHook       uintptr
	simpleReplace      replacer.Replacer
	combinationReplace replacer.Replacer

	isCalled = false
)

func setHook(simpleReplaces []simplereplacer.ReplaceKey, combinationReplaces []combinationReplacer.ReplaceKey) error {
	if isCalled {
		log.Fatalln("Hook is already set")
	}
	isCalled = true

	simpleReplace = simplereplacer.NewSimpleReplacer(simpleReplaces)

	combinationReplace = combinationReplacer.NewCombination(combinationReplaces)

	keyboardHook, _, _ := procSetWindowsHookExW.Call(
		WH_KEYBOARD_LL,                // The type of hook procedure to be installed.
		syscall.NewCallback(hookProc), // A pointer to the hook procedure.
		uintptr(0),                    // A handle to the DLL containing the hook procedure pointed to by the lpfn parameter. 0 means the current process.
		0,                             // Thread ID. 0 means the hook procedure is associated with all existing threads.
	)

	if keyboardHook == 0 {
		log.Fatalln("Failed to set hook.")
	}

	var msg MSG

	for {
		ret, _, _ := procGetMessageW.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0)

		if ret == 0 {
			break
		}

		procTranslateMessage.Call(uintptr(unsafe.Pointer(&msg)))
		procDispatchMessageW.Call(uintptr(unsafe.Pointer(&msg)))
	}

	return nil
}

func unsetHook() {
	procUnhookWindowsHookEx.Call(keyboardHook)
	isCalled = false
}

func hookProc(nCode int, wParam uintptr, lParam uintptr) uintptr {
	kbdStruct := (*KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam))

	if simpleReplace.Replace(kbdStruct.VkCode, Message(wParam)) {
		return 1
	}

	if combinationReplace.Replace(kbdStruct.VkCode, Message(wParam)) {
		return 1
	}

	ret, _, _ := procCallNextHookEx.Call(keyboardHook, uintptr(nCode), wParam, lParam)
	return ret
}
