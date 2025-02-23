package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

// func PrintHello()  {
// 	fmt.Println("Hello,world")
// }
// func PrintCurrentTime()  {
// 	currentTime:=time.Now().Format("15:04:05")
// 	fmt.Printf("Current time :%s\n",currentTime)
// }
// func main()  {
// 	gocron.Every(1).Second().Do(PrintHello)
// 	gocron.Every(10).Seconds().Do(PrintCurrentTime)
// 	<-gocron.Start()
// }


func main() {
	RunCron()
	initiateGin()
}

func RunCron() {
	c := cron.New()

	//@every 00h00m00s
	c.AddFunc("@every 00h00m10s", sentMessage)

	c.Start()
}

func sentMessage() {
	resp, err := http.Get("http://localhost:8001/sent")
	if err != nil {
		fmt.Println("err from cron", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err from read all in cron", err)
	}
	fmt.Println(string(body))
}

func initiateGin() {
	r := gin.Default()
	r.GET("/sent", SenHelloMessage)
	r.Run(":8001")
}

func SenHelloMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "heello"})
}