package main
import("fmt")
func main() {
	x, y, i, z :=0, 1, 0, 0

	for {
		r:=x+y
		x = y
		y = r
		if r > 4000000{
			break
		}
		fmt.Println("	",r)
		if r%2 == 0 {
			z = z + r
		}
		fmt.Println("sum of even numbers = ", z)
		i++
	}
}

