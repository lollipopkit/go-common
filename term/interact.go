package term

import (
	"os"
	"regexp"

	"atomicgo.dev/cursor"
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

var (
	doubleByteCharacterRegexp = regexp.MustCompile(`[^\x00-\xff]`)
	EmptyStringList           = []string{}
	promptRunes               = []rune(prompt)
)

const (
	prompt = "> "
)

func ReadLine(linesHistory []string, onCtrlC func()) string {
	os.Stdout.WriteString(prompt)
	rs := []rune{}
	linesIdx := len(linesHistory)
	runeIdx := 0

	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		switch key.Code {
		case keys.CtrlC:
			os.Exit(0)
		case keys.Escape:
			return true, nil
		case keys.RuneKey:
			runes := key.Runes
			rs = append(rs[:runeIdx], append(runes, rs[runeIdx:]...)...)
			runeIdx += len(runes)
			resetLine(rs, prompt)
		case keys.Enter:
			println()
			return true, nil
		case keys.Backspace:
			if len(rs) > 0 && runeIdx > 0 {
				rs = append(rs[:runeIdx-1], rs[runeIdx:]...)
				resetLine(rs, prompt)
				runeIdx--
			}
		case keys.Left:
			if runeIdx > 0 {
				runeIdx--
			}
		case keys.Right:
			if runeIdx < len(rs) {
				runeIdx++
			}
		case keys.Up:
			if linesIdx > 0 {
				linesIdx--
				rs = []rune(linesHistory[linesIdx])
				resetLine(rs, prompt)
				runeIdx = len(rs)
			}
		case keys.Down:
			if linesIdx < len(linesHistory)-1 {
				linesIdx++
				rs = []rune(linesHistory[linesIdx])
				resetLine(rs, prompt)
				runeIdx = len(rs)
			} else if linesIdx == len(linesHistory)-1 {
				linesIdx++
				rs = []rune("")
				resetLine(rs, prompt)
				runeIdx = 0
			}
		case keys.Space:
			if runeIdx == len(rs) {
				rs = append(rs, ' ')
				print(" ")
				runeIdx++
			} else {
				rs = append(rs[:runeIdx], append([]rune(" "), rs[runeIdx:]...)...)
				resetLine(rs, prompt)
				runeIdx++
			}
		case keys.Tab:
			if runeIdx == len(rs) {
				rs = append(rs, '\t')
				print("\t")
				runeIdx++
			} else {
				rs = append(rs[:runeIdx], append([]rune("\t"), rs[runeIdx:]...)...)
				resetLine(rs, prompt)
				runeIdx++
			}
		case keys.Delete:
			if runeIdx < len(rs) {
				rs = append(rs[:runeIdx], rs[runeIdx+1:]...)
				resetLine(rs, prompt)
			}
		}

		idx := calcIdx(rs, runeIdx)
		pIdx := calcIdx(promptRunes, len(promptRunes))
		cursor.HorizontalAbsolute(idx + pIdx)
		return false, nil
	})
	return string(rs)
}

func resetLine(rs []rune, prompt string) {
	cursor.ClearLine()
	cursor.StartOfLine()
	print(prompt + string(rs))
}

func calcIdx(rs []rune, runeIdx int) int {
	idx := 0
	for rIdx, r := range rs {
		if rIdx >= runeIdx {
			break
		}
		if isHan(r) {
			idx += 2
		} else {
			idx++
		}
	}
	return idx
}

func isHan(r rune) bool {
	return doubleByteCharacterRegexp.MatchString(string(r))
}
