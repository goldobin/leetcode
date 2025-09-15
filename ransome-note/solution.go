package ransome_note

func canConstruct(ransomNote string, magazine string) bool {
	letters := make(map[int32]int, len(magazine))
	for _, l := range magazine {
		count, ok := letters[l]
		if !ok {
			letters[l] = 1
		}
		letters[l] = count + 1
	}

	for _, l := range ransomNote {
		count, ok := letters[l]
		if !ok {
			return false
		}

		count -= 1
		if count == 0 {
			delete(letters, l)
		} else {
			letters[l] = count
		}
	}

	return true
}
