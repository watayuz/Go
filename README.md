# Goで遊ぶ

## WebAPI.goの使い方
`$ go run WebAPi.go`  
別Terminalで  
`$ curl http://localhost:8080/api/ -X POST -d '{"name":"john"}'`  
とかやるといい感じになる  

##  about templates/template.go
Irisを使った  
Iris強い  
描画されるHTMLは `templates/` の直下を参照するようになっている  

## Standard
外部ライブラリを使わないもの

## Go-Json-Rest
go-json-restライブラリを使ったもの

## chat
使うときには  
同じディレクトリにconfigを作成、その直下のconfig.goに以下を記述する  
```Go:config.go
package config  
import()  
var Key string = "api_key"
```

を追加する  