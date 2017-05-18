package operfile

import (
	"fmt"
	"os"
	"strconv"
)

type DBSet struct {
	filename string
	offstart int64
	offset int64
}

type ElemData struct {
	Data []byte
	Loc DBSet
}

type DB struct {
	Filename [100]string
	DbFile  [100]*os.File
	Filelen [100]int64
}

const (
	DB_FILE_SIZE = 128 * 1024 * 10
)

func (db *DB) CreateDbFile() bool{
	var err error
	find := false
	var n int64
	for i , v := range db.Filename {
		if v == "" {
			db.Filename[i] = fmt.Sprintf("./data/fix_db_%d.bin",i)
			db.DbFile[i], err = os.OpenFile(db.Filename[i], os.O_CREATE|os.O_APPEND|os.O_RDWR,
				os.ModePerm|os.ModeTemporary)
			//defer db.dbFile[i].Close() //打开文件出错处理
			if err != nil {
				panic(err)
			}
			for n = 0; n < DB_FILE_SIZE; n++ {
				db.DbFile[i].Write([]byte("0000000000000000000000000000000000000000000000000" +
					"000000000000000000000000000000000000000000000000000"))
			}
			db.Filelen[i] = DB_FILE_SIZE
			
			find = true
			break
		}
	}
	if find == false {
		return false
	}
	return true
}

func (db *DB) RecordDataIndex(elem *ElemData) bool {
	f ,err := os.OpenFile("dbdata.index", os.O_CREATE|os.O_APPEND|os.O_RDWR,
		os.ModePerm|os.ModeTemporary)
	defer f.Close() //打开文件出错处理
	if err != nil {
		panic(err)
	}
	
	if elem == nil {
		return false
	}
	
	loc_start := strconv.FormatInt(elem.Loc.offstart, 10)
	loc_offset := strconv.FormatInt(elem.Loc.offset, 10)
	wstring := fmt.Sprintf("%s\t%s\t%s\n", elem.Loc.filename, loc_start, loc_offset)
	
	_, ierr := f.WriteString(wstring)
	if ierr != nil {
		return false
	}
	
	return true
}

func (db *DB) WriteElemData(elem *ElemData) bool {
	wSuccess := false
	if elem == nil {
		return false
	}
	for i , f := range db.DbFile {
		fileInfo, _ := f.Stat()
		filesize := fileInfo.Size()
		if filesize  < DB_FILE_SIZE {
			wbyte, _ := db.DbFile[i].Write(elem.Data)
			elem.Loc.offstart = filesize
			elem.Loc.offset  =  filesize + int64(wbyte)
			elem.Loc.filename = db.Filename[i]
			wSuccess = true
			break
		}
	}
	return wSuccess
}


func Test(){
	db := new(DB)
	db.CreateDbFile()
	
	elem := new(ElemData)
	elem.Data = []byte("abcabcabcabcabc")
	db.WriteElemData(elem)
	db.RecordDataIndex(elem)
	
}