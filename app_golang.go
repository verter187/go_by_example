package main

import "fmt"

func main() {
	// nums := []int{2, 3, 4}

	// rand.Seed(time.Now().UnixNano())

	// for i := 0; i < 100000000; i++ {
	// 	z := rand.Intn(99)
	// 	nums = append(nums, z)
	// }
	// fmt.Println("nums = ", len(nums))
	// sum := 0
	// for _, num := range nums {
	// 	sum += num
	// }
	// fmt.Println("sum = ", sum)

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	//num - это копия значения, а не ссылка! Если нужно поменять значение, то
	//nums[i]= 5, а не num = 5
	for i, num := range nums {

		if num == 3 {
			fmt.Println("index:", i)
		}
		nums[i] = 3
	}
	fmt.Println("nums:", nums)

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k, v := range kvs {
		fmt.Println("key:", k, "value:", v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

	for _, c := range "go" {
		fmt.Printf("%c\n", c)
	}
}
