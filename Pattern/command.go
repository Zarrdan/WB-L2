package main

import "fmt"

/**
Команда — это поведенческий паттерн проектирования, который превращает запросы в объекты,
позволяя передавать их как аргументы при вызове методов, ставить запросы в очередь,
логировать их, а также поддерживать отмену операций.
*/

// команда
type button struct {
	command command
}

// функция нажатия кнпоки
func (b *button) press() {
	b.command.execute()
}

// интерфейс команды
type command interface {
	execute()
}

//
type onCommand struct {
	device device
}

// команда вкл
func (c *onCommand) execute() {
	c.device.on()
}

type offCommand struct {
	device device
}

// команда выкл
func (c *offCommand) execute() {
	c.device.off()
}

type device interface {
	on()
	off()
}
type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

func main() {
	tv := &tv{}
	fmt.Println(tv)

	onCommand := &onCommand{
		device: tv,
	}
	fmt.Println(onCommand)

	offCommand := &offCommand{
		device: tv,
	}

	onButton := &button{
		command: onCommand,
	}
	fmt.Println(onButton)
	onButton.press()

	offButton := &button{
		command: offCommand,
	}
	offButton.press()
}
