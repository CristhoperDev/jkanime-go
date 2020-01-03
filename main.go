package main

import (
	"github.com/labstack/echo"
	"github.com/subosito/gotenv"
	"io"
	"jkanime-go/internal/handler"
	"log"
	"math/rand"
	"os"
	"time"
)

var startTime string

//Error log
var (
	publicEndpoints = make(map[string]string)
	Error           *log.Logger
)

//Init app
func Init(errorHandle io.Writer) {
	gotenv.Load()
	rand.Seed(time.Now().UTC().UnixNano())
	//Trace = log.New(traceHandle,
	//  "TRACE: ",
	//  log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}


func main() {

	// startup code
	startTime = time.Now().String()
	programLog, err := os.OpenFile("main.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer programLog.Close()
	Init(programLog)
	log.SetOutput(programLog)

	e := echo.New()
	e.GET("/", handler.LastAnimeEcho)
	e.GET("/anime/:id", handler.GetContentInformation)
	e.Logger.Fatal(e.Start(":3000"))
}