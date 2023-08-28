package windows

type VKCode uint32

func (code VKCode) GetPointer() *VKCode {
	return &code
}

func VK_LBUTTON() *VKCode             { return VKCode(0x01).GetPointer() }
func VK_RBUTTON() *VKCode             { return VKCode(0x02).GetPointer() }
func VK_CANCEL() *VKCode              { return VKCode(0x03).GetPointer() }
func VK_MBUTTON() *VKCode             { return VKCode(0x04).GetPointer() }
func VK_XBUTTON1() *VKCode            { return VKCode(0x05).GetPointer() }
func VK_XBUTTON2() *VKCode            { return VKCode(0x06).GetPointer() }
func VK_BACK() *VKCode                { return VKCode(0x08).GetPointer() }
func VK_TAB() *VKCode                 { return VKCode(0x09).GetPointer() }
func VK_CLEAR() *VKCode               { return VKCode(0x0C).GetPointer() }
func VK_RETURN() *VKCode              { return VKCode(0x0D).GetPointer() }
func VK_SHIFT() *VKCode               { return VKCode(0x10).GetPointer() }
func VK_CONTROL() *VKCode             { return VKCode(0x11).GetPointer() }
func VK_MENU() *VKCode                { return VKCode(0x12).GetPointer() }
func VK_PAUSE() *VKCode               { return VKCode(0x13).GetPointer() }
func VK_CAPITAL() *VKCode             { return VKCode(0x14).GetPointer() }
func VK_KANA() *VKCode                { return VKCode(0x15).GetPointer() }
func VK_HANGEUL() *VKCode             { return VKCode(0x15).GetPointer() }
func VK_HANGUL() *VKCode              { return VKCode(0x15).GetPointer() }
func VK_IME() *VKCode                 { return VKCode(0x16).GetPointer() }
func VK_JUNJA() *VKCode               { return VKCode(0x17).GetPointer() }
func VK_FINAL() *VKCode               { return VKCode(0x18).GetPointer() }
func VK_HANJA() *VKCode               { return VKCode(0x19).GetPointer() }
func VK_KANJI() *VKCode               { return VKCode(0x19).GetPointer() }
func VK_IME_OFF() *VKCode             { return VKCode(0x1A).GetPointer() }
func VK_ESCAPE() *VKCode              { return VKCode(0x1B).GetPointer() }
func VK_CONVERT() *VKCode             { return VKCode(0x1C).GetPointer() }
func VK_NONCONVERT() *VKCode          { return VKCode(0x1D).GetPointer() }
func VK_ACCEPT() *VKCode              { return VKCode(0x1E).GetPointer() }
func VK_MODECHANGE() *VKCode          { return VKCode(0x1F).GetPointer() }
func VK_SPACE() *VKCode               { return VKCode(0x20).GetPointer() }
func VK_PRIOR() *VKCode               { return VKCode(0x21).GetPointer() }
func VK_NEXT() *VKCode                { return VKCode(0x22).GetPointer() }
func VK_END() *VKCode                 { return VKCode(0x23).GetPointer() }
func VK_HOME() *VKCode                { return VKCode(0x24).GetPointer() }
func VK_LEFT() *VKCode                { return VKCode(0x25).GetPointer() }
func VK_UP() *VKCode                  { return VKCode(0x26).GetPointer() }
func VK_RIGHT() *VKCode               { return VKCode(0x27).GetPointer() }
func VK_DOWN() *VKCode                { return VKCode(0x28).GetPointer() }
func VK_SELECT() *VKCode              { return VKCode(0x29).GetPointer() }
func VK_PRINT() *VKCode               { return VKCode(0x2A).GetPointer() }
func VK_EXECUTE() *VKCode             { return VKCode(0x2B).GetPointer() }
func VK_SNAPSHOT() *VKCode            { return VKCode(0x2C).GetPointer() }
func VK_INSERT() *VKCode              { return VKCode(0x2D).GetPointer() }
func VK_DELETE() *VKCode              { return VKCode(0x2E).GetPointer() }
func VK_HELP() *VKCode                { return VKCode(0x2F).GetPointer() }
func VK_0() *VKCode                   { return VKCode(0x30).GetPointer() }
func VK_1() *VKCode                   { return VKCode(0x31).GetPointer() }
func VK_2() *VKCode                   { return VKCode(0x32).GetPointer() }
func VK_3() *VKCode                   { return VKCode(0x33).GetPointer() }
func VK_4() *VKCode                   { return VKCode(0x34).GetPointer() }
func VK_5() *VKCode                   { return VKCode(0x35).GetPointer() }
func VK_6() *VKCode                   { return VKCode(0x36).GetPointer() }
func VK_7() *VKCode                   { return VKCode(0x37).GetPointer() }
func VK_8() *VKCode                   { return VKCode(0x38).GetPointer() }
func VK_9() *VKCode                   { return VKCode(0x39).GetPointer() }
func VK_A() *VKCode                   { return VKCode(0x41).GetPointer() }
func VK_B() *VKCode                   { return VKCode(0x42).GetPointer() }
func VK_C() *VKCode                   { return VKCode(0x43).GetPointer() }
func VK_D() *VKCode                   { return VKCode(0x44).GetPointer() }
func VK_E() *VKCode                   { return VKCode(0x45).GetPointer() }
func VK_F() *VKCode                   { return VKCode(0x46).GetPointer() }
func VK_G() *VKCode                   { return VKCode(0x47).GetPointer() }
func VK_H() *VKCode                   { return VKCode(0x48).GetPointer() }
func VK_I() *VKCode                   { return VKCode(0x49).GetPointer() }
func VK_J() *VKCode                   { return VKCode(0x4A).GetPointer() }
func VK_K() *VKCode                   { return VKCode(0x4B).GetPointer() }
func VK_L() *VKCode                   { return VKCode(0x4C).GetPointer() }
func VK_M() *VKCode                   { return VKCode(0x4D).GetPointer() }
func VK_N() *VKCode                   { return VKCode(0x4E).GetPointer() }
func VK_O() *VKCode                   { return VKCode(0x4F).GetPointer() }
func VK_P() *VKCode                   { return VKCode(0x50).GetPointer() }
func VK_Q() *VKCode                   { return VKCode(0x51).GetPointer() }
func VK_R() *VKCode                   { return VKCode(0x52).GetPointer() }
func VK_S() *VKCode                   { return VKCode(0x53).GetPointer() }
func VK_T() *VKCode                   { return VKCode(0x54).GetPointer() }
func VK_U() *VKCode                   { return VKCode(0x55).GetPointer() }
func VK_V() *VKCode                   { return VKCode(0x56).GetPointer() }
func VK_W() *VKCode                   { return VKCode(0x57).GetPointer() }
func VK_X() *VKCode                   { return VKCode(0x58).GetPointer() }
func VK_Y() *VKCode                   { return VKCode(0x59).GetPointer() }
func VK_Z() *VKCode                   { return VKCode(0x5A).GetPointer() }
func VK_LWIN() *VKCode                { return VKCode(0x5B).GetPointer() }
func VK_RWIN() *VKCode                { return VKCode(0x5C).GetPointer() }
func VK_APPS() *VKCode                { return VKCode(0x5D).GetPointer() }
func VK_SLEEP() *VKCode               { return VKCode(0x5F).GetPointer() }
func VK_NUMPAD0() *VKCode             { return VKCode(0x60).GetPointer() }
func VK_NUMPAD1() *VKCode             { return VKCode(0x61).GetPointer() }
func VK_NUMPAD2() *VKCode             { return VKCode(0x62).GetPointer() }
func VK_NUMPAD3() *VKCode             { return VKCode(0x63).GetPointer() }
func VK_NUMPAD4() *VKCode             { return VKCode(0x64).GetPointer() }
func VK_NUMPAD5() *VKCode             { return VKCode(0x65).GetPointer() }
func VK_NUMPAD6() *VKCode             { return VKCode(0x66).GetPointer() }
func VK_NUMPAD7() *VKCode             { return VKCode(0x67).GetPointer() }
func VK_NUMPAD8() *VKCode             { return VKCode(0x68).GetPointer() }
func VK_NUMPAD9() *VKCode             { return VKCode(0x69).GetPointer() }
func VK_MULTIPLY() *VKCode            { return VKCode(0x6A).GetPointer() }
func VK_ADD() *VKCode                 { return VKCode(0x6B).GetPointer() }
func VK_SEPARATOR() *VKCode           { return VKCode(0x6C).GetPointer() }
func VK_SUBTRACT() *VKCode            { return VKCode(0x6D).GetPointer() }
func VK_DECIMAL() *VKCode             { return VKCode(0x6E).GetPointer() }
func VK_DIVIDE() *VKCode              { return VKCode(0x6F).GetPointer() }
func VK_F1() *VKCode                  { return VKCode(0x70).GetPointer() }
func VK_F2() *VKCode                  { return VKCode(0x71).GetPointer() }
func VK_F3() *VKCode                  { return VKCode(0x72).GetPointer() }
func VK_F4() *VKCode                  { return VKCode(0x73).GetPointer() }
func VK_F5() *VKCode                  { return VKCode(0x74).GetPointer() }
func VK_F6() *VKCode                  { return VKCode(0x75).GetPointer() }
func VK_F7() *VKCode                  { return VKCode(0x76).GetPointer() }
func VK_F8() *VKCode                  { return VKCode(0x77).GetPointer() }
func VK_F9() *VKCode                  { return VKCode(0x78).GetPointer() }
func VK_F10() *VKCode                 { return VKCode(0x79).GetPointer() }
func VK_F11() *VKCode                 { return VKCode(0x7A).GetPointer() }
func VK_F12() *VKCode                 { return VKCode(0x7B).GetPointer() }
func VK_F13() *VKCode                 { return VKCode(0x7C).GetPointer() }
func VK_F14() *VKCode                 { return VKCode(0x7D).GetPointer() }
func VK_F15() *VKCode                 { return VKCode(0x7E).GetPointer() }
func VK_F16() *VKCode                 { return VKCode(0x7F).GetPointer() }
func VK_F17() *VKCode                 { return VKCode(0x80).GetPointer() }
func VK_F18() *VKCode                 { return VKCode(0x81).GetPointer() }
func VK_F19() *VKCode                 { return VKCode(0x82).GetPointer() }
func VK_F20() *VKCode                 { return VKCode(0x83).GetPointer() }
func VK_F21() *VKCode                 { return VKCode(0x84).GetPointer() }
func VK_F22() *VKCode                 { return VKCode(0x85).GetPointer() }
func VK_F23() *VKCode                 { return VKCode(0x86).GetPointer() }
func VK_F24() *VKCode                 { return VKCode(0x87).GetPointer() }
func VK_NUMLOCK() *VKCode             { return VKCode(0x90).GetPointer() }
func VK_SCROLL() *VKCode              { return VKCode(0x91).GetPointer() }
func VK_OEM_NEC_EQUAL() *VKCode       { return VKCode(0x92).GetPointer() }
func VK_OEM_FJ_JISHO() *VKCode        { return VKCode(0x92).GetPointer() }
func VK_OEM_FJ_MASSHOU() *VKCode      { return VKCode(0x93).GetPointer() }
func VK_OEM_FJ_TOUROKU() *VKCode      { return VKCode(0x94).GetPointer() }
func VK_OEM_FJ_LOYA() *VKCode         { return VKCode(0x95).GetPointer() }
func VK_OEM_FJ_ROYA() *VKCode         { return VKCode(0x96).GetPointer() }
func VK_LSHIFT() *VKCode              { return VKCode(0xA0).GetPointer() }
func VK_RSHIFT() *VKCode              { return VKCode(0xA1).GetPointer() }
func VK_LCONTROL() *VKCode            { return VKCode(0xA2).GetPointer() }
func VK_RCONTROL() *VKCode            { return VKCode(0xA3).GetPointer() }
func VK_LMENU() *VKCode               { return VKCode(0xA4).GetPointer() }
func VK_RMENU() *VKCode               { return VKCode(0xA5).GetPointer() }
func VK_BROWSER_BACK() *VKCode        { return VKCode(0xA6).GetPointer() }
func VK_BROWSER_FORWARD() *VKCode     { return VKCode(0xA7).GetPointer() }
func VK_BROWSER_REFRESH() *VKCode     { return VKCode(0xA8).GetPointer() }
func VK_BROWSER_STOP() *VKCode        { return VKCode(0xA9).GetPointer() }
func VK_BROWSER_SEARCH() *VKCode      { return VKCode(0xAA).GetPointer() }
func VK_BROWSER_FAVORITES() *VKCode   { return VKCode(0xAB).GetPointer() }
func VK_BROWSER_HOME() *VKCode        { return VKCode(0xAC).GetPointer() }
func VK_VOLUME_MUTE() *VKCode         { return VKCode(0xAD).GetPointer() }
func VK_VOLUME_DOWN() *VKCode         { return VKCode(0xAE).GetPointer() }
func VK_VOLUME_UP() *VKCode           { return VKCode(0xAF).GetPointer() }
func VK_MEDIA_NEXT_TRACK() *VKCode    { return VKCode(0xB0).GetPointer() }
func VK_MEDIA_PREV_TRACK() *VKCode    { return VKCode(0xB1).GetPointer() }
func VK_MEDIA_STOP() *VKCode          { return VKCode(0xB2).GetPointer() }
func VK_MEDIA_PLAY_PAUSE() *VKCode    { return VKCode(0xB3).GetPointer() }
func VK_LAUNCH_MAIL() *VKCode         { return VKCode(0xB4).GetPointer() }
func VK_LAUNCH_MEDIA_SELECT() *VKCode { return VKCode(0xB5).GetPointer() }
func VK_LAUNCH_APP1() *VKCode         { return VKCode(0xB6).GetPointer() }
func VK_LAUNCH_APP2() *VKCode         { return VKCode(0xB7).GetPointer() }
func VK_OEM_1() *VKCode               { return VKCode(0xBA).GetPointer() }
func VK_OEM_PLUS() *VKCode            { return VKCode(0xBB).GetPointer() }
func VK_OEM_COMMA() *VKCode           { return VKCode(0xBC).GetPointer() }
func VK_OEM_MINUS() *VKCode           { return VKCode(0xBD).GetPointer() }
func VK_OEM_PERIOD() *VKCode          { return VKCode(0xBE).GetPointer() }
func VK_OEM_2() *VKCode               { return VKCode(0xBF).GetPointer() }
func VK_OEM_3() *VKCode               { return VKCode(0xC0).GetPointer() }
func VK_OEM_4() *VKCode               { return VKCode(0xDB).GetPointer() }
func VK_OEM_5() *VKCode               { return VKCode(0xDC).GetPointer() }
func VK_OEM_6() *VKCode               { return VKCode(0xDD).GetPointer() }
func VK_OEM_7() *VKCode               { return VKCode(0xDE).GetPointer() }
func VK_OEM_8() *VKCode               { return VKCode(0xDF).GetPointer() }
func VK_OEM_AX() *VKCode              { return VKCode(0xE1).GetPointer() }
func VK_OEM_102() *VKCode             { return VKCode(0xE2).GetPointer() }
func VK_ICO_HELP() *VKCode            { return VKCode(0xE3).GetPointer() }
func VK_ICO_00() *VKCode              { return VKCode(0xE4).GetPointer() }
func VK_PROCESSKEY() *VKCode          { return VKCode(0xE5).GetPointer() }
func VK_ICO_CLEAR() *VKCode           { return VKCode(0xE6).GetPointer() }
func VK_OEM_RESET() *VKCode           { return VKCode(0xE9).GetPointer() }
func VK_OEM_JUMP() *VKCode            { return VKCode(0xEA).GetPointer() }
func VK_OEM_PA1() *VKCode             { return VKCode(0xEB).GetPointer() }
func VK_OEM_PA2() *VKCode             { return VKCode(0xEC).GetPointer() }
func VK_OEM_PA3() *VKCode             { return VKCode(0xED).GetPointer() }
func VK_OEM_WSCTRL() *VKCode          { return VKCode(0xEE).GetPointer() }
func VK_OEM_CUSEL() *VKCode           { return VKCode(0xEF).GetPointer() }
func VK_OEM_ATTN() *VKCode            { return VKCode(0xF0).GetPointer() }
func VK_OEM_FINISH() *VKCode          { return VKCode(0xF1).GetPointer() }
func VK_OEM_COPY() *VKCode            { return VKCode(0xF2).GetPointer() }
func VK_OEM_AUTO() *VKCode            { return VKCode(0xF3).GetPointer() }
func VK_OEM_ENLW() *VKCode            { return VKCode(0xF4).GetPointer() }
func VK_OEM_BACKTAB() *VKCode         { return VKCode(0xF5).GetPointer() }
func VK_ATTN() *VKCode                { return VKCode(0xF6).GetPointer() }
func VK_CRSEL() *VKCode               { return VKCode(0xF7).GetPointer() }
func VK_EXSEL() *VKCode               { return VKCode(0xF8).GetPointer() }
func VK_EREOF() *VKCode               { return VKCode(0xF9).GetPointer() }
func VK_PLAY() *VKCode                { return VKCode(0xFA).GetPointer() }
func VK_ZOOM() *VKCode                { return VKCode(0xFB).GetPointer() }
func VK_NONAME() *VKCode              { return VKCode(0xFC).GetPointer() }
func VK_PA1() *VKCode                 { return VKCode(0xFD).GetPointer() }
func VK_OEM_CLEAR() *VKCode           { return VKCode(0xFE).GetPointer() }

