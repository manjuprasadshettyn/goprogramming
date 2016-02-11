package main 
import "fmt" 
func main() { 
    var a [5]int 

/* to see the contents of the array after declaration */
    fmt.Println("empty array:", a) 
    a[4] = 100 
    fmt.Println("after setting a value in an array:", a) 
    fmt.Println("getting the value of the position which is set:", a[4]) 
    fmt.Println("lenght of the array is:", len(a)) 

/* declaration of a new array*/

    b := [5]int{1, 2, 3, 4, 5} 
    fmt.Println("values of array whose values are initialized during declaration is:", b) 
  
/* two dimensional array*/
 
 var twoD [2][3]int 
    for i := 0; i < 2; i++ { 
        for j := 0; j < 3; j++ { 
            twoD[i][j] = i + j 
        } 
    } 
    fmt.Println("two dimensional array is: ", twoD) 
} 
