# gu2en

An algorithm to transliterate phonetically Gujarati Unicode strings to Roman / Latin characters.

## Example

### Input

આ ગુજરાતી ભાષા જેવી સરળતા બીજે ક્યાક જો જોવા મળે તો તમે નવી શોધ કરી છે, એમ માનજો.

### Output

Ā gujarātī bhāṣhā jevī saraḷatā bīje kyāk jo jovā maḷe to tame navī shodh karī chhe, em mānajo.

## Usage

```js
import { transliterate } from 'gu2en';

console.log(
  transliterate(
    'આ ગુજરાતી ભાષા જેવી સરળતા બીજે ક્યાક જો જોવા મળે તો તમે નવી શોધ કરી છે, એમ માનજો.'
  )
);
```

*UMD bundle adds the `gu2en` namespace in the global scope.*

```html
<script src="dist/gu2en.umd.min.js"></script>
<script>
  console.log(
    gu2en.transliterate(
      'આ ગુજરાતી ભાષા જેવી સરળતા બીજે ક્યાક જો જોવા મળે તો તમે નવી શોધ કરી છે, એમ માનજો.'
    )
  );
</script>
```

## Todo

- [ ] Identify or fix transliteration mistakes
- [ ] Add [IPA](https://www.wikiwand.com/en/International_Phonetic_Alphabet) conversion mode
- [ ] Add normal english characters mode for better readability conversions (Gujlish) text

### License

GPL V3
