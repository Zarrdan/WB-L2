package main

import "fmt"

// last service

type Cashier struct {
	next department
}

func (c *Cashier) Execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (c *Cashier) SetNext(next department) {
	c.next = next
}

// Интерфейс
type department interface {
	Execute(*Patient)
	SetNext(department)
}

// Департамент
type DepartmentBase struct {
	NextDepartment department
}

// Еще один сервис
type Doctor struct {
	next department
}

func (d *Doctor) Execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.Execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.Execute(p)
}

func (d *Doctor) SetNext(next department) {
	d.next = next
}

// сервис
type Medical struct {
	next department
}

func (m *Medical) Execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.Execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.Execute(p)
}

func (m *Medical) SetNext(next department) {
	m.next = next
}

// объект
type Patient struct {
	Name              string
	RegistrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

//  Сервис
type Reception struct {
	next department
}

func (r *Reception) Execute(p *Patient) {
	if p.RegistrationDone {
		fmt.Println("Patient registration already done")
		r.next.Execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.RegistrationDone = true
	r.next.Execute(p)
}

func (r *Reception) SetNext(next department) {
	r.next = next
}
