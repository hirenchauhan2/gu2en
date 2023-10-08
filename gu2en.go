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

func Transliterate(text string) string {
	// Gujarati consonants
	guj_consonants := []string{
		"ક",
		"ખ",
		"ગ",
		"ઘ",
		"ચ",
		"છ",
		"જ",
		"ઝ",
		"ટ",
		"ઠ",
		"ડ",
		"ઢ",
		"ણ",
		"ત",
		"થ",
		"દ",
		"ધ",
		"ન",
		"પ",
		"ફ",
		"બ",
		"ભ",
		"મ",
		"ય",
		"ર",
		"લ",
		"વ",
		"શ",
		"ષ",
		"સ",
		"હ",
		"ળ",
		"ઞ",
	}

	// Equivalents for gujarati consonants
	eng_consonants := []string{
		"ka",
		"kha",
		"ga",
		"gha",
		"cha",
		"chha",
		"ja",
		"jha",
		"ṭa",
		"ṭha",
		"ḍa",
		"ḍha",
		"ṇa",
		"ta",
		"tha",
		"da",
		"dha",
		"na",
		"pa",
		"fa",
		"ba",
		"bha",
		"ma",
		"ya",
		"ra",
		"la",
		"va",
		"sha",
		"ṣha",
		"sa",
		"ha",
		"ḷa",
		"nya",
	}

	// Gujarati half letters, which would be joined with another character to form
	// conjunct letters
	guj_half := []string{
		"ક્",
		"ખ્",
		"ગ્",
		"ઘ્",
		"ચ્",
		"છ્",
		"જ્",
		"ઝ્",
		"ટ્",
		"ઠ્",
		"ડ્",
		"ઢ્",
		"ણ્",
		"ત્",
		"થ્",
		"દ્",
		"ધ્",
		"ન્",
		"પ્",
		"ફ્",
		"બ્",
		"ભ્",
		"મ્",
		"ય્",
		"ર્",
		"લ્",
		"વ્",
		"શ્",
		"ષ્",
		"સ્",
		"હ્",
		"ળ્",
		"ઞ્",
	}

	// Equivalent characters for gujarati half letters
	eng_half := []string{
		"k",
		"kh",
		"g",
		"gh",
		"ch",
		"chh",
		"j",
		"z",
		"ṭ",
		"ṭh",
		"ḍ",
		"ḍh",
		"ṇ",
		"t",
		"th",
		"d",
		"dh",
		"n",
		"p",
		"f",
		"b",
		"bh",
		"m",
		"y",
		"r",
		"l",
		"v",
		"sh",
		"ṣh",
		"s",
		"h",
		"ḷ",
		"ny",
	}

	// Gujarati vowels
	guj_vowels := []string{"ઓ", "ઔ", "આ", "ઇ", "ઈ", "ઉ", "ઊ", "એ", "ઐ", "ઍ", "ઑ", "ૠ", "અ"}

	// Equivalents for gujarati vowels
	eng_vowels := []string{
		"o",
		"au",
		"ā",
		"i",
		"ī",
		"u",
		"ū",
		"e",
		"ai",
		"ĕ",
		"ŏ",
		"ṛu",
		"a",
	}

	// Gujarati diacrits
	guj_diacritics := []string{
		"િ",
		"ી",
		"ુ",
		"ૂ",
		"ે",
		"ૈ",
		"ો",
		"ૌ",
		"ૅ",
		"ૉ",
		"ં",
		"ૃ",
		"્",
		"ઃ",
		"ા",
	}

	// Equivalents for gujarati diacrits
	eng_diacrits := []string{
		"i",
		"ī",
		"u",
		"ū",
		"e",
		"ai",
		"o",
		"au",
		"ĕ",
		"ŏ",
		"an",
		"ṛu",
		"",
		"ah",
		"ā",
	}

	// Gujarati digits
	guj_digits := []string{"૦", "૧", "૨", "૩", "૪", "૫", "૬", "૭", "૮", "૯"}

	// English digits
	eng_digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	// gujarati consonants and diacrit "ઃ" (h)
	for i := 0; i < len(guj_consonants); i++ {
		for j := 0; j < len(guj_diacritics); j++ {
			pat := regexp.MustCompile(guj_consonants[i] + guj_diacritics[j] + "ઃ ")
			replacer := eng_half[i] + eng_diacrits[j] + "h" + eng_diacrits[j] + " "
			text = pat.ReplaceAllString(text, replacer)
		}
	}

	for i := 0; i < len(guj_consonants); i++ {
		for j := 0; j < len(guj_diacritics); j++ {
			pat := regexp.MustCompile(guj_consonants[i] + guj_diacritics[j] + "ં")
			replacer := eng_half[i] + eng_diacrits[j] + "n"
			text = pat.ReplaceAllString(text, replacer)
		}
	}

	for i := 0; i < len(guj_vowels); i++ {
		pat := regexp.MustCompile(guj_vowels[i] + "ં")
		replacer := eng_vowels[i] + "n"
		text = pat.ReplaceAllString(text, replacer)
	}

	for i := 0; i < len(guj_consonants); i++ {
		for j := 0; j < len(guj_diacritics); j++ {
			pat := regexp.MustCompile(guj_consonants[i] + guj_diacritics[j])
			replacer := eng_half[i] + eng_diacrits[j]
			text = pat.ReplaceAllString(text, replacer)
		}
	}

	text = replaceArray(guj_half, eng_half, text)
	text = replaceArray(guj_diacritics, eng_diacrits, text)
	text = replaceArray(guj_vowels, eng_vowels, text)
	text = replaceArray(guj_consonants, eng_consonants, text)
	text = replaceArray(guj_digits, eng_digits, text)

	for i := 0; i < len(eng_consonants); i++ {
		pat := regexp.MustCompile(eng_consonants[i] + " ")
		replacer := eng_half[i] + " "

		text = pat.ReplaceAllString(text, replacer)
	}

	for i := 0; i < len(eng_half); i++ {
		for j := 0; j < len(eng_consonants); j++ {
			pat := regexp.MustCompile(eng_half[i] + eng_half[j] + " ")
			replacer := eng_half[i] + eng_consonants[j] + " "

			text = pat.ReplaceAllString(text, replacer)
		}
	}

	eng_comb := []string{
		"kh",
		"gh",
		"ch",
		"chh",
		"ṭh",
		"ḍh",
		"th",
		"dh",
		"bh",
		"sh",
		"ṣh",
		"ny",
	}

	for i := 0; i < len(eng_vowels); i++ {
		for j := 0; j < len(eng_comb); j++ {
			pat := regexp.MustCompile(eng_vowels[i] + eng_comb[j] + "a ")
			replacer := eng_vowels[i] + eng_comb[j] + " "
			text = pat.ReplaceAllString(text, replacer)
		}
	}

	lowercase_alphabets := []string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
		"h",
		"i",
		"j",
		"k",
		"l",
		"m",
		"n",
		"o",
		"p",
		"q",
		"r",
		"s",
		"t",
		"u",
		"v",
		"w",
		"x",
		"y",
		"z",
		"ā",
		"ḍ",
		"ĕ",
		"ī",
		"ḷ",
		"ṇ",
		"ŏ",
		"ṛ",
		"ṣ",
		"ṭ",
		"ū",
	}
	uppercase_alphabets := []string{
		"A",
		"B",
		"C",
		"D",
		"E",
		"F",
		"G",
		"H",
		"I",
		"J",
		"K",
		"L",
		"M",
		"N",
		"O",
		"P",
		"Q",
		"R",
		"S",
		"T",
		"U",
		"V",
		"W",
		"X",
		"Y",
		"Z",
		"Ā",
		"Ḍ",
		"Ĕ",
		"Ī",
		"Ḷ",
		"Ṇ",
		"Ŏ",
		"Ṛ",
		"Ṣ",
		"Ṭ",
		"Ū",
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
