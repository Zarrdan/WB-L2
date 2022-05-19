package main

import (
	"fmt"
)

// Паттерн Фабричный метод является пораждающим паттерном. Он определяет общий интерфейс поведения
// для создаваемых объектов. Решает проблему дополнения бизнес-логики.

// Плюсы:
// 1. Избавляет от привязки к конкретному типу объекта при помощи конструктора;
// 2. Упрощает добавление новых объектов от базового класса;
// 3. Реализует принцип open-closed.

// Минусы:
// 1. Может привести к созданию больших иерапрхий объектов (большое количество структур, которые сложны при сопровождении);
// 2. Появляется один главный конструктор, к которому будет привязана вся логика программы.

const (
	ServerType           = "server"
	PersonalComputerType = "personal"
	LaptopType           = "laptop"
)

type Computer interface {
	GetType() string
	PrintDetails()
}

// +++ Инициализация фабричного метода
// При возникновении нового типа электроники, достаточно добавить
// инициализацию, в соответствии с интерфейсом.
func New(typeName string) Computer {
	switch typeName {
	default:
		fmt.Printf("%s Несуществующий тип объекта.\n", typeName)
		return nil
	case ServerType:
		return NewServer()
	case PersonalComputerType:
		return NewPersonalComputer()
	case LaptopType:
		return NewLaptop()
	}
}

// ---

// +++ Инициализация сервера
type Server struct {
	Type   string
	Core   int
	Memory int
}

// Создание базового конструктора для фабричного метода
func NewServer() Computer {
	return Server{
		Type:   ServerType,
		Core:   16,
		Memory: 256,
	}
}

func (pc Server) GetType() string {
	return pc.Type
}

func (pc Server) PrintDetails() {
	fmt.Printf("%s Core: [%d] Memory: [%d]\n", pc.Type, pc.Core, pc.Memory)
}

// ---

// +++ Инициализация персонального компьютера
type PersonalComputer struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

// Создание базового конструктора для фабричного метода
func NewPersonalComputer() Computer {
	return PersonalComputer{
		Type:    PersonalComputerType,
		Core:    8,
		Memory:  16,
		Monitor: true,
	}
}

func (pc PersonalComputer) GetType() string {
	return pc.Type
}

func (pc PersonalComputer) PrintDetails() {
	fmt.Printf("%s Core: [%d] Memory: [%d] Monitor: [%v]\n",
		pc.Type, pc.Core, pc.Memory, pc.Monitor)
}

// ---

// +++ Инициализация ноутбука
type Laptop struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

// Создание базового конструктора для фабричного метода
func NewLaptop() Computer {
	return Laptop{
		Type:    LaptopType,
		Core:    4,
		Memory:  8,
		Monitor: true,
	}
}

func (pc Laptop) GetType() string {
	return pc.Type
}

func (pc Laptop) PrintDetails() {
	fmt.Printf("%s Core: [%d] Memory: [%d] Monitor: [%v]\n",
		pc.Type, pc.Core, pc.Memory, pc.Monitor)
}
