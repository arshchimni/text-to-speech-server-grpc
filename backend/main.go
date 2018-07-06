package main

import (
	"io/ioutil"
	"os/exec"
	"google.golang.org/grpc"
	"fmt"
	"net"
	"flag"
	"log"
	"golang.org/x/net/context"
	pb "github.com/say-gpc/api"
)

func main(){
	port := flag.Int("p",8080,"port to listen to")
	flag.Parse()
	log.Printf("listening to port %d",*port)
	listner,err:=net.Listen("tcp",fmt.Sprintf(":%d",*port))
	if err !=nil{
		log.Fatal(err)
	}
	srv:=grpc.NewServer()
	pb.RegisterTextToSpeechServer(srv,server{})
	err = srv.Serve(listner)
	if err != nil{
		log.Fatal(err)
	}

}

type server struct{}

func (server) Say(ctx context.Context, text *pb.Text) (*pb.Speech, error){
	
	f,err:=ioutil.TempFile("","")
	if err !=nil{
		return nil,fmt.Errorf("could not create file %v",err)
	}
		if err :=f.Close();err!=nil{
			return nil,fmt.Errorf("could not close file %s : %v",f.Name(),err)
		}

	

	cmd := exec.Command("flite","-t",text.Text,"-o",f.Name())
	
	if data,err := cmd.CombinedOutput(); err != nil{
		return nil,fmt.Errorf("command execution err %s",data)
	}
	data,err:=ioutil.ReadFile(f.Name())
	if err !=nil{
		return nil,fmt.Errorf("file read error %v",err)
	}
	return &pb.Speech{Audio: data},nil
}

