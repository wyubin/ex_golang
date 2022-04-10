/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	pbCalc "ex_golang/script/grpc_ex/pkg/calculate"
	pbGopher "ex_golang/script/grpc_ex/pkg/gopher"
)

const (
	address = "localhost:9000"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Query the gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("連線錯誤: %s", err)
		}
		defer conn.Close()
		// ctx
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		clientGopher := pbGopher.NewGopherClient(conn)
		clientCalc := pbCalc.NewCalcClient(conn)

		if len(os.Args) > 3 {
			switch os.Args[2] {
			case "gopher":
				r, _ := clientGopher.GetGopher(ctx, &pbGopher.GopherRequest{Name: os.Args[3]})
				log.Printf("Response: %s", r.GetMessage())
			case "calc":
				num, _ := strconv.Atoi(os.Args[3])
				r, _ := clientCalc.GetCalc(ctx, &pbCalc.CalcRequest{Input: int32(num)})
				log.Printf("Response: %v", r.GetSum())
			}
		} else {
			log.Fatal("參數不足")
		}
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
