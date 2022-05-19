package main

import "fmt"

// Паттерн Стратегия является поведенческим паттерном. Он определяет схожие алгоритмы
// и помещает каждый из них в отдельную структуру. После чего, алгоритмы могут
// взаимодействовать в исполняемой программе.
// Решает проблему часто расширширяющихся и изменяющихся алгоритмов, путем
// выноса их в собственный объект. Применяется в случае необходимости использования
// разных вариантов одного алгоритма внутри одного объекта.

// Плюсы:
// 1. Замена алгоритмов налету;
// 2. Изоляция кода и данных алгоритмов от остальных объектов бизнес-логики;
// 3. Уход от наследования;
// 4. Реализует принцип open-closed.

// Минусы:
// 1. Усложнение программы за счет дополнительных объектов;
// 2. Необходимость клиенту знать, в чем состоит разница между стратегиями, чтобы выбрать подходящую.

// +++ Инициализация Стратегии
type Strategy interface {
	Route(startPoint int, endPoint int)
}

// ---

// +++ Инициализация общего контекста, в котором будет применяться
// та или иная стратегия, в зависимости от выбора пользователя.

type Navigator struct {
	Strategy
}

func (nav *Navigator) SetStrategy(stg Strategy) {
	nav.Strategy = stg
}

// ---

// +++ Реализация алгоритмов построений маршрута
type RoadStrategy struct {
}

func (r *RoadStrategy) Route(startPoint int, endPoint int) {
	avgSpeed := 30
	trafficJam := 2
	total := endPoint - startPoint
	totalTime := total * 40 * trafficJam / 100
	fmt.Printf("Road A:[%d] to B:[%d] Avg speed:[%d] Traffic jam:[%d] Total:[%d] Total time:[%d] min\n",
		startPoint, endPoint, avgSpeed, trafficJam, total, totalTime)
}

type PublicTransportStrategy struct {
}

func (r *PublicTransportStrategy) Route(startPoint int, endPoint int) {
	avgSpeed := 40
	total := endPoint - startPoint
	totalTime := total * 40 / 100
	fmt.Printf("Public transport A:[%d] to B:[%d] Avg speed:[%d] Total:[%d] Total time:[%d] min\n",
		startPoint, endPoint, avgSpeed, total, totalTime)
}

type WalkStrategy struct {
}

func (r *WalkStrategy) Route(startPoint int, endPoint int) {
	avgSpeed := 4
	total := endPoint - startPoint
	totalTime := total * 60 / 100
	fmt.Printf("Walk A:[%d] to B:[%d] Avg speed:[%d] Total:[%d] Total time:[%d] min\n",
		startPoint, endPoint, avgSpeed, total, totalTime)
}

// ---

var (
	srart      = 10
	end        = 100
	strategies = []Strategy{
		&PublicTransportStrategy{},
		&RoadStrategy{},
		&WalkStrategy{},
	}
)

func main() {
	nav := Navigator{}
	for _, strategy := range strategies {
		nav.SetStrategy(strategy)
		nav.Route(srart, end)
	}
}
