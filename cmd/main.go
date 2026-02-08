package main

import (
	"fmt"
	"os"
	"strings"
	"task_tracker/internal/service"
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
	fmt.Println("  update              Update existing task.(Example: description id)")
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
