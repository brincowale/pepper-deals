package main

import (
	"log"
	"pepper-deals/config"
	"pepper-deals/pepper"
	"pepper-deals/storage"
	"pepper-deals/telegram"
)

func main() {
	cfg := config.ReadConfig()
	t := telegram.New(cfg.TelegramApiKey)
	for _, deal := range pepper.GetNewDeals(cfg).Data {
		if storage.InsertDeal(deal.DealURI) && pepper.Matches(deal, cfg.Filters) {
			err := t.SendMessage(cfg.TelegramChannel, t.CreateMessage(deal))
			if err != nil {
				log.Println(err)
			}
		}
	}
}
