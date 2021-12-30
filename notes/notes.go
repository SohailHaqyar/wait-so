package notes

import (
	"github.com/SohailHaqyar/wait-so/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
}

func SetupRoutes(app *fiber.App) {
	app.Get("/notes", getNotes)
	app.Post("/notes", addNote)
	app.Delete("/notes/:id", deleteNote)
	app.Put("/notes/:id", updateNote)
}

func getNotes(c *fiber.Ctx) error {
	var notes []Note
	db := database.DatabaseConfig
	db.Find(&notes)
	return c.JSON(notes)
}

func addNote(c *fiber.Ctx) error {
	db := database.DatabaseConfig
	note := new(Note)

	if err := c.BodyParser(note); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	db.Create(&note)
	return c.JSON(note)
}

// Delete a note with the id param
func deleteNote(c *fiber.Ctx) error {
	db := database.DatabaseConfig
	id := c.Params("id")
	var note Note
	db.First(&note, id)
	// if not found return error
	if note.ID == 0 {
		return c.Status(fiber.StatusNotFound).SendString("Note not found")
	}

	db.Delete(&note)
	return c.SendString("Note with id " + id + " deleted")
}

// Update a note with the id param
func updateNote(c *fiber.Ctx) error {
	db := database.DatabaseConfig
	id := c.Params("id")

	body := new(Note)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	var note Note

	db.First(&note, id)
	if note.ID == 0 {
		return c.Status(fiber.StatusNotFound).SendString("Note not found")
	}

	if body.Title != "" {
		note.Title = body.Title
	}

	if body.Content != "" {
		note.Content = body.Content
	}

	db.Save(&note)
	return c.JSON(fiber.Map{
		"message": "Note updated",
		"note":    note,
	})

}
