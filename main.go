package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Fatal .env Load!...")
	}

	dbUrl := os.Getenv("DB_URL")
	connect, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		log.Fatalf("Gagal terhubung ke postgre : %v \n", err)
	}
	defer connect.Close(context.Background())

	err = connect.Ping(context.Background())
	if err != nil {
		log.Fatalf("Postgres mati/tidak bisa terhubung: %v \n", err)
	}
	fmt.Println("SUKSES TERHUBUNG KE DATABASE")

	redisAddr := os.Getenv("REDIS_ADDR")
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	ping, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Gagal terhubung ke redis: %v \n", err)
	}

	fmt.Printf("Berhasil terhubung ke redis: %v \n", ping)
}
