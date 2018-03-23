package main

import (
	"github.com/go-up/go-example/protocol-buffer-http/helloworld"

	"log"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
)

func main() {
	route := gin.Default()

	route.Handle("GET", "/get/binary", func(c *gin.Context) {
		test := &helloworld.Test{
			Label: proto.String("hello"),
			Type:  proto.Int32(17),
			Reps:  []int64{1, 2, 3},
			Optionalgroup: &helloworld.Test_OptionalGroup{
				RequiredField: proto.String("good bye"),
			},
		}

		data, err := proto.Marshal(test)
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}

		newTest := &helloworld.Test{}
		err = proto.Unmarshal(data, newTest)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}

		log.Println(newTest)

		c.Writer.Header().Set("Content-Type", "application/octet-stream")

		c.Writer.Write(data)

		c.Status(http.StatusOK)

	})

	route.Run(":19999")
}
