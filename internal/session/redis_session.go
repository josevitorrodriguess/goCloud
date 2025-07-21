package session

import (
	"context"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

var ctx = context.Background()
var rdb *redis.Client

func InitRedisSession(addr, password string, db int) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}

func SetSession(w http.ResponseWriter, email string) (string, error) {
	sessionID := uuid.New().String()
	err := rdb.Set(ctx, sessionID, email, 24*time.Hour).Err()
	if err != nil {
		return "", err
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "mysession",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // true em produção
		SameSite: http.SameSiteLaxMode,
	})
	return sessionID, nil
}

func GetSession(r *http.Request) (string, bool) {
	c, err := r.Cookie("mysession")
	if err != nil {
		return "", false
	}
	email, err := rdb.Get(ctx, c.Value).Result()
	if err != nil {
		return "", false
	}
	return email, true
}

func DeleteSession(r *http.Request) {
	c, err := r.Cookie("mysession")
	if err == nil {
		rdb.Del(ctx, c.Value)
	}
}
