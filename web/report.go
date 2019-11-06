package web

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/alecbcs/gatekey/config"
	"github.com/alecbcs/gatekey/database"
	"github.com/alecbcs/gatekey/token"
)

// Report authenticates POST requests from foriegn agents with temp tokens
// against the internal database, then copies the POST data to temp storage.
func Report(w http.ResponseWriter, r *http.Request) {
	token, err := checkAuth(w, r)
	if err != nil || token.Value == "" {
		http.Error(w, "Not authorized", 401)
		log.Println(err)
		return
	}
	file, _, err := r.FormFile("myFile")
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
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func checkAuth(w http.ResponseWriter, r *http.Request) (token.Token, error) {
	w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
	username, password, authOK := r.BasicAuth()
	if authOK == false {
		return token.Token{}, errors.New("authOk Failed")
	}

	db := database.Open(config.Global.Database.Location)
	defer db.Close()

	authKey, err := database.Get(db, password)
	if err != nil {
		return token.Token{}, err
	}
	if authKey.Value == "" {
		return token.Token{}, errors.New("Key Not Found in Database")
	}
	if password == authKey.Value && username == authKey.MachineID {
		return authKey, nil
	}
	return token.Token{}, errors.New("MachineID doesn't match")
}
