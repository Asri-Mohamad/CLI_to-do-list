package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Asri-Mohamad/Master_Function"
)

type taskStruct struct {
	Task string `json:"task"`
	Date string `json:"data"`
	Time string `json:"time"`
}

func main() {
	var taskMaster []taskStruct
	var Task = taskStruct{}
	var command string
	var num int
	var error bool
	for {
		command, Task.Task, Task.Date, Task.Time, num, error = mainMenuProcess(startReadCommand(), &taskMaster)
		if command == "exit" {
			return
		} else {
			if command == "none" {
				continue
			} else {
				if !error {
					fmt.Println(command)

				} else {
					switch command {
					case "add":
						addTask(&taskMaster, Task.Task, Task.Date, Task.Time)

						deleteTask(&taskMaster, num)
					case "edit":
						editTask(taskMaster, Task.Task, Task.Date, Task.Time, num)
					case "save":
						saveList(taskMaster, Task.Task)
					case "load":
						taskMaster = loadList(taskMaster, Task.Task)

					}
				}
			}
		}

	}
}

// ------------------------------------------------

// ------------ Read start -------------
func startReadCommand() string {

	fmt.Print("Command>>")

	newcommand := bufio.NewReader(os.Stdin)
	command, _ := newcommand.ReadString('\n')
	command = strings.TrimSpace(command)
	return command
}

// -------------- process Read ----------
func mainMenuProcess(readData string, tasks *[]taskStruct) (string, string, string, string, int, bool) {
	var command, task, date, time string
	var nu int

	switch readData {
	case "exit", "x":
		fmt.Println("Thanks for use this program ...")

		command = "exit"
	case "cls":
		Master_Function.Cls()
		command = "none"
	case "list":
		showList(tasks)
		command = "none"
	case "help":
		fmt.Println("This is a To-Do-List program you can Add, Delete, Edit, List, Save to file and Load from file on your list.")
		fmt.Println()
		fmt.Println("      add  [Task] [Date] [Time]                            For add task in your list.        ")
		fmt.Println("      list                                                 For show your list.               ")
		fmt.Println("      delete [Task number]                                 For delete task by number of task.")
		fmt.Println("      edit [Task number] [New Task] [New Date] [New Time]  For edit task by number of task.  ")
		fmt.Println("      save [filename.json]                                 For save to file.                 ")
		fmt.Println("      load [filename.json]                                 For load to file.                 ")
		fmt.Println("      cls                                                  For clear screen.                 ")
		fmt.Println("      exit                                                 For exit this program.            ")
		command = "none"
	default:

		pr := strings.SplitN(readData, " ", 2)
		command = pr[0]

		if checkCommant(command) && len(pr) > 1 {

			switch command {
			//add ------------------------------------------
			case "add":
				parts := []string{}
				in := 0
				err := 0
				parts = append(parts, command)
				cut := false
				pr[1] = strings.TrimSpace(pr[1])
				for _, char := range pr[1] {
					if char == '"' {
						in++

						if !cut {
							cut = true
							parts = append(parts, "")
						} else {
							cut = false
							in--
						}
					}
					if cut {
						if char != '"' {
							if in > 3 {
								err = 1
								break
							}
							parts[in] += string(char)
						}

					}
				}

				if len(parts) == 4 && err != 1 {
					return parts[0], parts[1], parts[2], parts[3], nu, true
				} else {
					return "Sintax error !!! add command not fund ", "", "", "", 0, false
				}
				//delete ------------------------------------------
			case "delete":
				pr[1] = strings.TrimSpace(pr[1])
				index, err := strconv.Atoi(pr[1])

				if err == nil {
					nu = index
					return command, "", "", "", nu, true
				} else {
					return "Sintax error !!! delete command not fund ", "", "", "", 0, false
				}
				//Edit ------------------------------------------
			case "edit":
				parts := []string{}
				in := 0
				err := 0
				parts = append(parts, command)
				cut := false

				pr = strings.SplitN(pr[1], " ", 2)

				if len(pr) < 2 {
					return "Sintax error !!! edit command not fund ", "", "", "", 0, false
				}

				pr[0] = strings.TrimSpace(pr[0])
				index, errN := strconv.Atoi(pr[0])

				if errN == nil {
					nu = index
					pr[1] = strings.TrimSpace(pr[1])
					for _, char := range pr[1] {
						if char == '"' {
							in++

							if !cut {
								cut = true
								parts = append(parts, "")
							} else {
								cut = false
								in--
							}
						}
						if cut {
							if char != '"' {
								if in > 3 {
									err = 1
									break
								}
								parts[in] += string(char)
							}

						}
					}
					//fmt.Println("len =", len(parts))
					if len(parts) == 4 && err != 1 {
						return parts[0], parts[1], parts[2], parts[3], nu, true
					} else {
						return "Sintax error !!! edit command not fund ", "", "", "", 0, false
					}

				} else {
					return "Sintax error !!! edit command not fund ", "", "", "", 0, false
				}
				//save ------------------------------------------
			case "save":
				invalidCh := []rune{'<', '>', ':', '"', '/', '\\', '|', '?', '*'}
				for _, char := range invalidCh {
					if strings.ContainsRune(pr[1], char) {

						return fmt.Sprintln("Invalid charakters in file name:", string(char)), "", "", "", 0, false
					}
				}
				if !strings.HasSuffix(pr[1], ".json") {
					return "Invalid Extention in file name ...", "", "", "", 0, false
				}

				return command, pr[1], "", "", 0, true
				//load ------------------------------------------
			case "load":
				invalidCh := []rune{'<', '>', ':', '"', '/', '\\', '|', '?', '*'}
				for _, char := range invalidCh {
					if strings.ContainsRune(pr[1], char) {

						return fmt.Sprintln("Invalid charakters in file name:", string(char)), "", "", "", 0, false
					}
				}
				if !strings.HasSuffix(pr[1], ".json") {
					return "Invalid Extention in file name ...", "", "", "", 0, false
				}

				return command, pr[1], "", "", 0, true

			default:
				return "Sintax error !!! command not fund ", "", "", "", 0, false
			}

		} else {
			command = fmt.Sprint("Sintax error !!! (\"", command, "\")")
			return command, task, date, time, nu, false
		}
	}
	return command, task, date, time, nu, false
}

