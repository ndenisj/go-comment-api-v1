package main

import (
	"fmt"
	"github.com/ndenisj/go-comment-api-v1/internal/comment"
	"github.com/ndenisj/go-comment-api-v1/internal/database"
	transportHTTP "github.com/ndenisj/go-comment-api-v1/internal/transport/http"
	"net/http"
)

// App - the struct which contains things like pointers to db connections
type App struct {
}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting up our App")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to setup server!")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go Rest API")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our Rest API")
		fmt.Println(err)
	}

}
