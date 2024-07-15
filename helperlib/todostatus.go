package helperlib

import (
	"errors"
	"slices"
	"strconv"
)

type Todostatus string

const (
	Active    Todostatus = "active"
	Inactive  Todostatus = "inactive"
	OnHold    Todostatus = "onhold"
	Done      Todostatus = "done"
	Cancelled Todostatus = "cancelled"
)

type Todo struct {
	id      int
	message string
	status  Todostatus
}

type TodoStore struct {
	name    string
	maxId   int
	counter int
	todos   map[int]Todo
}

func NewTodoStore(name string) *TodoStore {
	return &TodoStore{
		name:  name,
		todos: make(map[int]Todo),
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
func (tds *TodoStore) CountTodos() int {
	return tds.counter
}
func (tds *TodoStore) GetTodo(id int) [3]string {
	todo := [3]string{
		strconv.Itoa(tds.todos[id].id),
		tds.todos[id].message,
		string(tds.todos[id].status),
	}
	return todo
}

func (tds *TodoStore) GetTodos() [][]string {
	todos := make([][]string, 0)
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
		todo_strings := []string{
			strconv.Itoa(tds.todos[idx].id),
			message,
			string(tds.todos[idx].status),
		}
		todos = append(todos, todo_strings)
	}
	return todos
}

func (tds *TodoStore) AddTodo(message string) {
	newId := tds.maxId + 1
	tds.todos[newId] = Todo{
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
	return nil
}
