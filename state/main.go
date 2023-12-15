package state

import (
	"fmt"
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
