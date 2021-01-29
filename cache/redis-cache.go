package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/nicobianchetti/Go-CleanArchitecture/model"
)

type redisCache struct {
	host    string
	db      int           // es un indice entre 0 y 15
	expires time.Duration //tiempo de expiracion
}

//NewRedisCache .
func NewRedisCache(host string, db int, exp time.Duration) PermisoCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisCache) Set(key string, value *model.Permiso) {

	client := cache.getClient()

	// ping, err := client.Ping(client.Context()).Result()
	// fmt.Println(ping, err)

	json, err := json.Marshal(value)

	if err != nil {
		fmt.Println(err)
	}

	err = client.Set(client.Context(), key, json, cache.expires*time.Second).Err()

	if err != nil {
		fmt.Println(err)
	}

}

func (cache *redisCache) Get(key string) *model.Permiso {
	// fmt.Println("\n key:", key)
	client := cache.getClient()

	val, err := client.Get(client.Context(), key).Result()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	permiso := model.Permiso{}
	err = json.Unmarshal([]byte(val), &permiso)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &permiso
}
