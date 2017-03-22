package resultio

import (
	"log"
	"testing"
	"time"

	"github.com/teeratpitakrat/hora/adm"
	"github.com/teeratpitakrat/hora/cfp"
	"github.com/teeratpitakrat/hora/fpm"

	"github.com/spf13/viper"
)

func TestWriteCfpResult(t *testing.T) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("toml")
	viper.AddConfigPath("../.")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Print("Fatal error config file: %s \n", err)
	}

	viper.SetDefault("influxdb.addr", "http://localhost:8086")
	viper.SetDefault("influxdb.username", "root")
	viper.SetDefault("influxdb.password", "root")

	writer, err := New(
		viper.GetString("influxdb.addr"),
		viper.GetString("influxdb.username"),
		viper.GetString("influxdb.password"),
	)
	if err != nil {
		t.Error(err)
	}
	a := adm.Component{"A", "host1", "responsetime", 0}
	b := adm.Component{"B", "host2", "responsetime", 0}
	resulta := cfp.Result{
		a,
		time.Now(),
		time.Now().Add(10 * time.Minute),
		0.555,
	}
	resultb := cfp.Result{
		b,
		time.Now(),
		time.Now().Add(10 * time.Minute),
		0.666,
	}
	writer.WriteCfpResult(resulta)
	writer.WriteCfpResult(resultb)
}

func TestWriteFpmResult(t *testing.T) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("toml")
	viper.AddConfigPath("../.")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Print("Fatal error config file: %s \n", err)
	}

	viper.SetDefault("influxdb.addr", "http://localhost:8086")
	viper.SetDefault("influxdb.username", "root")
	viper.SetDefault("influxdb.password", "root")

	writer, err := New(
		viper.GetString("influxdb.addr"),
		viper.GetString("influxdb.username"),
		viper.GetString("influxdb.password"),
	)
	if err != nil {
		t.Error(err)
	}

	a := adm.Component{"A", "host1", "responsetime", 0}
	b := adm.Component{"B", "host2", "responsetime", 0}
	failProbs := make(map[adm.Component]float64)
	failProbs[a] = 0.2
	failProbs[b] = 0.3

	result := fpm.Result{
		failProbs,
		time.Now(),
		time.Now().Add(10 * time.Minute),
	}
	writer.WriteFpmResult(result)
}
