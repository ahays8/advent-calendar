package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
	"slices"
  "strings"
)

func sliceAtoi(sa []string) []int {
  si := make([]int, 0, len(sa))
  for _, a := range sa {
    i, err := strconv.Atoi(a)
    if err == nil {
      si = append(si, i)
    }

  }
  return si
}
func intAbs(x int) int {
  if x < 0 {
    return -x
  }
  return x
}
func isSafe(list []int) bool{
  if(len(list)<2){
    return true
  }
  curr:=list[0]
  inc:=(list[1]-list[0])>0
  for _, next := range list[1:] {
    dif:=next-curr
    if(((dif>0)!=inc)||(intAbs(dif)<1)||(intAbs(dif)>3)){
      return false
    }
    curr=next
  }
  return true
}
func parseLists() int {
  res := 0
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
		okay:=false
    dummy := sliceAtoi(strings.Split(scanner.Text(), " "))
    //fmt.Println(dummy,"=",isSafe(dummy))
		for i:=range dummy{
			if(isSafe(slices.Concat(dummy[:i],dummy[i+1:]))){
				okay=true
			}
		}
    if(okay){
      res += 1
    }
  }
  return res
}
func main() {
  fmt.Println(parseLists())
}
