package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strconv"
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
	if ok,err := regexp.Match(`^\s*\S+\s+\S+\s*$`,[]byte(s)); !ok || err != nil{
		return false, ""
	}
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

func getDocument(prefix string, param string) (string,error){
	cmd := exec.Command("kubectl","explain",param)
	stdout, err := cmd.StdoutPipe()
	if err != nil{
		return "",nil
	}
	cmd.Start()
	reader := bufio.NewReader(stdout)
	document := ""
	preDoc := ""
	for{
		line, err := reader.ReadString('\n')
		if err != nil || err == io.EOF{
			break
		}
		if strings.HasPrefix(line, "KIND:"){
			kind := strings.TrimPrefix(line, "KIND:")
			fmt.Println("## "+strings.TrimSpace(kind))
			document += fmt.Sprintf("%s %s\n",prefix,kind)
			continue
		}
		if ok, s := isParamClaim(line); ok{
			if preDoc != ""{
				document += fmt.Sprintf("%s",preDoc)
				preDoc = ""
			}
			paramName := strings.TrimSuffix(line, s)
			paramName = strings.TrimSpace(paramName)
			if strings.Contains(s,"Object"){

			}
		}else{
			preDoc += line
		}
	}
	cmd.Wait()
	return "",nil
}

func getAllParams(depth int, params string, lastType string) {
	if !strings.Contains(lastType,"Object"){
		return
	}
	cmd := exec.Command("kubectl","explain",params)
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	if err := cmd.Run(); err != nil{
		log.Println(stdErr.String())
		return
	}
	result := stdOut.String()
	lines := strings.Split(result,"\n")
	for i := 0; i< len(lines); i++{
		if ok, s := isParamClaim(lines[i]); ok {
			paramName := strings.TrimSuffix(lines[i], s)
			paramName = strings.TrimSpace(paramName)
			fmt.Println(strconv.Itoa(depth) + " " + params + "." + paramName)
			getAllParams(depth+1, params+"."+paramName, s)
		}
	}
}


func getParamNameFromParamPath(paramsPath string) string{
	pathList := strings.Split(paramsPath,".")
	if len(pathList) == 0 {
		log.Fatal("invalid path: " + paramsPath)
	}
	return pathList[len(pathList)-1]
}

// 深度遍历所有的参数，并获取其说明文档
func parseDocument(titlePrefix string, paramPath string,lastType string, doc *string){
	if !strings.Contains(lastType,"Object"){
		return
	}
	if len(titlePrefix) > 6{
		titlePrefix = "- "
	}

	// 执行命令，并获取执行结果
	cmd := exec.Command("kubectl","explain",paramPath)
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	if err := cmd.Run(); err != nil{
		log.Println(stdErr.String())
		return
	}
	result := stdOut.String()
	// 将结果按行分割，方便后续解析出参数
	lines := strings.Split(result,"\n")

	//preContent := titlePrefix + " " +
	//	getParamNameFromParamPath(paramPath) +
	//	" " + QuoteMeta(lastType) + "\n"
	preContent := fmt.Sprintf(
		"%s %s %s\n**PATH:**  %s\n\n",
		titlePrefix, getParamNameFromParamPath(paramPath),QuoteMeta(lastType),paramPath)
	fmt.Println(titlePrefix + " " + paramPath)

	preParamName := ""
	preParamType := ""
	for i := 0; i< len(lines); i++{
		if strings.HasPrefix(lines[i],"FIELDS:"){
			continue
		}
		if ok, s := isParamClaim(lines[i]); ok {
			paramName := strings.TrimSuffix(lines[i], s)
			paramName = strings.TrimSpace(paramName)
			if preParamName != ""{
				// 如果之前存在一个参数还没有处理，现在进行处理，这么做主要是保证先记录参数的简介，然后在加入更详细的说明信息
				parseDocument(titlePrefix + "#", paramPath+"."+ preParamName, s,doc)
			}
			*doc += preContent
			preContent = fmt.Sprintf(
				"%s %s\n**PATH:**  %s\n\n",
				titlePrefix + "#", strings.TrimSpace(QuoteMeta(lines[i])),paramPath+"."+ paramName)
			//preContent = titlePrefix + " " + strings.TrimSpace(lines[i]) + "\n"
			preParamName = paramName
			preParamType = s
		}else {
			preContent += lines[i] + "\n"
		}
	}
	parseDocument(titlePrefix + "#", paramPath+"."+ preParamName, preParamType,doc)
}

func saveDocument(doc string, filename string) {
	if err := ioutil.WriteFile(filename,[]byte(doc),0644); err != nil {
		log.Println(err)
	}
}

func main(){
	allResources := []string{"pod","deployment","daemonset","statefulset","Service"}
	for i :=0 ; i < len(allResources); i++{
		document := ""
		parseDocument("#",allResources[i],"Object",&document)
		saveDocument(document,"doc/"+allResources[i] + ".md")
	}

}
