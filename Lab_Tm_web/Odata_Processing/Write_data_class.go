package Odata_Processing
import (

	"log"
	"io"
	"fmt"
"Lab_Tm_web/Global_Struct"
)


func Write_data_class(P_data_file ,N_data_file string,P_da_M,N_da_M map [string]Global_Struct.MapPrefixPath)  {

	file_P, err := Openfile(P_data_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file_P.Close()

	file_N, err := Openfile(N_data_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file_N.Close()



	//写入
	for key,value:=range P_da_M{

		//fmt.Println("k.....",key,value.Class)
		if value.Class==1{
			onehash:=key+"||||"+value.Data+"||||"+"\n"

			_, errWrit := io.WriteString(file_P, onehash) //写入文件(字符串)

			if errWrit !=nil{
				fmt.Println("写入失败")
			}
		}else {

			onehash:=key+"||||"+value.Data+"\n"



			_, errWrit := io.WriteString(file_N, onehash) //写入文件(字符串)

			if errWrit !=nil{
				fmt.Println("写入失败")
			}
		}
	}

	//写入
	for key,value:=range N_da_M{
		//fmt.Println("k2.....",key,value.Class)
		if value.Class==1{
			onehash:=key+"||||"+value.Data+"||||"+"\n"
			_, errWrit := io.WriteString(file_P, onehash) //写入文件(字符串)


			if errWrit !=nil{
				fmt.Println("写入失败")
			}
		}else {

			onehash:=key+"||||"+value.Data+"\n"

			_, errWrit := io.WriteString(file_N, onehash) //写入文件(字符串)

			if errWrit !=nil{
				fmt.Println("写入失败")
			}
		}
	}
}