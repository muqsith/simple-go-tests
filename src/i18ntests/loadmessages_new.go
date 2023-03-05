package main

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func customMarshalFn(data []byte, v any) error {
	re := regexp.MustCompile("\"defaultMessage\"")
	updatedData := re.ReplaceAll(data, []byte("\"other\""))
	return json.Unmarshal(updatedData, v)
}

func load_messages_from_file_new() {
	// Unmarshaling from files

	bundle := i18n.NewBundle(language.Spanish)

	bundle.RegisterUnmarshalFunc("json", customMarshalFn)

	bundle.MustLoadMessageFile("src/i18ntests/lang-out/es.json")

	loc := i18n.NewLocalizer(bundle, "es")

	t1 := loc.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "HelloWorld",
		},
	})

	fmt.Println(t1)

	t2 := loc.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "OneMessage",
		},
		TemplateData: map[string]interface{}{
			"Name":  "Alex",
			"Count": 1,
		},
	})

	fmt.Println(t2)

	messagesCount := 10

	t3 := loc.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "MoreMessages",
		},
		TemplateData: map[string]interface{}{
			"Name":  "Alex",
			"Count": messagesCount,
		},
	})

	fmt.Println(t3)
}
