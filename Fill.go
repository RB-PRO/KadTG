package KadArbitr

// Заполнить форму для снятия данных
func (core *CoreReq) FillForm() error {
	// А17-5639/2022
	LocatorFamily, ErrorLocator := core.page.Locator("[placeholder=\"фамилия судьи\"]")
	if ErrorLocator != nil {
		return ErrorLocator
	}

	ClickError := LocatorFamily.Click()
	if ClickError != nil {
		return ClickError
	}

	FillError := LocatorFamily.Fill("А17-5639/2022")
	if FillError != nil {
		return FillError
	}

	return nil
}
