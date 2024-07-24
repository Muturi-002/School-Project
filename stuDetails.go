package students
import (
	"fmt"
	"strings"
	"strconv"
	"time"
)
var firstName, lastName, contactNo, emailAddress string
//var dateReg, dateBirth int64
func Details(){
	fmt.Println("Enter First Name: ")
	fmt.Scanln(&firstName)
	fmt.Println("Enter Last Name: ")
	fmt.Scanln(&lastName)
	/* fmt.Println("Enter date of birth(dd/mm/yyyy): ")
	fmt.Scanln(&dateBirth)
	dateBirth= date.Format("12/03/2000") */
	fmt.Println("Enter Contact No: ")
	fmt.Scanln(&contactNo)
	fmt.Println("Enter Email Address: ")
	fmt.Scanln(&emailAddress)

}
func dates(){
	var year,month, date int16
	fmt.Println("Enter date of Birth: ")
	fmt.Scanln(&date)
	fmt.Println("Enter month of Birth: ")
	fmt.Scanln(&month)
	fmt.Println("Enter year of Birth: ")
	fmt.Scanln(&year)
	dateBirth := time.Date(year, month, date)
	fmt.Println(dateBirth)	
}