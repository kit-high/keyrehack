package simplereplacer

import (
	"keyrehack/internal/replacer"
	. "keyrehack/internal/windows"
	"syscall"
	"unsafe"
)

type SimpleReplacer struct {
	ReplaceKeys   []ReplaceKey
	procSendInput *syscall.Proc
}

func NewSimpleReplacer(arr []ReplaceKey) *SimpleReplacer {
	modUser32, _ := syscall.LoadDLL("user32.dll")
	procSendInput, _ := modUser32.FindProc("SendInput")

	return &SimpleReplacer{
		ReplaceKeys:   arr,
		procSendInput: procSendInput,
	}
}

func (r *SimpleReplacer) Replace(vkCode VKCode, message Message) replacer.IsReplaced {
	for _, key := range r.ReplaceKeys {
		if key.Source == nil || key.Target == nil {
			continue
		}

		if vkCode == *key.Source {

			getDwFlags := func(message Message) DwFlags {
				switch message {
				case WM_KEYDOWN:
					return KEYEVENTF_KEYDOWN
				case WM_KEYUP:
					return KEYEVENTF_KEYUP
				default:
					panic("invalid message")
				}
			}

			input := []INPUT{
				{Type: 1, Ki: KEYBDINPUT{Vk: uint16(*key.Target), DwFlags: getDwFlags(message)}},
			}

			r.procSendInput.Call(uintptr(uint32(len(input))), uintptr(unsafe.Pointer(&input[0])), uintptr(int(unsafe.Sizeof(input[0]))))

			return replacer.IsReplaced(true)
		}
	}

	return replacer.IsReplaced(false)
}

var _ replacer.Replacer = (*SimpleReplacer)(nil)
