package user

import (
	"github.com/pocketbase/pocketbase"
)

type UserResult struct {
    Id       string `db:"id" json:"id"`
    Username string `db:"username" json:"username"`
	Description string `db:"description" json:"description"`
}

func GetUser(app *pocketbase.PocketBase) ([]UserResult, error) {
    res := []UserResult{}

    query := app.Dao().DB().
        NewQuery(`
			SELECT u.id, d.username, d.description
			FROM USERS u
			INNER JOIN DIM_USERS d ON u.id = d.id
		`)
    
    err := query.All(&res) 
    if err != nil {
        return nil, err
    }

    return res, nil
}