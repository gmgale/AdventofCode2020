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
		p1 := string(pass[nums[0]-1])
		p2 := string(pass[nums[1]-1])

		internalCount := 0

		// Non-zero indexing!
		if p1 == char {
			internalCount++
		}
		if p2 == char {
			internalCount++
		}

		if internalCount == 1 {
			count++
		}
	}
	fmt.Printf("The answer is %d.", count)
}
