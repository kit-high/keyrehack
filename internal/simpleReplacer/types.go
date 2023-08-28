package simplereplacer

import . "keyrehack/internal/windows"

type ReplaceKey struct {
	Source *VKCode // nil means no replacement
	Target *VKCode // nil means no replacement
}
