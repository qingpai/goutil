package util

import (
	"github.com/mozillazg/go-pinyin"
	"strings"
)

// GetPinyinFirstLetter @title 获取拼音首字母
// @param text 输入文字
func GetPinyinFirstLetter(text string) string {
	return GetPinyin(text, "firstLetter", "")
}

// GetPinyin @title 获取给定文字的拼音
// @param text 输入文字
// @param mode 模式: tone(声调), heteronym(多音字), firstLetter(首字母)
// @param upper 是否大写
func GetPinyin(text string, mode string, upper string) string {
	args := pinyin.NewArgs()
	args.Fallback = func(r rune, a pinyin.Args) []string {
		return []string{string(r)}
	}

	if mode == "tone" {
		args.Style = pinyin.TONE
	} else if mode == "heteronym" {
		args.Heteronym = true
	} else if mode == "firstLetter" {
		args.Style = pinyin.FirstLetter
	}

	text = strings.Trim(text, ",")

	return textToPinyin(text, args, upper)
}

func textToPinyin(text string, args pinyin.Args, upper string) string {
	pinyinResult := pinyin.Pinyin(text, args)
	result := make([]string, 0)

	for _, s := range pinyinResult {
		sub := strings.Join(s, "")
		if upper == "title" {
			sub = strings.Title(sub)
		} else if upper == "all" {
			sub = strings.ToUpper(sub)
		}
		result = append(result, sub)
	}

	return strings.Join(result, "")
}
