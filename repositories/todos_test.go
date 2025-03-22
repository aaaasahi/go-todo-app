package repositories

import (
	"go-todo-app/models"
	"testing"
)

func TestGetTodo(t *testing.T) {
	db, err := connectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	testCases := []struct {
		name     string
		expected models.Todo
	}{
		{
			name: "test1",
			expected: models.Todo{
				ID:      1,
				Title:   "todo1 title",
				Content: "todo1 content",
			},
		},
		{
			name: "test2",
			expected: models.Todo{
				ID:      2,
				Title:   "todo2 title",
				Content: "todo2 content",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetTodo(db, tc.expected.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.ID != tc.expected.ID {
				t.Errorf("ID: got %v, expected %v", got.ID, tc.expected.ID)
			}

			if got.Title != tc.expected.Title {
				t.Errorf("Title: got %v, expected %v", got.Title, tc.expected.Title)
			}

			if got.Content != tc.expected.Content {
				t.Errorf("Content: got %v, expected %v", got.Content, tc.expected.Content)
			}
		})
	}
}

func TestListTodos(t *testing.T) {
	db, err := connectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	todos, err := ListTodos(db)
	if err != nil {
		t.Fatal(err)
	}

	// 少なくとも2つのTodoが存在することを確認
	if len(todos) < 2 {
		t.Errorf("Expected minimum todos: 2, got: %d", len(todos))
	}

	// 最初の2つのTodoが期待通りであることを確認
	expectedTodos := []models.Todo{
		{
			ID:      1,
			Title:   "todo1 title",
			Content: "todo1 content",
		},
		{
			ID:      2,
			Title:   "todo2 title",
			Content: "todo2 content",
		},
	}

	for i, expected := range expectedTodos {
		if i >= len(todos) {
			t.Fatalf("Not enough todos returned")
		}

		if todos[i].ID != expected.ID {
			t.Errorf("ID: got %v, expected %v", todos[i].ID, expected.ID)
		}

		if todos[i].Title != expected.Title {
			t.Errorf("Title: got %v, expected %v", todos[i].Title, expected.Title)
		}

		if todos[i].Content != expected.Content {
			t.Errorf("Content: got %v, expected %v", todos[i].Content, expected.Content)
		}
	}
}

func TestCreateTodo(t *testing.T) {
	db, err := connectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	newTodo := models.Todo{
		Title:   "newTodo title",
		Content: "newTodo content",
	}

	createdTodo, err := CreateTodo(db, newTodo)
	if err != nil {
		t.Fatal(err)
	}

	// IDが割り当てられていることを確認
	if createdTodo.ID == 0 {
		t.Error("Created todo has no ID assigned")
	}

	// 作成されたTodoをデータベースから取得して確認
	fetchedTodo, err := GetTodo(db, createdTodo.ID)
	if err != nil {
		t.Fatal(err)
	}

	if fetchedTodo.Title != newTodo.Title {
		t.Errorf("Title: got %v, expected %v", fetchedTodo.Title, newTodo.Title)
	}

	if fetchedTodo.Content != newTodo.Content {
		t.Errorf("Content: got %v, expected %v", fetchedTodo.Content, newTodo.Content)
	}

	// テスト後にデータを削除
	err = DeleteTodo(db, createdTodo.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateTodo(t *testing.T) {
	db, err := connectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// まず新しいTodoを作成
	newTodo := models.Todo{
		Title:   "updateTodo title",
		Content: "updateTodo content",
	}

	createdTodo, err := CreateTodo(db, newTodo)
	if err != nil {
		t.Fatal(err)
	}

	updatedTodo := models.Todo{
		ID:      createdTodo.ID,
		Title:   "updatedTodo title",
		Content: "updatedTodo content",
	}

	_, err = UpdateTodo(db, updatedTodo)
	if err != nil {
		t.Fatal(err)
	}

	// 更新されたTodoをデータベースから取得して確認
	fetchedTodo, err := GetTodo(db, createdTodo.ID)
	if err != nil {
		t.Fatal(err)
	}

	if fetchedTodo.Title != updatedTodo.Title {
		t.Errorf("Title: got %v, expected %v", fetchedTodo.Title, updatedTodo.Title)
	}

	if fetchedTodo.Content != updatedTodo.Content {
		t.Errorf("Content: got %v, expected %v", fetchedTodo.Content, updatedTodo.Content)
	}

	// テスト後にデータを削除
	err = DeleteTodo(db, createdTodo.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteTodo(t *testing.T) {
	db, err := connectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// まず新しいTodoを作成
	newTodo := models.Todo{
		Title:   "deleteTodo title",
		Content: "deleteTodo content",
	}

	createdTodo, err := CreateTodo(db, newTodo)
	if err != nil {
		t.Fatal(err)
	}

	err = DeleteTodo(db, createdTodo.ID)
	if err != nil {
		t.Fatal(err)
	}

	_, err = GetTodo(db, createdTodo.ID)
	if err == nil {
		t.Error("Todo was not deleted")
	}
}
