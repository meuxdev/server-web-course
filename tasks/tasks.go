package main

import "fmt"

type TaskList struct {
	tasks []*Task
}

func NewTaskList() *TaskList {
	return &TaskList{}
}

func (tl *TaskList) AddTaskToList(task *Task) {
	tl.tasks = append(tl.tasks, task)
}

func (tl *TaskList) RemoveTaskFromIndex(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)

}

func (tl *TaskList) PrintAllTasks() {
	var chain string
	chain += "-----------------\nTasks List \n------------------\n---------------\n"
	for _, task := range tl.tasks {
		chain += task.String()
	}
	fmt.Println(chain)
}

func (tl *TaskList) PrintCompletedTasks() {
	var chain string
	chain += "-----------------\nCompleted Tasks List \n------------------\n---------------\n"
	for _, task := range tl.tasks {
		if task.completed {
			chain += task.String()
		}
	}

	fmt.Println(chain)
}

func (tl *TaskList) PrintTasksWithOddIndex() {
	var chain string
	chain += "-----------------\nOdd Tasks List \n------------------\n---------------\n"
	for index, task := range tl.tasks {
		if (index+1)%2 == 0 {
			continue
		}
		chain += task.String()
	}
	fmt.Println(chain)
}

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

	l := NewTaskList()
	t := NewTask("Task Name", "Some random description to this task.", false)

	// adding 3 tasks

	for i := 0; i < 3; i++ {
		taskName := fmt.Sprintf("Task Number %d", i+1)
		tempTask := NewTask(taskName, "Some random description to this task.", false)
		l.AddTaskToList(tempTask)
	}

	l.AddTaskToList(t)

	fmt.Printf(t.String())
	fmt.Println("----")

	t.UpdateName("Finish Go Course!!!")
	t.UpdateDescription("Complete the go course as soon as possible")
	t.MarkCompleted()
	fmt.Printf(t.String())

	l.PrintAllTasks()
	l.RemoveTaskFromIndex(len(l.tasks) - 2)
	l.PrintAllTasks()
	l.PrintTasksWithOddIndex()
	l.PrintCompletedTasks()

}
