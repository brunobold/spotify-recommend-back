package hellocron

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/tools/cron"

	user "spotify-recommend-back/internal/database/user"
	"spotify-recommend-back/internal/spotify"
)

// HelloWorld prints "Hello, World!" to the log
func HelloWorld(app *pocketbase.PocketBase) {
    scheduler := cron.New()
    scheduler.MustAdd("hello_world", "*/1 * * * *", func() {
        log.Println("Hello, World!")

        users, err := user.GetUser(app)
        if err != nil {
            log.Println("Error fetching users:", err)
            return
        }
        log.Printf("Fetched users: %+v\n", users)

		token, err := user.GetToken(app, "")
        if err != nil {
            log.Println("Error fetching token:", err)
            return
        }
        log.Printf("Fetched token: %+v\n", token)

        // profile, err := spotify.GetProfile(token)
        // if err != nil {
        //     log.Println("Error fetching profile:", err)
        //     return
        // }
        // log.Println("Fetched profile:", profile)

        listening, err := spotify.GetListeningData(token)
        if err != nil {
            log.Println("Error fetching profile:", err)
            return
        }
        log.Println("\nFetched listening:", listening)
    })

    scheduler.Start()
}