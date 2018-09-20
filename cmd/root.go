package cmd 

import (
	"context"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	logger *log.Logger
	dbPool *sql.DB

	handler http.Handler
	
)

var RootCmd = &cobra.Command{
	Use: "InfluxDB Api",
	Short: "Influx",
	Long: `nyobi InfluxDb with cobra and viper`,
	PreRun: func(cmd *cobra.Command, args []string) {
		// Pre-Run specific context timeout.
		//Note: Not sure if it's necessary to "config-ize" the timeout.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
	}

	fmt.Println(`
		Server InfluxDB
	`)

	fmt.Println("Current server config:")
	keys := viper.AllKeys()
	sort.Strings(keys)
	for _, key :=  range keys {
		fmt.Println(fmt.Spintf("%s: %+v", key, viper.Get(key)))
	}

	initDB()
}
