package main

import (
	"fmt"
	"time"
)

func main()  {
	myLunar , err := SolarTimeToLunar(time.Now().AddDate(0,1,0))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%#v\n", myLunar)

	myLunar2 , err := SolarToLunar(2018, 5,20)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%#v\n", myLunar2)
}