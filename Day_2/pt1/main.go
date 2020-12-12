package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	cookie "github.com/gmgale/AdventofCode2020/secrets"
	utils "github.com/gmgale/AdventofCode2020/utils"
)

const taskURL = "https://adventofcode.com/2020/day/2/input"

func main() {
	resp, err := utils.GetData(cookie.MyCookie, taskURL)
	if err != nil {
		log.Printf("Error fetching data. %v\n", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	bodyString := string(body)

	item := strings.Split(bodyString, "\n")

	var count = 0

	for i := 0; i < len(item)-1; i++ {
		j := strings.Split(item[i], " ")
		numSr := strings.Split(j[0], "-")
		var nums [2]int
		nums[0], _ = strconv.Atoi(numSr[0])
		nums[1], _ = strconv.Atoi(numSr[1])

		char := j[1][0:1]
		pass := j[2]

		k := strings.Count(pass, char)
		if k < nums[0] || k > nums[1] {
			continue
		} else {
			count++
		}
	}
	fmt.Printf("The answer is %d.", count)
}
