package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

func main(){
	url := "https://github.com/probloys/webCatchForGit/file-list/master"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	ioutil.WriteFile("response.txt",body,0644)
	//fmt.Println(res)
	//fmt.Println(string(body))
//https://github.com/probloys/webCatchForGit/commit/f2a506ff81b9f1ee026e379d25227085c062bea3
//github.githubassets.com js-repo-pjax-container
	x:=MatchLine(string(body))
	for _,v:=range x{
		println(v)
	}


}

func MatchLine(partern string)[]string{
	regexp.QuoteMeta("*href*")
	regex,err:=regexp.Compile("href.*")
	if err !=nil{
		print("compile pattern error")
		return nil
	}
	resultString:=regex.FindAllString(partern,10)

	return resultString
}
