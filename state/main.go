package state

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

var lstId int64
var initId int64

type St *map[int64]*Store

type Store struct {
	Job      []Job      `json:"job"`
	Workflow []Workflow `json:"workflow"`
	Server   []Server   `json:"server"`
	Script   []Script   `json:"script"`
}

func Init() *Store {
	fmt.Println("store")
	StateZero := Store{}
	fmt.Println("store", StateZero)
	return &StateZero
}

func GetLstId() int64 {
	return lstId
}

func Make(Id int64) St {
	lstId = Id
	initId = Id
	thisState := map[int64]*Store{}
	thisState[initId] = Init()
	return &thisState
}

func CloneStore(s *Store) Store {
	clone := *s
	return clone
}

func Mutate(state St, new Store) St {
	Lchng := time.Now().UnixNano()

	(*state)[Lchng] = &new
	lstId = Lchng
	return state
}

func saveState(f *os.File, s Store) error {
	content, err := json.Marshal(s)
	if err != nil {
		return err
	}
	f.WriteString(string(content))
	return nil
}

func Save(state Store) error {
	dist := fmt.Sprintf("./%d_%v_state.json", lstId, time.UTC)

	if _, err := os.Stat(fmt.Sprintf(dist, lstId, time.UTC)); errors.Is(err, os.ErrNotExist) {
		File, err := os.Create(dist)
		if err != nil {
			return err
		}
		defer File.Close()
		saveState(File, state)
	} else {
		File, err := os.Open(dist)
		if err != nil {
			return err
		}
		defer File.Close()
		saveState(File, state)
	}

	return nil
}

//TODO: load previous state of program

func Load() {

}
