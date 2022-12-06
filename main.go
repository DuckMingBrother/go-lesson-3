package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("view/*")
	r.GET("/", func(c *gin.Context) {
		images := []string{}
		co := colly.NewCollector()
		co.OnHTML("img", func(e *colly.HTMLElement) {
			// images = append(images, "<img src='"+e.Attr("src")+"'/>")
			images = append(images, e.Attr("src"))
		})
		co.OnScraped(func(r *colly.Response) {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"images": images,
			})
		})
		co.OnError(func(res *colly.Response, err error) {
			fmt.Println(err)
		})
		co.OnResponse(func(res *colly.Response) {
			// fmt.Println(string(res.Body))
		})
		co.OnRequest(func(req *colly.Request) {
			fmt.Println("Visiting", req.URL)
		})
		co.Visit("https://tw.yahoo.com/")

		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// }

		// c.JSON(http.StatusOK, "This is your token")
	})
	r.Run(":8080")
}

type Login struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}
