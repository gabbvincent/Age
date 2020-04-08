//Name: Vincent G.
//Date: 4/7/2020
//Descritpion: Age



package main

import "fmt"
//create function age(year int)
func age(year int){
  //Declare variable for Date as integer with value of 2020
  var Date int = 2020
  //Declare variable for Age as integer with a value of Date - year
  var Age int = Date - year
  //Print message telling user their age
  fmt.Println("If you were born in", year,"you would be",Age,"in",Date,".")
}






func main() {
  //declare variable for i 
  var i int
  //ask user for birth year,scan for i, call age(i)
  fmt.Println("Enter in a year before 2020 to find out how old someone would be today if they were born in that year.")
  fmt.Scanln(&i)
  age(i)
}