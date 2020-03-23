package scraper

import (
	"fmt"
	"strings"

	"strconv"
	"time"

	"github.com/gocolly/colly"
	// "regexp"
)

// CSGOteam holds the information for the csgo team
type CSGOteam struct {
	ID         int    `json:"ID"`
	TeamName   string `json:"team_name"`
	Points     string `json:"points"` // Points need to be int
	Ranking    int    `json:"ranking"`
	URL        string `json:"url"`
	Date       string `json:"date"`
	PlayerList string `json:"player_list"`
	// Player1  string `json:"player_1"`
	// Player2  string `json:"player_2"`
	// Player3  string `json:"player_3"`
	// Player4  string `json:"player_4"`
	// Player5  string `json:"player_5"`
}

// ScrapeHltvTeams scrapes the top teams, needs automation currently.
func ScrapeHltvTeams() []CSGOteam {
	url := "https://www.hltv.org/ranking/teams"
	c := colly.NewCollector()
	var header = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1913.47 Safari/537.36"
	colly.Async(true)
	fmt.Println(url)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
		r.Headers.Set("User-Agent", header)

	})
	var csgoteams = []CSGOteam{}
	c.OnHTML("div.ranked-team", func(e *colly.HTMLElement) {
		e.ForEach("div.ranking-header", func(_ int, e *colly.HTMLElement) {
			playerString := e.ChildText("div.playersLine")
			players := strings.Split(playerString, "\n")
			for i := range players {
				players[i] = strings.TrimSpace(players[i])
			}
			const layoutISO = "2006-01-02" // Format the date for golang
			ranking, err := strconv.Atoi(strings.ReplaceAll(e.ChildText("span.position"), "#", ""))
			if err != nil {
				ranking = 0
			}
			var team = CSGOteam{
				TeamName:   e.ChildText("span.name"),
				Ranking:    ranking,
				Points:     e.ChildText("span.points"),
				PlayerList: strings.Join(players, " "),
				Date:       time.Now().Format(layoutISO),
			}
			// fmt.Println(players)
			csgoteams = append(csgoteams, team)
		})
		// Hard assign the players
	})
	c.Visit(url)
	fmt.Println(csgoteams)
	return csgoteams
}
