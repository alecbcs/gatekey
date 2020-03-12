package web

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/alecbcs/gatekey/config"
	"github.com/alecbcs/gatekey/database"
	"github.com/alecbcs/gatekey/token"
)

// Report authenticates POST requests from foriegn agents with temp tokens
// against the internal database, then copies the POST data to temp storage.
// Report Syntax: http://server/report/TOKEN
// Report Example: curl -F "upload=@FILENAME" http://server/report/TOKEN
func Report(w http.ResponseWriter, r *http.Request) {
	token, err := checkAuth(w, r)        // Check token authentication
	if err != nil || token.Value == "" { // If token is incorrect or nil return not authorized.
		http.Error(w, "Not authorized", 401)
		log.Println(err)
		return
	}
	file, _, err := r.FormFile("upload")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	err = os.MkdirAll(config.Global.Relay.TempFileLocation, os.ModePerm)
	if err != nil {
		reportError(w, r, "web.Report", err)
		return
	}
	tempFile, err := ioutil.TempFile(config.Global.Relay.TempFileLocation, "upload")
	if err != nil {
		reportError(w, r, "web.Report", err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		reportError(w, r, "web.Report", err)
	}
	_, err = tempFile.Write(fileBytes)
	if err != nil {
		reportError(w, r, "web.Report", err)
	}
	defer tempFile.Close()

	fmt.Fprintf(w, "Successfully Uploaded File\n")

	// Generate
	client := &http.Client{}
	req, err := http.NewRequest("POST", config.Global.Relay.Location, tempFile)
	req.SetBasicAuth(config.Global.Relay.User, config.Global.Relay.Password)

	res, err := client.Do(req)
	if err != nil {
		reportError(w, r, "web.Relay", err)
	}

	if res != nil {
		fmt.Println(res)
		defer res.Body.Close()
	}

	defer os.Remove(tempFile.Name())
}

func checkAuth(w http.ResponseWriter, r *http.Request) (token.Token, error) {
	args := strings.Split(r.RequestURI, "/")
	passcode := args[2] // Parse token from url.

	db := database.Open(config.Global.Database.Location)
	defer db.Close()

	authKey, err := database.Get(db, passcode) // If error invalidate auth for security.
	if err != nil {
		return token.Token{}, err
	}
	if passcode == authKey.Value && authKey.Value != "" {
		err = database.Remove(db, authKey)
		if err != nil {
			return token.Token{}, err
		}
		return authKey, nil
	}
	return token.Token{}, errors.New("Key Not Found in Database")
}
