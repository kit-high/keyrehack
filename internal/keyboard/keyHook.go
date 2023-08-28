package keyboard

import (
	"keyrehack/internal/combinationReplacer"
	simplereplacer "keyrehack/internal/simpleReplacer"
)

func SetHack(simpleReplaces []simplereplacer.ReplaceKey, combinationReplaces []combinationReplacer.ReplaceKey) error {
	return setHook(simpleReplaces, combinationReplaces)
}

func UnsetHack() {
	unsetHook()
}
