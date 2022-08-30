package postgres

import (
	"database/sql"
	"fmt"
)

type WBModel struct {
	DB *sql.DB
}

func (m WBModel) SelectData() (map[string][]byte, error) {
	MapDATA := make(map[string][]byte)
	rows, err := m.DB.Query("select * from items;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var data []byte
		err := rows.Scan(&id, &data)
		if err != nil {
			fmt.Println(err)
			continue
		}
		MapDATA[id] = data
	}

	return MapDATA, nil
}

func (m WBModel) InsertData(orderUid string, b []byte) error {
	_, err := m.DB.Exec("INSERT INTO wbweek (order_uid, data) VALUES ($1, $2);", orderUid, b)
	if err != nil {
		return err
	}
	return nil
}
