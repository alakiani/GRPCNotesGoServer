package main

import (
	"context"
	"log"
	"net"

	pb "github.com/alakiani/GRPCNotesGoServer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

func (s *server) List(ctx *pb.Empty, in pb.NoteService_ListServer) error {
	note1 := pb.Note{Id: "2", Title: "E", Content: "DDD"}
	note2 := pb.Note{Id: "1", Title: "E", Content: "DDD"}
	notes := []*pb.Note{&note1, &note2}
	noteList := pb.NoteList{Notes: notes}
	in.Send(&noteList)
	log.Printf("NoteId")
	return nil
}

func (s *server) Get(ctx context.Context, in *pb.NoteRequestId) (*pb.Note, error) {
	log.Printf("NoteId: %v", in.Id)
	return &pb.Note{}, nil
}

func (s *server) Insert(ctx context.Context, in *pb.Note) (*pb.Note, error) {
	log.Printf("NoteId:  %v", in.Id)
	return &pb.Note{}, nil
}

func (s *server) Update(ctx context.Context, in *pb.Note) (*pb.Note, error) {
	log.Printf("NoteId:  %v", in.Id)
	return &pb.Note{}, nil
}

func (s *server) Delete(ctx context.Context, in *pb.NoteRequestId) (*pb.Empty, error) {
	log.Printf("NoteId:  %v", in.Id)
	return &pb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNoteServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
