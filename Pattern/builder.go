package main

import "fmt"

/**
Паттерн Builder относится к порождающим паттернам уровня объекта.
Паттерн Builder определяет процесс поэтапного построения сложного продукта.
После того как будет построена последняя его часть, продукт можно использовать.
*/

type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

func getBuilder(builderType string) iBuilder {
	if builderType == "normal" {
		return &normalBuilder{}
	}

	if builderType == "igloo" {
		return &iglooBuilder{}
	}
	return nil
}

type normalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (b *normalBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *normalBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *normalBuilder) setNumFloor() {
	b.floor = 2
}

func (b *normalBuilder) getHouse() house {
	return house{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (b *iglooBuilder) setWindowType() {
	b.windowType = "Snow Window"
}

func (b *iglooBuilder) setDoorType() {
	b.doorType = "Snow Door"
}

func (b *iglooBuilder) setNumFloor() {
	b.floor = 1
}

func (b *iglooBuilder) getHouse() house {
	return house{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

type house struct {
	windowType string
	doorType   string
	floor      int
}

type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) buildHouse() house {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

func main() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")
	normalBuilder.setNumFloor()
	normalBuilder.setWindowType()
	normalBuilder.setDoorType()
	a := normalBuilder.getHouse()

	fmt.Println("Normal House : ", a)
	iglooBuilder.setNumFloor()
	iglooBuilder.setWindowType()
	iglooBuilder.setDoorType()
	i := iglooBuilder.getHouse()
	fmt.Println("Igloo House : ", i)

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	director1 := newDirector(iglooBuilder)
	newHouse := director1.buildHouse()

	fmt.Println(" Director make new House", normalHouse)
	fmt.Println(" Director make new Igloo", newHouse)
}
