package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"context"
	"google.golang.org/grpc"
	"log"
	"flag"
	pb "github.com/say-gpc/api"
)

func main(){
	backend := flag.String("b","localhost:9001","backend address")
	output := flag.String("o","output.wav","output the audio file")
	flag.Parse()
	if flag.NArg()<2{
		fmt.Printf("useage text to speech %s",os.Args[0])
		os.Exit(1)
	}

	conn,err:=grpc.Dial(*backend,grpc.WithInsecure())
	if err != nil{
		log.Fatal(err)
	}
	defer conn.Close()
	client :=pb.NewTextToSpeechClient(conn)
	text := &pb.Text{Text:flag.Arg(0)}
	res,err:=client.Say(context.Background(),text)
	if err != nil{
		log.Fatal(err)
	}

	if err:=ioutil.WriteFile(*output,res.Audio,0666);err != nil{
		log.Fatal(err)
	}

	

}