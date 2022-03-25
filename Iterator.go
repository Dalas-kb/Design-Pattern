package main

import "testing"


type AgendaIterator interface {
	FirstMatiere()
	NextMatiere()
	MatiereFaite() bool
	CurrentMatiere() string
}

//agenda c'est ConcreteAggregate
type agenda struct {
	modules []string
}

//GetIterator returne un Iterateur
func (n agenda) GetIterator() AgendaIterator {
	return &Iterator{n, 0}
}

//Iterator c'est le ConcreteIterator
type Iterator struct {
	_agendas agenda
	_index   int
}

// positionner l iterateur a la premiere matiere
func (i *Iterator) FirstMatiere() {
	i._index = 0
}

// se deplacer d 'une matiere dans notre tableau
func (i *Iterator) NextMatiere() {
	i._index++
}

//MatiereFaite explique si la matiere est faite ou non
func (i *Iterator) MatiereFaite() bool {
	return i._index >= len(i._agendas.modules)
}

//CurrentMatiere() retourne la matiere courante
func (i *Iterator) CurrentMatiere() string {
	return i._agendas.modules[i._index]
}

func TestCurrentMatiere(t *testing.T) {
	ag := agenda{[]string{"GLC", "IMB"}}
	iterator := ag.GetIterator()
	v := iterator.CurrentMatiere()
	if v != "GLC" {
		t.Error("le premiere matiere doit etre GLC")
	}
}

func TestMatiereFaite(t *testing.T) {
	ag := agenda{[]string{"GLC"}}
	iterator := ag.GetIterator()
	iterator.NextMatiere()
	v := iterator.MatiereFaite()
	if v == false {
		t.Error("le matiere GLC n a pas ete fait")
	}
}

func TestFirstMatiere(t *testing.T) {
	ag := agenda{[]string{"GLC", "IMB"}}
	iterator := ag.GetIterator()
	iterator.FirstMatiere()
	v := iterator.CurrentMatiere()
	if v != "GLC" {
		t.Error("le matiere GLC doit etre en 1er")
	}
}

func TestNextMatiere(t *testing.T) {
	ag := agenda{[]string{"GLC", "IMB"}}
	iterator := ag.GetIterator()
	iterator.NextMatiere()
	v := iterator.CurrentMatiere()
	if v != "IMB" {
		t.Error("le prochaine matiere apres IMB doit etre GLC ")
	}
}

