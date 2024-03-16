package auth

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type SessionRepository struct {
	RedisConn *redis.Client
}

func NewSessionRepository() *SessionRepository {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return &SessionRepository{RedisConn: client}
}

func (sr *SessionRepository) ValidateSession(token string) (bool, error) {
	_, err := sr.RedisConn.Get(context.TODO(), token).Result()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (sr *SessionRepository) StoreSession(token string, userID int) error {
	return sr.RedisConn.Set(context.TODO(), token, userID, time.Minute*time.Duration(15)).Err()
}
