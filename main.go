/**
 * @apiDefine Ping Ping
 *
 * Ping is a simple endpoint to check if the server is up and running.
 */
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/**
 * @api {get} /ping 只是一个 Ping ，用来测试是否连通
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "message": "Pong!"
 *     }
 */

func main() {
	loadConfig("config.json")

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		var req map[string]string = map[string]string{"message": "Pong!"}
		// convert req to json string
		res, err := json.Marshal(req)
		if err != nil {
			panic(err)
		}
		// write response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", res)
	})

	http.ListenAndServe(":8080", nil)
}
