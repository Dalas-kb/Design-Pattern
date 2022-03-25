package main

import (
	"fmt"
	"testing"
)

// Interface agendastrategy

type agendastrategy interface {
	Affiche()
}

//coursStrategy est ConcreteStrategy

type coursStrategy struct {
	jour      string
	module    string
	nom_pof   string
	num_salle int
}

func (b coursStrategy) Affiche() {
	fmt.Println("Vous avez cours ")
}

//tdStrategy est ConcreteStrategy

type tdStrategy struct {
	jour      string
	module    string
	nom_pof   string
	num_salle int
}

func (m tdStrategy) Affiche() {
	fmt.Println("\nVous avez td")
}

//agendaContext est Context

type agendaContext struct {
	Strategy agendastrategy
}

func (c agendaContext) Exec() {
	c.Strategy.Affiche()
}

func TestAffiche(t *testing.T) {
	agenda1 := coursStrategy{
		jour:      "lundi",
		module:    "EDP",
		nom_pof:   "Soraya",
		num_salle: 15,
	}

	var ctx agendaContext
	ctx = agendaContext{Strategy: coursStrategy{}}
	ctx.Exec()
	fmt.Println(agenda1)

}
func TestExec(t *testing.T) {

	agenda2 := tdStrategy{
		jour:      "mardi",
		module:    "compilation",
		nom_pof:   "pablo",
		num_salle: 26,
	}
	var ctx agendaContext

	ctx = agendaContext{Strategy: tdStrategy{}}
	ctx.Exec()
	fmt.Println(agenda2)
}

