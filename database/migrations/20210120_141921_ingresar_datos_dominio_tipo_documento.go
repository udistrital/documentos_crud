package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type IngresarDatosDominioTipoDocumento_20210120_141921 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &IngresarDatosDominioTipoDocumento_20210120_141921{}
	m.Created = "20210120_141921"

	migration.Register("IngresarDatosDominioTipoDocumento_20210120_141921", m)
}

// Run the migrations
func (m *IngresarDatosDominioTipoDocumento_20210120_141921) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210120_141921_ingresar_datos_dominio_tipo_documento_up.sql")

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
func (m *IngresarDatosDominioTipoDocumento_20210120_141921) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210120_141921_ingresar_datos_dominio_tipo_documento_down.sql")

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
