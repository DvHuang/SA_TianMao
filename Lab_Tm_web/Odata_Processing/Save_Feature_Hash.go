package Odata_Processing
import (
	"strconv"
	"io"

	"os"
	"log"
	"bufio"
	"strings"

	"sort"
	"fmt"
"Lab_Tm_web/Global_Struct"
)

const(
		Feature1="D:/D-go/src/Lab_Tm/Global_Struct/P_Feature.txt"
		Feature2="D:/D-go/src/Lab_Tm/Global_Struct/N_Feature.txt"

)



type PsliceList_F []Global_Struct.L_Fea_str

func (list PsliceList_F) Len() int {
	return len(list)
}
func (list PsliceList_F) Less(i, j int) bool {
	if list[i].Frequency > list[j].Frequency {
		return true
	} else if list[i].Frequency < list[j].Frequency {
		return false
	} else {
		return list[i].Frequency < list[j].Frequency
	}
}
func (list PsliceList_F) Swap(i, j int) {
	var temp Global_Struct.L_Fea_str = list[i]
	list[i] = list[j]
	list[j] = temp
}


func Write_Feature_Hash( Ev_Feature_map map [string]Global_Struct.Fea_str) {

	file, err := Openfile(Feature1)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file2, err := Openfile(Feature2)
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()



	//排序

	var IndexList PsliceList_F
	fmt.Println("Ev_Feature_map len=",len(Ev_Feature_map))
	for key,efvalue:=range Ev_Feature_map {

		var oneF Global_Struct.L_Fea_str
		oneF.Name=key
		oneF.Frequency=efvalue.Frequency
		oneF.Lexical_category=efvalue.Lexical_category
		oneF.M_E_classification=efvalue.M_E_classification
		oneF.M_E_Polarity=efvalue.M_E_Polarity
		oneF.M_E_Strength=efvalue.M_E_Strength
		oneF.Pos=efvalue.Pos
		IndexList = append(IndexList, oneF)
	}
	fmt.Println("IndexList len =",len(IndexList))

	sort.Sort(IndexList)
	//写入
	for number ,value:=range IndexList{

		var s string

		if len(value.M_E_classification)>1 {
			s = value.M_E_classification[0:1]

		}else{
			continue
		}

		onehash:=strconv.Itoa(number)+"||||"+value.Name+"||||"+strconv.Itoa(value.Frequency)+"||||"+value.M_E_classification+"\n"

		if s=="P" {
			_, errWrit := io.WriteString(file, onehash) //写入文件(字符串)

			if errWrit != nil {
				fmt.Println("写入失败")
			}
		}else{
			_, errWrit := io.WriteString(file2, onehash) //写入文件(字符串)

			if errWrit != nil {
				fmt.Println("写入失败")
			}
		}
	}

}
func Read_Feature_Hash(FeatureHashFile string) (FeatureHashTable map [string]Global_Struct.Fea_str) {

	file, err := os.Open(FeatureHashFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	FeatureHashTable=make(map [string]Global_Struct.Fea_str)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "||||")
		if len(data) != 3 {
			continue
		}
		var linshi Global_Struct.Fea_str
		linshi.Lexical_category=data[2]


		//linshi.Frequency=strconv.ParseUint(data[1], 10, 64)
		linshi.Frequency=1
		//oneitem.class=strconv.ParseUint(data[2], 1, 1)
		//FeatureHashTable[data[0]]= strconv.ParseInt(data[1],10,64)
		FeatureHashTable[data[0]]= linshi
	}

	return FeatureHashTable
}

