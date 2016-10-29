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