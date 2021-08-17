package main

func main() {
	// resp, err := http.Get("https://adventofcode.com/2020/day/1/input")
	// if err != nil {
	// 	fmt.Println("Error getting input")
	// }
	// defer resp.Body.Close()
	// input, err := ioutil.ReadAll(resp.Body)
	// inputSlice := strings.Split(string(input), "\n")

	// for _, i := range inputSlice {
	// 	fmt.Printf("%v", i)
	// }
	input := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	result := make(map[int]int)
	max := len(input)
	// we need to keep adding the numbers together only once 
	for i, v := input {
		result[v] = 0
	}

}
