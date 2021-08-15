package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "gopkg.in/gcfg.v1"
    "gopkg.in/yaml.v3"
    "io/ioutil"
    "log"
    "os"
)

func main_read_txt() {
    // data, err := ioutil.ReadFile("F:/Go/my_project/file_go/text.txt")
    // if err != nil {
    //     fmt.Println("file reading error", err)
    // }
    // fmt.Println("contents of file: ", string(data))
    // 可以从命令行获取文件路径作为输入，然后读取其内容。
    // 这个函数返回存储 flag 值的字符串变量的地址
    fptr := flag.String("fpath", "test.txt", "file path to read from")
    flag.Parse()
    // 在程序访问 flag 之前，必须先调用 flag.Parse()。
    fmt.Println("value of fpath is ", *fptr)
    data, err := ioutil.ReadFile(*fptr)
    if err != nil {
        fmt.Println("file reading error ", err)
        return
    }
    fmt.Println("contents of file: ", string(data))
}

func main_read_ini() {
    config := struct {
        Section struct {
            Enabled bool
            Path    string
        }
    }{}
    err := gcfg.ReadFileInto(&config, "conf.ini")
    if err != nil {
        fmt.Println("Failed to parse config file: %s ", err)
    }
    fmt.Println(config.Section.Enabled)
    fmt.Println(config.Section.Path)
}

type configuration struct {
    Enabled bool
    Path    string
}

func main_read_json() {
    // 打开文件
    file, err := os.Open("conf.json")
    if err != nil {
        fmt.Println(err)
    }
    // 关闭文件
    defer file.Close()

    // NewDecoder创建一个从file读取并解码json对象的*Decoder,
    // 解码器有自己的缓冲,并可能超前读取部分json数据
    decoder := json.NewDecoder(file)
    conf := configuration{}
    // Decode从输入流读取一个json解码值并保存在v指向的值里
    err2 := decoder.Decode(&conf)
    if err2 != nil {
        fmt.Println("Error", err)
    }
    fmt.Println("Path: " + conf.Path)

}

type conf struct {
    Enabled bool   `yaml:"enabled"`
    Path    string `yaml: "path"`
}

func (c *conf) getConf() *conf {
    yamlFile, err := ioutil.ReadFile("conf.yaml")
    if err != nil {
        fmt.Println("yamlFile.Get err #%v", err)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }
    return c
}

func main() {
    var c conf
    c.getConf()
    fmt.Println("Path: " + c.Path)
}
