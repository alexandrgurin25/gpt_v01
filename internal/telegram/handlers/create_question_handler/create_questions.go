package handlers


type handler struct {
 
}


// if update.Message == nil {
// 				continue
// 			}

// 			gigachatResponses, err := gigachat.Request(update.Message.Text)

// 			if err != nil {
// 				fmt.Println(err)
// 				continue // Отобразить в чате
// 			}

// 			for _, response := range gigachatResponses {
// 				msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
// 				msg.ReplyToMessageID = update.Message.MessageID

// 				if _, err := bot.Send(msg); err != nil {
// 					if _, err := bot.Send(msg); err != nil {
// 						log.Println("Failed to send message via Telegram bot:", err)
// 						continue
// 					}
// 				}
// 			}