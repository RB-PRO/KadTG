package tg

import (
	"fmt"
	"log"

	"github.com/RB-PRO/KadArbitr"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start() {

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
			if update.Message.Caption == "" {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Не вижу текста.\nНужно отправить фотографию вместе с текстом."))
				continue
			}

			continue
		}

		switch update.Message.Command() {
		case "example":
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, `1. [ИНН или компания]; [сторона( "0" - Истец, "1" - Ответчик,"2" - Третье лицо, "3" - Иное лицо)]
2. [судья]; [инстанция]
3. [номер дела]
4. [Дата регистрации С]
5. [Дата регистрации ДО]
6. [Параметр поиска("a" - Административные,"c" - Гражданские, "b" - Банкротные, "o" - Найти обычным поиском)]`))

			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, `1. ООО М4 Б2Б МАРКЕТПЛЕЙС; 0
2. Снегур А. А.; Суд по интеллектуальным правам
3. СИП-344/2023
4. 14.04.2023
5. 14.04.2023
6. c`))
			req, errorunwrap := unwrap(update.Message.Text())
			if errorunwrap != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, errorunwrap.Strings()))
			}

			// Создаём ядро
			core, ErrorCore := KadArbitr.NewCore()
			if ErrorCore != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, ErrorCore.Strings()))
			}

			// Заполнение формы поиска
			ErrorReq := core.FillReqestOne(req)
			if ErrorReq != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, ErrorReq.Strings()))
			}

			ErrorSearch := core.Search(req)
			if ErrorSearch != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, ErrorSearch.Strings()))
			}

			pr, ErrorAll := core.ParseAll()
			if ErrorAll != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, ErrorAll.Strings()))
			}

			// Сохраняем и отравляем ему данные
			fmt.Println(len(pr.Data))

		default:
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Я не знаю такую команду\nПопробуй /start"))
			continue
		}
	}

	//

}