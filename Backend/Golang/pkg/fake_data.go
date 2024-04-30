package pkg

import (
	"math/rand"

	"time"

	"gorm.io/gorm"
)

var FakeData = func(db *gorm.DB) {
	ticker := time.NewTicker(10 * time.Second)
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)
	for range ticker.C {

		//random memory
		memory := 200 + rnd.Intn(4025-200)
		//random cpu
		cpu := 20 + rnd.Intn(100-20)
		timestamp := time.Now().Format(time.RFC3339)

		sql := "INSERT INTO measurements (metric_id, value, timestamp) VALUES (?, ?, ?), (?, ?, ?)"
		db.Exec(sql, 1, memory, timestamp, 2, cpu, timestamp)

		//log.Printf("FAKE DATA Inserted new metric memory: %v , cpu: %v at %v", memory, cpu, timestamp)
	}
}
