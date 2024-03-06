package handlers

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func validationErrorResponse(c echo.Context, err error) error {
	// สร้าง map หรือ struct สำหรับข้อผิดพลาดที่ต้องการส่งกลับ
	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		// สร้าง key ใหม่จากชื่อ field ที่ผิดพลาด โดยแปลงจาก camel case เป็น snake case
		field := camelToSnake(err.Field())
		errors[field] = validationErrorToMessage(err)
	}

	// ส่ง response กลับไปยัง client พร้อมกับข้อผิดพลาด
	return c.JSON(400, map[string]interface{}{
		"error":   "Validation failed",
		"details": errors,
	})
}

func validationErrorToMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "max":
		return fmt.Sprintf("The value exceeds the maximum allowed limit of %s", err.Param())
	case "min":
		return fmt.Sprintf("The value is less than the minimum required limit of %s", err.Param())
	case "gender":
		return "The value must be either 'female' or 'male'"
	default:
		return fmt.Sprintf("Validation failed on the '%s' tag", err.Tag())
	}
}

func camelToSnake(str string) string {
	// ใช้ strings.Builder เพื่อประสิทธิภาพในการสร้างสตริง
	var result strings.Builder
	for i, r := range str {
		// ถ้าเป็นตัวพิมพ์ใหญ่และไม่ใช่ตัวแรก, ใส่ _ ก่อนตัวอักษร
		if unicode.IsUpper(r) && i > 0 {
			result.WriteRune('_')
		}
		// แปลงตัวอักษรเป็นพิมพ์เล็กและเพิ่มลงใน result
		result.WriteRune(unicode.ToLower(r))
	}
	return result.String()
}
