package state

import "fmt"

type Store struct {
	Job      []Job
	Workflow []Workflow
	Server   []Server
	Script   []Script
}

func Make() *Store {
	thisState := Store{}
	go func(thisState *Store) {
		for {
			fmt.Println(&thisState)
			s
		}

	}(&thisState)
	return &thisState
}
