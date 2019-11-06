package web

import (
	"errors"
	"log"
	"net/http"

	"github.com/alecbcs/gatekey/config"
)

// authenticate checks if the http request has the correct auth tokens.
func authenticate(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
	username, password, authOK := r.BasicAuth()
	if authOK == false {
		return errors.New("Authentication Incorrect")
	}
	if username != config.Global.Authentication.User ||
		password != config.Global.Authentication.Password {
		return errors.New("Authentication Incorrect")
	}
	return nil
}

// NotFound returns an error message for all invalid requests.
// For example, if any unknown URI is used this message is returned.
func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not authorized", 401)
}

// ReportError reports an encountered error back to the user.
func reportError(w http.ResponseWriter, r *http.Request, location string, err error) {
	report := "[ERROR Encountered]: " +
		"[Location: " + location + "] " +
		"[Error: " + err.Error() + "]"
	http.Error(w, report, 401)
	log.Println(report)
}
