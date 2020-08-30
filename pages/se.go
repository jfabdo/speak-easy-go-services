package pages

import (
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

func redisClient() redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDISADDRESS") + ":" + os.Getenv("REDISPORT"),
	})
}

func postify(re *http.Request) PostMessage {
	var PostMessage newPost
	reader := buffio.Reader()
	postBody := re.ReadRequest()
}

//HandleSe handles messages going to the root of speak-easy
func HandleSe(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	rc := redisClient()
	isBlocked, err := rc.Get(&r.RemoteAddr) //change placeholder string to actual struct
	if isBlocked == nil {
		return
	}
	post := postify(r)
	rc.Publish("ingress")
}
