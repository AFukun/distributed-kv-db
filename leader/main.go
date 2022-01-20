package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/AFukun/distributed-kv-db/pb"
	"google.golang.org/grpc"
)

func main() {
	ports := []string{"9000", "9001", "9002"}
	for {
		reader := bufio.NewReader(os.Stdin)
		requestString, _ := reader.ReadString('\n')
		fields := strings.Fields(requestString)
		method := fields[0]
		key := fields[1]
		var value int
		if len(fields) > 2 {
			value, _ = strconv.Atoi(fields[2])
		}
		request := &pb.Request{Method: method, Key: key, Value: int64(value)}
		var responseStatus string
		var responseValue int

		for _, port := range ports {
			conn, err := grpc.Dial("127.0.0.1:"+port, grpc.WithInsecure())
			if err != nil {
				continue
			}
			client := pb.NewDatabaseServiceClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			response, err := client.Query(ctx, request)
			if err != nil {
				continue
			}
			if port == "9000" {
				responseStatus = response.GetStatus()
				responseValue = int(response.GetValue())
			}
		}
		if responseStatus == "VALUE" {
			if responseValue == 0 {
				fmt.Println("null")
			} else {
				fmt.Println(responseValue)
			}
		} else {
			fmt.Println(responseStatus)
		}
	}
}
