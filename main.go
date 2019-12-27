package main

import (
        "os"
        tb "gopkg.in/tucnak/telebot.v2"
        "log"
        "encoding/json"
        "fmt"
        "net/http"
)

type Movie struct {
        Title        string   `json:"Title"`
        Year         string   `json:"Year"`
        Rated        string   `json:"Rated"`
        Released     string   `json:"Released"`
        Runtime      string   `json:"Runtime"`
        Genre        string   `json:"Genre"`
        Director     string   `json:"Director"`
        Writer       string   `json:"Writer"`
        Actors       string   `json:"Actors"`
        Plot         string   `json:"Plot"`
        Language     string   `json:"Language"`
        Country      string   `json:"Country"`
        Awards       string   `json:"Awards"`
        Poster       string   `json:"Poster"`
        Ratings      []Rating `json:"Ratings"`
        Metascore    string   `json:"Metascore"`
        ImdbRating   string   `json:"imdbRating"`
        ImdbVotes    string   `json:"imdbVotes"`
        ImdbID       string   `json:"imdbID"`
        Type         string   `json:"Type"`
        TotalSeasons string   `json:"totalSeasons"`
        Response     string   `json:"Response"`
}

type Rating struct {
        Source string `json:"Source"`
        Value  string `json:"Value"`
}

func searchMovie(movie string) string {
        endpoint := fmt.Sprintf("www.omdbapi.com/?apikey=18506062&t=%s", movie)

        resp, err := http.Get(endpoint)
        if err != nil {
                log.Fatal(err)
        }

        result := &Movie{}
        err = json.NewDecoder(resp.Body).Decode(&result)
        if err != nil {
                log.Fatal(err)
        }

        return result.Plot
}

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
                b.Send(m.Sender, "You entered "+m.Payload)
        })

        b.Handle("/movie", func(m *tb.Message) {
                finalPlot := searchMovie(m.Payload)
                b.Send(m.Sender, "Plot: "+finalPlot)
        })

        b.Start()
}

