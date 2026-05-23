package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var myName string
	var myName01 string
	var myName02 string
	println("Nhập họ của bạn: ")
	if scanner.Scan() {
		myName01 = scanner.Text()
	}
	println("Nhập tên của bạn: ")
	if scanner.Scan() {
		myName02 = scanner.Text()
	}
	myName = myName01 + " " + myName02
	fmt.Println("- Tên tôi là: ", myName)
	fmt.Println("Số ký tự: ", utf8.RuneCountInString(myName))

	// println("Độ dài chuỗi của myName: ", utf8.RuneCountInString(myName))
	// for i := 0; i < utf8.RuneCountInString(myName); i++ {
	// 	fmt.Println("Ký tự thứ", i, "là:", string(myName[i]))
	// }
}

func BaiTap01() {
	// Câu 1: Tạo map Literal
	inventory := map[string]int{
		"Laptop":   15,
		"iphone":   8,
		"Bàn phím": 50,
	}

	// Câu 2: Cập nhật và xuất dữ liệu
	priceMap := make(map[string]float64)
	for k := range inventory {
		fmt.Println("Hãy nhập giá tiền cho sản phẩm -", k, ": ")
		var cost float64
		fmt.Scan(&cost)
		priceMap[k] = cost
	}
	fmt.Println("===== Tình trạng kho hàng hiện tại =====")
	for i := range inventory {
		fmt.Println("Sản phẩm: ", i, " - Giá tiền: ", priceMap[i], "- Số lượng trong kho: ", inventory[i])
	}

	// Câu 3.1: Lọc ra sản phẩm cần bổ sung gấp (số lượng tồn kho dưới 10)
	// Câu 3.2 Lọc ra sản phẩm cao cấp (Giá tiền > 500)
	warningProducts := []string{}
	luxuryProducts := []string{}
	for key := range inventory {
		if inventory[key] < 10 {
			warningProducts = append(warningProducts, key)
		}
		if priceMap[key] > 500 {
			luxuryProducts = append(luxuryProducts, key)
		}
	}
	fmt.Println("- Sản phẩm cần bổ sung gấp (warningProducts): ", warningProducts)
	fmt.Println("- Sản phẩm sa sỉ (luxuryProducts): ", luxuryProducts)

	// Câu 3.3: Xóa sản phẩm ra khỏi cửa hàng
	var tmp string
	fmt.Println("===== Bạn có muốn xóa sản phẩm không? yes/no =====")
	fmt.Scan(&tmp)
	if tmp == "yes" {
		var tempProduct string
		fmt.Println("===== Nhập tên sản phẩm cần xóa =====")
		fmt.Scan(&tempProduct)
		_, ok := inventory[tempProduct]
		for ok != true {
			fmt.Println("===== Sản phẩm không tồn tại, nhập lại! =====")
			fmt.Scan(&tempProduct)
			_, ok = inventory[tempProduct]
		}
		delete(inventory, tempProduct)
		delete(priceMap, tempProduct)
		fmt.Println("===== Sản phẩm sau khi xóa =====")
		for i := range inventory {
			fmt.Println("Sản phẩm: ", i, " - Giá tiền: ", priceMap[i], "- Số lượng trong kho: ", inventory[i])
		}
	}
}
