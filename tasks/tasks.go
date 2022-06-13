package main

import "fmt"

type Task struct {
	name        string
	description string
	completed   bool
}

func NewTask(name string, description string, completed bool) *Task {
	return &Task{
		name:        name,
		description: description,
		completed:   completed,
	}
}

func (task *Task) String() string {
	return fmt.Sprintf("Type Task ---\nTask Name: %s\nTask Description: %s\nTask Completed: %t\n", task.name, task.description, task.completed)
}

func (t *Task) MarkCompleted() {
	t.completed = true
}

func (t *Task) UpdateDescription(newDescription string) {
	t.description = newDescription
}

func (t *Task) UpdateName(newName string) {
	t.name = newName
}

func main() {
	t := NewTask("Task Name", "Some random description to this task.", false)

	fmt.Printf(t.String())
	fmt.Println("----")

	t.UpdateName("Finish Go Course!!!")
	t.UpdateDescription("Complete the go course as soon as possible")
	t.MarkCompleted()
	fmt.Printf(t.String())
}
