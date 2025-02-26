package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
)

func main(){
  scanner := bufio.NewScanner(os.Stdin)
  for ;; {
    fmt.Print("Pokedex > ")
    open := scanner.Scan()
    text := scanner.Text()
    firstWord := strings.Fields(text)[0]
    fmt.Println("Your command was:",strings.ToLower(firstWord))
    if open == false {
      break
    }
  }
  
}
