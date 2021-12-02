package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ModificarTipoDocumentoNuxeo_20210923_143814 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ModificarTipoDocumentoNuxeo_20210923_143814{}
	m.Created = "20210923_143814"

	migration.Register("ModificarTipoDocumentoNuxeo_20210923_143814", m)
}

// Run the migrations
func (m *ModificarTipoDocumentoNuxeo_20210923_143814) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210923_143814_modificar_tipo_documento_nuxeo_up.sql")

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
func (m *ModificarTipoDocumentoNuxeo_20210923_143814) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210923_143814_modificar_tipo_documento_nuxeo_down.sql")

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
