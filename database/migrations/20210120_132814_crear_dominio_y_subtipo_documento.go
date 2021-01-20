package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearDominioYSubtipoDocumento_20210120_132814 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearDominioYSubtipoDocumento_20210120_132814{}
	m.Created = "20210120_132814"

	migration.Register("CrearDominioYSubtipoDocumento_20210120_132814", m)
}

// Run the migrations
func (m *CrearDominioYSubtipoDocumento_20210120_132814) Up() {
	//  use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210120_132814_crear_dominio_y_subtipo_documento_up.sql")

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
func (m *CrearDominioYSubtipoDocumento_20210120_132814) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210120_132814_crear_dominio_y_subtipo_documento_down.sql")

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
