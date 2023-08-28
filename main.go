package main

import (
	"database/sql"
	"fmt"
	"os"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "password"
	dbname = "postgres"
)

type Item struct {
	Id string `json:"id"`
	Title string `json:"title"`
	ImageUrl string `json:"image_url"`
	Category string `json:"category"`
	Reviews string `json:"reviews"`
	Price string `json:"price"`
	PreviousPrice string `json:"previous_price"`
	Discount string `json:"discount"`
}

type Items struct {
	Items []Item `json:"employees"`
}

func Connect() error {
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := Connect(); err !=nil {
		log.Fatal(err)
	}

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	
	app.Get("/", getHome)
	app.Get("/feed", getFeed)

	log.Fatal(app.Listen(":5000"))
}


func getFeed(c *fiber.Ctx) error {
		rows, err := db.Query("SELECT id, title, category, image_url, reviews, price, previous_price, discount FROM feed;")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		defer rows.Close()
		result := Items{}

		for rows.Next() {
			item := Item{}
			if err := rows.Scan(&item.Id, &item.Title, &item.Category, &item.ImageUrl, &item.Reviews, &item.Price, &item.PreviousPrice, &item.Discount); err != nil {
				return err
			}
			result.Items = append(result.Items, item)
		}
		return c.Render("index", fiber.Map{
			"Items": result.Items,
		})
}


func getHome(c *fiber.Ctx) error {
	filePath := "files/kadec-bg.png"
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading file: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	// Set response headers for file download
	c.Set("Content-Disposition", "attachment; filename=kadec-bg.png")
	c.Set("Content-Type", "application/octet-stream")

	// Send file content as response
	return c.Send(fileContent)
}

func changeId(id *string) {
	*id = strings.Replace(*id, "magazineluiza.com.br/", "magazinevoce.com.br/magazinepromothunder/", -1)
	
	if (strings.Contains(*id, "amazon.com")) {
		*id = fmt.Sprintf("%s%s", *id, "?tag=promothunder-20&language=pt_BR&ref_=as_li_ss_tl")
	}
}