type Message uintptr

const (
	WM_LBUTTONDOWN Message = 0x0201
	WM_LBUTTONUP   Message = 0x0202
	WM_MOUSEMOVE   Message = 0x0200
	WM_MOUSEWHEEL  Message = 0x020A
	WM_MOUSEHWHEEL Message = 0x020E
	WM_RBUTTONDOWN Message = 0x0204
	WM_RBUTTONUP   Message = 0x0205
	WM_KEYDOWN     Message = 0x0100
	WM_KEYUP       Message = 0x0101
	WM_SYSKEYDOWN  Message = 0x0104
	WM_SYSKEYUP    Message = 0x0105
)

// const (
// 	VK_LBUTTON    VKCode = 0x01
// 	VK_RBUTTON    VKCode = 0x02
// 	VK_CANCEL     VKCode = 0x03
// 	VK_MBUTTON    VKCode = 0x04
// 	VK_XBUTTON1   VKCode = 0x05
// 	VK_XBUTTON2   VKCode = 0x06
// 	VK_BACK       VKCode = 0x08
// 	VK_TAB        VKCode = 0x09
// 	VK_CLEAR      VKCode = 0x0C
// 	VK_RETURN     VKCode = 0x0D
// 	VK_SHIFT      VKCode = 0x10
// 	VK_CONTROL    VKCode = 0x11
// 	VK_MENU       VKCode = 0x12
// 	VK_PAUSE      VKCode = 0x13
// 	VK_CAPITAL    VKCode = 0x14
// 	VK_KANA       VKCode = 0x15
// 	VK_HANGEUL    VKCode = 0x15
// 	VK_HANGUL     VKCode = 0x15
// 	VK_IME        VKCode = 0x16
// 	VK_JUNJA      VKCode = 0x17
// 	VK_FINAL      VKCode = 0x18
// 	VK_HANJA      VKCode = 0x19
// 	VK_KANJI      VKCode = 0x19
// 	VK_IME_OFF    VKCode = 0x1A
// 	VK_ESCAPE     VKCode = 0x1B
// 	VK_CONVERT    VKCode = 0x1C
// 	VK_NONCONVERT VKCode = 0x1D
// 	VK_ACCEPT     VKCode = 0x1E
// 	VK_MODECHANGE VKCode = 0x1F
// 	VK_SPACE      VKCode = 0x20
// 	VK_PRIOR      VKCode = 0x21
// 	VK_NEXT       VKCode = 0x22
// 	VK_END        VKCode = 0x23
// 	VK_HOME       VKCode = 0x24
// 	VK_LEFT       VKCode = 0x25
// 	VK_UP         VKCode = 0x26
// 	VK_RIGHT      VKCode = 0x27
// 	VK_DOWN       VKCode = 0x28
// 	VK_SELECT     VKCode = 0x29
// 	VK_PRINT      VKCode = 0x2A
// 	VK_EXECUTE    VKCode = 0x2B
// 	VK_SNAPSHOT   VKCode = 0x2C
// 	VK_INSERT     VKCode = 0x2D
// 	VK_DELETE     VKCode = 0x2E
// 	VK_HELP       VKCode = 0x2F

