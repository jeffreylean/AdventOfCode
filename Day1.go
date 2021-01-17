package main 

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"sort"
)
const target int = 2020
func check(e error){
	if e!=nil{
		panic(e)
	}
}

func GetInput(path string) (result []int){
	file,err:=os.Open(path)
	check(err)
	defer file.Close()
	var lines []int
	scanner:=bufio.NewScanner(file)
	for scanner.Scan(){
		lineStr := scanner.Text()
		num,_:=strconv.Atoi(lineStr)
		lines=append(lines,num)
	}
	return lines
}

func PartOne(input[] int){	
	//create hash map
	hash:=make(map[int]int)
	for i:=0;i<len(input);i++{
		complement:=target - input[i]
		// find if the complement exist in the hashmap 
		if _, ok:=hash[complement];ok{
			fmt.Printf("The Two entries found in the input which sum out is equal to 2020 are %d and %d \n The multiplication between them is %d \n\n",complement,input[i],complement*input[i])
		}
		//if not exist in hashmap, insert into the hash for future reference
		hash[input[i]]=i
	}	
}

func PartTwo(input[] int){
	// sort the slice 
	sort.Ints(input)
	for i:=0;i<len(input)-2;i++{
		var a int = input[i]
		// check if a is larger than 2020. if so move to the next index
		if a>target{
			continue
		}
		// index of second element, start from 1 index away from a's index
		var indexB int = i+1
		// index of third element, start from end of the slice
		var indexC int = len(input)-1


		//use while loop to find the b and c element which able to add up with a which equal to 2020
		for indexC>indexB{
			var b int = input[indexB]
			var c int = input[indexC]
			if a+b+c ==target{
				fmt.Printf("The Three entries found in the input which sum out is equal to 2020 are %d ,%d and %d \n The multiplication between them is %d \n\n",a,b,c,a*b*c)
			}

			// if it is larger than 2020 reduce it by move c's index negatively 
			if a+b+c >=target{
				for input[indexC -1 ]==c {
					indexC--
				}
				indexC--
			}
			// if it is smaller than 2020 increase it by move b's index positively 
			if a+b+c <=target{
				for input[indexB + 1 ]==b {
					indexB++
				}
				indexB++
			}
		}

	}
}

func main(){
	var input []int = GetInput("C:/projects/AdventOfCode/Day1/input.txt")
	PartOne(input)
	PartTwo(input)
}