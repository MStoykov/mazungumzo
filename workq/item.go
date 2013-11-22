package workq

import (
	"fmt"
	"github.com/Vladimiroff/mazungumzo/translator"
)

type Item struct {
	Sender     string
	Message    string
	Translated *string
	Src        string
	Dest       string
}

func (i Item) Work() {
	var err error

	if i.Src == i.Dest {
		*i.Translated = i.Message
	} else {
		*i.Translated, err = translator.Translate(i.Src, i.Dest, i.Message)
		if err != nil {
			*i.Translated = fmt.Sprintf("(untranslated) %s", i.Message)
		}
	}
}