// 	VK_0 VKCode = 0x30
// 	VK_1 VKCode = 0x31
// 	VK_2 VKCode = 0x32
// 	VK_3 VKCode = 0x33
// 	VK_4 VKCode = 0x34
// 	VK_5 VKCode = 0x35
// 	VK_6 VKCode = 0x36
// 	VK_7 VKCode = 0x37
// 	VK_8 VKCode = 0x38
// 	VK_9 VKCode = 0x39
// 	VK_A VKCode = 0x41
// 	VK_B VKCode = 0x42
// 	VK_C VKCode = 0x43
// 	VK_D VKCode = 0x44
// 	VK_E VKCode = 0x45
// 	VK_F VKCode = 0x46
// 	VK_G VKCode = 0x47
// 	VK_H VKCode = 0x48
// 	VK_I VKCode = 0x49
// 	VK_J VKCode = 0x4A
// 	VK_K VKCode = 0x4B
// 	VK_L VKCode = 0x4C
// 	VK_M VKCode = 0x4D
// 	VK_N VKCode = 0x4E
// 	VK_O VKCode = 0x4F
// 	VK_P VKCode = 0x50
// 	VK_Q VKCode = 0x51
// 	VK_R VKCode = 0x52
// 	VK_S VKCode = 0x53
// 	VK_T VKCode = 0x54
// 	VK_U VKCode = 0x55
// 	VK_V VKCode = 0x56
// 	VK_W VKCode = 0x57
// 	VK_X VKCode = 0x58
// 	VK_Y VKCode = 0x59
// 	VK_Z VKCode = 0x5A

