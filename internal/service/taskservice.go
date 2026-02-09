package service

import (
	"encoding/json"
	"os"
	"task_tracker/internal/generals"
	"time"

	"github.com/google/uuid"
)

type RegisterTask struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

const file_name = "task.json"

func TaskRegister(description string) (bool, string) {
	var task_dto []RegisterTask

	if generals.FileExists(file_name) {
		content, err := os.ReadFile(file_name)
		if err == nil {
			json.Unmarshal(content, &task_dto)
		}
	} else {
		generals.CreateFile(file_name, []byte("[]"))
	}

	task := RegisterTask{
		ID:          uuid.New(),
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	task_dto = append(task_dto, task)

	registr, _ := json.MarshalIndent(task_dto, "", "  ")
	result := generals.WriteFile(file_name, registr)

	if !result {
		return false, "Failed to add task"
	}

	return true, "Task added successfully"
}
