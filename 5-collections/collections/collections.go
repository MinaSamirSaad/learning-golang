package collections

var Countries [10]string
var Slice []string
var Map map[string]string

func init() {
	Countries = [10]string{"USA", "Canada", "Mexico"}
	Countries[3] = "Brazil"
}
func init() {
	Slice = []string{"USA", "Canada", "Mexico"}
	Slice = append(Slice, "Brazil")
}

func init() {
	Map = map[string]string{
		"US": "United States",
		"CA": "Canada",
		"MX": "Mexico",
	}
	Map["BR"] = "Brazil"
}
