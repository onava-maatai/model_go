package model_go

import (
	"fmt"
	"github.com/go-errors/errors"
	"log"
)

var conectornamess []string
var conectors map[string]ModelConnector

func RegisterModel(conectorname string, conn ModelConnector) {
	if conectors == nil {
		conectors = make(map[string]ModelConnector)
	}
	conectors[conectorname] = conn
}

func Open(conectorname string) (ModelConnector, error) {
	val, ok := conectors[conectorname]
	if !ok {
		return nil, fmt.Errorf("Conector doesn't exists")
	}
	return val, nil
}

/*
set defaults
*/
var default_modelname string
var default_conectionstring string

func SetDefaultConectors(model, conectionstring string) {
	// TODO VERIFY VALIDITY OF MODEL NAMES
	default_conectionstring = conectionstring
	default_modelname = model
}

func OpenDefault() (*ModelConnector, error) {

	log.Printf("Conectores Disponibles \n%+v\nmodel: '%s'\nconnectg string: %s\n", conectors, default_modelname, default_conectionstring)

	val, ok := conectors[default_modelname]
	if !ok {
		return nil, fmt.Errorf("Conector doesn't exists")
	}

	if default_conectionstring == "" {
		return nil, errors.New("Not default setted")
	}
	val.SetConnectionString(default_conectionstring)

	log.Printf("Returging val: '%+v'", val)
	return &val, nil
}
