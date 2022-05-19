package main

import (
	"fmt"
	"strings"
)

/**
Фасад — это структурный паттерн проектирования,
Предоставляет единый интерфейс к группе интерфейсов подсистемы. Определяет высокоуровеный интерфейс,
делая подсистему проще для использования

Преимущества и недостатки :
Изолирует клиентов от компонентов сложной подсистемы.
Фасад рискует стать божественным объектом, привязанным ко всем классам программы.
*/

type Engine struct {
}

type Transmissiom struct {
}

type Fuel struct {
}

func (e *Engine) checkEngine() string {
	return "engine is good"
}

func (t *Transmissiom) checkTrandmission() string {
	return "transmission is good"
}

func (f *Fuel) checkFuel() string {
	return "fuel is good"
}

type startCar struct {
	engine       Engine
	fuel         Fuel
	transmissiom Transmissiom
}

func newCar() *startCar {
	return &startCar{
		engine:       Engine{},
		fuel:         Fuel{},
		transmissiom: Transmissiom{},
	}

}
func (s *startCar) Start() string {
	result := []string{
		s.engine.checkEngine(),
		s.fuel.checkFuel(),
		s.transmissiom.checkTrandmission(),
	}
	return strings.Join(result, "\n")
}

func main() {
	car := newCar()
	result := car.Start()
	fmt.Println(result)
}
