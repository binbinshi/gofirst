package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"time"
)

func main() {
	//gocron.Every(1).Second().Do(task)
	gocron.Every(1).Minute().Do(task1)

	//_, time := gocron.NextRun()
	//fmt.Println(time)
	<-gocron.Start()

	//s := gocron.NewScheduler()
	//s.Every(2).Seconds().Do(task)
	//<-s.Start()
}

func task(){
	fmt.Println(time.Now(), " task go go go")
}

func task1(){
	fmt.Println(time.Now(), " task jo jo jo")
}