package Global_Struct



type OneItemStruct struct  {
	Class int
	Data  string
}

type Fea_str struct {
	Frequency int					//该元总频率
	Lexical_category  string		//字元属性


	Pos	  string						//词性种类,

	M_E_classification  string			//主情感分类,
	M_E_Strength string					//强度,
	M_E_Polarity string					//极性,

	Auxiliary_emotion_classification string		//辅助情感分类,
	A_E_Strength string							//强度,
	A_E_polarity	string						//极性

}
type L_Fea_str struct {

	Name string
	Frequency int					//该元总频率
	Lexical_category  string		//字元属性


	Pos	  string						//词性种类,

	M_E_classification  string			//主情感分类,
	M_E_Strength string					//强度,
	M_E_Polarity string					//极性,

	//Auxiliary_emotion_classification string		//辅助情感分类,
	//A_E_Strength string							//强度,
	//A_E_polarity	string						//极性

}

type MapPrefixPath struct {
	Path  map[string]int	//事务词元Map 及事务内该元频率
	Class int				//事务属性
	Data string				//事务日期
}

type Emotional_vocabulary struct {

			Pos	  string						//词性种类,

			M_E_classification  string			//主情感分类,
			M_E_Strength string					//强度,
			M_E_Polarity string					//极性,

			Auxiliary_emotion_classification string		//辅助情感分类,
			A_E_Strength string							//强度,
			A_E_polarity	string						//极性


}