Translation
===========

This is a translation package that uses in your project to translate your text to any language.

## Installation

```bash
$ go get github.com/mohsenabedy1991/translation
```


## Set environment variables

```bash
export TRANSLATION_LOCALE=en
export TRANSLATION_FALLBACK_LOCALE=en
export TRANSLATION_PATH_LOCALE=translation
```

or you can copy to .env file
    
```
TRANSLATION_LOCALE=en
TRANSLATION_FALLBACK_LOCALE=en
TRANSLATION_PATH_LOCALE=translation
```


you can change the name of the path locale to any name you want.

## Usage

```go
package main

import (
	"fmt"
	"github.com/mohsenabedy1991/translation"
)

func main() {
	// Create a new translation instance.
	t := translation.NewTranslation()
	t.GetLocalizer("el")

	// Print the translated string.
	fmt.Println(t.Lang("hello", nil, nil))

	// Print the translated string with variables.
	fmt.Println(t.Lang("hello", map[string]interface{}{
		"name": "John Doe",
	}, nil))

	// Print the translated string with variables and locale.
	lang := "en"
	fmt.Println(t.Lang("hello", map[string]interface{}{
		"name": "John Doe",
	}, &lang))
}
```
