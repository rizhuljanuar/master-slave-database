# MariaDB Master-Slave Configuration

Proyek ini menunjukkan implementasi MariaDB Master-Slave replication dengan aplikasi Laravel yang terintegrasi untuk membaca dari slave dan menulis ke master.

## Struktur Proyek

```
mariadb-master-slave/
├── master-slave-databases/          # Konfigurasi database
│   ├── conf/                       # File konfigurasi MariaDB
│   │   ├── master.cnf             # Konfigurasi master database
│   │   ├── slave1.cnf             # Konfigurasi slave 1
│   │   ├── slave2.cnf             # Konfigurasi slave 2
│   │   └── slave3.cnf             # Konfigurasi slave 3
│   ├── data/                       # Direktori data database
│   └── docker-compose.yml          # Konfigurasi Docker untuk database
├── master-slave-apps/              # Konfigurasi aplikasi
│   └── master-slave-app-laravel/   # Aplikasi Laravel
│       ├── app/                    # Source code Laravel
│       ├── php/                    # Konfigurasi PHP
│       ├── nginx/                 # Konfigurasi Nginx
│       └── docker-compose.yml      # Konfigurasi Docker aplikasi
├── db.sql                         # Script SQL inisialisasi database
└── .gitignore                     # File ignore untuk Git
```

## Arsitektur Sistem

### Database Layer
- **MariaDB Master**: Server utama untuk operasi tulis
- **MariaDB Slaves (3 instances)**: Server backup untuk operasi baca
- **Network**: Docker network terpisah untuk komunikasi database

### Application Layer
- **Laravel Application**: Web application dengan read/write splitting
- **Nginx**: Web server untuk Laravel
- **PHP-FPM**: PHP processor untuk aplikasi Laravel

## Tahap Pembuatan Proyek

### 1. Tahap Persiapan Awal
1. **Clone Repository**: Menyiapkan struktur proyek awal
2. **Docker Installation**: Memastikan Docker dan Docker Compose terinstall
3. **MariaDB LTS**: Menggunakan MariaDB LTS versi stable

### 2. Konfigurasi Database Master-Slave

#### Konfigurasi Master (`master.cnf`)
```ini
[mariadb]
server-id=1
log-bin=mariadb-bin
binlog-format=ROW
expire_logs_days=7
max_binlog_size=100M
bind-address=0.0.0.0
innodb_buffer_pool_size=512M
default_storage_engine=InnoDB
innodb_autoinc_lock_mode=2
innodb_flush_log_at_trx_commit=1
log_error=/var/log/mysql/error.log
```

**Penjelasan:**
- `server-id=1`: ID unik untuk master database
- `log-bin=mariadb-bin`: Aktifkan binary logging
- `binlog-format=ROW`: Format replication row-based
- `expire_logs_days=7`: Hapus log setelah 7 hari
- `max_binlog_size=100M`: Maksimal ukuran binary log

#### Konfigurasi Slave (`slave*.cnf`)
```ini
[mariadb]
server-id=2
read-only=1
default_storage_engine=InnoDB
innodb_autoinc_lock_mode=2
innodb_flush_log_at_trx_commit=1
log_error=/var/log/mysql/error.log
```

**Penjelasan:**
- `server-id=2,3,4`: ID unik untuk setiap slave
- `read-only=1`: Hanya operasi bisa dilakukan
- Konfigurasi lain sama dengan master untuk konsistensi

### 3. Docker Compose Configuration

#### Database Services
```yaml
services:
  mariadb-master:
    image: mariadb:lts-noble
    container_name: mariadb-master
    environment:
      MARIADB_ROOT_PASSWORD: secret
    ports:
      - "3386:3306"
    volumes:
      - ./conf/master.cnf:/etc/mysql/conf.d/master.cnf
      - ./data/db-master-data:/var/lib/mysql

  mariadb-slave1:
    image: mariadb:lts-noble
    container_name: mariadb-slave1
    environment:
      MARIADB_ROOT_PASSWORD: secret
    ports:
      - "3387:3306"
    volumes:
      - ./conf/slave1.cnf:/etc/mysql/conf.d/slave.cnf
      - ./data/db-slave1-data:/var/lib/mysql
```

#### Application Services
```yaml
services:
  master-slave-app-laravel:
    build: .
    container_name: master-slave-app-laravel
    environment:
      DB_CONNECTION=mariadb
      DB_NAME=simple_app
      # Write connections
      DB_WRITE_HOST=mariadb-master
      DB_WRITE_PORT=3306
      DB_WRITE_USER=root
      DB_WRITE_PASSWORD=secret
      # Read connections
      DB_READ_HOST=mariadb-slave3
      DB_READ_PORT=3306
      DB_READ_USER=readonlyuser
      DB_READ_PASSWORD=readonlypass
```

### 4. Aplikasi Laravel Configuration

#### Database Setup
Aplikasi Laravel dikonfigurasi dengan read/write splitting:
- **Write Operations**: Ke master database (`mariadb-master`)
- **Read Operations**: Ke slave database (`mariadb-slave3`)

