package Positive_Negative


import (


"Lab_Tm_web/Global_Struct"
)

func Predict (feature_Hash_S map[string]Global_Struct.Fea_str,  data_Hash_S map[string]Global_Struct.MapPrefixPath)(map[string]Global_Struct.MapPrefixPath,map[string]Global_Struct.MapPrefixPath){

	P_data_Hash :=make (map[string]Global_Struct.MapPrefixPath)
	N_data_Hash :=make (map[string]Global_Struct.MapPrefixPath)


	for data_key,data_value:=range data_Hash_S{

		var linshi Global_Struct.MapPrefixPath

		linshi.Data=data_value.Data
		linshi.Path=data_value.Path


		for path_key,_ :=range data_value.Path{

			if Fhvalue,ok:=feature_Hash_S[path_key];ok{

				var Me_c string
				//fmt.Println("Fhvalue.M_E_classification",Fhvalue.M_E_classification)
				if len(Fhvalue.M_E_classification)>1 {
					Me_c = Fhvalue.M_E_classification[0:1]
					//fmt.Println("Me_c",Me_c)
				}else{
					break
				}

				if Me_c=="P"{
					linshi.Class=1
					P_data_Hash[data_key]=linshi
				}else {
					linshi.Class=2
					N_data_Hash[data_key]=linshi
				}
			}else{
				continue
			}
		}
	}

	return  P_data_Hash,N_data_Hash

}