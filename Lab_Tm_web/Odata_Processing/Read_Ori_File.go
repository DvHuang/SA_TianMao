package Odata_Processing

import (
	"fmt"
	"strconv"

	"bufio"

	"os"
	"log"
"strings"


"Lab_Tm_web/segment"
"Lab_Tm_web/Global_Struct"
)

const 	(
//SaveUnTestDimen="C:/Users/JTdavy/IdeaProjects/src/DMforTmC/readfile/SaveUnTestDimen.txt"
	SaveUnTestDimen="D:/D-go/src/DMforTmC-0/readfile/Testdata/SaveUnTestDimen.txt"

)

var (
	segmenter sego.Segmenter
	
)


func Openfile (file string) (f *os.File ,err error){

	infile :=file

	var err1 error

	if _,err:=os.Stat(infile);os.IsNotExist(err){
		fmt.Println("文件不存在,创建")
		f,err1 =os.Create(infile)
		//fmt.Printf("f is %v ",f)

	}else{
		fmt.Println("文件存在")
		f,err1 =os.OpenFile(infile,os.O_APPEND,0666)
		//fmt.Printf("f is %v ",f)

	}

	return f,err1
}




func Read_ordata_File(OriginalModefile string)  (OriginalModeSets map[string]Global_Struct.OneItemStruct) {

	// 读入数据
	file, err := os.Open(OriginalModefile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	OriginalModeSets=make(map[string]Global_Struct.OneItemStruct)

	var oneitem Global_Struct.OneItemStruct
	scanner := bufio.NewScanner(file)


	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "||||")

		//fmt.Println("data.........",data)
		if len(data) != 3 {
			continue
		}
		i, err := strconv.ParseInt(data[2], 10, 64)
		if err != nil {
			panic(err)
		}
		oneitem.Class=int(i)
		oneitem.Data= data[1]
		//fmt.Println("oneitem ........",oneitem)
		OriginalModeSets[data[0]] = oneitem
	}
	return OriginalModeSets
}
//create hash table for feature


func Create_Fea_Hash(modeDaSet map[string]Global_Struct.OneItemStruct, Ev_map map [string]Global_Struct.Emotional_vocabulary)(map[string]Global_Struct.Fea_str) {

	feature_Hash :=make(map[string]Global_Struct.Fea_str)

	segmenter.LoadDictionary("D:/WBfile/dictionary.txt")

	//分词，得到fh
	for key,_ := range modeDaSet {

		text := []byte(key)

		segments := segmenter.Segment(text)

		feature_Hash = sego.DavysegmentsTo_Create_Fea_Hash(segments, feature_Hash)
	}


	fmt.Println("Ev_map len =", len(Ev_map))

	//fh情感化
	for key,F_value:=range feature_Hash{

		if Ev_value,ok:=Ev_map[key];ok{

			var linshi Global_Struct.Fea_str

			linshi.Lexical_category=F_value.Lexical_category
			linshi.M_E_Polarity=Ev_value.M_E_Polarity
			linshi.M_E_Strength=Ev_value.M_E_Strength
			linshi.M_E_classification=Ev_value.M_E_classification
			linshi.Frequency=F_value.Frequency
			linshi.Pos=Ev_value.Pos


			feature_Hash[key]=linshi
		}

	}

	return  feature_Hash
}

func Create_Data_Hash(modeDaSet map[string]Global_Struct.OneItemStruct) (map[string]Global_Struct.MapPrefixPath) {

	data_Hash_S := make(map[string]Global_Struct.MapPrefixPath)
	segmenter.LoadDictionary("D:/WBfile/dictionary.txt")


	for key,value := range modeDaSet {

		text := []byte(key)

		segments := segmenter.Segment(text)

		item_m, stringMapPrefixPath := sego.DavysegmentsToString_allstr(segments)

		item_S, _ := data_Hash_S[stringMapPrefixPath]
		item_S.Path = make(map[string]int)

		for key, invalue := range item_m {
			item_S.Path[key] = invalue
			item_S.Data=value.Data
			item_S.Class = 1			//该值作为事务预设类别
		}
		
		data_Hash_S[stringMapPrefixPath] = item_S

	}

	return  data_Hash_S

}

func Creat_fea_data_hash( modeDaSet map[string]Global_Struct.OneItemStruct, Ev_map map [string]Global_Struct.Emotional_vocabulary)(map[string]Global_Struct.Fea_str,map[string]Global_Struct.MapPrefixPath){

	feature_Hash :=make(map[string]Global_Struct.Fea_str)

	segmenter.LoadDictionary("D:/WBfile/dictionary.txt")

	//分词，得到fh
	for key,_ := range modeDaSet {

		text := []byte(key)

		segments := segmenter.Segment(text)

		feature_Hash = sego.DavysegmentsTo_Create_Fea_Hash(segments, feature_Hash)
	}


	fmt.Println("Ev_map len =", len(Ev_map))

	//fh情感化
	for key,F_value:=range feature_Hash{

		if Ev_value,ok:=Ev_map[key];ok{

			var linshi Global_Struct.Fea_str

			linshi.Lexical_category=F_value.Lexical_category
			linshi.M_E_Polarity=Ev_value.M_E_Polarity
			linshi.M_E_Strength=Ev_value.M_E_Strength
			linshi.M_E_classification=Ev_value.M_E_classification
			linshi.Frequency=F_value.Frequency
			linshi.Pos=Ev_value.Pos


			feature_Hash[key]=linshi
		}

	}



	data_Hash_S := make(map[string]Global_Struct.MapPrefixPath)
	//segmenter.LoadDictionary("D:/WBfile/dictionary.txt")


	for key,value := range modeDaSet {

		text := []byte(key)

		segments := segmenter.Segment(text)

		item_m, stringMapPrefixPath := sego.DavysegmentsToString_allstr(segments)

		item_S, _ := data_Hash_S[stringMapPrefixPath]
		item_S.Path = make(map[string]int)

		for key, invalue := range item_m {
			item_S.Path[key] = invalue
			item_S.Data=value.Data
			item_S.Class = 1			//该值作为事务预设类别
		}

		data_Hash_S[stringMapPrefixPath] = item_S

	}

	return     feature_Hash,data_Hash_S


}
