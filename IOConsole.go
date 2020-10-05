package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	fmt.Println("Begin IO console...")
	//for {
	//	fmt.Print("->")
	//	text, _ := reader.ReadString('\n')
	//	text = strings.Replace(text, "\n", "", -1)
	//	fmt.Println("You have enter: " + text)
	//}
	giaiPhuongTrinhBac2()
}

func giaiPhuongTrinhBac2() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Giai phuong trinh bac 2:")
	fmt.Print("Nhap he so a ->")
	aString, _ := reader.ReadString('\n')
	a, _ := strconv.ParseFloat(strings.Replace(aString, "\n", "", -1), 32)
	fmt.Print("Nhap he so b ->")
	bString, _ := reader.ReadString('\n')
	fmt.Print("Nhap he so c ->")
	b, _ := strconv.ParseFloat(strings.Replace(bString, "\n", "", -1), 32)
	cString, _ := reader.ReadString('\n')
	c, _ := strconv.ParseFloat(strings.Replace(cString, "\n", "", -1), 32)
	denta := math.Pow(b, 2) - 4*a*c
	if a == 0 {
		fmt.Println("Input dau va khong dung")
		return
	}
	if denta < 0 {
		fmt.Println("Phuong trinh vo nghiem")
	} else if denta == 0 {
		fmt.Println("Phuong trinh co nghiem kep")
		fmt.Println("x1=x2= ", -b/(2*a))
	} else if denta > 0 {
		fmt.Println("Phuong trinh co 2 nghiem phan biet")
		x1 := (-b + math.Sqrt(denta)) / (2 * a)
		x2 := (-b - math.Sqrt(denta)) / (2 * a)
		fmt.Println("x1= ", x1)
		fmt.Println("x2= ", x2)
	}
}
