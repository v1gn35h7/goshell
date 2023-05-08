package respository

import (
	"github.com/kristoiv/gocqltable"
	"github.com/spf13/viper"
)

var (
	KeySpace gocqltable.Keyspace
)

// Now we're ready to create our first keyspace. We start by getting a keyspace object
//d := gocqltable.NewKeyspace("gocqltable_test")

func Base() gocqltable.Keyspace {
	return gocqltable.NewKeyspace(viper.GetString("cassandra.keyspace"))
}
