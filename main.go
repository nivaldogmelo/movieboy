package main

import (
        "os"
        tb "gopkg.in/tucnak/telebot.v2"
        "log"
)

func main() {
        var (
                port      = os.Getenv("PORT")
                publicURL = os.Getenv("PUBLIC_URL")
                token     = os.Getenv("TOKEN")
        )

        webhook := &tb.Webhook{
                Listen: ":" + port,
                Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
        }

        pref := tb.Settings{
                Token: token,
                Poller: webhook,
        }

        b, err := tb.NewBot(pref)
        if err != nil {
                log.Fatal(err)
        }

        b.Handle("/hello", func(m *tb.Message) {
                b.Send(m.Sender, "Hi!")
        })
}