// -------------------------------------------

// --------------------------------------------------
func checkCommant(comand string) bool {

	if comand == "add" || comand == "delete" || comand == "edit" || comand == "save" || comand == "load" {
		comand = ""
		return true
	}
	return false

}

// ---------------------------------------------------
func addTask(tasks *[]taskStruct, readTask string, readDate string, readTime string) {

	read := taskStruct{
		Task: readTask,
		Date: readDate,
		Time: readTime}
	*tasks = append(*tasks, read)
	fmt.Println("New task added...")

}

//-------------------------------------------------

func deleteTask(tasks *[]taskStruct, index int) {

	if len(*tasks) <= 0 {
		fmt.Println("List is empty...")
		//_ = Master_Function.CharGetKey()
		return
	}
	//showList(tasks)

	if index >= 0 && index < len(*tasks) {
		fmt.Printf("\n%v) Task: %s   Date: %s   Time: %s\nAre you shore for delete ?(Y/N)", index, (*tasks)[index].Task, (*tasks)[index].Date, (*tasks)[index].Time)

	outLoop:
		for {

			switch Master_Function.CharGetKey() {
			case 'Y', 'y':
				*tasks = append((*tasks)[:index], (*tasks)[index+1:]...)
				fmt.Println("Y")
				fmt.Println("Deleted Complet...")
				break outLoop

			case 'N', 'n':
				fmt.Println("N")
				break outLoop

			}
		}

	} else {

		fmt.Printf("\nError to Enter! select over betwin 0 to %v: \n", len(*tasks)-1)

	}

}

// ---------------------------------------------------
func showList(tasks *[]taskStruct) {
	if len(*tasks) > 0 {
		fmt.Println()

		for key, data := range *tasks {
			fmt.Printf("%s  %s %s   %s %s   %s %s\n%s\n", Master_Function.ColorText("read", strconv.Itoa(key)+")"),
				Master_Function.ColorText("yellow", "Task :"), Master_Function.ColorText("green", data.Task),
				Master_Function.ColorText("yellow", "Date:"), Master_Function.ColorText("cyan", data.Date),
				Master_Function.ColorText("yellow", "Time:"), Master_Function.ColorText("blue", data.Time),
				Master_Function.ColorText("magenta", "-----------------------"))

		}
	} else {
		fmt.Println(" This list is empty ...")
	}

}

// ---------------------------------------------------
func saveList(tasks []taskStruct, fileName string) {

	showList(&tasks)
	print("Do you want to save this list(Y/N)?")

	for {
		yn := Master_Function.CharGetKey()
		if yn == 'Y' || yn == 'y' || yn == 'n' || yn == 'N' {
			if yn == 'n' || yn == 'N' {
				return
			}
			break
		}
	}

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Create file have problem...")

		return

	} else {

		defer file.Close()

		encode := json.NewEncoder(file)
		err := encode.Encode(tasks)
		if err != nil {
			fmt.Println("Write file have problem....")

		}

		fmt.Printf("\nSave is completly...%v data saved\n", len(tasks))

	}

}

// ---------------------------------------------------
func loadList(tasks []taskStruct, fileName string) []taskStruct {
	if len(tasks) > 0 {
		fmt.Printf("\n%s", Master_Function.ColorText("yellow", "The list in not empty do you want over write?(Y/N)"))
		for {
			yn := Master_Function.CharGetKey()
			if yn == 'Y' || yn == 'y' || yn == 'n' || yn == 'N' {
				if yn == 'n' || yn == 'N' {
					return tasks
				}
				break
			}
		}
	}
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Printf("%s\n", Master_Function.ColorText("read", "I Can't open file ...."))

		return tasks
	} else {
		defer file.Close()
		decode := json.NewDecoder(file)
		err := decode.Decode(&tasks)
		if err != nil {
			fmt.Printf("\n%s", Master_Function.ColorText("read", "Decoding file problem...."))

			return tasks
		}

	}
	fmt.Printf("%s\n", Master_Function.ColorText("yellow", "The list loded OK...."))

	return tasks
}

// ---------------------------------------------------
func editTask(tasks []taskStruct, readTask string, readDate string, readTime string, index int) []taskStruct {

	if len(tasks) <= 0 {
		fmt.Println("List is empty...")

	} else {

		if index >= 0 && index < len(tasks) {

			fmt.Printf("\n%v) Task: %s   Date: %s   Time: %s\nAre you shore for Edit ?(Y/N)", index, tasks[index].Task,
				tasks[index].Date, tasks[index].Time)

		outLoop:
			for {

				switch Master_Function.CharGetKey() {
				case 'Y', 'y':

					tasks[index] = taskStruct{readTask, readDate, readTime}
					fmt.Println("\nEdit Complet...")
					break outLoop

				case 'N', 'n':
					fmt.Println("N")
					break outLoop

				}
			}

		} else {
			fmt.Printf("\nError to Enter! Over betwin 0 to %v: \n", len(tasks)-1)

		}
	}
	return tasks
}

// ---------------------------------------------------
