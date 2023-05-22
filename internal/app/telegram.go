package app

import (
	"fmt"
	"log"
	"strconv"

	"github.com/RB-PRO/KadTG/pkg/KadArbitr"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/playwright-community/playwright-go"
)

func Start() {
	playwright.Install()

	token, ErrorFile := dataFile("token")
	if ErrorFile != nil {
		log.Fatal(ErrorFile)
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Произошла авторизация %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	//

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		log.Println(update.Message.Chat.UserName, "-", update.Message.Text, ">", update.Message.Caption)

		// Игнорируем НЕкоманды
		if !update.Message.IsCommand() {
			// Проверка наличия текста в сообщении
			if update.Message.Text == "" {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Не вижу текста.\nНужно отправить фотографию вместе с текстом."))
				continue
			}

			req, errorunwrap := unwrap(update.Message.Text)
			if errorunwrap != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, errorunwrap.Error()))
				continue
			}
			fmt.Printf("Запрос: %+v\n", req)
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Получил Ваш запрос.")) //fmt.Sprintf("Запрос: %+v\n", req)

			// Создаём ядро
			core, ErrorCore := KadArbitr.NewCore()
			if ErrorCore != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, ErrorCore.Error()))
				continue
			}

			// Заполнение формы поиска
			ErrorReq := core.FillReqestOne(req)
			if ErrorReq != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, ErrorReq.Error()))
				continue
			}

			fmt.Printf("Нажимаем на кнопку поиска\n")
			// Нажимаем на кнопку поиска
			ErrorSearch := core.Search(req)
			if ErrorSearch != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, ErrorSearch.Error()))
				continue
			}

			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Открываю первую страницу Kad.Arbitr"))

			// Получаем настройки
			settings, ErrorSettings := core.Settings()
			if ErrorSettings != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, ErrorSettings.Error()))
				continue
			}
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf(`Я ещё тестовая версия, но уже могу что-то показать.
Дли вашего запроса найдено всего %v записей и я вижу всего %v страниц, на каждой из которых максмум %v записей.
Начинаю парсинг`, settings.DocumentsTotalCount, settings.DocumentsPagesCount, settings.DocumentsPageSize)))
			if core.Screen("screen.jpg") == nil {
				photo := tgbotapi.NewPhoto(update.Message.From.ID, tgbotapi.FilePath("screen.jpg"))
				if _, err = bot.Send(photo); err != nil {
					log.Fatalln(err)
				}
			}

			// Парсим всё
			pr, ErrorAll := core.ParseAll()
			if ErrorAll != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, ErrorAll.Error()))
				continue
			}
			// 			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf(`Всего найдено записей %v.
			// Запускаю парсинг по всем карточкам. Это будет несколько долго.`, len(pr.Data))))

			// for index := range pr.Data {
			// 	fmt.Println("Парсинг каждой карточки", index+1, "из", len(pr.Data))
			// 	bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Парсинг каждой карточки "+strconv.Itoa(index+1)+" из "+strconv.Itoa(len(pr.Data))))
			// 	pr.Data[index].Card, _ = core.ParseCard(pr.Data[index].UrlNumber)
			// }

			// Сохраняем и отравляем ему данные
			filename, ErrorSave := saveXlsx(pr)
			if ErrorSave != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, ErrorSave.Error()))
				continue
			}

			// Отправить все ссылки
			var StrLinks string
			for IndexLink, ValueLink := range pr.Data {
				StrLinks += strconv.Itoa(IndexLink) + ". " + ValueLink.UrlNumber + "\n"
			}
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Найденные ссылки судебных дел:\n"+StrLinks))

			// отправляем файл
			file := tgbotapi.FilePath(filename)
			bot.Send(tgbotapi.NewDocument(update.Message.Chat.ID, file))

			core.Stop()

			fmt.Println(len(pr.Data))

			continue
		}

		switch update.Message.Command() {
		case "example":
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, `Ввод должен быть примерно такой. Некотоыре записи можно опустить, главное пишите цифру точку и пробел перед интересующим Вас параметром
1. [ИНН или компания]; [сторона( "0" - Истец, "1" - Ответчик,"2" - Третье лицо, "3" - Иное лицо)]
2. [судья]; [инстанция]
3. [номер дела]
4. [суд]
5. [Дата регистрации С]
6. [Дата регистрации ДО]
7. [Параметр поиска("a" - Административные,"c" - Гражданские, "b" - Банкротные, "o" - Найти обычным поиском)]`))

			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, `1. ООО М4 Б2Б МАРКЕТПЛЕЙС; 0
2. Снегур А. А.; Суд по интеллектуальным правам
3. СИП-344/2023
5. 14.04.2023
6. 14.04.2023
7. c`))

			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, `1. 7714030726; 1`))
		default:
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Я не знаю такую команду\nПопробуй /start"))
			continue
		}
	}

	//

}
