package main

import (
	"context"
	gateway "gateway/kitex_gen/gateway"
	"gateway/module"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"strings"
)

type BizServiceImpl struct{ db *gorm.DB }

func (s *BizServiceImpl) Register(ctx context.Context, req *gateway.BizRequest) (*gateway.BizResponse, error) {
	log.Println(req)
	if err := s.db.Table("students").Create(student2Model(req.Student)).Error; err != nil {
		return nil, err
	}
	return createSuccessResponse("register success"), nil
}

func (s *BizServiceImpl) Query(ctx context.Context, req *gateway.BizRequest) (*gateway.BizResponse, error) {
	log.Println(req.ItemId)
	var stuRes module.Student
	if err := s.db.Table("students").First(&stuRes, req.ItemId).Error; err != nil {
		return nil, err
	}
	return createSuccessResponseWithStudent("query success", model2Student(&stuRes)), nil
}

func (s *BizServiceImpl) InitDB() {
	db, err := gorm.Open(sqlite.Open("students.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&module.Student{}); err != nil {
		panic(err)
	}
	s.db = db
}

func student2Model(student *gateway.Student) *module.Student {
	return &module.Student{
		Id:             student.Id,
		Name:           student.Name,
		Email:          strings.Join(student.Email, ","),
		CollegeName:    student.College.Name,
		CollegeAddress: student.College.Address,
	}
}

func model2Student(student *module.Student) *gateway.Student {
	return &gateway.Student{
		Id:      student.Id,
		Name:    student.Name,
		Email:   strings.Split(student.Email, ","),
		College: &gateway.College{Name: student.CollegeName, Address: student.CollegeAddress},
	}
}

func createSuccessResponse(message string) *gateway.BizResponse {
	log.Println(message)
	return &gateway.BizResponse{
		Success: ptrBool(true),
		Message: ptrString(message),
	}
}

func createSuccessResponseWithStudent(message string, student *gateway.Student) *gateway.BizResponse {
	resp := createSuccessResponse(message)
	resp.Student = student
	return resp
}

func ptrBool(b bool) *bool       { return &b }
func ptrString(s string) *string { return &s }
