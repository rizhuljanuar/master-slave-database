package config

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	MasterDB *sql.DB
	SlaveDB  *sql.DB
}

func NewConfig() (*Config, error) {
	// Konfigurasi koneksi database master
	masterDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_WRITE_USER"),
		os.Getenv("DB_WRITE_PASSWORD"),
		os.Getenv("DB_WRITE_HOST"),
		os.Getenv("DB_WRITE_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Konfigurasi koneksi database slave
	slaveDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_READ_USER"),
		os.Getenv("DB_READ_PASSWORD"),
		os.Getenv("DB_READ_HOST"),
		os.Getenv("DB_READ_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Inisialisasi koneksi database
	masterDB, err := sql.Open("mysql", masterDSN)
	if err != nil {
		return nil, fmt.Errorf("gagal koneksi ke database master: %v", err)
	}

	slaveDB, err := sql.Open("mysql", slaveDSN)
	if err != nil {
		return nil, fmt.Errorf("gagal koneksi ke database slave: %v", err)
	}

	// Uji koneksi
	if err := masterDB.Ping(); err != nil {
		return nil, fmt.Errorf("gagal ping database master: %v", err)
	}

	if err := slaveDB.Ping(); err != nil {
		return nil, fmt.Errorf("gagal ping database slave: %v", err)
	}


	// Ambil pengaturan connection pool dari .env
	masterMaxOpenConns, err := strconv.Atoi(os.Getenv("DB_WRITE_MAX_OPEN_CONNS"))
	if err != nil {
		masterMaxOpenConns = 25 // Default jika tidak ada di .env
	}
	masterMaxIdleConns, err := strconv.Atoi(os.Getenv("DB_WRITE_MAX_IDLE_CONNS"))
	if err != nil {
		masterMaxIdleConns = 25
	}
	masterConnMaxLifetime, err := time.ParseDuration(os.Getenv("DB_WRITE_CONN_MAX_LIFETIME"))
	if err != nil {
		masterConnMaxLifetime = 5 * time.Minute
	}
	masterConnMaxIdleTime, err := time.ParseDuration(os.Getenv("DB_WRITE_CONN_MAX_IDLE_TIME"))
	if err != nil {
		masterConnMaxIdleTime = 10 * time.Minute
	}

	slaveMaxOpenConns, err := strconv.Atoi(os.Getenv("DB_READ_MAX_OPEN_CONNS"))
	if err != nil {
		slaveMaxOpenConns = 50
	}
	slaveMaxIdleConns, err := strconv.Atoi(os.Getenv("DB_READ_MAX_IDLE_CONNS"))
	if err != nil {
		slaveMaxIdleConns = 50
	}
	slaveConnMaxLifetime, err := time.ParseDuration(os.Getenv("DB_READ_CONN_MAX_LIFETIME"))
	if err != nil {
		slaveConnMaxLifetime = 5 * time.Minute
	}
	slaveConnMaxIdleTime, err := time.ParseDuration(os.Getenv("DB_READ_CONN_MAX_IDLE_TIME"))
	if err != nil {
		slaveConnMaxIdleTime = 10 * time.Minute
	}

	// Konfigurasi connection pool untuk master
	masterDB.SetMaxOpenConns(masterMaxOpenConns)
	masterDB.SetMaxIdleConns(masterMaxIdleConns)
	masterDB.SetConnMaxLifetime(masterConnMaxLifetime)
	masterDB.SetConnMaxIdleTime(masterConnMaxIdleTime)

	// Konfigurasi connection pool untuk slave
	slaveDB.SetMaxOpenConns(slaveMaxOpenConns)
	slaveDB.SetMaxIdleConns(slaveMaxIdleConns)
	slaveDB.SetConnMaxLifetime(slaveConnMaxLifetime)
	slaveDB.SetConnMaxIdleTime(slaveConnMaxIdleTime)


	return &Config{
		MasterDB: masterDB,
		SlaveDB:  slaveDB,
	}, nil
}