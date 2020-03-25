package scraper

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

// CSGOteam holds the information for the csgo team
type CSGOteam struct {
	ID         int    `json:"ID"`
	TeamName   string `json:"teamName"`
	Points     int    `json:"points"` // Points need to be int
	Ranking    int    `json:"ranking"`
	URL        string `json:"url"`
	Date       string `json:"date"`
	PlayerList string `json:"playerList"`
}

// ScrapeHltvTeams scrapes the top teams, needs automation currently.
func ScrapeHltvTeams() []CSGOteam {
	url := "https://www.hltv.org/ranking/teams"
	c := colly.NewCollector()
	var header = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1913.47 Safari/537.36"
	colly.Async(true)
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
			// Format the points
			re := regexp.MustCompile("[0-9]+")
			points, err := strconv.Atoi((re.FindAllString(e.ChildText("span.points"), 1)[0]))
			if err != nil {
				points = 0
			}
			var team = CSGOteam{
				TeamName:   e.ChildText("span.name"),
				Ranking:    ranking,
				Points:     points,
				PlayerList: strings.Join(players, " "),
				Date:       time.Now().Format(layoutISO),
			}
			csgoteams = append(csgoteams, team)
		})
		// Hard assign the players
	})
	c.Visit(url)
	return csgoteams
}
