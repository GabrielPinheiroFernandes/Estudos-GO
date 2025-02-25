package main

func returnMapNameAndAge(name string, age int8) map[string]int8 {
	temp := make(map[string]int8)
	temp[name] = age
	return temp
}