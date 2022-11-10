# O que é o ip-cli?
**ip-cli é um cli que retorna o ip público de um site.**

## Como usar?

1. Para executar o programa é necessário ter o go [instalado](https://go.dev/doc/install) na sua máquina.
2. Após ter instalado o go, dê um git clone no repositório.
3. Entre na pasta ip-cli e rode o comando `go run main.go` que irá exibir as informações.

Para executar o código abaixo basta passar o comando ip seguido da flag --host e o url do site.

```go
go run main.go ip --host url-do-site
```
- Exemplo real:

```go
go run main.go ip --host google.com
```

- Exibe:

```bash
go run main.go ip --host google.com
Iniciando a aplicação...
142.251.129.206
2800:3f0:4001:831::200e
```

**Note que temos o comando ip que retorna por padrão o ip público do google.com, a flag --host serve para passar o url do site que você deseja obter o ip público.**