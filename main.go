package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
//	"net/http"
)

func main(){
	
	type Input struct {
	    Result []struct{
	    	Rtt float64 `json:"rtt"`
	    } `json:"result"`	    
	}
	
    t := 0
	
	input, err := ioutil.ReadFile("/Users/nick/Desktop/test.json")
	
	output := []Input{}
	
	err = json.Unmarshal([]byte(input), &output)	
	
	if err!=nil {
		panic(err)
	}
	
	for _, i := range output{
		for _, j := range i.Result{
			if j.Rtt > 60 {
				t++
			}
		}
	}
	fmt.Println(t)
}