#### Environment Variables
```env
DB_CONNECTION=mariadb
DB_NAME=simple_app
DB_WRITE_HOST=mariadb-master
DB_WRITE_PORT=3306
DB_WRITE_USER=root
DB_WRITE_PASSWORD=secret
DB_READ_HOST=mariadb-slave3
DB_READ_PORT=3306
DB_READ_USER=readonlyuser
DB_READ_PASSWORD=readonlypass
```

### 5. Inisialisasi Database

#### SQL Schema
```sql
CREATE DATABASE IF NOT EXISTS simple_app;
USE simple_app;
CREATE TABLE IF NOT EXISTS items (
  id INT AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  qty INT NOT NULL,
  price DECIMAL(10,2) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

## Setup dan Deployment

### Langkah-langkah Setup

1. **Clone Repository**
```bash
git clone <repository-url>
cd mariadb-master-slave
```

2. **Build Docker Images**
```bash
# Build database services
cd master-slave-databases
docker-compose up -d

# Build application services
cd ../master-slave-apps/master-slave-app-laravel
docker-compose up -d --build
```

3. **Konfigurasi Replication**
```bash
# Login ke master
docker exec -it mariadb-master mysql -u root -p

# Buat user replication di master
CREATE USER 'replication_user'@'%' IDENTIFIED BY 'replication_password';
GRANT REPLICATION SLAVE ON *.* TO 'replication_user'@'%';

# Export binlog position
SHOW MASTER STATUS;
```

4. **Setup Slaves**
```bash
# Login ke masing-masing slave
docker exec -it mariadb-slave1 mysql -u root -p
docker exec -it mariadb-slave2 mysql -u root -p
docker exec -it mariadb-slave3 mysql -u root -p

# Configure replication untuk setiap slave
CHANGE REPLICATION SOURCE TO SOURCE_HOST='mariadb-master',
SOURCE_PORT=3306, SOURCE_USER='replication_user',
SOURCE_PASSWORD='replication_password',
SOURCE_LOG_FILE='mariadb-bin.000001',
SOURCE_LOG_POS=154;

START REPLICA;
```

### Verifikasi Setup

1. **Cek Status Master**
```bash
docker exec mariadb-master mysql -u root -p -e "SHOW MASTER STATUS;"
```

2. **Cek Status Slaves**
```bash
docker exec mariadb-slave1 mysql -u root -p -e "SHOW REPLICA STATUS\G"
docker exec mariadb-slave2 mysql -u root -p -e "SHOW REPLICA STATUS\G"
docker exec mariadb-slave3 mysql -u root -p -e "SHOW REPLICA STATUS\G"
```

3. **Test Replication**
```bash
# Insert data ke master
docker exec mariadb-master mysql -u root -p simple_app -e "INSERT INTO items (title, description, qty, price) VALUES ('Test Item', 'Description', 10, 99.99);"

# Cek data di slaves
docker exec mariadb-slave1 mysql -u root -p simple_app -e "SELECT * FROM items;"
docker exec mariadb-slave2 mysql -u root -p simple_app -e "SELECT * FROM items;"
docker exec mariadb-slave3 mysql -u root -p simple_app -e "SELECT * FROM items;"
```

## Port Configuration

### Database Ports
- **Master**: `3386` -> `3306`
- **Slave 1**: `3387` -> `3306`
- **Slave 2**: `3388` -> `3306`
- **Slave 3**: `3389` -> `3306`

### Application Ports
- **Laravel + Nginx**: `8003` -> `80`

## Security Considerations

1. **Password Management**: Gunakan environment variables untuk sensitive data
2. **User Permissions**: Batasi permissions untuk user replication
3. **Network Security**: Gunakan Docker network untuk isolasi
4. **Regular Backups**: Implementasikan backup strategy untuk data

## Monitoring dan Maintenance

### Log Monitoring
```bash
# Master logs
docker logs mariadb-master

# Slave logs
docker logs mariadb-slave1
docker logs mariadb-slave2
docker logs mariadb-slave3
```

### Performance Tuning
- Adjust `innodb_buffer_pool_size` berdasarkan memory available
- Monitor binary log size and retention
- Optimize query untuk read/write splitting

## Troubleshooting

### Common Issues

1. **Replication Lag**
   - Cek slave status dengan `SHOW REPLICA STATUS\G`
   - Verifikasi network connectivity antar containers

2. **Connection Issues**
   - Cek Docker network configuration
   - Verify port mappings

3. **Permission Errors**
   - Ensure replication user has proper permissions
   - Check firewall/security group settings

### Debug Commands
```bash
# Container status
docker-compose ps

# Container logs
docker-compose logs -f [service-name]

# Enter container
docker exec -it [container-name] bash
```

## Kesimpulan

Proyek ini menunjukkan implementasi MariaDB Master-Slave replication dengan:
- **High Availability**: Multiple slave instances
- **Load Balancing**: Read operations distributed ke slaves
- **Data Consistency**: Replication real-time dari master ke slaves
- **Scalability**: Easy scale dengan menambah slave instances

Setup ini ideal untuk aplikasi web yang membutuhkan performa tinggi dan data consistency yang baik.