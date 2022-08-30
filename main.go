package main

import (
	"1week/models"
	"1week/models/postgres"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
	"log"
	"os"
)

var (
	Sc stan.Conn
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	db       *postgres.WBModel
}

var cache = map[string][]byte{}

func (a *application) caching() {
	mapData, err := a.db.SelectData()
	if err != nil {
		a.errorLog.Println(err.Error())
	}
	for i := range mapData {
		cache[i] = mapData[i]
	}
}

func main() {
	dbFlag := flag.String("db", "user=postgres password=admin dbname=wbweek sslmode=disable", "Название POSTGRES источника данных")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dbFlag)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
		db:       &postgres.WBModel{DB: db},
	}

	app.caching()

	go app.startSubNats()

	app.startServer()

	//quit := make(chan struct{})
	//startSub("nats-example")

	//data := []byte("this is a test message")
	//PublishNats(data, "test")

	//PrintMessage("test", "test", "test-1")
	//<-quit
}

func (a *application) parseJsonFile(byteValue []byte) {
	var orders models.Orders
	err := json.Unmarshal(byteValue, &orders)
	if err != nil {
		switch e := err.(type) {
		case *json.UnmarshalTypeError:
			fmt.Printf("UnmarshalTypeError %s, %s\n", e.Type, e.Value)
		case *json.InvalidUnmarshalError:
			fmt.Printf("InvalidUnmarshalError %s\n", e.Type)
		default:
			fmt.Println(err.Error())
		}
	}
	a.saveData(orders.OrderUid, byteValue)
}

func (a *application) saveData(orderUid string, b []byte) {
	cache[orderUid] = b

	err := a.db.InsertData(orderUid, b)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}

func openDB(psql string) (*sql.DB, error) {
	db, err := sql.Open("postgres", psql)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func PrintMessage(subject, qgroup, durable string) {
	mcb := func(msg *stan.Msg) {
		if err := msg.Ack(); err != nil {
			log.Printf("failed to ACK msg:%v", err)
		}

		fmt.Println(string(msg.Data))
	}

	_, err := Sc.QueueSubscribe(subject,
		qgroup, mcb,
		stan.DeliverAllAvailable(),
		stan.SetManualAckMode(),
		stan.DurableName(durable))
	if err != nil {
		log.Println(err)
	}
}
