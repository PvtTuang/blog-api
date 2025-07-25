# 📝 Blog API ด้วย Go

RESTful API สำหรับจัดการโพสต์บทความ เขียนด้วยภาษา **Go** โดยใช้ **Gorilla Mux**, **GORM** และ **PostgreSQL**

## 🚀 ความสามารถของ API

- 📄 เพิ่ม, อ่าน, แก้ไข, ลบ โพสต์ (CRUD)
- 🗃 จัดการข้อมูลด้วย PostgreSQL
- 🔄 ส่งและรับข้อมูลเป็น JSON
- 🛠 พัฒนาแบบโมดูล: `main.go`, `handler.go`, `database.go`, `models.go`

---

## 🧰 เทคโนโลยีที่ใช้

- **ภาษา**: Go
- **Router**: Gorilla Mux
- **ORM**: GORM
- **ฐานข้อมูล**: PostgreSQL
- **ENV Loader**: godotenv

---

## 📁 โครงสร้างโปรเจกต์

blog-api/
├── main.go # จุดเริ่มต้นของโปรแกรม
├── handler.go # จัดการ HTTP requests
├── database.go # เชื่อมต่อฐานข้อมูล
├── models.go # โครงสร้างข้อมูล Post
├── .env # ตัวแปรสภาพแวดล้อม
├── go.mod / go.sum # Go module dependencies


---

## ⚙️ วิธีเริ่มต้นใช้งาน

### 1. Clone และเข้าโปรเจกต์

```bash
git clone https://github.com/PvtTuang/blog-api.git
cd blog-api
