package middlewares

import (
	"fmt"

	database "github.com/champseo2017/go_project_learn_2024/src/database"
	models "github.com/champseo2017/go_project_learn_2024/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5" // เปลี่ยนเป็น v5
)

// IsAuthenticated ตรวจสอบว่าผู้ใช้ล็อกอินหรือไม่
func IsAuthenticated(c *fiber.Ctx) error {
    cookie := c.Cookies("jwt")

    // เปลี่ยนวิธีการใช้งาน jwt ให้เข้ากับ v5
    token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte("secret"), nil
    })

    if err != nil || !token.Valid {
        c.Status(fiber.StatusUnauthorized)
        return c.JSON(fiber.Map{
            "message": "กรุณาเข้าสู่ระบบ",
        })
    }

    return c.Next()
}

// IsAdmin ตรวจสอบว่าผู้ใช้เป็นแอดมินหรือไม่
func IsAdmin(c *fiber.Ctx) error {
    cookie := c.Cookies("jwt")

    // เปลี่ยนวิธีการใช้งาน jwt ให้เข้ากับ v5
    token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte("secret"), nil
    })

    if err != nil {
        c.Status(fiber.StatusUnauthorized)
        return c.JSON(fiber.Map{
            "message": "กรุณาเข้าสู่ระบบ",
        })
    }

    // ดึง claims แบบ v5
    claims := token.Claims.(*jwt.RegisteredClaims)

    // ค้นหาผู้ใช้จาก ID ใน token
    var user models.User
    database.DB.Where("id = ?", claims.Subject).First(&user)

    if !user.IsAdmin {
        c.Status(fiber.StatusUnauthorized)
        return c.JSON(fiber.Map{
            "message": "คุณไม่มีสิทธิ์เข้าถึงส่วนนี้",
        })
    }

    return c.Next()
}


// GetUserId - ดึงรหัสผู้ใช้จาก JWT token
func GetUserId(c *fiber.Ctx) (uint, error) {
    // ดึง token จาก cookie
    cookie := c.Cookies("jwt")

    // ตรวจสอบและถอดรหัส token
    token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, 
        func(token *jwt.Token) (interface{}, error) {
            return []byte("secret"), nil
        },
    )

    if err != nil {
        return 0, err
    }

    // ดึงข้อมูลจาก token
    claims := token.Claims.(*jwt.RegisteredClaims)

    // แปลง ID จาก string เป็น uint
    var id uint
    fmt.Sscanf(claims.Subject, "%d", &id)

    return id, nil
}