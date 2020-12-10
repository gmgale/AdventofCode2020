package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gmgale/AdventofCode2020/secrets"
)

const URL = "https://adventofcode.com/2020/day/1/input"

func main() {

	client := &http.Client{}
	req, _ := http.NewRequest("GET", URL, nil)
	req.AddCookie(&secrets.MyCookie)

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	bodyString := string(body)
	stringNums := strings.Split(bodyString, "\n")

	var numbers = make([]int, len(stringNums))

	for i := range stringNums {
		i := i
		x, _ := strconv.Atoi(stringNums[i])
		numbers = append(numbers, x)
	}

	for i := 0; i < len(numbers)-1; i++ {
		i := i
		for j := i + 1; j < len(numbers)-1; j++ {
			if (numbers[i] + numbers[j]) == 2020 {
				fmt.Printf("The answer is %d.", (i * j))
			}
		}
	}
}
