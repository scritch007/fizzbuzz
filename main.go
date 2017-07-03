/*
Start a fizzbuzz server on port 8080

Client should call http://IP:8080/fizzbuzz?string1=fizz&string2=buzz&int1=3&int2=5&limit=100

*/
package main

import (
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

const maxLimit = 20000 //Limit the max value for limit.

//fizzbuzz is done in memory. It could be changed to a stream mechanism to be able to handle more
func fizzbuzz(s1, s2 string, i1, i2, l int) string {
	res := make([]string, l)
	for i := 1; i <= l; i++ {
		value := ""

		//Check if i is a multiple of (i1) int1
		if 0 == math.Mod(float64(i), float64(i1)) {
			//Append string1 (s1) value
			value += s1
		}
		//Check if i is a multiple of (i2) int2
		if 0 == math.Mod(float64(i), float64(i2)) {
			//Append string2 (s2) value
			value += s2
		}

		//It didn't match anything store int value
		if len(value) == 0 {
			value = strconv.Itoa(i)
		}
		res[i-1] = value
	}
	return strings.Join(res, ",")
}

func invalidInput(c echo.Context) error {
	return c.String(http.StatusBadRequest, "Invalid Input")
}

func main() {
	e := echo.New()

	//Create endpoint for the fizzbuzz api
	e.GET("/fizzbuzz", func(c echo.Context) error {
		string1 := c.QueryParam("string1")
		string2 := c.QueryParam("string2")
		strint1 := c.QueryParam("int1")
		strint2 := c.QueryParam("int2")
		strlimit := c.QueryParam("limit")
		//Ensure that every expected fields where provided
		if len(string1) == 0 || len(string2) == 0 || len(strint1) == 0 || len(strint2) == 0 || len(strlimit) == 0 {
			return invalidInput(c)
		}

		//Check that all int fields are actual int
		int1, err := strconv.Atoi(strint1)
		if err != nil {
			return invalidInput(c)
		}
		int2, err := strconv.Atoi(strint2)
		if err != nil {
			return invalidInput(c)
		}
		limit, err := strconv.Atoi(strlimit)
		if err != nil {
			return invalidInput(c)
		}
		if limit < 0 || limit > maxLimit {
			//Only positive limit will work and ensure limit < maxLimit
			return invalidInput(c)
		}
		//Call the fizzbuzz method
		return c.String(http.StatusOK, fizzbuzz(string1, string2, int1, int2, limit))
	})

	//Start on port 8080
	e.Start(":8080")

}
