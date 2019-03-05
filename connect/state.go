package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"

	"github.com/mitchellh/go-homedir"

	"github.com/abrander/garmin-connect"
)

var client = connect.NewClient(
	connect.AutoRenewSession(true),
)

func stateFilename() string {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatalf("Could not detect home directory: %s", err.Error())
	}

	return path.Join(home, ".garmin-connect.json")
}

func loadState() {
	data, err := ioutil.ReadFile(stateFilename())
	if err != nil {
		log.Printf("Could not open state file: %s", err.Error())
		return
	}

	err = json.Unmarshal(data, client)
	if err != nil {
		log.Fatalf("Could not unmarshal state: %s", err.Error())
	}
}

func storeState() {
	b, err := json.MarshalIndent(client, "", "  ")
	if err != nil {
		log.Fatalf("Could not marshal state: %s", err.Error())
	}

	err = ioutil.WriteFile(stateFilename(), b, 0600)
	if err != nil {
		log.Fatalf("Could not write state file: %s", err.Error())
	}
}