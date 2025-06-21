package utils

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func StringInSlice(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}

// สำหรับเช็คค่า slice ว่ามี string ของเราอยู่ในนั้นหรือไม่
func Contains[V string | []string](s []string, val V) bool {
	if val, ok := any(val).(string); ok {
		for _, v := range s {
			if v == val {
				return true
			}
		}
	}

	if val, ok := any(val).([]string); ok {
		check := []string{}
		for _, str := range val {
			for _, v := range s {
				if v == str {
					check = append(check, str)
					continue
				}
			}
		}
		if len(check) == len(val) {
			return true
		}
	}

	return false
}
func IsUrl(str string) bool {
	rule := regexp.MustCompile(`(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`)
	return rule.MatchString(str)
}

// สำหรับเช็คว่าข้อความนั้นมีรูปแบบ example@mail.com
func IsEmail(str string) bool {
	rule := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return rule.MatchString(str)
}

// สำหรับเช็คว่าข้อความนั้นไม่มี white space
func NoSpace(str string) bool {
	rule := regexp.MustCompile(`\s+`)
	return rule.MatchString(str)
}

func IsCid(str string) bool {
	// ref https://mee-ci.blogspot.com/2021/03/blog-post.html
	// ! จะมีปัญหาเมื่อ ค่าในหลักที่ x ผิด "11x11111111"
	rulr := regexp.MustCompile(`\d{13}`)
	if rulr.MatchString(str) {
		cid11 := str[:12]
		sum := 0

		for i := 0; i < len(cid11); i++ {
			n, _ := strconv.Atoi(string(cid11[i]))
			sum = sum + ((13 - i) * n)
		}

		mod11 := sum % 11
		subtract11 := 11 - mod11
		firstDigit := strconv.Itoa(subtract11 % 10)
		number_card := cid11 + firstDigit
		return number_card == str
	} else {
		return false
	}
}

// สำหรับเช็คว่า String นั้นมี White Space หรือไม่หากมี Space จะ Return True
func IsSpace(str string) (result bool) {
	rule := regexp.MustCompile(`\s+`)
	return rule.MatchString(str)
}

// สำหรับเช็คว่า String นั้นเป็นตัวเลขทั้งหมดหรือไม่ Return True
func IsNumberOnly(str string) (result bool) {
	rule := regexp.MustCompile(`^\d+$`)
	return rule.MatchString(str)
}

// สำหรับเช็คว่า String นั้นเป็นรูปแบบ uuid หรือไม่
func IsUuid(str string) (result bool) {
	rule := regexp.MustCompile(`^[0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{12}$`)
	return rule.MatchString(str)
}

// สำหรับเช็คว่า String นั้นเป็นรูปแบบ Phone number หรือไม่ ex.0956658596
func IsPhoneNumber(str string) (result bool) {
	rule := regexp.MustCompile(`^[0-9]{10}$`)
	return rule.MatchString(str)
}

func IsStringOnly(str string) (result bool) {
	return true
}

func IsDateTimeFormat(str string) (result bool) {
	rule := regexp.MustCompile(`^\d{4}\-(0[1-9]|1[012])\-(0[1-9]|[12][0-9]|3[01]) ([01][0-9]|2[0-4]):([0-5][0-9]):([0-5][0-9])$`)
	return rule.MatchString(str)
}

func ExtractMailDescription(val string) string {
	val = strings.ReplaceAll(val, "\n", " ")
	if strings.Contains(val, `\n`) {
		val = strings.ReplaceAll(val, `\n`, " ")
	}

	// select first description mail
	val = regexp.MustCompile("From:|From :").ReplaceAllString(val, "\nFrom:")
	val = regexp.MustCompile(`From:.*\n`).FindString(val + "\n")

	for _, c := range []string{"To", "Cc", "Sent", "Subject", "ถึง", "สำเนา", "วันที่", "เรื่อง"} {
		val = regexp.MustCompile(fmt.Sprintf(" %s:| %s :", c, c)).ReplaceAllString(val, fmt.Sprintf("\n%s:", c))
	}
	return regexp.MustCompile(`From:.*\n|To:.*\n|Cc:.*\n|จาก:.*\n|ถึง:.*\n|สำเนา:.*\n`).ReplaceAllString(val, "")
}

// PrintJson ทำการแปลงข้อมูล (data) ให้เป็น JSON ที่จัดรูปแบบ (Pretty JSON) และพิมพ์ออกหน้าจอ
// หากเกิดข้อผิดพลาดระหว่างการแปลง จะพิมพ์ข้อความแสดงข้อผิดพลาดแทน
func PrintJson(data interface{}) {
	// ใช้ json.MarshalIndent เพื่อแปลง struct หรือ data เป็น JSON พร้อมจัดรูปแบบให้อ่านง่าย
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		// ถ้าแปลงไม่สำเร็จ พิมพ์ข้อความแสดงข้อผิดพลาด
		fmt.Println(err)
		return // จบการทำงานเพื่อไม่ให้ไปพิมพ์ค่า val ที่เป็นค่าเริ่มต้น
	}

	// แปลง JSON ที่ได้จาก []byte เป็น string แล้วพิมพ์ออกหน้าจอ
	fmt.Println(string(val))
}

// SprintJson ทำการแปลงข้อมูล (data) ให้เป็น JSON ที่จัดรูปแบบ (Pretty JSON)
// หากเกิดข้อผิดพลาดระหว่างการแปลง จะคืนค่าข้อความแสดงข้อผิดพลาดแทน
func SprintJson(data interface{}) string {
	// ใช้ json.MarshalIndent เพื่อแปลง struct หรือ data เป็น JSON พร้อมจัดรูปแบบให้อ่านง่าย
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		// ถ้าแปลงไม่สำเร็จ คืนค่าข้อความแสดงข้อผิดพลาด
		return fmt.Sprintf("Error: %v", err)
	}
	// แปลง JSON ที่ได้จาก []byte เป็น string และคืนค่า
	return string(val)
}

// LoggingFormat ใช้สำหรับจัดรูปแบบข้อความ Log ให้อยู่ในรูป JSON
// Input:
//   - userId: รหัสผู้ใช้งานที่ทำ action นั้น
//   - msg: ข้อความหรือรายละเอียดของ log message
//
// Output:
//   - string: ข้อความ JSON ที่จัดรูปแบบแล้ว เช่น {"user_id":"1234","message":"some action"}
func LoggingFormat(userId string, msg string) string {
	log := map[string]string{
		"user_id": userId,
		"message": msg,
	}
	bytes, err := json.Marshal(log)
	if err != nil {
		// fallback เผื่อ marshal มีปัญหา เช่น string มี character แปลก ๆ
		return fmt.Sprintf(`{"user_id":"%s","message":"%s"}`, userId, msg)
	}
	return string(bytes)
}
