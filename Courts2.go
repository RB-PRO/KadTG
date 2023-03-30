package KadArbitr

import (
	"fmt"
	"log"

	"github.com/playwright-community/playwright-go"
)

func Couters2() (map[string]string, error) {

	//var ErrorURL error
	couters := make(map[string]string)

	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	if _, err = page.Goto("https://kad.arbitr.ru/"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	entries, err := page.QuerySelectorAll("select[name='Courts'] > option[value]")
	if err != nil {
		log.Fatalf("could not get entries: %v", err)
	}
	for i, entry := range entries {
		//fmt.Printf("%d: %s\n", i+1, title)
		fmt.Print(i, " ")
		fmt.Println(entry.TextContent())
	}
	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
	return couters, nil
}
