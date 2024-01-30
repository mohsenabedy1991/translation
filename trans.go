package translation

import (
	"encoding/json"
	"fmt"
	"github.com/mohsenabedy1991/translation/configs"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"os"
	"path/filepath"
)

var (
	AcceptLanguage *i18n.Localizer
	bundle         *i18n.Bundle
)

// Init initializes the localizer with the desired language.
func init() {
	config, err := configs.LoadConfig()
	if err != nil {
		return
	}

	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	root := config.PathLocale
	createLocaleDirectory(root)

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		// Check if the current path is a JSON file
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			if _, err = bundle.LoadMessageFile(path); err != nil {
				fmt.Printf("Failed to load message file %s: %v\n", path, err)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

type Translation struct{}

// Translator is an interface that defines the methods needed to translate messages.
type Translator interface {
	GetLocalizer(lang string) *i18n.Localizer
	Lang(key string, args map[string]interface{}) string
}

func NewTranslation() *Translation {
	return &Translation{}
}

// GetLocalizer initializes the localizer with the desired language.
func (t *Translation) GetLocalizer(lang string) *i18n.Localizer {
	config, err := configs.LoadConfig()
	if err != nil {
		return nil
	}

	if lang == "" {
		lang = config.Locale
	}
	tag, err := language.Parse(lang)
	if err != nil {
		fmt.Println("Failed to parse language tag:", err)
		tag = language.English
	}

	AcceptLanguage = i18n.NewLocalizer(bundle, tag.String())

	return AcceptLanguage
}

// Lang is a helper function that translates a message.
func (t *Translation) Lang(key string, args map[string]interface{}, lang *string) string {
	config := &i18n.LocalizeConfig{
		MessageID:    key,
		TemplateData: args,
	}

	if lang != nil {
		AcceptLanguage = t.GetLocalizer(*lang)
	}

	message, err := AcceptLanguage.Localize(config)
	if err != nil {
		defaultLang := i18n.NewLocalizer(bundle, os.Getenv("TRANSLATION_FALLBACK_LOCALE"))
		message, err = defaultLang.Localize(config)
		if err != nil {
			return key
		}
	}

	return message
}

// createLocaleDirectory create locale directory in root project if not exists
func createLocaleDirectory(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return
		}
	}
	// if en.json not exists create it
	if _, err := os.Stat(path + "/en.json"); !os.IsNotExist(err) {
		return
	}
	_, err := os.Create(path + "/en.json")
	if err != nil {
		return
	}
}
