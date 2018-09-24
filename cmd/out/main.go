package main

import (
	"encoding/json"
	"os"

	resource "github.com/concourse/mock-resource"
	"github.com/sirupsen/logrus"
)

type OutRequest struct {
	Source  resource.Source    `json:"source"`
	Version resource.Version   `json:"version"`
	Params  resource.PutParams `json:"params"`
}

type OutResponse struct {
	Version  resource.Version         `json:"version"`
	Metadata []resource.MetadataField `json:"metadata"`
}

func main() {
	logrus.SetOutput(os.Stderr)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	decoder := json.NewDecoder(os.Stdin)
	decoder.DisallowUnknownFields()

	var req OutRequest
	err := decoder.Decode(&req)
	if err != nil {
		logrus.Errorf("invalid payload: %s", err)
		os.Exit(1)
		return
	}

	if len(os.Args) < 2 {
		logrus.Errorf("source path not specified")
		os.Exit(1)
		return
	}

	logrus.Printf("pushing version: %s", req.Params.Version)

	json.NewEncoder(os.Stdout).Encode(OutResponse{
		Version:  resource.Version{Version: req.Params.Version},
		Metadata: []resource.MetadataField{},
	})
}
