package main

import (
    "fmt"
    "log"
    "os"

    "github.com/AbsaOSS/env-binder/env"
    "gopkg.in/yaml.v3"
)

type AppConfig1 struct {
    ServerConfig struct {
        Host        string  `yaml:"host"`
        Port        int     `yaml:"port"`
        TlsPort     int     `yaml:"tlsPort"`
        DummyPort   int     `yaml:"dummyPort"`
    } `yaml:"server"`
}

type AppConfig2 struct {
    ServerConfig struct {
        Host        string  `yaml:"host" env:"MY_HOST"`
        Port        int     `yaml:"port" env:"MY_PORT"`
        TlsPort     int     `yaml:"tlsPort" env:"MY_TLS_PORT"`
        DummyPort   int     `yaml:"dummyPort" env:"MY_DUMMY_PORT"`
    } `yaml:"server"`
}

type AppConfig3 struct {
    ServerConfig struct {
        Host        string  `env:"MY_HOST" yaml:"host"`
        Port        int     `env:"MY_PORT" yaml:"port"`
        TlsPort     int     `env:"MY_TLS_PORT" yaml:"tlsPort"`
        DummyPort   int     `env:"MY_DUMMY_PORT" yaml:"dummyPort"`
    } `yaml:"server"`
}

func processError(err error) {
    fmt.Println(err)
    os.Exit(2)
}

func readFile(cfg interface{}) {
    configFile := "config.yaml"
    log.Printf("Reading App config from %s ...\n", configFile)
    f, err := os.Open(configFile)
    if err != nil {
        processError(err)
    }
    defer f.Close()

    decoder := yaml.NewDecoder(f)
    err = decoder.Decode(cfg)
    if err != nil {
        processError(err)
    }
}

func main() {
    cfg1 := &AppConfig1{}
    readFile(cfg1)

    cfg2 := &AppConfig2{}
    readFile(cfg2)

    cfg3 := &AppConfig3{}
    readFile(cfg3)

    if err := env.Bind(cfg1); err != nil {
        fmt.Println(err)
        return
    }
    log.Printf("App config 1: %+v\n", cfg1)

    if err := env.Bind(cfg2); err != nil {
        fmt.Println(err)
        return
    }
    log.Printf("App config 2: %+v\n", cfg2)
    
    if err := env.Bind(cfg3); err != nil {
        fmt.Println(err)
        return
    }
    log.Printf("App config 3: %+v\n", cfg3)
}