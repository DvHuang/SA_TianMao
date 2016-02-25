package main

import (
	"net/http"
	"log"
	"io"
	"encoding/json"
	"Lab_Tm_web/Estimate/Positive_Negative"
	"Lab_Tm_web/TMspider"
	"Lab_Tm_web/Odata_Processing"
	"fmt"
	"Lab_Tm_web/Global_Struct"

)

const (
	odataFile="D:/D-go/src/Lab_Tm_web/Estimate/o_data_File.txt"

	P_data_file="D:/D-go/src/Lab_Tm_web/http/P_data_file.txt"
	N_data_file="D:/D-go/src/Lab_Tm_web/http/N_data_file.txt"

	Ev_file="D:/D-go/src/Lab_Tm_web/Estimate/Positive_Negative/情感词汇.txt"

)

var (

	Ev_map map[string] Global_Struct.Emotional_vocabulary
	P_data_map map[string]Global_Struct.MapPrefixPath
	N_data_map map[string]Global_Struct.MapPrefixPath
)

func JsonRpcServer(w http.ResponseWriter, req *http.Request) {

	query := req.URL.Query().Get("query")

	if len(query)>50{
		fmt.Println("进入json",query)


		//TMspider.Spider(query,odataFile)
		//o_data_Map:= Odata_Processing.Read_ordata_File(odataFile)

		//获取用户地址评论的原始数据
		o_data_Map:=TMspider.Spider_map(query)
		Ev_Feature_map,Data_map:=Odata_Processing.Creat_fea_data_hash(o_data_Map,Ev_map)
		//Ev_Feature_map:= Odata_Processing.Create_Fea_Hash(o_data_Map,Ev_map)
		//Data_map:= Odata_Processing.Create_Data_Hash(o_data_Map)


		//Odata_Processing.Write_Feature_Hash(Ev_Feature_map)
		//Odata_Processing.Write_data_class(P_data_file,N_data_file,P_data_map,N_data_map)
		P_data_map,N_data_map=Positive_Negative.Predict(Ev_Feature_map,Data_map)


		var docs []string
		for key, value  := range N_data_map {

			one_item := key+"|||||||||"+value.Data

			docs = append(docs,one_item)
		}
		response, _ := json.Marshal(docs)



		w.Header().Set("Content-Type", "application/json")

		io.WriteString(w, string(response))

	}


}

func main(){


	//获取积极词语列表
	Ev_map=Positive_Negative.Creat_EV(Ev_file)

	http.HandleFunc("/json", JsonRpcServer)
	http.Handle("/", http.FileServer(http.Dir("http")))
	log.Print("服务器启动")
	http.ListenAndServe(":9888", nil)

}
