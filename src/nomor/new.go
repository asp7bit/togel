package nomor

import (
	"log"
	"net/http"
	"database/sql"
	"github.com/drone/routes"
	"github.com/rtulus/togel/src/conf"
)
var query preparedQueryMessage

type preparedQueryMessage struct {
	selectUser	*sql.Stmt
}

func New(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")
	

	dbCore := conf.DB.Core
	// prepared SQL statement
	var errPrepared error
	query.selectUser, errPrepared = dbCore.Prepare(`
		SELECT 
			user_id, full_name
		FROM 
			ws_user
		WHERE 
			user_id = $1`)
	if errPrepared != nil {
		log.Printf(errPrepared.Error())
		return	
	}

	var response Response

	userId := 2800
	errSelect := query.selectUser.QueryRow(userId).Scan(
			&response.UserId,
			&response.FullName)

	if errSelect != nil {
		log.Println(errSelect)
		return
	}

	w.WriteHeader(http.StatusOK)
	routes.ServeJson(w, response)

	return

}
