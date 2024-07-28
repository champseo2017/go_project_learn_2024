package models

type Cart struct {
    Id       uint      `gorm:"primaryKey"` // ไอดีของตะกร้า เป็น primary key
    Products []Product `gorm:"many2many:cart_product; ForeignKey:Id"` // สินค้าในตะกร้า เชื่อมกับ Product ด้วย many to many
}  