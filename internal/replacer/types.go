package replacer

type IsReplaced bool

type Result struct {
	IsReplaced  bool
	InputLen    uintptr
	InputArray  uintptr
	SizeOfInput uintptr
}
