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

func parseDocument(depth int, paramPath string,lastType string, doc *string){
	if !strings.Contains(lastType,"Object"){
		return
	}
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
	lines := strings.Split(result,"\n")
	preContent := ""
	preParamName := ""
	preParamType := ""
	for i := 0; i< len(lines); i++{
		if ok, s := isParamClaim(lines[i]); ok {
			if preParamName != ""{
				*doc += preContent
				preContent = ""
				parseDocument(depth+1, paramPath+"."+ preParamName, s,doc)
			}
			paramName := strings.TrimSuffix(lines[i], s)
			paramName = strings.TrimSpace(paramName)
			preParamName = paramName
			preParamType = s
		}else {
			preContent += lines[i] + "\n"
		}
	}
	parseDocument(depth+1, paramPath+"."+ preParamName, preParamType,doc)
}

func saveDocument(doc string, filename string) {
	//f, err := os.Create(filename)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer f.Close()
	////if n,err := f.WriteString(doc); err != nil{
	////	log.Println("error")
	////	log.Println(err)
	////}else{
	////	log.Println("success")
	////	log.Println(n)
	////}
	//f.WriteString(doc)
	//f.Sync()
	if err := ioutil.WriteFile(filename,[]byte(doc),0644); err != nil {
		log.Println(err)
	}
}
func main(){
	//fmt.Println(strings.Split("abc\ncdf\n","'\n'"))
	document := ""
	//getAllParams(1, "pod", "Object")
	parseDocument(1,"pod","Object",&document)
	//fmt.Println(document)
	//saveDocument(&document,"temp.md")
	saveDocument(document,"temp.md")
	//cmd := exec.Command("kubectl","explain","deployment")
	//stdout, err := cmd.StdoutPipe()
	//
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//
	//cmd.Start()
	//reader := bufio.NewReader(stdout)
	//for{
	//	line, err := reader.ReadString('\n')
	//	if err != nil || err == io.EOF{
	//		break
	//	}
	//	if strings.HasPrefix(line, "KIND:"){
	//		kind := strings.TrimPrefix(line, "KIND:")
	//		fmt.Println("## "+strings.TrimSpace(kind))
	//		continue
	//	}
	//
	//	if ok, s := isParamClaim(line); ok{
	//		paramName := strings.TrimSuffix(line, s)
	//		paramName = strings.TrimSpace(paramName)
	//
	//	}
	//}
	//
	//cmd.Wait()

}
