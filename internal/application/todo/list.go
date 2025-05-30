package todo

import "time"

// ListUseCase represents the use case for listing TODO items.
type ListUseCase struct{}

// NewListUseCase creates a new instance of ListUseCase.
func NewListUseCase() *ListUseCase {
	return &ListUseCase{}
}

// Execute retrieves a list of TODO items.
func (uc *ListUseCase) Execute() ([]TODO, error) {
	todos := []TODO{
		{
			ID:        1,
			Title:     "Learn Go",
			IsDone:    false,
			CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:        2,
			Title:     "Build a web app",
			IsDone:    false,
			CreatedAt: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
		},
	}

	return todos, nil
}
