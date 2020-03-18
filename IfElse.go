package main
import ("fmt")
func main(){

	var i int
	fmt.Print("\nEnter a value")
	fmt.Scanf("%d",&i) //fmt.Scan(&i)
	if (i<10){//We Can also use () Outside the Condition
		fmt.Print("If Statement is  running = ",i)
	}else if i>=10 && i<20{
		fmt.Print("Else If Statement Execute ")
		}else{
		fmt.Print("Else")
	}
	
}