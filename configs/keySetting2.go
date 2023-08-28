package configs

import (
	"keyrehack/internal/combinationReplacer"           // Delete this line to eliminate import from internal.
	simplereplacer "keyrehack/internal/simpleReplacer" // Delete this line to eliminate import from internal.
	. "keyrehack/internal/windows"                     // Delete this line to eliminate import from internal.
)

func GetKeySetting2() ([]simplereplacer.ReplaceKey, []combinationReplacer.ReplaceKey) {

	yourFunctionKey := func() *VKCode { return VK_F13() }

	simpleReplace := []simplereplacer.ReplaceKey{
		{Source: yourFunctionKey(), Target: VK_LCONTROL()},
	}

	combinationReplace := []combinationReplacer.ReplaceKey{
		{Source: combinationReplacer.Combination{Hold: VK_LMENU(), Press: &[]*VKCode{VK_OEM_PLUS()}},
			Target: combinationReplacer.Combination{Hold: VK_LCONTROL(), Press: &[]*VKCode{VK_LEFT()}}},
		{Source: combinationReplacer.Combination{Hold: VK_LMENU(), Press: &[]*VKCode{VK_OEM_1()}},
			Target: combinationReplacer.Combination{Hold: VK_LCONTROL(), Press: &[]*VKCode{VK_RIGHT()}}},
	}

	return simpleReplace, combinationReplace
}
