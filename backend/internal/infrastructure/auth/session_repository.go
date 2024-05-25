package auth

import (
	"MelkOnline/internal/infrastructure"
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type SessionRepository struct {
	RedisConn *redis.Client
}

func NewSessionRepository() *SessionRepository {
	return &SessionRepository{RedisConn: infrastructure.GetRedis()}
}

func (sr *SessionRepository) ValidateSession(token string) (int, error) {
	uid, err := sr.RedisConn.Get(context.TODO(), token).Result()
	if err != nil {
		return 0, err
	}
	uidInt, _ := strconv.Atoi(uid)
	return uidInt, nil
}

func (sr *SessionRepository) StoreSession(token string, userID int) error {
	return sr.RedisConn.Set(context.TODO(), token, userID, time.Minute*time.Duration(15)).Err()
}
