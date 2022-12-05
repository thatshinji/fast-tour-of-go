package main

import (
	trippb "fast-tour-of-go/src/proto_buf/gen/go"
	"fmt"
	"google.golang.org/protobuf/proto"
)

func main() {
	trip := trippb.Trip{
		Start:       "shanghai",
		End:         "beijing",
		FeeCent:     100,
		DurationSec: 3600,
	}
	marshaled, err := proto.Marshal(&trip)
	if err != nil {
		panic(err)
	}
	fmt.Println(marshaled)
}
