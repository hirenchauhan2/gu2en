/**
 * Replace the characters from a text with
 * @param {string[]} sourceCharsList Source array of characters to be replaced in the text
 * @param {string[]} targetCharsList Replacements characters array to get replaced with source in the text
 * @param {string} text The text from which the characters needs to be replaced
 * @returns {string} text with replaced characters from source characters
 */
export default function replaceChars(sourceCharsList, targetCharsList, text) {
  if (sourceCharsList.length !== targetCharsList.length) {
    return '';
  }

  for (let i = 0; i < sourceCharsList.length; i++) {
    const regex = new RegExp(sourceCharsList[i], 'g');
    text = text.replace(regex, targetCharsList[i]);
  }
  return text;
}
