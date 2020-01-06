package web

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/alecbcs/gatekey/config"
	"github.com/alecbcs/gatekey/database"
	"github.com/alecbcs/gatekey/token"
)

// Create generates a new one time token for use in the database.
// Create Syntax: http://server/create/ with authentication header credentials.
func Create(w http.ResponseWriter, r *http.Request) {
	err := authenticate(w, r)
	if err != nil {
		http.Error(w, "Not authorized", 401)
		return
	}
	db := database.Open(config.Global.Database.Location)
	args := strings.Split(r.RequestURI, "/")
	if len(args) != 4 {
		reportError(w, r, "web.Create", errors.New("Invalid Command Arguments"))
	}
	result := token.Token{
		Value:     token.GenValue(config.Global.Tokens.Length),
		MachineID: args[3],
		JobID:     args[2],
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
