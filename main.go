package main

import (
	"github.com/MajsterApp/Backend/application"
	"context"
	"fmt"
)

func main() {
    app := application.New();
    err := app.Start(context.TODO())
    if err != nil {
        fmt.Println("failed to Start app: %w", err)
    }

    fmt.Println("Serv runnig on localhosr:3000")
}


