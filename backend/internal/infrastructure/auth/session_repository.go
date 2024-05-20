package auth

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type SessionRepository struct {
	RedisConn *redis.Client
}

func NewSessionRepository() *SessionRepository {
	err := godotenv.Load("/home/ssaeidifarzad/ssfdata/ssaeidifarzad/Classes/S8/SE/Project/SE_project/backend/.env")
	if err != nil {
		panic(err)
	}
	redisPass := os.Getenv("REDIS_PASS")
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: redisPass,
	})
	return &SessionRepository{RedisConn: client}
}

func (sr *SessionRepository) ValidateSession(token string) (int, error) {
	uid, err := sr.RedisConn.Get(context.TODO(), token).Result()
	if err != nil {
		return 0, err
	}
	uidInt, err := strconv.Atoi(uid)
	return uidInt, nil
}

func (sr *SessionRepository) StoreSession(token string, userID int) error {
	return sr.RedisConn.Set(context.TODO(), token, userID, time.Minute*time.Duration(15)).Err()
}
