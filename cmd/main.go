package main

import (
	"log"
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"

	hellocron "spotify-recommend-back/internal/cron"
)

func main() {
    app := pocketbase.New()

    // serves static files from the provided public dir (if exists)
    app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
        
        hellocron.HelloWorld(app)

        e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
        return nil
    })

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}