#  Electromart – Fullstack E-commerce Platform  

For Backend documentation, see the README.md file in the code under:
/backend/README.md  

**IDATG2204 – Group 18**

**Group Members:**  

**Amir Ali Moaddeli**  (studentnummer: 485503)  
**Mateusz Tomasz Kaczmarek** (studentnummer: 588805)  
**Bao Mark Nguyen**  (studentnummer: 116191)  
**Simen Hauger Wilberg**  (studentnummer: 116218)  
**Emil Gullstrand**  (studentnummer: 116260)


A complete educational web application that simulates an online electronics store, built with:

-  **Frontend**: React + Vite + Tailwind CSS  
-  **Backend**: Go REST API server (with layered architecture)  
-  **Database**: MySQL 8+, initialized via SQL scripts and Go

---

##  Project Structure

```
IDATG2204-Group-18-main/
├── frontend/                  # React + Vite frontend
├── backend/                   # Go REST API backend
├── sql/                       # MySQL schema & seed logic
│   ├── ProjectTables.sql
│   └── sampledatabase/
│       ├── databasetables.sql
│       ├── sampleData.sql
│       ├── DATABASE_NAME.txt
│       └── populatedb.go
```

---

##  Requirements (All Platforms)

| Tool            | Used For               |
| --------------- | ---------------------- |
| Node.js (18+)   | Frontend (Vite)        |
| MySQL 8+        | Data persistence       |
| Go 1.19+        | Backend server         |
| XAMPP (Windows) | MySQL GUI              |
| Git / VS Code   | Editing and versioning |

---

##  MySQL Setup

###  Windows (XAMPP)

1. Launch **XAMPP**, start **MySQL**.
2. Open [http://localhost/phpmyadmin](http://localhost/phpmyadmin)
3. Create a database named `idatg2204`.

###  macOS /  Linux

```bash
# macOS
brew install mysql
brew services start mysql

# Linux
sudo apt install mysql-server
sudo service mysql start

mysql -u root -p
CREATE DATABASE idatg2204;
```

---

##  Database Initialization

###  Option 1: Manual SQL Import

Use CLI or phpMyAdmin to run:

```sql
SOURCE sql/ProjectTables.sql;
SOURCE sql/sampledatabase/databasetables.sql;
SOURCE sql/sampledatabase/sampleData.sql;
```

###  Option 2: Go Seeder (REQUIRED)

####  Configuration
- Set database name in `sql/sampledatabase/DATABASE_NAME.txt`
- Default DB: `idatg2204`

####  Run Seeder

```bash
cd sql/sampledatabase
go run populatedb.go
```

If configured properly, it will:
- Read the DB name
- Connect to MySQL
- Create & populate all tables

---

##  Backend Setup (Go API Server)

### Installation
#### Requirements:
- A MySQL database solution running on `root:@tcp(127.0.0.1:3306)/idatg2204` with the correct SQL tables and a user called root without credentials
- Or set the environmental variable `DSN` to the desired path with credentials

####  Run the Server
To start the backend, run the executable for your appropriate OS and system architecture (most commonly amd64) located in the `/backend/` directory

### Alternativly, to run the code without the executable:
#### Requirements
- GO version `1.24.1` or higher

####  Run the Server
Open a terminal window, and navigate to the `/backend/` directory. Then run this command: `go run main.go`

To stop the program, close the terminal window or press `ctr + c`

>  API server runs on: `http://localhost:8080`

---

##   Frontend Setup (React + Vite)

###  Location: `frontend/`

####  Steps

```bash
cd frontend
npm install
npm run dev
```

>  Frontend runs at: `http://localhost:5173`

####  API Integration
Update fetch/axios URLs to call:
```
http://localhost:8080/api/...
```

---

##  Folder Highlights

```
backend/
├── main.go
├── handlers/              # Routes and controllers
├── constants/             # SQL queries, DB config
├── structs/               # Data types
├── utility/               # Helper logic

frontend/
├── src/components/        # UI components
├── src/pages/             # Page views
├── src/context/           # Global state providers
├── index.html
├── vite.config.js

sql/
├── ProjectTables.sql
└── sampledatabase/
    ├── databasetables.sql
    ├── sampleData.sql
    ├── populatedb.go
    ├── DATABASE_NAME.txt
```

---

##  Verify Everything Works

| Component       | Command                | Expected Result                |
| --------------- | ---------------------- | ------------------------------ |
| Frontend        | `npm run dev`          | Runs on `localhost:5173`       |
| Backend         | `go run main.go`       | Runs on `localhost:8080`       |
| Database Seeder | `go run populatedb.go` | Tables + data created in MySQL |

---

##  User for testing
admin bruker:  
Username: helloWorld  
email: john_doe@gmail.com  
password:  helloWorld  

---

##  Auth & Security

- JWT is used for authentication (see `usersHandler`).

---

##  Deployment Notes

- Production builds (frontend):
  ```bash
  npm run build
  ```
- Backend can be compiled with:
  ```bash
  go build -o electromart-server main.go
  ```

