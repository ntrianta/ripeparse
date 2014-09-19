package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
)


func main(){
	
	
	// type Rtt struct{
// 		Val float64
// 	}
	
	type Input struct {
		Name int64 `json:"af"`
	    Result []struct{
	    	Val float64 `json:"rtt"`
	    } `json:"result"`	    
	}
    t := 0
	fi, err := ioutil.ReadFile("/Users/nick/Desktop/test.json")
	out := []Input{}
	err = json.Unmarshal([]byte(fi), &out)	
	for _, i := range out{
		for _, j := range i.Result{
			if j.Val > 60 {
				t++
			}
		}
	}
	fmt.Println(t)
	fmt.Println(err)
}