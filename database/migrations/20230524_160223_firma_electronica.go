package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type FirmaElectronica_20230524_160223 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &FirmaElectronica_20230524_160223{}
	m.Created = "20230524_160223"

	migration.Register("FirmaElectronica_20230524_160223", m)
}

// Run the migrations
func (m *FirmaElectronica_20230524_160223) Up() {
	file, err := ioutil.ReadFile("../scripts/20230524_160223_firma_electronica_up.sql")

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
func (m *FirmaElectronica_20230524_160223) Down() {
	file, err := ioutil.ReadFile("../scripts/20230524_160223_firma_electronica_down.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
