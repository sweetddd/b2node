package bitcoin

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func openDB(cfg StateConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Pass, cfg.DBName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}

type VerifiedBatch struct {
	BatchNum   int64
	TxHash     string
	Aggregator string
	StateRoot  string
	BlockNum   int64
	IsTrusted  bool
}

func GetStateRoot(cfg StateConfig, index int64) ([]*VerifiedBatch, error) {
	db, err := openDB(cfg)
	userSql := "select state_root, max(block_num) block_num from verified_batch where block_num > $1 and is_trusted=true  group by state_root order by block_num desc"
	rows, err := db.Query(userSql, index)
	if err != nil {
		return nil, fmt.Errorf("vin parse err:%w", err)
	}
	var batchs []*VerifiedBatch

	for rows.Next() {
		var stateRoot string
		var blockNum int64
		err = rows.Scan(&stateRoot, &blockNum)
		if err != nil {
			return nil, fmt.Errorf("vin parse err:%w", err)
		}
		batch := &VerifiedBatch{
			StateRoot: stateRoot,
			BlockNum:  blockNum,
		}
		batchs = append(batchs, batch)
	}
	return batchs, nil
}
