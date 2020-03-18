package main
import "fmt" //We Can't use '' outside fmt it shows Error
func main(){
fmt.Print("Switch Condition")
var name string
fmt.Print("Enter Name = ")
fmt.Scanf("%s",&name)
switch name{
case "JAIN":
	fmt.Print("Hii Sarthak ")
	
case "SarthakJAIN"	:
	fmt.Print("Jaijinendra JAIN")
}
}