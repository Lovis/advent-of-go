package main

import "fmt"

var data = []int{1,1,1,4,99,5,6,0,99}



func main() {
	fmt.Println(data)
	i := 0
	for {
		fmt.Println(i)
		if (i >= len(data)) {
			break
		}
		if (data[i] == 1 || data[i] == 2) {
			fmt.Println("found a 1 or 2", data[i])
			fmt.Println("appending!", data[i], data[i]+3)

			if (i+3 <= len(data)) {
				fmt.Println("code serie", data[i], data[i+1], data[i+2], data[i+3])
				if (data[i] == 1) {
					data[data[i+3]] = data[data[i+1]] + data[data[i+2]]
				}
				if (data[i] == 2) {
					data[data[i+3]] = data[data[i+1]] * data[data[i+2]]
				}
				i=i+4
				continue
			}
		}
		if (i == 99) {
			fmt.Println("found a 99:", i)
			break
		}
		
		if (i < len(data)) {
			fmt.Println("reached end, incrementing by 1")
			i++	
		} 
	}


	fmt.Println(data)
}

