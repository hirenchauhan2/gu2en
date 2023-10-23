package gu2en

import (
	"regexp"
	"strings"
)

func replaceArray(a []string, b []string, text string) string {
	if len(a) != len(b) {
		return ""
	}
	for i := 0; i < len(a); i++ {
		pat := regexp.MustCompile(a[i])
		text = pat.ReplaceAllString(text, b[i])
	}
	return text
}

// It generates transliteration of Gujarati Unicode text to English/Roman text.
//
// For example, the Gujarati text "ગુજરાતી ભાષા" will be transliterated to "Gujarātī bhāṣhā".
func Transliterate(text string) string {
	// gujarati consonants and diacrit "ઃ" (h)
	for i := 0; i < len(guj_consonants); i++ {
		for j := 0; j < len(guj_diacritics); j++ {
			pat := regexp.MustCompile(guj_consonants[i] + guj_diacritics[j] + "ઃ ")
			replacer := eng_halant[i] + eng_diacritics[j] + "h" + eng_diacritics[j] + " "
			text = pat.ReplaceAllString(text, replacer)
		}
	}
	// check if the text has anusvara in it
	anusvara_regex := regexp.MustCompile("ં")
	has_anusvara := anusvara_regex.MatchString(text)
	if has_anusvara {
		// replace the combo of consonant + diacrit + anusvara
		for i := 0; i < len(guj_consonants); i++ {
			for j := 0; j < len(guj_diacritics); j++ {
				pat_m := regexp.MustCompile(guj_consonants[i] + guj_diacritics[j] + "ં[મપબભ]")
				res := pat_m.MatchString(text)
				if res {
					pat_m := regexp.MustCompile(guj_consonants[i] + guj_diacritics[j] + "ં")
					replacer_m := eng_halant[i] + eng_diacritics[j] + "m"
					text = pat_m.ReplaceAllString(text, replacer_m)
				} else {
					pat_n := regexp.MustCompile(guj_consonants[i] + guj_diacritics[j] + "ં")
					replacer_n := eng_halant[i] + eng_diacritics[j] + "n"
					text = pat_n.ReplaceAllString(text, replacer_n)
				}
			}
		}

		// replace the combo of vowel + anusvara
		for i := 0; i < len(guj_vowels); i++ {
			pat_m := regexp.MustCompile(guj_vowels[i] + "ં[મપબભ]")
			res := pat_m.MatchString(text)
			if res {
				pat_m := regexp.MustCompile(guj_vowels[i] + "ં")
				replacer_m := eng_vowels[i] + "m"
				text = pat_m.ReplaceAllString(text, replacer_m)
			} else {
				pat_n := regexp.MustCompile(guj_vowels[i] + "ં")
				replacer_n := eng_vowels[i] + "n"
				text = pat_n.ReplaceAllString(text, replacer_n)
			}
		}

		// replace the combo of consonant + anusvara
		for i := 0; i < len(guj_consonants); i++ {
			pat_m := regexp.MustCompile(guj_consonants[i] + "ં[મપબભ]")
			res := pat_m.MatchString(text)
			if res {
				pat_m := regexp.MustCompile(guj_consonants[i] + "ં")
				replacer_m := eng_consonants[i] + "m"
				text = pat_m.ReplaceAllString(text, replacer_m)
			} else {
				pat_n := regexp.MustCompile(guj_consonants[i] + "ં")
				replacer_n := eng_consonants[i] + "n"
				text = pat_n.ReplaceAllString(text, replacer_n)
			}
		}
	}

	// replace the combo of consonant + diacrit
	for i := 0; i < len(guj_consonants); i++ {
		for j := 0; j < len(guj_diacritics); j++ {
			pat := regexp.MustCompile(guj_consonants[i] + guj_diacritics[j])
			replacer := eng_halant[i] + eng_diacritics[j]
			text = pat.ReplaceAllString(text, replacer)
		}
	}

	text = replaceArray(guj_halant, eng_halant, text)
	text = replaceArray(guj_diacritics, eng_diacritics, text)
	text = replaceArray(guj_vowels, eng_vowels, text)
	text = replaceArray(guj_consonants, eng_consonants, text)
	text = replaceArray(guj_digits, eng_digits, text)

	for i := 0; i < len(eng_consonants); i++ {
		pat := regexp.MustCompile(eng_consonants[i] + " ")
		replacer := eng_halant[i] + " "

		text = pat.ReplaceAllString(text, replacer)
	}

	for i := 0; i < len(eng_halant); i++ {
		for j := 0; j < len(eng_consonants); j++ {
			pat := regexp.MustCompile(eng_halant[i] + eng_halant[j] + " ")
			replacer := eng_halant[i] + eng_consonants[j] + " "

			text = pat.ReplaceAllString(text, replacer)
		}
	}

	for i := 0; i < len(eng_vowels); i++ {
		for j := 0; j < len(eng_comb); j++ {
			pat := regexp.MustCompile(eng_vowels[i] + eng_comb[j] + "a ")
			replacer := eng_vowels[i] + eng_comb[j] + " "
			text = pat.ReplaceAllString(text, replacer)
		}
	}

	//capitalize first word in sentence
	for i := 0; i < len(lowercase_alphabets); i++ {
		p1 := regexp.MustCompile(`(\.|\?|!)\s` + lowercase_alphabets[i])
		p2 := regexp.MustCompile(`,\s("|“)` + lowercase_alphabets[i])
		p3 := regexp.MustCompile(`(\."|\.”)\s` + lowercase_alphabets[i])
		p4 := regexp.MustCompile(`(\.|\?|!)\s("|“)` + lowercase_alphabets[i])
		p5 := regexp.MustCompile(`(\?”+\?”|!”|!")\s` + lowercase_alphabets[i])
		p6 := regexp.MustCompile(`("|”)\s("|“)` + lowercase_alphabets[i])
		p7 := regexp.MustCompile(`^` + lowercase_alphabets[i])
		p8 := regexp.MustCompile(`\n` + lowercase_alphabets[i])

		text = p1.ReplaceAllString(text, "$1 "+uppercase_alphabets[i])

		text = p2.ReplaceAllString(text, ", $1"+uppercase_alphabets[i])

		text = p3.ReplaceAllString(text, "$1 "+uppercase_alphabets[i])
		text = p4.ReplaceAllString(text, "$1 $2"+uppercase_alphabets[i])
		text = p5.ReplaceAllString(text, "$1 "+uppercase_alphabets[i])
		text = p6.ReplaceAllString(text, "$1 $2"+uppercase_alphabets[i])
		text = p7.ReplaceAllString(text, uppercase_alphabets[i])
		text = p8.ReplaceAllString(text, "\n"+uppercase_alphabets[i])
	}

	// post-processing fixes and trailing y
	text = strings.ReplaceAll(text, "yvar ", "ya")
	text = strings.ReplaceAll(text, "y ", "ya ")
	text = strings.ReplaceAll(text, "jnyaa", "gnaa")
	text = strings.ReplaceAll(text, "jnya", "gna")
	text = strings.ReplaceAll(text, "jny ", "gna")
	text = strings.ReplaceAll(text, "svaa", "swaa")
	text = strings.ReplaceAll(text, "svā", "swā")

	return text
}
