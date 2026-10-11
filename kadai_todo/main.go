package main

import (
	"fmt"
)

type ToDo struct {
	ID        int
	Title     string
	Completed bool
}

func (t *ToDo) Complete() {
	fmt.Printf("ID：%dのToDoを完了に更新します\n", t.ID)
	t.Completed = true
}

func printToDos(td []*ToDo) {
	var cmp string
	for i := range td {
		if td[i].Completed {
			cmp = "完　了"
		} else {
			cmp = "未完了"
		}
		fmt.Printf("[%s]（ID：%d）%s\n", cmp, td[i].ID, td[i].Title)
	}
}

func main() {

	todo1 := ToDo{
		ID:        1,
		Title:     "学習計画",
		Completed: false,
	}

	todo2 := ToDo{
		ID:        2,
		Title:     "環境構築",
		Completed: false,
	}

	todo3 := ToDo{
		ID:        3,
		Title:     "基礎文法",
		Completed: false,
	}

	todos := []*ToDo{&todo1, &todo2, &todo3}

	fmt.Println("--- ToDoリスト（初期状態） ---")
	printToDos(todos)
	fmt.Println()

	todo1.Complete()
	todo2.Complete()

	fmt.Println()
	fmt.Println("--- ToDoリスト（最終状態） ---")
	printToDos(todos)
}
