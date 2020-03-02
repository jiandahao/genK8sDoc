package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os/exec"
	"strings"
)

func execCommand(name string, args ...string) (string, error){
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	cmd := exec.Command(name, args...)
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	if err := cmd.Run(); err != nil{
		return "", errors.New(stdErr.String())
	}
	return stdOut.String(), nil
}


//
//func execBlockedCommand(name string, args ...string) (string, error){
//
//}
var paramType  =  []string{
	"<string>",
	"<[]string>",
	"<Object>",
	"<[]Object>",
	"<integer>",
	"<boolean>",
	"<map[string]string>",
}
func isParamClaim(s string) (bool,string){
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "RESOURCE:"){
		return false, ""
	}
	//if ok,err := regexp.Match(`^\s*\S+\s+\S+\s*$`,[]byte(s)); !ok || err != nil{
	//	return false, ""
	//}
	for i := 0; i < len(paramType) ; i++{
		if strings.Contains(s, paramType[i]){
			return true, paramType[i]
		}
	}
	return false,""
}

// 转义符号< >,以便markdown标题中能正常显示变量类型
func QuoteMeta(s string) string{
	p := strings.ReplaceAll(s,"<",fmt.Sprintf("\\<"))
	p = strings.ReplaceAll(p, ">",fmt.Sprintf("\\>"))
	return p
}

func getParamNameFromParamPath(paramsPath string) string{
	pathList := strings.Split(paramsPath,".")
	if len(pathList) == 0 {
		log.Fatal("invalid path: " + paramsPath)
	}
	return pathList[len(pathList)-1]
}

func getParamNameFromParamClaim( claim string) string{
	claim = strings.TrimSpace(claim)
	s := strings.ReplaceAll(claim, " ","")
	t := strings.Split(s,"<")
	s = t[0]
	s = strings.TrimSpace(s)
	return s
}

func makeTitlePrefix( prefix string) string{
	if len(prefix) > 6{
		//newPrefix := ""
		//length := len(prefix) - 7
		//for ; length>0; length--{
		//	newPrefix += "\t"
		//}
		//return newPrefix + "- "
		return "-"
	}
	return prefix
}
// 深度遍历所有的参数，并获取其说明文档
func parseDocument(titlePrefix string, paramPath string,lastType string, doc *string) error {
	if !strings.Contains(lastType,"Object"){
		return nil
	}
	//if len(titlePrefix) > 6{
	//	titlePrefix = ""
	//	length := len(titlePrefix) - 7
	//	for ;length > 0; length--{
	//		titlePrefix +=  "\t"
	//	}
	//	titlePrefix += "- "
	//}

	// 执行命令，并获取执行结果
	cmd := exec.Command("kubectl","explain",paramPath)
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	if err := cmd.Run(); err != nil{
		//log.Println(stdErr.String())
		return errors.New(stdErr.String())
	}
	result := stdOut.String()
	// 将结果按行分割，方便后续解析出参数
	lines := strings.Split(result,"\n")

	//preContent := titlePrefix + " " +
	//	getParamNameFromParamPath(paramPath) +
	//	" " + QuoteMeta(lastType) + "\n"
	preContent := fmt.Sprintf(
		"%s %s %s\n**PATH:**  %s\n\n",
		makeTitlePrefix(titlePrefix), getParamNameFromParamPath(paramPath),QuoteMeta(lastType),paramPath)
	fmt.Println(makeTitlePrefix(titlePrefix) + " " + paramPath)

	preParamName := ""
	preParamType := ""
	for i := 0; i< len(lines); i++{
		if strings.HasPrefix(lines[i],"FIELDS:"){
			continue
		}
		if ok, paramType := isParamClaim(lines[i]); ok {
			paramName := getParamNameFromParamClaim(lines[i])
			if preParamName != ""{
				// 如果之前存在一个参数还没有处理，现在进行处理，这么做主要是保证先记录参数的简介，然后在加入更详细的说明信息
				if err := parseDocument(titlePrefix + "#", paramPath+"."+ preParamName, preParamType ,doc); err != nil{
					return err
				}
			}
			*doc += preContent
			//fmt.Println(preContent)
			preContent = fmt.Sprintf(
				"%s %s\n**PATH:**  %s\n\n",
				makeTitlePrefix(titlePrefix + "#"), strings.TrimSpace(QuoteMeta(lines[i])),paramPath+"."+ paramName)
			//preContent = titlePrefix + " " + strings.TrimSpace(lines[i]) + "\n"
			preParamName = paramName
			preParamType = paramType
		}else {
			preContent += lines[i] + "\n"
		}
	}
	if err := parseDocument(titlePrefix + "#", paramPath+"."+ preParamName, preParamType,doc); err != nil{
		return err
	}
	return nil
}

func saveDocument(doc string, filename string) {
	if err := ioutil.WriteFile(filename,[]byte(doc),0644); err != nil {
		log.Println(err)
	}
}

var s = flag.String("s", "","resource name")

func main(){
	flag.Parse()
	allResources := []string{}
	if *s != ""{
		resources :=  strings.Split(*s, " ")
		for i:=0 ; i< len(resources) ; i++{
			allResources = append(allResources, strings.TrimSpace(resources[i]))
		}
	}else{
		allResources = []string{"pod","deployment","daemonset","statefulset","Service","ClusterRoleBinding","ServiceAccount"}
	}
	for i :=0 ; i < len(allResources); i++{
		document := "# " + allResources[i] + "\n"
		if err := parseDocument("",allResources[i],"Object",&document); err != nil{
			log.Println(err)
			return
		}
		saveDocument(document,"doc/"+allResources[i] + ".md")
	}
}
