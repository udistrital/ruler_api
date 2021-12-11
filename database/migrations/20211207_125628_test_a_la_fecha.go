package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/migration"
	"github.com/astaxie/beego/orm"
)

// DO NOT MODIFY
type TestALaFecha_20211207_125628 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &TestALaFecha_20211207_125628{}
	m.Created = "20211207_125628"

	migration.Register("TestALaFecha_20211207_125628", m)
}

type dbSchema struct {
	SchemaName string
}

// Run the migrations
func (m *TestALaFecha_20211207_125628) Up() {
	o := orm.NewOrm()
	var names []dbSchema
	rows, err := o.Raw("SELECT schema_name FROM information_schema.schemata WHERE schema_name = 'ruler'").QueryRows(&names)
	if err != nil {
		panic(err)
	}
	logs.Info(names)
	if rows != 0 {
		return
	}

	file, err := ioutil.ReadFile("../scripts/20211207_125628_test_a_la_fecha.up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		m.SQL(request)
		// do whatever you need with result and error
	}

}

// Reverse the migrations
func (m *TestALaFecha_20211207_125628) Down() {
	file, err := ioutil.ReadFile("../scripts/20211207_125628_test_a_la_fecha.down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		m.SQL(request)
		// do whatever you need with result and error
	}
}
