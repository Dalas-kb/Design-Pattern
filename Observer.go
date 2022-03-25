package main

import (
	"fmt"
	"testing"
)

type observer interface {
	update(state string)
}

type student struct {
	student_name string
}

func (stu student) update(crs string) {
	fmt.Println("student", stu.student_name, "has joined couse", crs)
}

type subject struct {
	_observers []observer
}

func (sub *subject) Attach(obs observer) {
	sub._observers = append(sub._observers, obs)
}

func (sub *subject) Detach(obs observer) {
	index := 0
	for i := range sub._observers {
		if sub._observers[i] == obs {
			index = i
			break
		}
	}
	sub._observers = append(sub._observers[0:index], sub._observers[index+1:]...)
}

func (sub subject) notify(state string) {
	for _, observer := range sub._observers {
		observer.update(state)
	}
}

type course struct {
	subject
	course_name string
}

func (crs course) Setcourse(text string) {
	crs.course_name = text
	crs.subject.notify(text)
}

func TestUpdate(t *testing.T) {

	student := student{"Faouzi"}
	student.update("GLC")
}

func TestAttach(t *testing.T) {

	student := student{"Hamid"}
	student.update("IMB")
	crs := course{}
	crs.Attach(student)
}

func TestDetach(t *testing.T) {

	student := student{"Hery"}
	student.update("GLC")
	crs := course{}
	crs.Attach(student)
	crs.Detach(student)
}

func TestNotify(t *testing.T) {

	student := student{"Rami"}
	crs := course{}
	crs.Attach(student)
	crs.subject.notify("AOC")
	crs.Detach(student)
}

func TestSetcourse(t *testing.T) {

	student := student{"Rami"}

	crs := course{}

	crs.Attach(student)
	crs.Setcourse("COA")
	crs.Detach(student)
}

