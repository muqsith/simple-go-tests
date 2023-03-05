package main

import (
	"encoding/json"
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func load_messages_from_file() {
	// Unmarshaling from files

	bundle := i18n.NewBundle(language.English)

	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	bundle.MustLoadMessageFile("src/i18ntests/lang-out/en.json")

	bundle.MustLoadMessageFile("src/i18ntests/lang-out/el.json")

	loc := i18n.NewLocalizer(bundle, "en")

	messagesCount := 10

	translation := loc.MustLocalize(&i18n.LocalizeConfig{

		MessageID: "messages",

		TemplateData: map[string]interface{}{

			"Name": "Alex",

			"Count": messagesCount,
		},

		PluralCount: messagesCount,
	})

	fmt.Println(translation)
}
