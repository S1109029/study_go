package iftest


func Iftest(x int,y int) string{
	var result string
	if x > y {
		result = "Big"
	} else if x < y {
		result = "Small"
	} else {
		result = "Equal"
	}
	return result
}