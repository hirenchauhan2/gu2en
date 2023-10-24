import {
  eng_comb,
  eng_consonants,
  eng_diacritics,
  eng_digits,
  eng_halant,
  eng_vowels,
  guj_consonants,
  guj_diacritics,
  guj_digits,
  guj_halant,
  guj_vowels,
  lowercase_alphabets,
  uppercase_alphabets,
} from './constants';

import replaceChars from './utils/replace-chars';

/**
 * Transliterate Gujarati unicode string to english/roman characters string
 * @param {string} text String with Gujarati unicode characters
 * @returns English/Roman characters transliterated from gujarati
 */
export function transliterate(text) {
  // gujarati consonants and diacrit "ઃ" (h)
  for (let gucIndex = 0; gucIndex < guj_consonants.length; gucIndex++) {
    for (let gudIndex = 0; gudIndex < guj_diacritics.length; gudIndex++) {
      const regex = new RegExp(
        guj_consonants[gucIndex] + guj_diacritics[gudIndex] + 'ઃ ',
        'g'
      );
      const replacer =
        eng_halant[gucIndex] +
        eng_diacritics[gudIndex] +
        'h' +
        eng_diacritics[gudIndex] +
        ' ';
      text = text.replace(regex, replacer);
    }
  }

  const hasAnusvara = text.indexOf('ં') !== -1;
  if (hasAnusvara) {
    // replace the combo of consonant + diacrit + anusvara
    for (let gucIndex = 0; gucIndex < guj_consonants.length; gucIndex++) {
      for (let gudIndex = 0; gudIndex < guj_diacritics.length; gudIndex++) {
        const regexAnusvaraM = new RegExp(
          guj_consonants[gucIndex] + guj_diacritics[gudIndex] + 'ં[મપબભ]',
          'g'
        );
        if (regexAnusvaraM.test(text)) {
          const regexM = new RegExp(
            guj_consonants[gucIndex] + guj_diacritics[gudIndex] + 'ં'
          );
          const replacerM =
            eng_halant[gucIndex] + eng_diacritics[gudIndex] + 'm';
          text = text.replace(regexM, replacerM);
        } else {
          const regexN = new RegExp(
            guj_consonants[gucIndex] + guj_diacritics[gudIndex] + 'ં',
            'g'
          );
          const replacerN =
            eng_halant[gucIndex] + eng_diacritics[gudIndex] + 'n';
          text = text.replace(regexN, replacerN);
        }
      }
    }

    // replace the combo of vowel + anusvara
    for (let gujVIndex = 0; gujVIndex < guj_vowels.length; gujVIndex++) {
      const regexAnusvaraM = new RegExp(guj_vowels[gujVIndex] + 'ં[મપબભ]', 'g');
      if (regexAnusvaraM.test(text)) {
        const regexM = new RegExp(guj_vowels[gujVIndex] + 'ં', 'g');
        const replacerM = eng_vowels[gujVIndex] + 'm';
        text = text.replace(regexM, replacerM);
      } else {
        const regexN = new RegExp(guj_vowels[gujVIndex] + 'ં', 'g');
        const replacerM = eng_vowels[gujVIndex] + 'n';
        text = text.replace(regexN, replacerM);
      }
    }

    // replace the combo of consonant + anusvara
    for (let gucIndex = 0; gucIndex < guj_consonants.length; gucIndex++) {
      const regexAnusvaraM = new RegExp(
        guj_consonants[gucIndex] + 'ં[મપબભ]',
        'g'
      );

      if (regexAnusvaraM.test(text)) {
        const regexM = new RegExp(guj_consonants[gucIndex] + 'ં', 'g');
        const replacerM = eng_consonants[gucIndex] + 'm';
        text = text.replace(regexM, replacerM);
      } else {
        const regexM = new RegExp(guj_consonants[gucIndex] + 'ં', 'g');
        const replacerN = eng_consonants[gucIndex] + 'n';
        text = text.replace(regexM, replacerN);
      }
    }
  }

  // replace the combo of consonant + diacrit
  for (let i = 0; i < guj_consonants.length; i++) {
    for (let j = 0; j < guj_diacritics.length; j++) {
      const comboRegex = new RegExp(guj_consonants[i] + guj_diacritics[j], 'g');
      const replacer = eng_halant[i] + eng_diacritics[j];
      text = text.replace(comboRegex, replacer);
    }
  }

  // replace direct characters for vowels, consonants, diacritics digits, halant
  text = replaceChars(guj_halant, eng_halant, text);
  text = replaceChars(guj_diacritics, eng_diacritics, text);
  text = replaceChars(guj_vowels, eng_vowels, text);
  text = replaceChars(guj_consonants, eng_consonants, text);
  text = replaceChars(guj_digits, eng_digits, text);

  for (let i = 0; i < eng_consonants.length; i++) {
    const pat = new RegExp(eng_consonants[i] + ' ', 'g');
    const replacer = eng_halant[i] + ' ';
    text = text.replace(pat, replacer);
  }

  for (let i = 0; i < eng_halant.length; i++) {
    for (let j = 0; j < eng_consonants.length; j++) {
      const pat = new RegExp(eng_halant[i] + eng_halant[j] + ' ', 'g');
      const replacer = eng_halant[i] + eng_consonants[j] + ' ';
      text = text.replace(pat, replacer);
    }
  }

  for (let i = 0; i < eng_vowels.length; i++) {
    for (let j = 0; j < eng_comb.length; j++) {
      const pat = new RegExp(eng_vowels[i] + eng_comb[j] + 'a ');
      const replacer = eng_vowels[i] + eng_comb[j] + ' ';
      text = text.replace(pat, replacer);
    }
  }

  // Capitalize first word in sentence
  for (let i = 0; i < lowercase_alphabets.length; i++) {
    const p1 = new RegExp('(\\.|\\?|!)\\s' + lowercase_alphabets[i], 'g');
    const p2 = new RegExp(',\\s("|“)' + lowercase_alphabets[i], 'g');
    const p3 = new RegExp('(\\."|\\.”)\\s' + lowercase_alphabets[i], 'g');
    const p4 = new RegExp('(\\.|\\?|!)\\s("|“)' + lowercase_alphabets[i], 'g');
    const p5 = new RegExp('(\\?”+\\?"|!”|!")\\s' + lowercase_alphabets[i], 'g');
    const p6 = new RegExp('("|”)\\s("|“)' + lowercase_alphabets[i], 'g');
    const p7 = new RegExp('^' + lowercase_alphabets[i], 'g');
    const p8 = new RegExp('\n' + lowercase_alphabets[i], 'g');
    text = text.replace(p1, '$1 ' + uppercase_alphabets[i]);
    text = text.replace(p2, ', $1' + uppercase_alphabets[i]);
    text = text.replace(p3, '$1 ' + uppercase_alphabets[i]);
    text = text.replace(p4, '$1 $2' + uppercase_alphabets[i]);
    text = text.replace(p5, '$1 ' + uppercase_alphabets[i]);
    text = text.replace(p6, '$1 $2' + uppercase_alphabets[i]);
    text = text.replace(p7, uppercase_alphabets[i]);
    text = text.replace(p8, '\n' + uppercase_alphabets[i]);
  }

  // post-processing for gna and trailing y
  text = text.replace(/yvar\s/g, 'ya');
  text = text.replace(/y\s/g, 'ya ');
  text = text.replace(/jnyaa/g, 'gnaa');
  text = text.replace(/jnya/g, 'gna');
  text = text.replace(/jny\s/g, 'gna ');
  text = text.replace(/svaa/g, 'swaa');
  text = text.replace(/svā/g, 'swā');

  // ignore the anusvara for the word નાં (nān) as mostly it's written as nā only
  text = text.replace(/nān\s/g, 'nā ');

  return text;
}
