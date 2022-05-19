package main

import "fmt"

//_______________________________________________________________________
// Интерфейс стратегии

type state interface {
	AddItem(int) error
	RequestItem() error
	InsertMoney(money int) error
	DispenseItem() error
}

//_______________________________________________________________________
// Конкретная стратегия

type noItemState struct {
	vendingMachine *vendingMachine
}

func (i *noItemState) RequestItem() error {
	return fmt.Errorf("item out of stock")
}

func (i *noItemState) AddItem(count int) error {
	i.vendingMachine.incrementItemCount(count)
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	return nil
}

func (i *noItemState) InsertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}
func (i *noItemState) DispenseItem() error {
	return fmt.Errorf("item out of stock")
}

//_______________________________________________________________________
// Конкретная стратегия

type hasItemState struct {
	vendingMachine *vendingMachine
}

func (i *hasItemState) RequestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return fmt.Errorf("no item present")
	}
	fmt.Printf("Item requestd\n")
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	return nil
}

func (i *hasItemState) AddItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.vendingMachine.incrementItemCount(count)
	return nil
}

func (i *hasItemState) InsertMoney(money int) error {
	return fmt.Errorf("please select item first")
}
func (i *hasItemState) DispenseItem() error {
	return fmt.Errorf("please select item first")
}

//_______________________________________________________________________
// Конкретная стратегия
type itemRequestedState struct {
	vendingMachine *vendingMachine
}

func (i *itemRequestedState) RequestItem() error {
	return fmt.Errorf("item already requested")
}

func (i *itemRequestedState) AddItem(count int) error {
	return fmt.Errorf("item Dispense in progress")
}

func (i *itemRequestedState) InsertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("inserted money is less. Please insert %d", i.vendingMachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}
func (i *itemRequestedState) DispenseItem() error {
	return fmt.Errorf("please insert money first")
}

//_______________________________________________________________________
// Конкретная стратегия

type hasMoneyState struct {
	vendingMachine *vendingMachine
}

func (i *hasMoneyState) RequestItem() error {
	return fmt.Errorf("item dispense in progress")
}

func (i *hasMoneyState) AddItem(count int) error {
	return fmt.Errorf("item dispense in progress")
}

func (i *hasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}
func (i *hasMoneyState) DispenseItem() error {
	fmt.Println("Dispensing Item")
	i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.setState(i.vendingMachine.hasItem)
	}
	return nil
}

//_______________________________________________________________________
// Контекст

type vendingMachine struct {
	hasItem       state
	itemRequested state
	hasMoney      state
	noItem        state

	currentState state

	itemCount int
	itemPrice int
}

func NewVendingMachine(itemCount, itemPrice int) *vendingMachine {
	v := &vendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	hasItemState := &hasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &itemRequestedState{
		vendingMachine: v,
	}
	hasMoneyState := &hasMoneyState{
		vendingMachine: v,
	}
	noItemState := &noItemState{
		vendingMachine: v,
	}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

func (v *vendingMachine) RequestItem() error {
	return v.currentState.RequestItem()
}

func (v *vendingMachine) AddItem(count int) error {
	return v.currentState.AddItem(count)
}

func (v *vendingMachine) InsertMoney(money int) error {
	return v.currentState.InsertMoney(money)
}

func (v *vendingMachine) DispenseItem() error {
	return v.currentState.DispenseItem()
}

func (v *vendingMachine) setState(s state) {
	v.currentState = s
}

func (v *vendingMachine) incrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.itemCount = v.itemCount + count
}
