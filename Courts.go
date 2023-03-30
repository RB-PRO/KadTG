package KadArbitr

import (
	"log"
)

func (core *CoreReq) ParseCouters() error {

	//var ErrorURL error
	couters := make(map[string]string)

	if _, err := core.page.Goto("https://kad.arbitr.ru/"); err != nil {
		return err
	}

	entries, err := core.page.QuerySelectorAll("select[name='Courts'] > option[value]")
	if err != nil {
		log.Fatalf("could not get entries: %v", err)
	}
	for _, entry := range entries {
		Text, ErrorText := entry.TextContent()
		if ErrorText == nil {
			Value, ErrorValue := entry.GetAttribute("value")
			if ErrorValue == nil {
				couters[Text] = Value
			}
		}
	}

	// Сохраняем структуру в ядро парсингаы
	core.Couters = couters

	return nil
}
