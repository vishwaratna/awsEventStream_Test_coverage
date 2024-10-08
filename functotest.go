package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Cookie represents a cookie with a name, price, and donation amount
type Cookie struct {
	name     string
	price    int
	donation int
}

// parseInput parses the user input and returns the maximum money and list of cookies
func parseInput(scanner *bufio.Scanner) (int, []Cookie) {
	fmt.Println("Enter the maximum money:")
	scanner.Scan()
	maxMoney, _ := strconv.Atoi(scanner.Text())

	var cookies []Cookie
	fmt.Println("Enter the cookies (name,price,donation) and type 'done' when finished:")
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "done" {
			break
		}
		parts := strings.Split(line, ",")
		price, _ := strconv.Atoi(parts[1])
		donation, _ := strconv.Atoi(parts[2])
		cookie := Cookie{name: parts[0], price: price, donation: donation}
		cookies = append(cookies, cookie)
	}

	return maxMoney, cookies
}

// knapsack computes the maximum donation amount using dynamic programming
func knapsack(maxMoney int, cookies []Cookie) int {
	dp := make([]int, maxMoney+1)

	for i := 1; i <= maxMoney; i++ {
		for _, cookie := range cookies {
			if cookie.price <= i {
				dp[i] = max(dp[i], dp[i-cookie.price]+cookie.donation)
			}
		}
	}

	return dp[maxMoney]
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	maxMoney, cookies := parseInput(scanner)
	result := knapsack(maxMoney, cookies)
	fmt.Printf("Maximum donation amount: %d\n", result)
}
