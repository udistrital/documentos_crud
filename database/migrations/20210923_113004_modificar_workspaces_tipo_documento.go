package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ModificarWorkspacesTipoDocumento_20210923_113004 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ModificarWorkspacesTipoDocumento_20210923_113004{}
	m.Created = "20210923_113004"

	migration.Register("ModificarWorkspacesTipoDocumento_20210923_113004", m)
}

// Run the migrations
func (m *ModificarWorkspacesTipoDocumento_20210923_113004) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210923_113004_modificar_workspaces_tipo_documento_up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}

// Reverse the migrations
func (m *ModificarWorkspacesTipoDocumento_20210923_113004) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210923_113004_modificar_workspaces_tipo_documento_down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}
