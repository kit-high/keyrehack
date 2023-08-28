package replacer

import (
	. "keyrehack/internal/windows"
)

type Replacer interface {
	Replace(vkCode VKCode, message Message) IsReplaced
}