// 	VK_LWIN                VKCode = 0x5B
// 	VK_RWIN                VKCode = 0x5C
// 	VK_APPS                VKCode = 0x5D
// 	VK_SLEEP               VKCode = 0x5F
// 	VK_NUMPAD0             VKCode = 0x60
// 	VK_NUMPAD1             VKCode = 0x61
// 	VK_NUMPAD2             VKCode = 0x62
// 	VK_NUMPAD3             VKCode = 0x63
// 	VK_NUMPAD4             VKCode = 0x64
// 	VK_NUMPAD5             VKCode = 0x65
// 	VK_NUMPAD6             VKCode = 0x66
// 	VK_NUMPAD7             VKCode = 0x67
// 	VK_NUMPAD8             VKCode = 0x68
// 	VK_NUMPAD9             VKCode = 0x69
// 	VK_MULTIPLY            VKCode = 0x6A
// 	VK_ADD                 VKCode = 0x6B
// 	VK_SEPARATOR           VKCode = 0x6C
// 	VK_SUBTRACT            VKCode = 0x6D
// 	VK_DECIMAL             VKCode = 0x6E
// 	VK_DIVIDE              VKCode = 0x6F
// 	VK_F1                  VKCode = 0x70
// 	VK_F2                  VKCode = 0x71
// 	VK_F3                  VKCode = 0x72
// 	VK_F4                  VKCode = 0x73
// 	VK_F5                  VKCode = 0x74
// 	VK_F6                  VKCode = 0x75
// 	VK_F7                  VKCode = 0x76
// 	VK_F8                  VKCode = 0x77
// 	VK_F9                  VKCode = 0x78
// 	VK_F10                 VKCode = 0x79
// 	VK_F11                 VKCode = 0x7A
// 	VK_F12                 VKCode = 0x7B
// 	VK_F13                 VKCode = 0x7C
// 	VK_F14                 VKCode = 0x7D
// 	VK_F15                 VKCode = 0x7E
// 	VK_F16                 VKCode = 0x7F
// 	VK_F17                 VKCode = 0x80
// 	VK_F18                 VKCode = 0x81
// 	VK_F19                 VKCode = 0x82
// 	VK_F20                 VKCode = 0x83
// 	VK_F21                 VKCode = 0x84
// 	VK_F22                 VKCode = 0x85
// 	VK_F23                 VKCode = 0x86
// 	VK_F24                 VKCode = 0x87
// 	VK_NUMLOCK             VKCode = 0x90
// 	VK_SCROLL              VKCode = 0x91
// 	VK_OEM_NEC_EQUAL       VKCode = 0x92
// 	VK_OEM_FJ_JISHO        VKCode = 0x92
// 	VK_OEM_FJ_MASSHOU      VKCode = 0x93
// 	VK_OEM_FJ_TOUROKU      VKCode = 0x94
// 	VK_OEM_FJ_LOYA         VKCode = 0x95
// 	VK_OEM_FJ_ROYA         VKCode = 0x96
// 	VK_LSHIFT              VKCode = 0xA0
// 	VK_RSHIFT              VKCode = 0xA1
// 	VK_LCONTROL            VKCode = 0xA2
// 	VK_RCONTROL            VKCode = 0xA3
// 	VK_LMENU               VKCode = 0xA4
// 	VK_RMENU               VKCode = 0xA5
// 	VK_BROWSER_BACK        VKCode = 0xA6
// 	VK_BROWSER_FORWARD     VKCode = 0xA7
// 	VK_BROWSER_REFRESH     VKCode = 0xA8
// 	VK_BROWSER_STOP        VKCode = 0xA9
// 	VK_BROWSER_SEARCH      VKCode = 0xAA
// 	VK_BROWSER_FAVORITES   VKCode = 0xAB
// 	VK_BROWSER_HOME        VKCode = 0xAC
// 	VK_VOLUME_MUTE         VKCode = 0xAD
// 	VK_VOLUME_DOWN         VKCode = 0xAE
// 	VK_VOLUME_UP           VKCode = 0xAF
// 	VK_MEDIA_NEXT_TRACK    VKCode = 0xB0
// 	VK_MEDIA_PREV_TRACK    VKCode = 0xB1
// 	VK_MEDIA_STOP          VKCode = 0xB2
// 	VK_MEDIA_PLAY_PAUSE    VKCode = 0xB3
// 	VK_LAUNCH_MAIL         VKCode = 0xB4
// 	VK_LAUNCH_MEDIA_SELECT VKCode = 0xB5
// 	VK_LAUNCH_APP1         VKCode = 0xB6
// 	VK_LAUNCH_APP2         VKCode = 0xB7
// 	VK_OEM_1               VKCode = 0xBA
// 	VK_OEM_PLUS            VKCode = 0xBB
// 	VK_OEM_COMMA           VKCode = 0xBC
// 	VK_OEM_MINUS           VKCode = 0xBD
// 	VK_OEM_PERIOD          VKCode = 0xBE
// 	VK_OEM_2               VKCode = 0xBF
// 	VK_OEM_3               VKCode = 0xC0
// 	VK_OEM_4               VKCode = 0xDB
// 	VK_OEM_5               VKCode = 0xDC
// 	VK_OEM_6               VKCode = 0xDD
// 	VK_OEM_7               VKCode = 0xDE
// 	VK_OEM_8               VKCode = 0xDF
// 	VK_OEM_AX              VKCode = 0xE1
// 	VK_OEM_102             VKCode = 0xE2
// 	VK_ICO_HELP            VKCode = 0xE3
// 	VK_ICO_00              VKCode = 0xE4
// 	VK_PROCESSKEY          VKCode = 0xE5
// 	VK_ICO_CLEAR           VKCode = 0xE6
// 	VK_OEM_RESET           VKCode = 0xE9
// 	VK_OEM_JUMP            VKCode = 0xEA
// 	VK_OEM_PA1             VKCode = 0xEB
// 	VK_OEM_PA2             VKCode = 0xEC
// 	VK_OEM_PA3             VKCode = 0xED
// 	VK_OEM_WSCTRL          VKCode = 0xEE
// 	VK_OEM_CUSEL           VKCode = 0xEF
// 	VK_OEM_ATTN            VKCode = 0xF0
// 	VK_OEM_FINISH          VKCode = 0xF1
// 	VK_OEM_COPY            VKCode = 0xF2
// 	VK_OEM_AUTO            VKCode = 0xF3
// 	VK_OEM_ENLW            VKCode = 0xF4
// 	VK_OEM_BACKTAB         VKCode = 0xF5
// 	VK_ATTN                VKCode = 0xF6
// 	VK_CRSEL               VKCode = 0xF7
// 	VK_EXSEL               VKCode = 0xF8
// 	VK_EREOF               VKCode = 0xF9
// 	VK_PLAY                VKCode = 0xFA
// 	VK_ZOOM                VKCode = 0xFB
// 	VK_NONAME              VKCode = 0xFC
// 	VK_PA1                 VKCode = 0xFD
// 	VK_OEM_CLEAR           VKCode = 0xFE
// )
