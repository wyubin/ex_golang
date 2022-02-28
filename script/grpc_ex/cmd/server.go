/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	pbCalc "example.com/grpc_ex/pkg/calculate"
	pbGopher "example.com/grpc_ex/pkg/gopher"
)

const (
	port = ":9000"
)

// 繼承pb 的 UnimplementedGopherServer
type GopherServer struct {
	pbGopher.UnimplementedGopherServer
}

type CalcServer struct {
	pbCalc.UnimplementedCalcServer
}

// implement GetGopher 應該有的功能
func (s *GopherServer) GetGopher(ctx context.Context, req *pbGopher.GopherRequest) (*pbGopher.GopherReply, error) {
	res := &pbGopher.GopherReply{} // 直接串 pb定義好的 struct

	// Check request
	if req == nil {
		fmt.Println("沒輸入")
		return res, fmt.Errorf("request must not be nil")
	}

	if req.Name == "" {
		fmt.Println("名字不能是空的")
		return res, fmt.Errorf("name must not be empty in the request")
	}

	log.Printf("Received: %v", req.GetName())

	//直接組出回應
	nameUser := req.GetName()
	if nameUser == "wyubin" {
		res.Message = fmt.Sprintf("Hello, my master - %s", nameUser)
	} else {
		res.Message = fmt.Sprintf("Hi, %s", nameUser)
	}
	return res, nil
}

// implement GetGopher 應該有的功能
func (s *CalcServer) GetCalc(ctx context.Context, req *pbCalc.CalcRequest) (*pbCalc.CalcReply, error) {
	res := &pbCalc.CalcReply{}
	if req.Input == 0 {
		fmt.Println("input不能是0")
		return res, fmt.Errorf("input can not be zero")
	}
	Input := req.GetInput()
	log.Printf("CalcServer Received: %v", Input)

	// Resp
	res.Sum = Input + 1
	return res, nil
}

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the Schema gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		// listen :9000
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()

		// Register services
		pbGopher.RegisterGopherServer(grpcServer, &GopherServer{})
		pbCalc.RegisterCalcServer(grpcServer, &CalcServer{})

		log.Printf("GRPC server listening on %v", lis.Addr())

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
