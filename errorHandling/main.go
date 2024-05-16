package main

import "fmt"

func divide(a,b float64) (float64,error) {
	if b == 0 {
		return 0,fmt.Errorf("denominator must not be zero");
	}
	return a/b, nil;
}
func main() {
	// ans,err := divide(10,0);
	// fmt.Println(ans);
	// fmt.Printf("Data is %.4f\n",ans);
	// if(err != nil){
	// 	fmt.Println("Error Handling");
	// }

	ans,_ := divide(10,0);
	fmt.Println(ans);
	fmt.Printf("Data is %.4f\n",ans);
}