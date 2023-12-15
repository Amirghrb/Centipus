package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Amirghrb/octopus/handler"
	"github.com/Amirghrb/octopus/state"
)

func main() {
	// input("hi?")
	stateInitId := time.Now().UnixNano()
	mainStore := state.Make(stateInitId)
	fmt.Println(*(*mainStore)[stateInitId])
	wfHandler := handler.WorkflowHandler{State: mainStore}

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

	wfHandler.Add(workflow)

	fmt.Println(*(*mainStore)[state.GetLstId()])

	work = state.Work{
		Id: 2,
		Script: state.Script{
			Name:       "salam world",
			Path:       "./fifi56.sh",
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

	work2 := state.Work{
		Id: 2,
		Script: state.Script{
			Name:       "salam donia",
			Path:       "./kiki56.sh",
			Created_at: time.Now().GoString(),
			Updated_at: time.Now().GoString(),
		},
		Server: state.Server{
			Name:       "sh543",
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
	workflow = state.Workflow{
		Id:          55,
		Name:        "stop",
		Description: "jaleb Shod",
		Works:       []state.Work{work, work2},
		Created_at:  time.Now().GoString(),
		Updated_at:  time.Now().GoString(),
	}

	wfHandler.Add(workflow)

	for i, S := range *mainStore {
		fmt.Printf("_________________%d_________________", i)
		p, _ := json.MarshalIndent(*S, "$", "  |")
		fmt.Printf("%s \n \n", p)
		fmt.Println("__________________________________")
	}
	fmt.Println((*mainStore))
	// fmt.Println(">>>>>>> 11 >>>>>>>", workflow.Name)
	// fmt.Println(">>>>>>> 22 >>>>>>>", workflow.Description)
	// fmt.Println(">>>>>>> 33 >>>>>>>", workflow.Updated_at, "\t=", workflow.Created_at)
	// fmt.Println(">>>>>>> 44 >>>>>>>", workflow.Description)
	// fmt.Println(" >>>>>>>  55  >>>>>>>", workflow.Name)
	// fmt.Println(">>>>>>> 66 >>>>>>>", workflow.Updated_at, "\t>", workflow.Created_at)
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
