package main

import "fmt"

// ===========================================
// 1. Tính đóng gói (encapsulations)
// - Dùng struct thay cho class
//	+ Public: Viết HOA chữ cái đầu tiên trong tên biến
//	+ Private: Viết THƯỜNG chữ cái đầu tiên trong tên biến
// --> 1.1 Định nghĩa cấu trúc dữ liệu (Class)
type Vehicle struct {
	Brand string
	model string
	speed int
}

// --> 1.2 Hàm khởi tạo
func NewVehicle(brand, model string) Vehicle {
	return Vehicle{Brand: brand, model: model, speed: 0}
}

// --> 1.3 Method của vehicle
func (v *Vehicle) Accelerate(amount int) {
	v.speed += amount
	println("%s %s đang tăng tốc lên đến %d km/h", v.Brand, v.model, v.speed)
}

// ===========================================
// 2. Tính kế thừa (inheritance) --> KHÔNG TƯỜNG MINH
// - Golang KHÔNG có cấu trúc "extends", thay vào đó ta nhúng trực tiếp vào struct
type Car struct {
	Vehicle          // inheritance từ Vehicle
	Fueltype  string // Loại nhiên liệu
	SeatCount int    // Số chỗ ngồi
}
type Motobike struct {
	Vehicle
	CylinderCapacity int // Dung tích xi lanh
}

func (myCar Car) ShowInfor() {
	fmt.Printf(
		"Hãng xe: %s\nTên xe: %s\nLoại xe: %s\nChỗ ngồi: %d\n",
		myCar.Brand,
		myCar.model,
		myCar.Fueltype,
		myCar.SeatCount,
	)
}

// ===========================================
// 3. Tính đa hình (polymorphism)
// - Interface trong Golang không chứa thuộc tính, chỉ chứa danh sách các hàm
// - Không cần khai báo "implements"
// - Cách khởi tạo interface:
type Moveable interface {
	Move()
}

func StartJourney(m Moveable) {
	m.Move() // Tính đa hình ở đây
}
func (c Car) Move() {
	println("Xe oto đi bằng 4 bánh")
}
func (m Motobike) Move() {
	println("Xe máy đi bằng 2 bánh")
}

// ===========================================
func main() {
	myCar := Car{
		Vehicle:   NewVehicle("Toyota", "Camry"),
		Fueltype:  "Xăng",
		SeatCount: 7,
	}
	myMotobike := Motobike{
		Vehicle:          NewVehicle("Yamaha", "NVX"),
		CylinderCapacity: 155,
	}
	StartJourney(myCar)
	StartJourney(myMotobike)
	// myCar.ShowInfor()
}
