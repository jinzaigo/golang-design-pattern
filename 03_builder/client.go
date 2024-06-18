package builder

import (
	"fmt"
	"time"
)

func GetDBPool() {
	dbPool, err := Builder().DSN("localhost:3306").MaxOpenConn(50).MaxConnLifeTime(0 * time.Second).Build()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dbPool)
}
