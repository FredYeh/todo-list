package redis

import (
	"context"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type redisConfig struct {
	Host     string      `json:"host"`
	Port     json.Number `json:"port"`
	Password string      `json:"password"`
	DB       int         `json:"db"`
}

type Storage struct {
	Client *redis.Client
}

func NewRedisStorage() *Storage {
	config := new(redisConfig)
	if bson, err := json.Marshal(viper.GetStringMap("database")); err != nil {
		log.Fatal(err)
	} else if err := json.Unmarshal(bson, &config); err != nil {
		log.Fatal(err)
	}
	return &Storage{Client: NewClient(config)}
}

func NewClient(config *redisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + config.Port.String(),
		Password: "",
		DB:       0,
	})

	return client
}

func (s *Storage) Create(t any) (string, error) {
	id := uuid.New().String()
	if err := s.Client.HSet(context.TODO(), id, t).Err(); err != nil {
		return "", err
	}
	return id, nil
}

func (s *Storage) Read() []map[string]string {
	res := make([]map[string]string, 0)
	if allkeys, err := s.Client.Keys(context.TODO(), "*").Result(); err != nil {
		log.Fatal(err)
	} else {
		for _, key := range allkeys {
			if val, err := s.Client.HGetAll(context.TODO(), key).Result(); err != nil {
				log.Fatal(err)
			} else {
				val["id"] = key
				res = append(res, val)
			}
		}
	}
	return res
}

func (s *Storage) Update(id string, t any) error {
	if _, err := s.Client.HGetAll(context.TODO(), id).Result(); err != nil {
		return err
	}
	err := s.Client.HSet(context.TODO(), id, t).Err()
	return err
}

func (s *Storage) Delete(id string) error {
	err := s.Client.Del(context.TODO(), id).Err()
	return err
}
