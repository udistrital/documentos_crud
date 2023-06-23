package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type DocumentoIdFirma_20230623_144206 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &DocumentoIdFirma_20230623_144206{}
	m.Created = "20230623_144206"

	migration.Register("DocumentoIdFirma_20230623_144206", m)
}

// Run the migrations
func (m *DocumentoIdFirma_20230623_144206) Up() {
	file, err := ioutil.ReadFile("../scripts/20230623_144206_documento_id_firma_up.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}

// Reverse the migrations
func (m *DocumentoIdFirma_20230623_144206) Down() {
	file, err := ioutil.ReadFile("../scripts/20230623_144206_documento_id_firma_down.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
