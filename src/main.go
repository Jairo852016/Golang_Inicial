package main

import (
	pk "curso_golang_platzi/src/mypackage"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func normalFunction(message string) {
	fmt.Println(message)
}
func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}
func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}
func areaRectanguloFun(baseRentangulo float64, alturaRectangulo float64) float64 {

	return baseRentangulo * alturaRectangulo
}
func isPair(a int) bool {
	return a%2 == 0
}
func userValid(user, pass string) bool {
	return user == "Jairo" && pass == "Juan"
}
func isPalindromo(text string) {
	var textReverse string
	text = strings.ToLower(text)
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	if text == textReverse {
		fmt.Println("Es Palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}

type car struct {
	brand string
	year  int
}
type cuadrado struct {
	base float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

type rectangulo struct {
	base   float64
	altura float64
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

type figuras2D interface {
	area() float64
}

func calcular(f figuras2D) {
	fmt.Println("Area:", f.area())
}
func main() {
	//Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.1416

	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	//Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//Zero values
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)

	// Area de Cuadrado

	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area Cuadrado: ", areaCuadrado)

	x := 10
	y := 50
	//Suma
	result := x + y
	fmt.Println("Suma: ", result)

	//Resta
	result = y - x
	fmt.Println("Resta: ", result)

	//Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion: ", result)

	//Divicion

	result = y / x
	fmt.Println("Division: ", result)

	// Modulo
	result = y % x
	fmt.Println("Modulo: ", result)

	//Incremental
	x++
	fmt.Println("Incremental: ", x)

	//Decremental
	x--
	fmt.Println("Decremental: ", x)

	// Area Rentangulo
	areaRectangulo := areaRectanguloFun(10, 50)
	fmt.Println("Area Rectangulo:", areaRectangulo)

	// Area Trapecio
	var baseUno float64 = 6.5
	var baseDos float64 = 5.5
	var alturaTrapecio float64 = 25

	areaTrapecio := ((baseUno + baseDos) * alturaTrapecio) / 2

	fmt.Println("Area Trapecio: ", areaTrapecio)

	//Area Circulo
	const PI float64 = 3.14
	var radioCirculo float64 = 10

	areaCirculo := PI * (radioCirculo * radioCirculo)
	fmt.Println("Area Circulo: ", areaCirculo)

	//Declaracion de variables
	helloMessage := "Hello"
	worldMessage := "world"

	//Println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	//Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos \n", nombre, cursos)
	fmt.Printf("%v tiene mas de %d cursos \n", nombre, cursos)

	//Sprintf

	message := fmt.Sprintf("%s tiene mas de %d cursos ", nombre, cursos)
	fmt.Println(message)
	//Tipo dato
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T", cursos)

	//Funciones
	fmt.Println("Hola mundo")
	fmt.Println("Hola mundo 2")
	fmt.Println("Hola mundo 3")

	normalFunction("Hola mundo")
	tripleArgument(1, 2, "Hola")

	value := returnValue(2)

	fmt.Println("Value: ", value)
	value1, value2 := doubleReturn(2)
	fmt.Println("value1, value2:", value1, value2)
	value1, _ = doubleReturn(2)

	//for condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}
	//For forever
	//counterForever := 0
	//for {
	//	fmt.Println(counterForever)
	//	counterForever++
	//}
	valor1 := 1
	valor2 := 2
	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}
	// with and
	if valor1 == 1 && valor2 == 3 {
		fmt.Println("Es verdad")
	}
	//with or
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Es verdad, OR")
	}
	// convertir text a numero
	value, err := strconv.Atoi("4")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Es numero:", value)
	// funcion calcula par o impar
	if isPair(5) {
		fmt.Println("Es par")
	} else {
		fmt.Println("Es impar")
	}
	//funcion que valida usuario y pass
	if userValid("Jairo", "Juan") {
		fmt.Println("Valido")
	} else {
		fmt.Println("Invalido")
	}
	// switch

	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}
	// switch sin condicion
	value4 := 99
	switch {
	case value4 > 100:
		fmt.Println("Es mayor a 100")
	case value4 < 0:
		fmt.Print("Es menor a 0")
	default:
		fmt.Println("No condicion")
	}
	//defer
	fmt.Println("Hola")
	fmt.Println("Mundo")

	defer fmt.Println("Hola")
	fmt.Println("Mundo")
	//continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		//continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}
		//break
		if i == 8 {
			fmt.Println("Break")
			break
		}
	}
	//array
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))
	//slice
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println(slice, len(slice), cap(slice))
	//metodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])
	// append
	slice = append(slice, 88)
	fmt.Println(slice)
	// append nueva list
	newSlice := []int{9, 10, 11, 12}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
	// slice range
	slice1 := []string{"Hola", "que", "Haces"}
	for i, valor := range slice1 {
		fmt.Println(i, valor)
	}
	for _, valor := range slice1 {
		fmt.Println(valor)
	}
	for i := range slice1 {
		fmt.Println(i)
	}
	//palindromo ama, amor a roma, casa
	isPalindromo("Ama")
	// clave valor
	m := make(map[string]int)
	m["Jose"] = 14
	m["Pepito"] = 20
	fmt.Println(m)
	//Recorrer map range
	for i, v := range m {
		fmt.Println(i, v)
	}
	// encontrar un valor
	value6, ok := m["Jose"]
	fmt.Println(value6, ok)

	//Clases o estructuras
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)
	// otra forma
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
	//Modificadores de acceso
	var myCard pk.CarPublic
	myCard.Brand = "Corsa"
	myCard.Year = 2020
	fmt.Println(myCard)

	pk.PrintMessage("Platzi")
	// punteros
	a1 := 50
	b1 := &a1

	fmt.Println(b1)
	fmt.Println(*b1)

	*b1 = 100
	fmt.Println(a1)

	// instancias pc

	myPc := pk.Pc{Ram: 16, Disk: 200, Brand: "msi"}
	fmt.Println(myPc)

	myPc.Ping()
	fmt.Println(myPc)
	myPc.DuplicateRam()
	fmt.Println(myPc)
	myPc.DuplicateRam()
	fmt.Println(myPc)
	//Stringers

	fmt.Println(myPc, "Fin")

	//
	myCuadrado := cuadrado{base: 2}
	myRectangulo := rectangulo{base: 2, altura: 4}

	calcular(myCuadrado)
	calcular(myRectangulo)
	//lista de interfaces
	myInterface := []interface{}{"hola", 12, 4.9}
	fmt.Println(myInterface...)

}
