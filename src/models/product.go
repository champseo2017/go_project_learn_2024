package models

type Product struct {
    Id          uint `gorm:"primary_key"` // ไอดีสินค้า เป็น primary key
    Name        string // ชื่อสินค้า  
    Description string // รายละเอียดสินค้า
    Price       float64 // ราคาสินค้า 
}