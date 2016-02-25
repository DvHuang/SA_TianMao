package TMspider

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
"Lab_Tm_web/Global_Struct"
"Lab_Tm_web/Odata_Processing"
)



func getTotal(url string) int {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("http get response error=%s\n", err)

	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	stringbody := string(body)

	indexstart := strings.Index(stringbody, "total")
	indexend := strings.Index(stringbody, "used")

	fmt.Println("start and end number ", indexstart, indexend)

	count := body[indexstart+7 : indexend-2]
	fmt.Println("[]byte", count)

	fmt.Println("string", string(count))

	stringcount := string(count)

	pageint, error := strconv.Atoi(stringcount)
	if error != nil {
		fmt.Println("page 字符串转换成整数失败")
	}
	page := pageint / 20
	if page >= 99 {
		page = 99
	}
	return page
}


func get_selerid( errorUrl string ) string {

	fmt.Println("errorUrl.........",errorUrl)
	resp, err := http.Get(errorUrl)
	if err != nil {
		fmt.Printf("http get response error=%s\n", err)

	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)


	utf8alldata, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader(body), simplifiedchinese.GBK.NewDecoder()))
	stralldata := string(utf8alldata)

	index_iSid := strings.Index(stralldata, "&sellerId=")

	itemid := stralldata[index_iSid+10 : index_iSid+130]
	fmt.Println("itemid......",itemid)
	end_Sid := strings.Index(itemid, "&")
	fmt.Println("end_Sid......",end_Sid)

	itemid = itemid[:end_Sid]

	return  itemid


}
func get_itemid( errorUrl string ) string {

	index_itemid := strings.Index(errorUrl, "&id=")
	itemid := errorUrl[index_itemid+4 : index_itemid+30]
	end_itemid := strings.Index(itemid, "&")
	itemid = itemid[:end_itemid]
	return itemid
}


func real_Url(errorUrl string)string{

	selerId:=get_selerid(errorUrl)
	fmt.Println("selerId...=",selerId)
	itemid:=get_itemid(errorUrl)

	URL:="https://rate.tmall.com/list_detail_rate.htm?itemId="+itemid+"&sellerId="+selerId+"&currentPage="
	return URL
}


func Spider_map(errorURL string )( map[string]Global_Struct.OneItemStruct) {

	var alldata []byte

	//errorURL := "https://detail.tmall.com/item.htm?spm=a1z10.3-b.w4011-2062703342.105.yx9LnJ&id=22563195262&rn=a7492be1f04123b9cee4429b0bfcef8b&abbucket=9"


	URL:=real_Url(errorURL)

	fmt.Println("url....",URL)

	page := getTotal(URL)

	//抓
	for i := 1; i < page; i++ {


		url := URL
		istring := strconv.Itoa(i)
		url += istring

		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("http get response error=%s\n", err)
			break
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		alldata = append(alldata, body...)
	}



	utf8alldata, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader(alldata), simplifiedchinese.GBK.NewDecoder()))
	stralldata := string(utf8alldata)

	//fmt.Println(stralldata)

	stringcount := strings.Count(stralldata, "rateContent")

	//存

	OriginalModeSets:=make(map[string]Global_Struct.OneItemStruct)

	for stringcount > 1 {

		stringcount -= 1

		indexratecount := strings.Index(stralldata, "rateContent")
		indexratedate := strings.Index(stralldata, "rateDate")

		rateContent := stralldata[indexratecount+14 : indexratedate-3]
		rateDate := stralldata[indexratedate+11 : indexratedate+30]

		StringrateContent := string(rateContent)

		/*if len(StringrateContent) < 2 {

			stralldata = stralldata[indexratedate+31:]
			continue
		}*/

		StringrateDate := string(rateDate)
		stralldata = stralldata[indexratedate+31:]

		var oneitem Global_Struct.OneItemStruct
		oneitem.Class=1
		oneitem.Data= StringrateDate
		//fmt.Println("oneitem ........",oneitem)
		OriginalModeSets[StringrateContent] = oneitem

	}

	return  OriginalModeSets
}





// write file
func Spider(errorURL string,o_data_File string) {

	var alldata []byte

	//errorURL := "https://detail.tmall.com/item.htm?spm=a1z10.3-b.w4011-2062703342.105.yx9LnJ&id=22563195262&rn=a7492be1f04123b9cee4429b0bfcef8b&abbucket=9"


	URL:=real_Url(errorURL)

	fmt.Println("url....",URL)

	page := getTotal(URL)

	//抓
	for i := 1; i < page; i++ {


		url := URL
		istring := strconv.Itoa(i)
		url += istring

		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("http get response error=%s\n", err)
			break
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		alldata = append(alldata, body...)
	}

	Originalf, errFile := Odata_Processing.Openfile(o_data_File)
	if errFile != nil {
		fmt.Println("打开文件失败")
	}
	defer Originalf.Close()

	utf8alldata, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader(alldata), simplifiedchinese.GBK.NewDecoder()))
	stralldata := string(utf8alldata)

	//fmt.Println(stralldata)

	stringcount := strings.Count(stralldata, "rateContent")

	//存

	for stringcount > 1 {

		stringcount -= 1

		indexratecount := strings.Index(stralldata, "rateContent")
		indexratedate := strings.Index(stralldata, "rateDate")

		rateContent := stralldata[indexratecount+14 : indexratedate-3]
		rateDate := stralldata[indexratedate+11 : indexratedate+30]

		StringrateContent := string(rateContent)

		/*if len(StringrateContent) < 2 {

			stralldata = stralldata[indexratedate+31:]
			continue
		}*/

		StringrateDate := string(rateDate)
		stralldata = stralldata[indexratedate+31:]

		StringAdd := StringrateContent + "||||" + StringrateDate + "||||" + "1" + "\n"
		io.WriteString(Originalf, StringAdd)

	}
}
