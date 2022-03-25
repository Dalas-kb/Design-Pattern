/ singleton_agenda
package main

import (
	"fmt"
	"testing"
)

type Agenda struct {
	courses []string
}

var instance *Agenda

func GetInstance() *Agenda {
	if instance == nil {
		instance = &Agenda{}
	}
	return instance
}

func TestGetInstance(t *testing.T) {

	ag := GetInstance()

	ag.courses = append(ag.courses, "IMB")
	ag.courses = append(ag.courses, "GLC")

	ag_prime := GetInstance()
	fmt.Println(ag_prime.courses)
}

