package gaussianfunction

//"Sum of the First n Natural Numbers" formula or "Gauss's formula." 
func GuassianFunction(s []int) int {
	return int(len(s) * (2 * s[0] + (len(s) -1)) / 2)
}
