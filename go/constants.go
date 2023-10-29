package gu2en

// Gujarati consonants
var guj_consonants = []string{
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
var eng_consonants = []string{
	"ka",
	"kha",
	"ga",
	"gha",
	"cha",
	"chh",
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

// Halant form of Gujarati consonants, which would be joined with another character to form
// conjunct letters
var guj_halant = []string{
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

// Equivalent characters for gujarati halant letters
var eng_halant = []string{
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
var guj_vowels = []string{"ઓ", "ઔ", "આ", "ઇ", "ઈ", "ઉ", "ઊ", "એ", "ઐ", "ઍ", "ઑ", "ૠ", "અ"}

// Equivalents for gujarati vowels
var eng_vowels = []string{
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

// Gujarati diacritics
var guj_diacritics = []string{
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
	// "ં", -- This is being replaced with regex due to contextual representation with "m" or "n"
	"ૃ",
	"્",
	"ઃ",
	"ા",
}

// Equivalents for gujarati diacritics
var eng_diacritics = []string{
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
	// "an", -- This is being replaced with regex due to contextual representation with "m" or "n"
	"ṛu",
	"",
	"ah",
	"ā",
}

// Gujarati digits
var guj_digits = []string{"૦", "૧", "૨", "૩", "૪", "૫", "૬", "૭", "૮", "૯"}

// English digits
var eng_digits = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

var eng_comb = []string{
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

var lowercase_alphabets = []string{
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

var uppercase_alphabets = []string{
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
