package rmq

type Shard struct {
	consumer
	producer
}

func NewShard(
	consumerAddr string,
	consumerQueueName string,
	producerAddr string,
	producerQueueName string,
) (Shard, error) {
	shard := Shard{}
	var err error
	shard.consumer, err = newConsumer(consumerAddr, consumerQueueName)
	if err != nil {
		return Shard{}, err
	}

	shard.producer, err = newProducer(producerAddr, producerQueueName)
	if err != nil {
		return Shard{}, err
	}

	return shard, nil
}

func (s *Shard) Close() {
	s.consumer.close()
	s.producer.close()
}
