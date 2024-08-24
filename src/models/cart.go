package models

type Cart struct {
    Id       uint      `gorm:"primaryKey"` // Id ของ Cart เป็น primary key
    Products []Product `gorm:"many2many:cart_product; ForeignKey:Id; References:Id"` 
    // Cart เชื่อมกับ Product ด้วยความสัมพันธ์แบบ many-to-many ผ่านตารางกลางชื่อ cart_product
    // ForeignKey:Id หมายถึง Id ของ Cart ในตารางกลาง
    // References:Id หมายถึง Id ของ Product ในตารางกลาง
}  