package main

import (
	"database/sql"
	"fmt"
	"os"
	"log"

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
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	
	app.Get("/", getHome)
	app.Get("/kadec", getKadec)
	app.Get("/kadecfiles/:filename", getKadecDownload)

	app.Get("/thunder", getThunder)
	app.Get("/thunderfiles/:filename", getThunderDownload)

	log.Fatal(app.Listen(":5000"))
}

func getHome(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}


func getKadec(c *fiber.Ctx) error {
	// Read the list of files and subdirectories in the specified folder
	files, err := os.ReadDir("static/kadec")
	if err != nil {
		log.Fatalf("Error reading folder: %v", err)
		c.SendString("ERROR")
	}

	var folderfiles []string

	// Iterate through the list of file info objects
	for _, file := range files {
		// Check if the object is a regular file (not a directory)
		if !file.IsDir() {
			folderfiles = append(folderfiles, file.Name())
		}
	}

	return c.Render("kadec", fiber.Map{
		"Files": folderfiles,
	})
}


func getKadecDownload(c *fiber.Ctx) error {
	filename := c.Params("filename")
	fileContent, err := os.ReadFile(fmt.Sprintf("static/kadec/%s", filename))
	if err != nil {
		log.Printf("Error reading file: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	// Set response headers for file download
	c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Set("Content-Type", "application/octet-stream")

	// Send file content as response
	return c.Send(fileContent)
}



func getThunder(c *fiber.Ctx) error {
	// Read the list of files and subdirectories in the specified folder
	files, err := os.ReadDir("static/thunder")
	if err != nil {
		log.Fatalf("Error reading folder: %v", err)
		c.SendString("ERROR")
	}

	var folderfiles []string

	// Iterate through the list of file info objects
	for _, file := range files {
		// Check if the object is a regular file (not a directory)
		if !file.IsDir() {
			folderfiles = append(folderfiles, file.Name())
		}
	}

	return c.Render("thunder", fiber.Map{
		"Files": folderfiles,
	})
}


func getThunderDownload(c *fiber.Ctx) error {
	filename := c.Params("filename")
	fileContent, err := os.ReadFile(fmt.Sprintf("static/thunder/%s", filename))
	if err != nil {
		log.Printf("Error reading file: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	// Set response headers for file download
	c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Set("Content-Type", "application/octet-stream")

	// Send file content as response
	return c.Send(fileContent)
}