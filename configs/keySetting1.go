package configs

import (
	"keyrehack/internal/combinationReplacer"           // Delete this line to eliminate import from internal.
	simplereplacer "keyrehack/internal/simpleReplacer" // Delete this line to eliminate import from internal.
	. "keyrehack/internal/windows"                     // Delete this line to eliminate import from internal.
)

func GetKeySetting1() ([]simplereplacer.ReplaceKey, []combinationReplacer.ReplaceKey) {
	simpleReplace := []simplereplacer.ReplaceKey{
		{Source: VK_NONCONVERT(), Target: VK_IME_OFF()},
		{Source: VK_CONVERT(), Target: VK_IME()},
	}

	yourFunctionKey := func() *VKCode { return VK_F13() }

	combinationReplace := []combinationReplacer.ReplaceKey{
		{Source: combinationReplacer.Combination{Hold: yourFunctionKey(), Press: &[]*VKCode{VK_OEM_PLUS()}},
			Target: combinationReplacer.Combination{Hold: VK_LEFT()}},
		{Source: combinationReplacer.Combination{Hold: yourFunctionKey(), Press: &[]*VKCode{VK_OEM_1()}},
			Target: combinationReplacer.Combination{Hold: VK_RIGHT()}},
		{Source: combinationReplacer.Combination{Hold: yourFunctionKey(), Press: &[]*VKCode{VK_OEM_3()}},
			Target: combinationReplacer.Combination{Hold: VK_UP()}},
		{Source: combinationReplacer.Combination{Hold: yourFunctionKey(), Press: &[]*VKCode{VK_OEM_2()}},
			Target: combinationReplacer.Combination{Hold: VK_DOWN()}},

		{Source: combinationReplacer.Combination{Hold: yourFunctionKey(), Press: &[]*VKCode{VK_K()}},
			Target: combinationReplacer.Combination{Hold: VK_HOME()}},
		{Source: combinationReplacer.Combination{Hold: yourFunctionKey(), Press: &[]*VKCode{VK_OEM_COMMA()}},
			Target: combinationReplacer.Combination{Hold: VK_END()}},

		{Source: combinationReplacer.Combination{Hold: VK_LMENU(), Press: &[]*VKCode{VK_OEM_PLUS()}},
			Target: combinationReplacer.Combination{Hold: VK_LCONTROL(), Press: &[]*VKCode{VK_LEFT()}}},
		{Source: combinationReplacer.Combination{Hold: VK_LMENU(), Press: &[]*VKCode{VK_OEM_1()}},
			Target: combinationReplacer.Combination{Hold: VK_LCONTROL(), Press: &[]*VKCode{VK_RIGHT()}}},
	}

	return simpleReplace, combinationReplace
}
