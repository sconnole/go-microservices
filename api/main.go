package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type Job struct {
	ID        int       `json:"id"`
	Data      string    `json:"data"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	db := initDB()
	defer db.Close()

	router := setupRouter(db)
	router.Run()
}

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./jobs.db")
	if err != nil {
		log.Fatalf("failed to connect to SQLite: %v", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS jobs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		data TEXT NOT NULL,
		status TEXT NOT NULL,
		created_at DATETIME NOT NULL
	)`)
	if err != nil {
		log.Fatalf("Failed to create jobs table: %v", err)
	}

	return db
}

func setupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	router.GET("/random-job", func(c *gin.Context) {
		getRandomJobHandler(c, db)
	})

	router.POST("/jobs", func(c *gin.Context) {
		createJobHandler(c, db)
	})

	return router
}

func getRandomJobHandler(c *gin.Context, db *sql.DB) {
	var job Job
	var createdAt string

	err := db.QueryRow(`
		SELECT id, data, status, created_at
		FROM jobs
		ORDER BY RANDOM() ASC
		LIMIT 1
	`).Scan(&job.ID, &job.Data, &job.Status, &createdAt)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusOK, gin.H{"message": "No pending Jobs"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch job"})
		return
	}

	job.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
	c.JSON(http.StatusOK, job)
}

func createJobHandler(c *gin.Context, db *sql.DB) {
	var body struct {
		Data string `json:"data"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	createdAt := time.Now().UTC().Format(time.RFC3339)
	_, err := db.Exec("INSERT INTO jobs (data, status, created_at) VALUES (?, ?, ?)", body.Data, "PENDING", createdAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert the record"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Job created"})

}
