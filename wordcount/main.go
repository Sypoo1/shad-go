package main 


import (
	"fmt"
	"bufio"
	"os"
)



func helper(err error){
	if err != nil{
		panic(err)
	}
}

func closeFile(file *os.File){
	err := file.Close()
	if err != nil{
		panic(err)
	}

}

func main(){
	args := os.Args

	if len(args) < 2{
		panic("put paths")
	}

	lines_map := make(map[string]int)

	for _, filePath := range(args){
		file, err := os.Open(filePath)
		helper(err)
		
		defer closeFile(file)
		scanner := bufio.NewScanner(file)

		for scanner.Scan(){
			line := scanner.Text()
			lines_map[line] = lines_map[line] + 1
		}
		
		// err = scanner.Err()
		// if err != nil{
		// 	panic(err)
		// }
	}

	for line, count := range lines_map {
		if count >= 2 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}
