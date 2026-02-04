# Treinamento Go — Funções, Ponteiros, Recursão, Closure, Defer, Panic/Recover

Pequenos exemplos para estudar conceitos básicos de Go.

## Estrutura
- ponteiros/ponteiros.go — exemplo de passagem por referência (pointer)
- recursao/recursao.go — exemplo de fatorial recursivo
- funcao.go — exemplos de funções e uso geral (main)
- closure/closure.go — exemplo de closure (função que captura estado)
- deferpkg/deferpkg.go — exemplos de defer e ordem de execução
- panicpkg/panicpkg.go — exemplo de panic e recover
- go.mod, .gitignore

## Resumo rápido
- Ponteiros: usar *T e & para referenciar e desreferenciar; modificar valor via função.
- Recursão: função que chama a si mesma (ex.: fatorial).
- Função: declaração, parâmetros, valores de retorno e organização de código.
- Closure: função retornando outra função que captura variáveis do escopo externo.
- Defer: execução adiada até o fim da função; LIFO quando há múltiplos defer.
- Panic/Recover: gerar pânico com panic e capturar com recover em um defer.

## Como executar
No terminal (PowerShell ou cmd) dentro do diretório do módulo:
- Executar exemplo de ponteiros:
  go run ./ponteiros
- Executar exemplo de recursão:
  go run ./recursao
- Executar main geral:
  go run funcao.go
- Rodar um pacote específico com main:
  go run ./closure
  go run ./deferpkg
  go run ./panicpkg

Ou para executar um arquivo .go isolado:
  go run caminho\para\arquivo.go

## Notas
- Padronize nomes de arquivos em minúsculas para evitar problemas no Go.
- Use `go build` para compilar e `go test` para testes quando presentes.
