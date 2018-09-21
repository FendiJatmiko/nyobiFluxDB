package model

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/influxdata/influxdb/client"
	"src/github.com/google/uuid"
)

type (
	InfluxModel struct {
		ID uuid.UUID `json:"id"`
	}
)

func GetMeanCPU(c client.Client, cluster string) float64 {
	q := client.Query{
		Command:  fmt.Sprintf("select mean(cpu_usage) from node_status where cluster = '%s'", cluster),
		Database: database,
	}

	resp, err := c.Query(q)
	if err != nil {
		log.Fatalln("error: ", err)
	}

	if resp.Error() != nil {
		log.Fatalln("Error:", resp.Error())
	}

	res, err := resp.Result[0].Series[0].Values[0][1].(json.Number).Float64()
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	return res
}
