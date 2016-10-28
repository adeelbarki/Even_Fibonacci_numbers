package main
import("fmt")
//600851475143
func main(){
	x, f, i := 600851475143, 0, 2

	for i<=x{
		if x%i ==0{
			for x%i == 0{
				x/=i
			}
			f = i
		}
		i++
	}
	fmt.Println("the largest factor is", f)

}
