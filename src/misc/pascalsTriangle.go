// To execute Go code, please declare a func main() in a package "main"

// Pascal's triangle

// n = 2
// 1
//1 1

// n = 4
//   1
//  1 1
// 1 2 1
//1 3 3 1

// n = 5
//    1
//   1 1
//  1 2 1
// 1 3 3 1
//1 4 6 4 1

// given that you just printed lines 0 through x - 1, how do print the numbers on line x?

// row r, col c
// row 3, col 1 = (2, 0) + (2, 1)
// row r, col c => (r - 1, c - 1) + (r - 1, c)

/*
  bottom row = 0 spaces
  top row = n-1 spaces
  first row will always be 1
  always print 1 -> add 2 numbers above to get middle values -> always print 1 
  each row will have have as many values as that row number
  
  1 [1,1] => 2
  [1,2,1]
*/



package main

import(
  "fmt"
  "strconv"
)

func createPascalsTriangle(n int) { 
  var spaces string
  
  for i := 0; i < n-1; i++ {
    spaces += " "
  }
  
  var nums []int
  
  for i := 0; i < n; i++ {
    var tempNums []int
    fmt.Print(spaces)
    fmt.Print("1 ")
    tempNums = append(tempNums, 1)
    
    for j := 0; j < i - 1; j++ {
      fmt.Print(strconv.Itoa(nums[j] + nums[j+1]) + " ")
      tempNums = append(tempNums, nums[j] + nums[j+1])
    }
    
    if i != 0 {
      fmt.Print("1")
      tempNums = append(tempNums, 1)
    }
    fmt.Print("\n")
    nums = tempNums
    
    if len(spaces) != 0 {
      spaces = spaces[:len(spaces)-1]
    }
  }
}

func main() {
  createPascalsTriangle(6);
}