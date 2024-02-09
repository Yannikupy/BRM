package factory

import (
	"brm-core/internal/ports/rmq"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func ConnectToRabbitmq() ([]rmq.Shard, error) {
	coreRmqUrl := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		viper.GetString("rabbitmq-core.username"),
		os.Getenv("RABBITMQ_CORE_PASSWORD"),
		viper.GetString("rabbitmq-core.host"),
		viper.GetInt("rabbitmq-core.port"))

	amount := viper.GetInt("rabbitmq-core.rabbitmq-shards.amount")
	shards := make([]rmq.Shard, amount)
	for i := range shards {
		var err error
		shards[i], err = rmq.NewShard(
			coreRmqUrl,
			viper.GetString(fmt.Sprintf("rabbitmq-core.rabbitmq-shards.shard%02d.consumer", i+1)),
			coreRmqUrl,
			viper.GetString(fmt.Sprintf("rabbitmq-core.rabbitmq-shards.shard%02d.producer", i+1)),
		)
		if err != nil {
			return nil, err
		}
	}

	return shards, nil
}
