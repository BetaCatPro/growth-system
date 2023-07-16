package service

import (
	"context"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"

	"growth/conf"
	"growth/dbhelper"
	"growth/initialize"
	"growth/models"
)

func initConf() {
	// default UTC time location
	time.Local = time.UTC
	// Load global config
	conf.LoadConfigs()
	// Initialize Logger
	initialize.InitLogger()
	// Initialize db
	dbhelper.InitDb()
}

func TestCoinTaskService_Save(t *testing.T) {
	initConf()
	s := NewCoinTaskService(context.Background())
	data := models.TbCoinTask{
		Id:    0,
		Task:  "post article",
		Coin:  10,
		Limit: 10,
	}
	if err := s.Save(&data); err != nil {
		t.Errorf("Save(%+v) error=%v", data, err)
	} else {
		zap.S().Infof("Save data=%+v\n", data)
	}
}

func TestCoinTaskService_Get(t *testing.T) {
	initConf()
	s := NewCoinTaskService(context.Background())
	if data, err := s.Get(1); err != nil {
		t.Errorf("Get(1) error=%v", err)
	} else {
		zap.S().Infof("Get(1) data=%+v\n", data)
	}
}

func TestCoinTaskService_GetByTask(t *testing.T) {
	initConf()
	s := NewCoinTaskService(context.Background())
	task := "post article"
	if data, err := s.GetByTask(task); err != nil {
		t.Errorf("GetByTask(%s) error=%v", task, err)
	} else {
		zap.S().Infof("GetByTask(%s) data=%+v\n", task, data)
	}
}

func TestCoinTaskService_FindAllPager(t *testing.T) {
	initConf()
	s := NewCoinTaskService(context.Background())
	if datalist, err := s.FindAll(); err != nil {
		t.Errorf("FindAll(1, 10) error=%v", err)
	} else {
		zap.S().Infof("FindAll(1, 10) datalist=%+v\n", datalist)
	}
}

func TestGradeInfoService_Save(t *testing.T) {
	initConf()
	s := NewGradeInfoService(context.Background())
	data := models.TbGradeInfo{
		Id:          0,
		Title:       "初级",
		Description: "初级用户",
		Score:       0,
		Expired:     0,
	}
	if err := s.Save(&data); err != nil {
		t.Errorf("Save(%+v) error=%v", data, err)
	} else {
		zap.S().Infof("Save data=%+v\n", data)
	}
}

func TestGradeInfoService_Get(t *testing.T) {
	initConf()
	s := NewGradeInfoService(context.Background())
	if data, err := s.Get(1); err != nil {
		t.Errorf("Get(1) error=%v", err)
	} else {
		zap.S().Infof("Get(1) data=%+v\n", data)
	}
}

func TestGradeInfoService_FindAll(t *testing.T) {
	initConf()
	s := NewGradeInfoService(context.Background())
	if datalist, err := s.FindAll(); err != nil {
		t.Errorf("FindAll error=%v", err)
	} else {
		zap.S().Infof("FindAll datalist=%+v\n", datalist)
	}
}

func TestGradeInfoService_NowGrade(t *testing.T) {
	initConf()
	s := NewGradeInfoService(context.Background())
	if data, err := s.NowGrade(1); err != nil {
		t.Errorf("NowGrade error=%v", err)
	} else {
		zap.S().Infof("NowGrade v=%+v\n", data)
	}
}
