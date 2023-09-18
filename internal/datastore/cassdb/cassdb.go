package cassdb

import (
	"strings"

	"github.com/go-logr/zerologr"
	"github.com/gocql/gocql"
	"github.com/kristoiv/gocqltable"
	"github.com/spf13/viper"
)

var (
	cassdbSession *gocql.Session
)

func SetUpSession(logger zerologr.Logger) {
	chosts := viper.GetString("cassandra.hosts")
	dbhosts := strings.Split(chosts, ",")

	cluster := gocql.NewCluster(dbhosts...)
	cluster.Keyspace = viper.GetString("cassandra.keyspace")
	s, err := cluster.CreateSession()

	cassdbSession = s

	if err != nil {
		logger.Error(err, "Error connecting to cassdb")
	}

	gocqltable.SetDefaultSession(cassdbSession)

}
