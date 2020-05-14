package main

// Room Structure
type Room struct{
	array [][]string
}
// Count step to find the key
var count int = 0
// New factory to generate room
func New()*Room{
	return &Room{
		[][]string{
			{"#","#","#","#","#","#", "#", "#"},
			{"#",".",".",".",".",".", ".", "#"},
			{"#",".","#","#","#",".", ".", "#"},
			{"#",".",".",".","#",".", "#", "#"},
			{"#","x","#",".",".",".", ".", "#"},
			{"#","#","#","#","#","#", "#", "#"},	
		},
	}
}
