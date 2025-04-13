package main

type TodoItem struct {
	Title   string `json:"title"`
	Details string `json:"details"`
	Status  bool   `json:"status"`
}
