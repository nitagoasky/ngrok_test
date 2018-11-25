package util

import (
    "net/http"
    "io/ioutil"
    "bytes"
    "strconv"
)

func GetData() []byte {
	//生成client 参数为默认
    client := &http.Client{}
    
    //生成要访问的url
    url := "http://www.baidu.com"
        
    //提交请求
    reqest, err := http.NewRequest("GET", url, nil)
    
    if err != nil {
        panic(err)
    }

    //处理返回结果
    response, _ := client.Do(reqest)
    responseBody := response.Body;
    // 
	content, _ := ioutil.ReadAll(responseBody)
	l2:=bytes.Count([]byte(content),nil)-1
    length := "Content-Length: " + strconv.Itoa(l2)
    str := response.Proto + " " + response.Status + "\n" + length + "\n\n"
    

	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲

	buffer.Write([]byte(str))

	buffer.Write(content)

	b3 :=buffer.Bytes()
   //将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理

    return b3
}