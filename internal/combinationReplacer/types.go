package combinationReplacer

import (
	. "keyrehack/internal/windows"
)

type ReplaceKey struct {
	Source Combination // Source Press key is only one key.
	Target Combination
}

type Combination struct {
	Hold  *VKCode    // nil means no replacement
	Press *[]*VKCode // nil means no replacement
}
