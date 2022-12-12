**GRPC-CRUD**
---

## CRUD Opertions 

You’ll start by editing this README file to learn how to edit a file in Bitbucket.

1. **GET** Employee
2. **GET** Employees
3. **POST** Employee
4. **DELETE** Employee
5. **UPDATE** Employee

---

## Employee Proto File
```
syntax = "proto3";

package EmployeeRpc;


option go_package = "/pb";

message Employee {  
    string id = 1;
    string firstname = 2;
    string lastname = 3;
    string position = 4;
    float salary =  5;
}

message AllEmployee {
    repeated Employee emp = 1;
}
message EmployeeResponse {
    string status = 1;
    Employee created = 2;

}

message EmployeeRequest {
    string empid = 1;
}
message GetAllEmployes {
}

message  DeleteEmloyeReponse {
    string status = 1;
}
message UpdateEmployeeRequest {
    Employee  emp = 1;
} 

service EmployeeService {SS
    rpc CreateEmployee(Employee)returns (EmployeeResponse){}
    rpc GetEmploye(EmployeeRequest)returns (Employee){}
    rpc GetEmployees(GetAllEmployes)returns (AllEmployee){}
    rpc DeleteEmployee (EmployeeRequest) returns (DeleteEmloyeReponse){}
    rpc UpdateEmployee(UpdateEmployeeRequest)returns (Employee){}
}
```

---
