package user

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
)

type TokenResult struct {
    Token string `db:"token" json:"token"`
}

func GetToken(app *pocketbase.PocketBase, userId string) (string, error) {
    var res TokenResult

    query := app.Dao().DB().
        NewQuery(`
            SELECT token
            FROM USERS
            WHERE id = {:id}
        `).
        Bind(dbx.Params{"id": userId})
    
    err := query.One(&res)
    if err != nil {
		return "", err
    }

    return res.Token, nil
}