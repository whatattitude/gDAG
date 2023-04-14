// Built-in environment variable reader
package env

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type GOENV struct {
	Ar           string `json:"AR"`
	Cc           string `json:"CC"`
	CgoCflags    string `json:"CGO_CFLAGS"`
	CgoCppflags  string `json:"CGO_CPPFLAGS"`
	CgoCxxflags  string `json:"CGO_CXXFLAGS"`
	CgoEnabled   string `json:"CGO_ENABLED"`
	CgoFflags    string `json:"CGO_FFLAGS"`
	CgoLdflags   string `json:"CGO_LDFLAGS"`
	Cxx          string `json:"CXX"`
	Gccgo        string `json:"GCCGO"`
	Go111Module  string `json:"GO111MODULE"`
	Goarch       string `json:"GOARCH"`
	Gobin        string `json:"GOBIN"`
	Gocache      string `json:"GOCACHE"`
	Goenv        string `json:"GOENV"`
	Goexe        string `json:"GOEXE"`
	Goexperiment string `json:"GOEXPERIMENT"`
	Goflags      string `json:"GOFLAGS"`
	Gogccflags   string `json:"GOGCCFLAGS"`
	Gohostarch   string `json:"GOHOSTARCH"`
	Gohostos     string `json:"GOHOSTOS"`
	Goinsecure   string `json:"GOINSECURE"`
	Gomod        string `json:"GOMOD"`
	Gomodcache   string `json:"GOMODCACHE"`
	Gonoproxy    string `json:"GONOPROXY"`
	Gonosumdb    string `json:"GONOSUMDB"`
	Goos         string `json:"GOOS"`
	Gopath       string `json:"GOPATH"`
	Goprivate    string `json:"GOPRIVATE"`
	Goproxy      string `json:"GOPROXY"`
	Goroot       string `json:"GOROOT"`
	Gosumdb      string `json:"GOSUMDB"`
	Gotmpdir     string `json:"GOTMPDIR"`
	Gotooldir    string `json:"GOTOOLDIR"`
	Govcs        string `json:"GOVCS"`
	Goversion    string `json:"GOVERSION"`
	Gowork       string `json:"GOWORK"`
	PkgConfig    string `json:"PKG_CONFIG"`
}


func GetEnvDefault() (goEnv GOENV, err error) {
	// 通过exec.Command函数执行命令或者shell
	// 第一个参数是命令路径，当然如果PATH路径可以搜索到命令，可以不用输入完整的路径
	// 第二到第N个参数是命令的参数
	// 下面语句等价于执行命令: ls -l /var/
	cmd := exec.Command("go", "env", "-json")
	// 执行命令，并返回结果
	output,err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return 
	}
	
	err = json.Unmarshal(output,&goEnv)
	if err != nil{
		fmt.Println(err)
	}
	return 
}



