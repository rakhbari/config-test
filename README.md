# config-test
Demonstrating different struct tags used in app config structs.

## Build
```
go build .
```

## Run
```
./config-test
```
Output:
```
2022/09/27 00:22:23 Reading App config from config.yaml ...
2022/09/27 00:22:23 Reading App config from config.yaml ...
2022/09/27 00:22:23 Reading App config from config.yaml ...
2022/09/27 00:22:23 App config 1: &{ServerConfig:{Host:0.0.0.0 Port:8080 TlsPort:8443 DummyPort:1111}}
2022/09/27 00:22:23 App config 2: &{ServerConfig:{Host: Port:0 TlsPort:0 DummyPort:0}}
2022/09/27 00:22:23 App config 3: &{ServerConfig:{Host: Port:0 TlsPort:0 DummyPort:0}}
```
You'll notice that only the struct `App config 1` is populated with the values in `config.yaml`. The other AppConfig structs failed to load values, even though they have `yaml` struct tags.
