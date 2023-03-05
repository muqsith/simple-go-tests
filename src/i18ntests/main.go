// courtesy: https://phrase.com/blog/posts/internationalisation-in-go-with-go-i18n/

package main

import (
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func main() {

	/*
		To extract
		$ ~/go/bin/goi18n extract -outdir=src/i18ntests/lang-out -format=json src/i18ntests/main.go
		or
		$ ~/go/bin/goi18n merge -outdir=src/i18ntests/lang-out -format=json src/i18ntests/main.go
	*/

	bundle := i18n.NewBundle(language.English)
	localizer := i18n.NewLocalizer(bundle, "en")

	m1, _ := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "HelloWorld",
			Description: "A simple text",
			Other:       "Hello World!",
		},
	})
	fmt.Println(m1)

	catsCount := 2

	m2, _ := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "PersonCats",
			Description: "Just about cats",
			One:         "{{.Name}} has {{.Count}} cat.",
			Other:       "{{.Name}} has {{.Count}} cats.",
		},
		TemplateData: map[string]interface{}{
			"Name":  "Nick",
			"Count": catsCount,
		},
		PluralCount: catsCount,
	}) // Nick has 2 cats.

	fmt.Println(m2)

	// handle without Plural count, by default always use "Other". "Other" is equvivalent of "defaultMessage" in JS

	m3, _ := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "PersonCats_1-cat",
			Description: "Just about a cat",
			Other:       "{{.Name}} has {{.Count}} cat.",
		},
		TemplateData: map[string]interface{}{
			"Name":  "Nick",
			"Count": 1,
		},
	})

	fmt.Println(m3)

	m4, _ := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "PersonCats_plural-more-cats",
			Description: "Just about more cats",
			Other:       "{{.Name}} has {{.Count}} cats.",
		},
		TemplateData: map[string]interface{}{
			"Name":  "Nick",
			"Count": 4,
		},
	})

	fmt.Println(m4)

	fmt.Println("Loading messages from files ...")

	fmt.Println("Unmarshall ICU messages ... ")
	load_messages_from_file()
	fmt.Println("Unmarshall custom ICU messages ... ")
	load_messages_from_file_new()
}
