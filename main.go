package main

func data() (map[int]bool, int) {

	len := 100
	data := map[int]bool{
		7:  true,
		9:  true,
		30: true,
		35: true,
		40: true,
		60: true,
		80: true,
		90: true,
	}

	data[3] = true

	// fmt.Println(data, len)
	return data, len
}

func findRangeData(data map[int]bool, int){
	len := 
}

type master struct {
	root *node
}

type node struct {
	left     *node
	right    *node
	maxIndex int
	minIndex int
}

func rangeValue(valueMin int, valueMax int) (leftMin int, leftMax int, rightMin int, rightMax int) {
	leftMin = valueMin
	median := (valueMax - valueMin) / 2
	leftMax = leftMin + median
	rightMin = leftMax
	rightMax = valueMax
	return
}

func NewMaster(min, max int) master {
	var vmaster master
	vmaster.root.maxIndex = min
	vmaster.root.maxIndex = max
	return vmaster
}

func main() {
	// datas, y := data()
	root := NewMaster(0, 100)

	// fmt.Println(datas, y)
}
