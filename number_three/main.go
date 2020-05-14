package main

import(
	"fmt"
)
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
func main(){
	Room := New()	
	// itteration for walk to north
	for i := 3; i > 0; i--{		
		if Room.array[i][1] == "#"{
			break
		}
		count++
	}	
	fmt.Println(count)
}
