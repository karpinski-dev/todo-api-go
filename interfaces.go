package main

type DataRepository interface {
	GetAll() []TodoItem
	GetDone() []TodoItem
	GetPending() []TodoItem
}
