package utils

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
)

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func IsPerfect(n int) bool {
	sum := 0

	for i := 1; i < n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}

	return sum == n
}

func IsArmstong(n int) bool {
	sum, temp := 0, n
	digits := int(math.Log10(float64(n))) + 1

	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		temp /= 10
	}
	return sum == n
}

func ClassifyProperties(n int) []string {
	prop := []string{}

	if IsArmstong(n) {
		prop = append(prop, "armstrong")
	}
	if n%2 == 0 {
		prop = append(prop, "even")
	} else {
		prop = append(prop, "odd")
	}

	return prop

}

func CalculateDigitSum(n int) int {
	sum := 0
	n = int(math.Abs(float64(n)))
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func GetFuncFact(n int) (string, error) {
	url := fmt.Sprintf("http://numbersapi.com/%d/math", n)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("unable to fetch fun fact err:%v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", fmt.Errorf("could not read response")

	}

	return string(body), nil
}
