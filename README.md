Este projeto está sendo desenvolvido para aperfeiçoar minhas habilidades usando [Go](https://go.dev/), portanto, este README serve apenas como uma documentação pessoal afins de estudo.

### Ferramentas e bibliotecas utilizadas

#### [Go Api Gen](https://github.com/discord-gophers/goapi-gen)

O `goapi-gen` serve para gerarmos um código boilerplate para serviços baseados na OpenAPI, como o Swagger. Com ele, geramos rapidamente modelos e estruturas baseado no scheme.json/yaml
da API documentada. Inclusive, `goapi-gen` utiliza o framework [Chi](https://github.com/go-chi/chi) para fazer o gerenciamento de rotas e chamadas HTTP.

#### [Tern](https://github.com/jackc/tern)

Tern é uma ferramenta para gerar migrations para nosso banco Postgres de modo fácil e rápido. Com ele nós criamos um arquivo chamado `tern.conf`, onde atribuímos valores importantes como
dados para conexão com banco de dados. Depois, executamos um comando para gerar uma migration, que irá gerar um arquivo `.sql` e escreveremos nossas queries. 

#### [SQLC](https://github.com/jackc/tern)

O `sqlc` funciona muito bem com o `tern`, ele serve para gerarmos código SQL com type-safety. Seu fluxo se baseia no seguinte:

1. Codamos nossa query em SQL;
2. Rodamos `sqlc` para gerar código com type-safety paras essas queries;
3. Escrevemos código que utilize estes arquivos gerados, no nosso caso, usamos estas queries com o `tern` para rodar as migrations no nosso banco.