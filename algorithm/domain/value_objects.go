package domain

type Node struct {
	Id       int
	Adjacent map[int]int
}
