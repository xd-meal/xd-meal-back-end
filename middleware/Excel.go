package middleware

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"github.com/wxnacy/wgo/arrays"
	"github.com/xd-meal-back-end/middleware/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"strconv"
	"strings"
	"time"
)

func ExportTmp(arr [][]string, c *gin.Context) {
	axis := [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	f := excelize.NewFile()
	//index := f.NewSheet(sheet)
	for k, v := range arr {
		for k2, v2 := range v {
			categoryCell := []int{0, 1, 8, 9}
			if arrays.Contains(categoryCell, k) != -1 {
				style, err := f.NewStyle(`{"border":[{"type":"bottom","style":1,"color":"#000000"},{"type":"top","style":1,"color":"#000000"},{"type":"left","style":1,"color":"#000000"},{"type":"right","style":1,"color":"#000000"}],
"alignment":{"horizontal":"center","Vertical":"center"},"font":{"bold":true},"fill":{"type":"pattern","color":["#FF9900"],"pattern":1}}`)
				if err != nil {
					fmt.Println(err)
				}
				err = f.SetCellStyle("Sheet1", axis[k2]+strconv.Itoa(k+1), axis[k2]+strconv.Itoa(k+1), style)
			}

			//合并单元格
			if arrays.Contains([]int{0, 8}, k) != -1 {
				_ = f.MergeCell("Sheet1", "A"+strconv.Itoa(k+1), "D"+strconv.Itoa(k+1))
				_ = f.SetCellValue("Sheet1", "A"+strconv.Itoa(k+1), v2)
			} else {
				_ = f.SetCellValue("Sheet1", axis[k2]+strconv.Itoa(k+1), v2)
			}
		}
	}

	f.SetSheetName("Sheet1", time.Now().Format("2006-01-02"))
	// Save xlsx file by the given path.
	date := []string{time.Now().Add(24 * time.Hour).Format("2006-01-02"), time.Now().Add(2 * 24 * time.Hour).Format("2006-01-02"),
		time.Now().Add(3 * 24 * time.Hour).Format("2006-01-02"), time.Now().Add(4 * 24 * time.Hour).Format("2006-01-02"),
		time.Now().Add(5 * 24 * time.Hour).Format("2006-01-02"), time.Now().Add(6 * 24 * time.Hour).Format("2006-01-02")}
	for _, v3 := range date {
		index := f.NewSheet(v3)
		_ = f.CopySheet(1, index)
		f.SetActiveSheet(index)
	}
	f.SetActiveSheet(1)
	c.Header("Content-Type", "application/vnd.ms-excel")
	c.Header("Content-Disposition", "attachment; filename="+"菜单模板.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "cache, must-revalidate") // HTTP/1.1
	c.Header("Cache-Control", "max-age=1")
	_ = f.Write(c.Writer)
}

func ReadMenuExcel(r io.Reader) ([]mongo.DishesMongo, error) {
	f, err := excelize.OpenReader(r)
	if err != nil {
		return nil, err
	}
	sheetMap := f.GetSheetMap()
	sheetArr := make([]string, len(sheetMap))
	for index, name := range sheetMap {
		sheetArr[index-1] = name
	}

	data := make(map[string][][]string, len(sheetMap))
	for _, name := range sheetArr {
		data[name], err = f.GetRows(name)
	}

	res := make([]mongo.DishesMongo, len(sheetMap)*12)
	key := 0
	var TypeA, typeB int32
	currentTime := time.Now()
	for _, date := range sheetArr {
		for i, v := range data[date] {
			if arrays.Contains([]int{2, 3, 4, 5, 6, 7}, i) != -1 { //午餐
				TypeA = 1
			} else if arrays.Contains([]int{10, 11, 12, 13, 14, 15}, i) != -1 { //晚餐
				TypeA = 2
			}
			if arrays.Contains([]int{2, 3, 4, 5, 6, 7, 10, 11, 12, 13, 14, 15}, i) != -1 {
				if len(v) > 0 {
					if v[2] == "园沁餐厅" {
						typeB = 1
					} else {
						typeB = 2
					}
					mealNum, _ := strconv.Atoi(v[3])
					res[key] = mongo.DishesMongo{
						ID: primitive.ObjectID{}, Name: v[0], Supplier: v[2], TypeA: TypeA, TypeB: typeB, MealDay: date, MealNum: mealNum, CreateTime: currentTime, UpdateTime: currentTime, Status: 0,
					}
					key++
				}
			}
		}
	}
	res1 := res[0:key]
	return res1, err
}

func ImportUser(r io.Reader) ([]mongo.UserMongo, error) {
	f, err := excelize.OpenReader(r)
	if err != nil {
		return nil, err
	}
	key := 0
	res := make([]mongo.UserMongo, 65)
	currentTime := time.Now()
	data, err := f.GetRows("Sheet1")
	for i, v := range data {
		if i > 0 {
			res[key] = mongo.UserMongo{
				Name: strings.TrimSpace(v[1]), Email: strings.TrimSpace(v[2]), PassWord: "xd123456", Type: 2, Depart: strings.TrimSpace(v[0]), CreateTime: currentTime,
			}
			key++
		}
	}
	return res, err
}
