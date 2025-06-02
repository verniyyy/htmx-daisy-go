package todo

import (
	"encoding/gob"
	"os"
)

// TodoRepository provides methods to interact with the TODO items storage.
type TodoRepository struct {
	filename string
}

// TODO represents a task in the to-do list application.
func NewTodoRepository(filename string) *TodoRepository {
	return &TodoRepository{filename: filename}
}

// Add adds a new TODO item to the repository.
func (r *TodoRepository) Add(todo TODO) error {
	todos, err := r.List()
	if err != nil {
		return err
	}

	todos = append(todos, todo)
	return saveToFile(r.filename, todos)
}

// List retrieves all TODO items from the repository.
func (r *TodoRepository) List() ([]TODO, error) {
	todos, err := loadFromFile[[]TODO](r.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []TODO{}, nil // Return empty list if file does not exist
		}

		return nil, err
	}

	return todos, nil
}

func saveToFile(filename string, data any) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	return encoder.Encode(data)
}

func loadFromFile[T any](filename string) (T, error) {
	file, err := os.Open(filename)
	if err != nil {
		var zero T
		return zero, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)

	var data T
	if err := decoder.Decode(&data); err != nil {
		var zero T
		return zero, err
	}

	return data, nil
}
