package workq

import (
	"fmt"
	"time"

	"github.com/Vladimiroff/mazungumzo/translator"
)

type Item struct {
	Time       time.Time
	Sender     string
	Message    string
	Translated string
	Src        string
	Dest       string
	Done       chan bool // TODO: This could be empty
}

func (i *Item) Translate() {
	var err error

	if i.Src == i.Dest {
		i.Translated = i.Message
	} else {
		i.Translated, err = translator.Translate(i.Src, i.Dest, i.Message)
		if err != nil {
			i.Translated = fmt.Sprintf("(untranslated) %s", i.Message)
		}
	}

	i.Done <- true
	close(i.Done)
}
