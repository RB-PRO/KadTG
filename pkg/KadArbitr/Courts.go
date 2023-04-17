package KadArbitr

// Получить список судов в мапу Couters
func (core *CoreReq) ParseCouters() error {

	//var ErrorURL error
	couters := make(map[string]string)

	entries, err := core.page.QuerySelectorAll("select[name='Courts'] > option[value]")
	if err != nil {
		return err // could not get entries
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
