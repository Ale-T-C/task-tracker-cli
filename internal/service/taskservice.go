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

func TaskUpdate(description string, type_mark string, id uuid.UUID) (bool, string) {

	var task_list []RegisterTask

	result_task, _ := generals.ReadFile(file_name)

	if result_task == nil {
		return false, "File not found"
	}

	json.Unmarshal(result_task, &task_list)

	for i, x := range task_list {
		if x.ID == id {

			if type_mark == "mark-in-progress" {
				task_list[i].Status = "in-progress"
				task_list[i].UpdatedAt = time.Now()
			}

			if type_mark == "mark-done" {
				task_list[i].Status = "done"
				task_list[i].UpdatedAt = time.Now()
			}

			if description != "" {
				task_list[i].Description = description
				task_list[i].UpdatedAt = time.Now()
			}
			break
		}
	}

	update, err := json.MarshalIndent(task_list, "", "  ")

	if err != nil {
		panic(err)
	}

	result := generals.WriteFile(file_name, update)

	if !result {
		panic("Update failed")
	}
	return true, "Task updated successfully"
}

func TaskDelete(id uuid.UUID) (bool, string) {

	var task_list []RegisterTask

	result_task, _ := generals.ReadFile(file_name)

	if result_task == nil {
		return false, "File not found"
	}

	json.Unmarshal(result_task, &task_list)

	for i, x := range task_list {
		if x.ID == id {

			task_list = append(task_list[:i], task_list[i+1:]...)

			break
		}
	}

	update, err := json.MarshalIndent(task_list, "", "  ")

	if err != nil {
		panic(err)
	}

	result := generals.WriteFile(file_name, update)

	if !result {
		panic("Delete failed")
	}
	return true, "Task delete successfully"
}

func TaskAll() ([]RegisterTask, string) {
	var task_list []RegisterTask

	content, err := generals.ReadFile(file_name)

	if err != nil {
		return nil, "File not found or could not be read"
	}

	json.Unmarshal(content, &task_list)

	return task_list, ""
}

func TaskByStatus(status string) ([]RegisterTask, string) {
	var task_list []RegisterTask

	result_task, _ := generals.ReadFile(file_name)

	if result_task == nil {
		return nil, "File not found"
	}

	json.Unmarshal(result_task, &task_list)

	var filtered []RegisterTask

	for _, x := range task_list {
		if x.Status == status {
			filtered = append(filtered, x)
		}
	}
	return filtered, ""
}
