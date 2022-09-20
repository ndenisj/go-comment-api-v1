package main

import "fmt"

// App - the struct which contains things like pointers to db connections
type App struct {
}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting up our App")
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
