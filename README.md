# Estudando Go Land

### Comandos

#### Adicionar um novo comando

> go get golang.org/x/tools/cmd/godoc

#### Habilitando documentação offline

> godoc -http=:6060

#### Verificar variáveis do escopo Go

> go env

#### Validar código

> go vet FILE.go

#### Compilar arquivo

> go build FILE.go

#### Compilar e executar

> go run FILE.go


### Concorrência vs Paralelismo

Go é uma liguagem concorrente

> Paralelismo - executa código simultaneamente em processadores físicos diferentes.
>> Paralelismo exige muito mais do SO e do hardware.

> Concorrência - intercala (administra) vários processos ao mesmo tempo e isso pode ocorrer em um úvico processador físico.
>> Concorrência viabiliza paralelismo.
>
>
>> É possível que a concorrência seja melhor que o paralelismo!
>
>> Concorrência - é a forma de estruturar o seu programa.
>> É a composição de processos(tipicamente funções) que executam de forma independente. 

