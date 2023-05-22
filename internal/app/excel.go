package app

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/RB-PRO/KadTG/pkg/KadArbitr"
	"github.com/xuri/excelize/v2"
)

const sshet string = "main"

func saveXlsx(pr KadArbitr.Parse) (filename string, err error) {
	// Создаём файл
	f := excelize.NewFile()
	defer f.Close()

	f.NewSheet(sshet)
	f.DeleteSheet("Sheet1")
	save(f, pr.Data)

	// Сохраняем
	dt := time.Now()
	filename = "KadArbitr от " + dt.Format("15h04m 01-02-2006") + ".xlsx"
	if ErrSave := f.SaveAs(filename); ErrSave != nil {
		return "", ErrSave
	}

	return filename, nil
}

// Сохранить результаты по Avtoto
func save(f *excelize.File, data []KadArbitr.Data) {
	style, err := f.NewStyle(&excelize.Style{Fill: excelize.Fill{Type: "pattern", Pattern: 1, Color: []string{"00FF00"}}}) // `{"fill":{"type":"pattern","color":["#00FF00"],"pattern":1}}`
	if err != nil {
		fmt.Println(err)
	}
	ssheet := "main"
	f.SetCellValue(ssheet, "A1", "Номер дела")
	f.SetCellValue(ssheet, "B1", "Ссылка на дело")
	f.SetCellValue(ssheet, "C1", "Дата")

	f.SetCellValue(ssheet, "E1", "Судья")
	f.SetCellValue(ssheet, "F1", "Инстанция")

	f.SetCellValue(ssheet, "G1", "Истец, Имя")
	f.SetCellValue(ssheet, "H1", "Истец, ИНН")
	f.SetCellValue(ssheet, "I1", "Истец, Адрес")

	f.SetCellValue(ssheet, "K1", "Ответчик, Имя")
	f.SetCellValue(ssheet, "L1", "Ответчик, ИНН")
	f.SetCellValue(ssheet, "M1", "Ответчик, Адрес")

	f.SetCellValue(ssheet, "O1", "Сумма исковых требований")
	f.SetCellValue(ssheet, "P1", "Название последнего файла")
	f.SetCellValue(ssheet, "Q1", "Ссылка на последний файл")

	for index, value := range data {
		f.SetCellValue(ssheet, "A"+strconv.Itoa(index+2), value.Number)
		f.SetCellValue(ssheet, "B"+strconv.Itoa(index+2), value.UrlNumber)
		f.SetCellValue(ssheet, "C"+strconv.Itoa(index+2), value.Date)
		f.SetCellValue(ssheet, "E"+strconv.Itoa(index+2), value.Judge)
		f.SetCellValue(ssheet, "F"+strconv.Itoa(index+2), value.Instance)
		f.SetCellValue(ssheet, "G"+strconv.Itoa(index+2), value.Plaintiff.Name)
		f.SetCellValue(ssheet, "H"+strconv.Itoa(index+2), value.Plaintiff.INN)
		f.SetCellValue(ssheet, "I"+strconv.Itoa(index+2), value.Plaintiff.Adress)
		f.SetCellValue(ssheet, "K"+strconv.Itoa(index+2), value.Respondent.Name)
		f.SetCellValue(ssheet, "L"+strconv.Itoa(index+2), value.Respondent.INN)
		f.SetCellValue(ssheet, "M"+strconv.Itoa(index+2), value.Respondent.Adress)

		f.SetCellValue(ssheet, "O"+strconv.Itoa(index+2), value.Card.Coast)
		if len(value.Card.Slips) != 0 {
			f.SetCellValue(ssheet, "P"+strconv.Itoa(index+2), value.Card.Slips[0].Main.FileName)
			f.SetCellValue(ssheet, "Q"+strconv.Itoa(index+2), value.Card.Slips[0].Main.FileLink)
		}

		if len(strings.TrimSpace(value.Plaintiff.INN)) == 10 {
			f.SetCellStyle(ssheet, "H"+strconv.Itoa(index+2), "H"+strconv.Itoa(index+2), style)
		}
		if len(strings.TrimSpace(value.Respondent.INN)) == 10 {
			f.SetCellStyle(ssheet, "L"+strconv.Itoa(index+2), "L"+strconv.Itoa(index+2), style)
		}
	}
}
