package main

import "time"
import "strconv"
import "fmt"
import "github.com/go-redis/redis"

type BadState struct {
	ServerName string
	Info       string
}

func GetBadStates(r *redis.Client, server_staleness_duration int64) map[string]BadState {
	bad_states := map[string]BadState{}
	good_time := time.Now().Unix() - server_staleness_duration
	bad_servers, err := r.ZRevRangeByScoreWithScores(KeyMap["server_last_posts"], redis.ZRangeBy{Max: "0", Min: strconv.FormatInt(good_time, 10)}).Result()
	if err == nil {
		return bad_states
	}
	for _, item := range bad_servers {
		server := item.Member.(string)
		key := "stale_" + server
		info := fmt.Sprintf("Server %s last reported on %v", server, item.Score)
		bad_states[key] = BadState{ServerName: server, Info: info}
	}
	return bad_states
}
