# gu2en

An algorithm to transliterate phonetically Gujarati Unicode strings to Roman / Latin characters.


## Example

### Input

આ ગુજરાતી ભાષા જેવી સરળતા બીજે ક્યાક જો જોવા મળે તો તમે નવી શોધ કરી છે, એમ માનજો.

### Output

Ā gujarātī bhāṣhā jevī saraḷatā bīje kyāk jo jovā maḷe to tame navī shodh karī chhe, em mānajo.

## Usage

```go
package main

import (
	"fmt"

	gu2en "github.com/hirenchauhan2/gu2en/go"
)

func main() {
    fmt.Println(gu2en.Transliterate("આ ગુજરાતી ભાષા જેવી સરળતા બીજે ક્યાક જો જોવા મળે તો તમે નવી શોધ કરી છે, એમ માનજો.
"))
}
```

## Todo

- [x] Write tests
- [x] Use "m" instead of "n" for diacrit "ં" in certain scenarios such as "ભં" and "અં" and others
- [ ] Identify or fix transliteration
- [ ] Add [IPA](https://www.wikiwand.com/en/International_Phonetic_Alphabet) conversion mode
- [ ] Add normal english characters mode for better readability conversions (Gujlish) text

### License
GPL V3