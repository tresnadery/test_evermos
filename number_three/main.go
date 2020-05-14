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
		if Room.array[i][2] == "."{			
			for j := 2; j < 7; j++{				
				if Room.array[i][j] == "#"{
					break
				}
				count++					
				if Room.array[i + 1][j] == "."{			
					for k := (i+1); k < 6; k++{
						if Room.array[k][j] == "#"{
							break
						}
						count++
					}		
				}					
			} 
		}
	}	
	fmt.Println(count)
}
