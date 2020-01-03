package handler

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/labstack/echo"
	"jkanime-go/internal/model"
	"net/http"
	"strings"
)

func LastAnimeEcho(c echo.Context) error {
	item := model.LastAnimes{}
	var result []model.LastAnimes
	scrapper := colly.NewCollector()
	var id = ""
	var title = ""
	var poster = ""

	scrapper.OnHTML(".portada-box", func(e *colly.HTMLElement) {
		//e.DOM.Find("h2.portada-title a").Attr("href")

		href, _ := e.DOM.Find("h2.portada-title a").Attr("href")

		id = strings.Split(href, "/")[3]
		title,_ = e.DOM.Find("h2.portada-title a").Attr("title")
		poster,_ = e.DOM.Find("a").ChildrenFiltered("img").Attr("src")

		item.Id = id
		item.Title = title
		item.Poster = poster
		//item.Content = getContentInformation(id)
		result = append(result, item)
	})

	scrapper.Visit("https://jkanime.net")

	return c.JSON(http.StatusOK, result)
}

func GetContentInformation(c echo.Context) error {
	id := c.Param("id")
	scrapper := colly.NewCollector()
	var eps_temp_list []string
	var pages []string
	var episodes_last = ""
	var episodes = ""
	var animeType = ""
	var status = ""
	var synopsis = ""
	var gender []string
	content := model.ContentAnime{}

	scrapper.OnHTML("div#container div.left-container div.navigation a", func(e *colly.HTMLElement) {
		//fmt.Println(e.Attr("href"))
		//fmt.Println(e.Text)
		eps_temp_list = append(eps_temp_list, e.Text)
		pages = append(pages, e.Attr("href"))
		//fmt.Println(eps_temp_list)
	})

	scrapper.OnHTML("div#container div.serie-info", func(e *colly.HTMLElement) {
		animeType = strings.Trim(strings.Split(e.DOM.Find("div.info-content div.info-field span.info-value").First().Text(), "\n")[0], " ")
		e.DOM.Find("div.mobile a").Each(func(i int, selection *goquery.Selection) {
			gender = append(gender, selection.Text())
		})
		status = e.DOM.Find("div.info-content div.info-field span.info-value b").Last().Text()
		synopsis = strings.TrimSpace(strings.Split(e.DOM.Find("div.sinopsis-box p.pc").Text(), ":")[1])
	})

	scrapper.Visit("https://jkanime.net/" + id)

	if len(eps_temp_list) < 1 {
		episodes = ""
	} else {
		episodes_last = strings.Split(eps_temp_list[len(eps_temp_list) - 1], "-")[1]
		episodes = strings.Trim(episodes_last, " ")
	}

	content.Type = animeType
	content.Gender = gender
	content.Synopsis = synopsis
	content.Status = status
	content.Episodes = episodes
	content.Pages = pages

	return c.JSON(http.StatusOK, content)
}