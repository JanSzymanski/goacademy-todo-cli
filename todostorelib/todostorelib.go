package todostorelib

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
)

type Todostatus string

const (
	Active   Todostatus = "active"
	Inactive Todostatus = "inactive"
	Done     Todostatus = "done"
)

type todo struct {
	id      int
	message string
	status  Todostatus
}

type TodoStore struct {
	name    string
	maxId   int
	counter int
	todos   map[int]todo
}

func NewTodoStore(name string) *TodoStore {
	return &TodoStore{
		name:  name,
		todos: make(map[int]todo),
	}
}
func (tds *TodoStore) GetVaultName() string {
	return tds.name
}

func (tds *TodoStore) GetVaultInfo() map[string]string {
	vaultInfo := make(map[string]string)
	vaultInfo["name"] = tds.name
	vaultInfo["maxId"] = strconv.Itoa(tds.maxId)
	vaultInfo["counter"] = strconv.Itoa(tds.counter)
	return vaultInfo
}
func (tds *TodoStore) CountTodos() int { // possible to remove
	return tds.counter
}
func (tds *TodoStore) GetTodo(id int) (map[string]string, error) {
	todo, ok := tds.todos[id]
	todo_map := make(map[string]string)
	if !ok {
		return todo_map, errors.New(fmt.Sprint("No todo with id: ", id))
	}
	todo_map["id"] = strconv.Itoa(id)
	todo_map["message"] = todo.message
	todo_map["status"] = string(todo.status)

	return todo_map, nil
}

func (tds *TodoStore) GetTodos(page int, max int) []map[string]string {
	todos := make([]map[string]string, 0)
	keys := make([]int, 0)
	for k := range tds.todos {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, idx := range keys {
		var message string
		if len(tds.todos[idx].message) > 20 {
			message = tds.todos[idx].message[:17] + "..."
		} else {
			message = tds.todos[idx].message
		}
		todo := make(map[string]string)
		todo["id"] = strconv.Itoa(tds.todos[idx].id)
		todo["message"] = message
		todo["status"] = string(tds.todos[idx].status)
		todos = append(todos, todo)
	}
	if tds.counter > max {
		if tds.counter > max*(page+1) {
			return todos[max*page : max*(page+1)]
		}
		return todos[max*page:]
	}
	return todos
}

func (tds *TodoStore) AddTodo(message string) {
	newId := tds.maxId + 1
	tds.todos[newId] = todo{
		id:      newId,
		message: message,
		status:  Inactive,
	}
	tds.maxId = newId
	tds.counter++
}
func (tds *TodoStore) DeleteTodo(id int) error {
	if _, ok := tds.todos[id]; !ok {
		return errors.New("no Todo with given ID in the vault")
	}
	delete(tds.todos, id)
	tds.counter--
	return nil
}

func (tds *TodoStore) ChangeTodoStatus(id int, new_status Todostatus) {
	tds.todos[id] = todo{
		id:      tds.todos[id].id,
		message: tds.todos[id].message,
		status:  new_status,
	}
}
func (tds *TodoStore) ChangeTodoMessagge(id int, new_message string) {
	tds.todos[id] = todo{
		id:      tds.todos[id].id,
		message: new_message,
		status:  tds.todos[id].status,
	}
}
