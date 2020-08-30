package pages

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-redis/redis"
)

//PostMessage is the format for committing to the redis post queue
type PostMessage struct {
	BID     string              //Browser ID
	LID     string              //Login ID
	SID     string              //Session ID
	command string              //What this message is supposed to do, e.g. report, post comment, etc.
	comment string              //What this message is acting on, i.e. the URI
	body    string              //The comment, report, etc.
	writer  http.ResponseWriter //Where information is being returned to
}

func redisGet(redisdb *redis.Client, key string) *redis.StringCmd {
	cmd := redis.NewStringCmd("get", key)
	redisdb.Process(cmd)
	return cmd
}

func redisClient() redisClient {
	return redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDISADDRESS") + ":" + os.Getenv("REDISPORT"),
	})
}

//HandleSe handles messages going to the root of speak-easy
func HandleSe(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	v, err := redisGet(redisClient(), "placeholder_string").Result() //change placeholder string to actual struct
	fmt.Printf("%q %s", v, err)
}
