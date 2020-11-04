package main

import (
	"errors"
	"fmt"
)

//User stric
type User struct {
	ID        int
	firstname string
	lastName  string
}

func main() {
	// var nama string = "Kato"
	// age := 20
	// age = 19
	// fmt.Println(nama)
	// fmt.Println(age)
	// if age > 10 {
	// 	fmt.Println("boleh main")
	// }

	// arr := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	// for i, v := range arr {
	// 	fmt.Println(i, v)
	// }

	// var myMap map[string]string
	// myMap = make(map[string]string)
	// myMap["key"] = "value"

	// luas, keliling := calculate(10, 5)
	// luas2, _ := calculate(10, 2)
	// fmt.Println(luas)
	// fmt.Println(keliling)
	// fmt.Println(luas2)

	// fmt.Println(sum([]int{1, 2, 3, 4, 5}))
	// result, err := kalkulator(1, 4, "e")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(result)

	// user := User{ID: 1,
	// 	firstname: "alu",
	// 	lastName:  "card"}

	// // fmt.Println(displayUser(user))
	// fmt.Println(user.display())

	// var varA int = 5
	// var varB *int = &varA

	// fmt.Println(varA)
	// fmt.Println(varB)
	// fmt.Println(*varB)

	// varA = 20

	// fmt.Println(varA)
	// fmt.Println(varB)
	// fmt.Println(*varB)

	//case pointer
	// number := 5
	// fmt.Println("alamat memory: ", &number)
	// fmt.Println("nilai awal: ", number)

	// change(&number, 100)

	// fmt.Println("alamat memory: ", &number)
	// fmt.Println("nilai akhir: ", number)

	//struct pinter method
	user := User{ID: 1,
		firstname: "alu",
		lastName:  "card"}
	fmt.Println(user.firstname)
	user.graduate()
	fmt.Println(user.firstname)

}

//---------------------------------------------------------------------------------

func (user *User) graduate() {
	user.firstname = user.firstname + "okesip"
}

func change(old *int, new int) {
	*old = new
	fmt.Println("alamat memory: ", old)
	fmt.Println("nilai didalam function: ", *old)
}

//struct as parameter
func displayUser(user User) string {
	return fmt.Sprintf("nama: %s %s", user.firstname, user.lastName)
}

//struct extension method
func (user User) display() string {
	return fmt.Sprintf("nama: %s %s", user.firstname, user.lastName)
}

//function with input argument and output
func calculate(panjang int, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	return luas, keliling
}

//predefine returen value
func calculatePredefineReturnValue(panjang int, lebar int) (luas int, keliling int) {
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)
	return
}

func sum(numbers []int) int {
	var total int
	for _, number := range numbers {
		total += number
	}
	return total
}

func kalkulator(angka1 int, angka2 int, operator string) (result int, err error) {
	switch operator {
	case "+":
		result = angka1 + angka2
	case "-":
		result = angka1 - angka2
	case "*":
		result = angka1 * angka2
	case "/":
		result = angka1 / angka2
	default:
		err = errors.New("operator salah")
	}
	return
}
