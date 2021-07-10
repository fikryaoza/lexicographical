package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type Characters struct {
	FirstOccurence string
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		fmt.Println(`Reading File ....`)
		fileBytes, err := ioutil.ReadFile("gistfile12.txt")

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// sliceData := strings.Split(string(fileBytes), " ")
		sliceData := string(fileBytes)
		profit := firstOccur(sliceData)
		fmt.Println(`Done ....`)
		return c.JSON(http.StatusOK, profit)
	})

	s := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	e.Logger.Fatal(e.StartServer(s))
}

func firstOccur(data string) Characters {
	var characters Characters
	var occurence string
	var occur = [256]bool{}
	for i := 0; i < len(data); i++ {
		ca := data[i]
		if !occur[ca] {
			occurence += string(ca)
			occur[ca] = true
		}
	}
	characters.FirstOccurence = occurence
	return characters
}
