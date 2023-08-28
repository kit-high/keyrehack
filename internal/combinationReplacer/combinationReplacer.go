package combinationReplacer

import (
	"keyrehack/internal/replacer"
	. "keyrehack/internal/windows"
	"syscall"
	"unsafe"
)

type CombinationReplacer struct {
	CombinationKeys      []ReplaceKey
	procgetAsyncKeyState *syscall.Proc
	procSendInput        *syscall.Proc
}

func NewCombination(arr []ReplaceKey) *CombinationReplacer {
	modUser32, _ := syscall.LoadDLL("user32.dll")
	procgetAsyncKeyState, _ := modUser32.FindProc("GetAsyncKeyState")
	procSendInput, _ := modUser32.FindProc("SendInput")

	return &CombinationReplacer{
		CombinationKeys:      arr,
		procgetAsyncKeyState: procgetAsyncKeyState,
		procSendInput:        procSendInput,
	}
}

func (cr *CombinationReplacer) Replace(vkCode VKCode, message Message) replacer.IsReplaced {
	for _, key := range cr.CombinationKeys {
		if key.Source.Hold == nil || key.Source.Press == nil {
			continue
		}
		if key.Target.Hold == nil {
			continue
		}

		if (message == WM_KEYDOWN || message == WM_SYSKEYDOWN) &&
			vkCode == *(*key.Source.Press)[0] &&
			(cr.getAsyncKeyStateCall(*key.Source.Hold)&0x8000) != 0 {

			input := []INPUT{}
			input = append(input, INPUT{Type: 1, Ki: KEYBDINPUT{Vk: uint16(*(*key.Source.Press)[0]), DwFlags: KEYEVENTF_KEYUP}}) // Release key
			input = append(input, INPUT{Type: 1, Ki: KEYBDINPUT{Vk: uint16(*key.Source.Hold), DwFlags: KEYEVENTF_KEYUP}})        // Release key

			input = append(input, INPUT{Type: 1, Ki: KEYBDINPUT{Vk: uint16(*key.Target.Hold), DwFlags: KEYEVENTF_KEYDOWN}}) // Hold key down

			if key.Target.Press != nil {
				for i, _ := range *key.Target.Press {
					input = append(input, INPUT{Type: 1, Ki: KEYBDINPUT{Vk: uint16(*(*key.Target.Press)[i]), DwFlags: KEYEVENTF_KEYDOWN}}) // Press key down
					input = append(input, INPUT{Type: 1, Ki: KEYBDINPUT{Vk: uint16(*(*key.Target.Press)[i]), DwFlags: KEYEVENTF_KEYUP}})   // Press key up
				}
			}

			input = append(input, INPUT{Type: 1, Ki: KEYBDINPUT{Vk: uint16(*key.Target.Hold), DwFlags: KEYEVENTF_KEYUP}})   // Hold key up
			input = append(input, INPUT{Type: 1, Ki: KEYBDINPUT{Vk: uint16(*key.Source.Hold), DwFlags: KEYEVENTF_KEYDOWN}}) // Restore

			cr.procSendInput.Call(uintptr(uint32(len(input))), uintptr(unsafe.Pointer(&input[0])), uintptr(int(unsafe.Sizeof(input[0]))))

			return replacer.IsReplaced(true)
		}
	}
	return replacer.IsReplaced(false)
}

func (c *CombinationReplacer) getAsyncKeyStateCall(vKey VKCode) uint32 {
	ret, _, _ := c.procgetAsyncKeyState.Call(uintptr(vKey))
	return uint32(ret)
}

var _ replacer.Replacer = (*CombinationReplacer)(nil)
