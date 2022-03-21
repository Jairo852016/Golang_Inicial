package mypackage

import "fmt"

//CarPublic Car con acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

//PrintMessage imprimir un mensaje
func PrintMessage(text string) {
	fmt.Println("Hola ", text)
}

func printMessage1(text string) {
	fmt.Println("Hola ", text)
}

//Pc struct PC
type Pc struct {
	Ram   int
	Disk  int
	Brand string
}

//Ping retorna marca y pong
func (myPC Pc) Ping() {
	fmt.Println(myPC.Brand, "Pong")
}

//DuplicateRam dupllica el valor de la ram
func (myPC *Pc) DuplicateRam() {
	myPC.Ram = myPC.Ram * 2
}
func (myPC Pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPC.Ram, myPC.Disk, myPC.Brand)
}
