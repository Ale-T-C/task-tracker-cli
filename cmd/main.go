package main

import (
	"fmt"
	"os"
	"strings"
	"task_tracker/internal/service"

	"github.com/google/uuid"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		globalHelp()
		return
	}

	comando := args[0]

	switch comando {
	case "add", "-a":

		if len(args) < 2 {
			fmt.Println("Error: A description is required to add a task.")
			return
		}

		descripcion := strings.Join(args[1:], " ")

		service.TaskRegister(descripcion)

		fmt.Printf("Task successful %s\n", descripcion)
		break
	case "update", "-u":
		if len(args) < 3 {
			fmt.Println("Error: Correct usage: update <id> <new description>")
			return
		}

		idInput := args[1]
		nuevaDesc := strings.Join(args[2:], " ")

		success, msg := service.TaskUpdate(nuevaDesc, "", uuid.MustParse(idInput))
		if !success {
			fmt.Printf("%s\n", msg)
		} else {
			fmt.Printf("%s\n", msg)
		}
		break
	case "list", "-l":

		var result []service.RegisterTask
		var errStr string

		if len(args) == 2 {
			status := args[1]
			result, errStr = service.TaskByStatus(status)

		} else {
			result, errStr = service.TaskAll()
		}

		if errStr != "" {
			fmt.Printf("Error: %s\n", errStr)
			break
		}

		if len(result) == 0 {
			fmt.Println("No tasks registered.")
			break
		}

		for _, t := range result {
			fmt.Printf("-----------------------------------\n")

			fmt.Printf(" Id: %-7s \n", t.ID.String())
			fmt.Printf(" Description: %-25s\n", t.Description)
			fmt.Printf(" Status: %-10s\n ", t.Status)
		}

		break
	case "mark-in-progress", "-mp":
		if len(args) < 2 {
			fmt.Println("Error: Correct usage: mark-in-progress <id>")
			return
		}
		idInput := args[1]

		success, msg := service.TaskUpdate("", "mark-in-progress", uuid.MustParse(idInput))

		if !success {
			fmt.Printf("%s\n", msg)
		} else {
			fmt.Printf("%s\n", msg)
		}
	case "mark-done", "-md":
		if len(args) < 2 {
			fmt.Println("Error: Correct usage: mark-done <id>")
			return
		}
		idInput := args[1]

		success, msg := service.TaskUpdate("", "mark-done", uuid.MustParse(idInput))

		if !success {
			fmt.Printf("%s\n", msg)
		} else {
			fmt.Printf("%s\n", msg)
		}
	case "delete", "-d":
		if len(args) < 2 {
			fmt.Println("Error: Correct usage: delete <id>")
			return
		}
		idInput := args[1]

		success, msg := service.TaskDelete(uuid.MustParse(idInput))

		if !success {
			fmt.Printf("%s\n", msg)
		} else {
			fmt.Printf("%s\n", msg)
		}
		break
	case "help", "-h":
		globalHelp()
		break
	default:
		fmt.Printf("Error: Command '%s' not recognized.\n", comando)
	}
}

func globalHelp() {
	fmt.Println("TASK TRAKER CLI v1.0")
	fmt.Println("Use: task <command> [arguments]")
	fmt.Println("\n Enable commands:")
	fmt.Println("  add                 Create new task.")
	fmt.Println("  update              Update existing task.(Example: id new description)")
	fmt.Println("  delete              Delete a task by ID. (Example: delete id)")
	fmt.Println("  mark-in-progress    Mark task in status progress. (Example: mark-in-progress id)")
	fmt.Println("  mark-done           Mark task as done. (Example: mark-done id)")
	fmt.Println("  list                List all tasks.")
	fmt.Println("  list todo           List tasks with status todo.")
	fmt.Println("  list done           List tasks with status done.")
	fmt.Println("  list in-progress    List tasks with status in-progress.")
	fmt.Println("  help                Show this documentation.")

	fmt.Println("\n Show commands shortcut:")
	fmt.Println("  -a  add")
	fmt.Println("  -u  update")
	fmt.Println("  -d  delete")
	fmt.Println("  -mp mark-in-progress")
	fmt.Println("  -md mark-done")
	fmt.Println("  -l  list")
	fmt.Println("  -h  help")
}
