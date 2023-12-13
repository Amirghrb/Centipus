package main

import (
	"fmt"
	"time"

	"github.com/Amirghrb/octopus/handler"
	"github.com/Amirghrb/octopus/state"
)

func main() {
	// input("hi?")
	mainStore := state.Make()
	work := state.Work{
		Id: 0,
		Script: state.Script{
			Name:       "hello world",
			Path:       "./hel.sh",
			Created_at: time.Now().GoString(),
			Updated_at: time.Now().GoString(),
		},
		Server: state.Server{
			Name:       "ha3",
			Ipv4:       []string{"192.167.1.105"},
			Ipv6:       []string{},
			Hostname:   "amir",
			Created_at: time.Now().GoString(),
			Updated_at: time.Now().GoString(),
		},
		Rel:     -1,
		Inputs:  map[string]string{},
		Outputs: map[string]string{},
	}
	workflow := state.Workflow{
		Id:          0,
		Name:        "run",
		Description: "ajab baba",
		Works:       []state.Work{work},
		Created_at:  time.Now().GoString(),
		Updated_at:  time.Now().GoString(),
	}
	wfHandler := handler.WorkflowHandler{&mainStore}
	fmt.Println(">>>>>>> 11 >>>>>>>", workflow.Name)
	fmt.Println(">>>>>>> 22 >>>>>>>", workflow.Description)
	fmt.Println(">>>>>>> 33 >>>>>>>", workflow.Updated_at, "\t=", workflow.Created_at)
	wfHandler.Add("amir", "chetory???")
	fmt.Println(">>>>>>> 44 >>>>>>>", workflow.Description)
	fmt.Println(" >>>>>>>  55  >>>>>>>", workflow.Name)
	fmt.Println(">>>>>>> 66 >>>>>>>", workflow.Updated_at, "\t>", workflow.Created_at)
}

func input(prompt string) (string, error) {
	var command string
	fmt.Println(prompt)
	_, err := fmt.Scanf("%v", &command)
	if err != nil {
		return "", err

	}
	return command, nil
}
