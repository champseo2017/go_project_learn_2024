# Go Project Learn 2024

โปรเจคนี้เป็นตัวอย่างการสร้าง REST API ด้วย Go และ Fiber Framework

## โครงสร้างโปรเจค

go_project_learn_2024/
│
├── src/
│ ├── main.go # ไฟล์หลักของแอพ
│ │
│ ├── database/
│ │ ├── connect.go # เชื่อมต่อฐานข้อมูล
│ │ └── automigrate.go # สร้างตารางอัตโนมัติ
│ │
│ ├── models/
│ │ ├── user.go # โมเดลผู้ใช้
│ │ ├── cart.go # โมเดลตะกร้าสินค้า
│ │ └── product.go # โมเดลสินค้า
│ │
│ ├── middlewares/ # โฟลเดอร์ใหม่ที่ต้องสร้าง
│ │ └── auth.go # ไฟล์ middleware ตรวจสอบสิทธิ์
│ │
│ └── routes/
│ └── routes.go # จัดการเส้นทาง
│
├── tmp/ # โฟลเดอร์ชั่วคราว
│ └── build-errors.log # ล็อกข้อผิดพลาด
│
├── .air.toml # คอนฟิก air (hot reload)
├── .gitignore # ไฟล์ที่ไม่ต้องการให้ git ติดตาม
├── Dockerfile # สำหรับสร้าง container
├── docker-compose.yml # คอนฟิก docker
├── go.mod # รายการ dependencies
└── go.sum # lock file ของ dependencies
