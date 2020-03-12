package web

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/alecbcs/gatekey/config"
	"github.com/alecbcs/gatekey/database"
	"github.com/alecbcs/gatekey/token"
)

// Create generates a new one time token for use in the database.
// Create Syntax: http://server/create/?jid=SOMETHING&mid=SOMETHINGELSE with authentication header credentials.
// Curl example: curl -u "USERNAME":"PASSWORD" 'http://SERVER:PORT/create/?jid=20&mid=10'
func Create(w http.ResponseWriter, r *http.Request) {
	err := authenticate(w, r)
	if err != nil {
		http.Error(w, "Not authorized", 401)
		return
	}
	db := database.Open(config.Global.Database.Location)

	urldata, err := url.Parse(r.RequestURI)
	if err != nil {
		reportError(w, r, "web.Create", errors.New("Invalid Command Arguments"))
	}

	args, _ := url.ParseQuery(urldata.RawQuery)
	if args["mid"] == nil || args["jid"] == nil {
		reportError(w, r, "web.Create", errors.New("Invalid Command Arguments"))
	}

	result := token.Token{
		Value:     token.GenValue(config.Global.Tokens.Length),
		MachineID: args["mid"][0],
		JobID:     args["jid"][0],
	}
	sucess, err := database.Add(db, result)
	for !sucess {
		if err != nil {
			reportError(w, r, "database.Add", err)
			return
		}
		result.Value = token.GenValue(config.Global.Tokens.Length)
		sucess, err = database.Add(db, result)
	}
	fmt.Fprintln(w, result.Value)
}
