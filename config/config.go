package config

import (
	"log" // Thư viện để ghi log thông tin
	"os"  // Thư viện để thao tác với môi trường (environment variables)

	"context"                                   // Dùng để quản lý context trong các thao tác bất đồng bộ
	"github.com/joho/godotenv"                  // Thư viện dùng để tải và xử lý các biến môi trường từ file .env
	"go.mongodb.org/mongo-driver/mongo"         // Thư viện MongoDB driver dùng để kết nối với MongoDB
	"go.mongodb.org/mongo-driver/mongo/options" // Thư viện bổ trợ để cấu hình các tùy chọn kết nối MongoDB
	"gorm.io/driver/postgres"                   // Thư viện driver dùng để kết nối PostgreSQL với GORM
	"gorm.io/gorm"                              // Thư viện ORM (Object-Relational Mapping) GORM
)

var DB *gorm.DB           // Biến toàn cục để lưu kết nối đến cơ sở dữ liệu PostgreSQL
var MongoDB *mongo.Client // Biến toàn cục để lưu kết nối đến cơ sở dữ liệu MongoDB

// LoadEnv tải các biến môi trường từ file .env vào hệ thống
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file") // Ghi lỗi nếu không thể tải được file .env
	}
}

// ConnectPostgres thiết lập kết nối với cơ sở dữ liệu PostgreSQL
func ConnectPostgres() {
	var err error
	dsn := os.Getenv("POSTGRES_DSN")                        // Lấy chuỗi kết nối (DSN) từ biến môi trường
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) // Mở kết nối với PostgreSQL qua GORM
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err) // Ghi lỗi nếu không kết nối được
	}
	log.Println("Connected to PostgreSQL!") // Log thông báo kết nối thành công
}

// ConnectMongoDB thiết lập kết nối với cơ sở dữ liệu MongoDB
func ConnectMongoDB() {
	var err error
	MongoDB, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI"))) // Kết nối MongoDB với URI từ biến môi trường
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err) // Ghi lỗi nếu không kết nối được
	}
	log.Println("Connected to MongoDB!") // Log thông báo kết nối thành công
}
