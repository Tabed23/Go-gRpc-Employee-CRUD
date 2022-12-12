package service

import (
	"context"
	"employee-grpc/db"
	"employee-grpc/pb"
	"employee-grpc/types"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	pb.UnimplementedEmployeeServiceServer
}

var (
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
)

func (s *Server) CreateEmployee(ctx context.Context, in *pb.Employee) (*pb.EmployeeResponse, error) {
	emp := types.Employee{
		ID:        primitive.NewObjectID(),
		FirstName: in.GetFirstname(),
		Lastname:  in.GetLastname(),
		Position:  in.GetPosition(),
		Salary:    in.GetSalary(),
	}
	_, err := db.DB.Database("Employees").Collection("employees").InsertOne(ctx, &emp)
	if err != nil {
		return nil, err
	}
	return &pb.EmployeeResponse{
		Status: "Employee have been Created",
		Created: &pb.Employee{
			Id:        emp.ID.String(),
			Firstname: emp.FirstName,
			Lastname:  emp.Lastname,
			Position:  emp.Position,
			Salary:    emp.Salary,
		},
	}, nil

}


func (*Server) GetEmploye(ctx context.Context, in *pb.EmployeeRequest) (*pb.Employee, error) {
	oid, err := primitive.ObjectIDFromHex(string(in.GetEmpid()))
	if err != nil {
		return nil, err
	}
	res := db.DB.Database("Employees").Collection("employees").FindOne(ctx, bson.M{"_id": oid})
	e := types.Employee{}
	if err := res.Decode(&e); err != nil {
		return nil, err
	}
	return &pb.Employee{
		Id:        e.ID.Hex(),
		Firstname: e.FirstName,
		Lastname:  e.Lastname,
		Salary:    e.Salary,
		Position:  e.Position,
	}, nil

}

func (*Server) GetEmployees(in *pb.GetAllEmployes, stream pb.EmployeeService_GetEmployeesServer) error {

	coursor, err := db.DB.Database("Employees").Collection("employees").Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	for coursor.Next(ctx) {
		var e types.Employee
		if err := coursor.Decode(&e); err != nil {
			return err
		}
		stream.Send(&pb.Employee{
			Id:        e.ID.Hex(),
			Firstname: e.FirstName,
			Lastname:  e.Lastname,
			Position:  e.Lastname,
			Salary:    e.Salary,
		})
	}
	return nil
}

func (s *Server) DeleteEmployee(ctx context.Context, in *pb.EmployeeRequest) (*pb.DeleteEmloyeReponse, error) {
	oid, err := primitive.ObjectIDFromHex(string(in.GetEmpid()))
	if err != nil {
		return nil, err
	}
	_, err = db.DB.Database("Employees").Collection("employees").DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, err
	}
	return &pb.DeleteEmloyeReponse{
		Status: "Employee has been deleted",
	}, nil

}

func (s *Server) UpdateEmployee(ctx context.Context, in *pb.UpdateEmployeeRequest) (*pb.Employee, error) {
	oid, err := primitive.ObjectIDFromHex(string(in.GetEmp().GetId()))
	if err != nil {
		return nil, err
	}
	update := bson.M{
		"first_name": in.Emp.GetFirstname(),
		"last_name":  in.Emp.GetLastname(),
		"position":   in.Emp.GetPosition(),
		"salary":     in.Emp.GetSalary(),
	}
	res := db.DB.Database("Employees").Collection("employees").FindOneAndUpdate(ctx, bson.M{"_id": oid}, bson.M{"$set": update}, options.FindOneAndUpdate().SetReturnDocument(1))
	if err != nil {
		return nil, err
	}
	resp := types.Employee{}
	err = res.Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &pb.Employee{
		Id:        resp.ID.Hex(),
		Firstname: resp.FirstName,
		Lastname:  resp.Lastname,
		Position:  resp.Position,
		Salary:    resp.Salary,
	}, nil

}
