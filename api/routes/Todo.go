package routes

import "fmt"

type Todo struct {
	ID          int    `json:"id"`
	DESCRIPTION string `json:"description"`
	STATUS      bool   `json:"status"`
}

var Todos []Todo

func AddTododMethod(todo Todo) []Todo {
	Todos = append(Todos, todo)
	fmt.Println("a new Todo was added")
	fmt.Println(Todos)
	return Todos
}

func DeleteTodoMethod(id int) []Todo {
	for index, item := range Todos {
		if item.ID == id {
			Todos = append(Todos[:index], Todos[index+1:]...)
		}
	}
	return Todos
}
func UpdateTodoMethod(todo Todo, id int) []Todo {
	for index, item := range Todos {
		if item.ID == id {
			Todos[index].ID = todo.ID
			Todos[index].DESCRIPTION = todo.DESCRIPTION
			Todos[index].STATUS = todo.STATUS

		}
	}
	return Todos
}
