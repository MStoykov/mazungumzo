package server

import (
	"sync"
	"time"
)

type LanguagePool struct {
	mutex     sync.Mutex
	languages map[string]*Language
}

func NewLanguagePool() *LanguagePool {
	p := new(LanguagePool)
	p.languages = make(map[string]*Language)
	return p
}

func (lp *LanguagePool) Remove(l *Language) {
	lp.mutex.Lock()
	defer lp.mutex.Unlock()

	delete(lp.languages, l.Name)
}

func (lp *LanguagePool) Get(name string) *Language {
	lp.mutex.Lock()
	defer lp.mutex.Unlock()

	language, ok := lp.languages[name]
	if !ok {
		language = NewLanguage(name)
		lp.languages[language.Name] = language
	}
	return language
}

func (lp *LanguagePool) Broadcast(sender *Client, message []byte) {
	lp.mutex.Lock()
	defer lp.mutex.Unlock()

	now := time.Now()
	for _, language := range lp.languages {
		language.Send(sender, now, message)
	}
}
