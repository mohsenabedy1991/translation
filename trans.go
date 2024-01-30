package translation

import (
	"encoding/json"
	"fmt"
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
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	root := os.Getenv("PATH_LOCALE")
	createLocaleDirectory(root)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		// Check if the current path is a JSON file
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			if _, err := bundle.LoadMessageFile(path); err != nil {
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
	SetTranslationDirectory(path string)
	SetFallbackLocale(locale string)
	SetLocale(locale string)
}

func NewTranslation() *Translation {
	return &Translation{}
}

// GetLocalizer initializes the localizer with the desired language.
func (t *Translation) GetLocalizer(lang string) *i18n.Localizer {
	if lang == "" {
		lang = os.Getenv("APP_LOCALE")
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

	if *lang != "" {
		AcceptLanguage = t.GetLocalizer(*lang)
	}

	message, err := AcceptLanguage.Localize(config)
	if err != nil {
		// If the message is not found, fall back to default language (configs.GetConfig().AppFallbackLocale))
		defaultLang := i18n.NewLocalizer(bundle, os.Getenv("APP_FALLBACK_LOCALE"))
		message = defaultLang.MustLocalize(config)
	}

	return message
}

func (t *Translation) SetTranslationDirectory(path string) {
	if path == "" {
		path = "templates/translations/"
	}
	err := os.Setenv("PATH_LOCALE", path)
	if err != nil {
		return
	}
}

func (t *Translation) SetFallbackLocale(locale string) {
	if locale == "" {
		locale = "en"
	}
	err := os.Setenv("APP_FALLBACK_LOCALE", locale)
	if err != nil {
		return
	}
}

func (t *Translation) SetLocale(locale string) {
	if locale == "" {
		locale = "en"
	}
	err := os.Setenv("APP_LOCALE", locale)
	if err != nil {
		return
	}
}

// createLocaleDirectory create locale directory in root project if not exists
func createLocaleDirectory(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, 0755)
		if err != nil {
			return
		}
	}
}
