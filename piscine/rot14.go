package piscine

func Rot14(the_characters_channel string) string {
	the_Runes_channel := []rune(the_characters_channel)
	var result_after_rotation string

	for i := 0; i < len(the_Runes_channel); i++ {
		if the_Runes_channel[i] >= 'a' && the_Runes_channel[i] <= 'z' {
			if the_Runes_channel[i] >= 'm' {
				the_Runes_channel[i] -= 12
			} else {
				the_Runes_channel[i] += 14
			}
		} else if the_Runes_channel[i] >= 'A' && the_Runes_channel[i] <= 'Z' {
			if the_Runes_channel[i] >= 'M' {
				the_Runes_channel[i] -= 12
			} else {
				the_Runes_channel[i] += 14
			}
		}
		result_after_rotation += string(the_Runes_channel[i])
	}

	return result_after_rotation
}
