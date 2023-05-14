package cassdb

import (
	"strings"

	"github.com/gocql/gocql"
	"github.com/kristoiv/gocqltable"
	"github.com/spf13/viper"
)

var (
	cassdbSession *gocql.Session
)

func SetUpSession() {
	chosts := viper.GetString("cassandra.hosts")
	dbhosts := strings.Split(chosts, ",")

	cluster := gocql.NewCluster(dbhosts...)
	cluster.Keyspace = viper.GetString("cassandra.keyspace")
	s, err := cluster.CreateSession()

	cassdbSession = s

	if err != nil {
		panic("Error connecting to cassdb")
	}

	gocqltable.SetDefaultSession(cassdbSession)

}
