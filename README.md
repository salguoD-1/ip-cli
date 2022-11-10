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

**Para saber o nome do servidor em que o site está hospedado, basta usar o comando servidores seguido da flag --host e o url do site.**

- Exemplo

```
go run main.go servidores --host url-do-site
```

- Exemplo real:

```
go run main.go servidores --host google.com
```

- Exibe:

```
go run main.go servidores --host google.com
Iniciando a aplicação...
ns4.google.com.
ns2.google.com.
ns3.google.com.
ns1.google.com.
```

**Caso nenhum site seja passado, o comando servidores irá retornar o nome dos servidores do google.com por padrão.**

## Compilando o arquivo

**Para compilar o arquivo, basta executar `go build` dentro da pasta ip-cli que irá gerar um arquivo chamado linha-de-comando**

```
go build
```

**Para executar o arquivo compilado, digite:**

```
./linha-de-comando
```

Resultado:

```
./linha-de-comando
Iniciando a aplicação...
NAME:
   Aplicação de Linha de Comando - Busca IPs e Nomes de Servidores na Internet

USAGE:
   Aplicação de Linha de Comando [global options] command [command options] [arguments...]

COMMANDS:
   ip          Busca IPs de endereços na Internet
   servidores  Busca o nome dos servidores na internet
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```
