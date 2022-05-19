package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

/*
=== HTTP server ===
Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.
В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

// resultResponse - структура вывода результатов
type resultResponse struct {
	Result []event `json:"result"`
}

// ResultResponse - функция вывода результатов
func ResultResponse(w http.ResponseWriter, res []event) {
	jsonError, _ := json.Marshal(&resultResponse{res})
	_, err := w.Write(jsonError)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadGateway)
	}
}

// Структура вывода ошибок
type errorResponse struct {
	Err string `json:"error"`
}

// ErrorResponse - функция вывода ошибок
func ErrorResponse(w http.ResponseWriter, err string, status int) {
	jsonError, _ := json.Marshal(&errorResponse{err})
	http.Error(w, string(jsonError), status)
}

// Структура нашего события в календаре
type event struct {
	User_ID     string     `json:"User_ID"`
	Event_ID    string     `json:"Event_ID"`
	Title       string     `json:"Title"`
	Description string     `json:"Description"`
	Date        CustomDate `json:"Date"`
}

// создаем затычку вместо БД
var events = []event{}

// Создаем кастомную дату для работы с json
type CustomDate struct {
	time.Time
}

// формат даты
const layout = "2006-01-02"

// функция десериализации json для кастомной даты
func (c *CustomDate) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`) // remove quotes
	if s == "null" {
		return
	}
	c.Time, err = time.Parse(layout, s)
	return
}

// функция сериализации json для кастомной даты
func (c CustomDate) MarshalJSON() ([]byte, error) {
	if c.Time.IsZero() {
		return nil, nil
	}
	return []byte(fmt.Sprintf(`"%s"`, c.Time.Format(layout))), nil
}

func main() {

	mux := http.NewServeMux()
	// задаем обработку маршрутов
	mux.HandleFunc("/create_event", createEvent)
	mux.HandleFunc("/update_event", updateEvent)
	mux.HandleFunc("/delete_event", deleteEvent)
	mux.HandleFunc("/events_for_day", events_for_day)
	mux.HandleFunc("/events_for_week", events_for_week)
	mux.HandleFunc("/events_for_month", events_for_month)

	err := http.ListenAndServe(":4000", mux)
	log.Fatalln(err)
}

// функция создания события
func createEvent(w http.ResponseWriter, r *http.Request) {
	// задаем заголовки
	w.Header().Set("Content-Type", "application/json")
	// проверка метода
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		ErrorResponse(w, " incorret method", http.StatusMethodNotAllowed)
		return
	}
	// считывание информации из тела запроса
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ErrorResponse(w, " error with reading json", http.StatusNotAcceptable)
		return
	}
	// анмаршалим данные в нашу структуру
	json.Unmarshal(reqBody, &newEvent)
	// проверка, если событие уже создано
	eventFind := findEvent(newEvent)
	if eventFind > -1 {
		ErrorResponse(w, " Event id exist, if you want update use updateEvent", http.StatusNotAcceptable)
		return
	}
	// добавляем в затычку бд, новое событие
	events = append(events, newEvent)

	// отправляем данные
	w.WriteHeader(http.StatusCreated)
	log.Println("Event created: ", newEvent.Event_ID)
	ResultResponse(w, []event{newEvent})
}

// функция обновления события
func updateEvent(w http.ResponseWriter, r *http.Request) {
	// задаем заголовки
	w.Header().Set("Content-Type", "application/json")
	// проверка метода
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		ErrorResponse(w, " incorret method", http.StatusMethodNotAllowed)
		return
	}
	// считываем данные из тела запроса
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ErrorResponse(w, " error with reading json", http.StatusNotAcceptable)
		return
	}

	json.Unmarshal(reqBody, &newEvent)
	// проверка на существование события
	eventID := findEvent(newEvent)
	if eventID == -1 {
		ErrorResponse(w, " Event id not found", http.StatusNotAcceptable)
		return
	}
	// если найдено событие меняем его
	events[eventID] = newEvent

	w.WriteHeader(http.StatusCreated)

	ResultResponse(w, []event{newEvent})
}

// удаляем событие
func deleteEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		ErrorResponse(w, " incorret method", http.StatusMethodNotAllowed)
		return
	}
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ErrorResponse(w, " error with reading json", http.StatusNotAcceptable)
		return
	}
	// поиск события по event id
	json.Unmarshal(reqBody, &newEvent)
	eventID := findEvent(newEvent)
	if eventID == -1 {
		ErrorResponse(w, " Event id not found", http.StatusBadRequest)
		return
	}
	// удаляем событие
	events[eventID] = events[len(events)-1]
	events = events[:len(events)-1]

	w.WriteHeader(http.StatusCreated)

	fmt.Fprintf(w, `{"result": "event deleted"}`)

}

// события за день
func events_for_day(w http.ResponseWriter, r *http.Request) {
	// заголовки
	w.Header().Set("Content-Type", "application/json")
	//проверка метода
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		ErrorResponse(w, " incorret method", http.StatusMethodNotAllowed)
		return
	}
	// парсим id и дату
	id := r.URL.Query().Get("id")
	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	ResultResponse(w, findDay(id, date))

}

// события за неделю
func events_for_week(w http.ResponseWriter, r *http.Request) {
	// заголовки
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		ErrorResponse(w, " incorret method", http.StatusMethodNotAllowed)
		return
	}
	// парсим id и дату
	id := r.URL.Query().Get("id")

	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	ResultResponse(w, findWeek(id, date))

}

// события за месяц
func events_for_month(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		ErrorResponse(w, " incorret method", http.StatusMethodNotAllowed)
		return
	}
	// парсим id и дату
	id := r.URL.Query().Get("id")
	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse(layout, dateStr)

	if err != nil {
		http.NotFound(w, r)
		return
	}
	ResultResponse(w, findMonth(id, date))

}

// поиск события в определенные день. Принимает на вход id и время. Выводит список событий
func findDay(id string, t time.Time) (ev []event) {
	for _, v := range events {
		if v.User_ID == id && t == v.Date.Time {
			ev = append(ev, v)
		}
	}
	return ev
}

// поиск события в определенную неделю. Принимает на вход id и время. Выводит список событий
func findWeek(id string, t time.Time) (ev []event) {
	for _, v := range events {
		year1, week1 := t.ISOWeek()
		year2, week2 := v.Date.ISOWeek()
		if v.User_ID == id && year1 == year2 && week1 == week2 {
			ev = append(ev, v)
		}
	}
	return ev
}

// поиск события в определенный месяц. Принимает на вход id и время. Выводит список событий
func findMonth(id string, t time.Time) (ev []event) {
	for _, v := range events {
		if v.User_ID == id && t.Month() == v.Date.Month() && t.Year() == v.Date.Year() {
			ev = append(ev, v)
		}
	}
	return ev
}

// поиск события, в затычке БД. Принимает на вход событие и ищет его id в списке событий
func findEvent(event event) int {
	for i, v := range events {
		if event.Event_ID == v.Event_ID {
			return i
		}
	}
	return -1
}
