# go-websocket

## Resumo:
Microsserviço de API para captura de acessos em um determinado cliente. O backend do projeto é escrito em Go e o frontend é escrito em JavaScript puro.

## Instalação e Execução:
- Criação do arquivo .env no diretório backend com as seguintes variáveis: `DB_DATABASE`, `DB_USER`, `DB_PASSWORD`, `DB_HOST` e `DB_PORT`

- Instalação das dependências pelo arquivo Makefile na raiz do projeto:
`make setup`

- Execução do backend:
`make run`

- Build do front e backend:
`make build`

- Execução do binário do backend:
`make run-build`

- Execução do frontend para teste do script de captura de acessos:
`cd frontend && npm run start`

- Execução do handler para teste do intermediário entre o frontend e o backend
`cd handler && npm run start`

## Estrutura do Projeto:
O projeto é dividido em três partes fundamentais: Companion (backend), Comrade (frontend) e um handler intermediário.

```bash
|-- backend/ (Go)
|-- frontend/ (JS)
|-- handler/ (JS)
|-- Makefile
```

Cada uma das partes será descrita a seguir, iniciando pela captura de acessos no frontend, passando pelo handler até a finalização no backend.

## Frontend:
O frontend do projeto se resume a um script customizado (Comrade) responsável pelo envio de três parâmetros para o handler: window (js window), document (js document) e ecommerce_hash (string).

O funcionamento do script se baseia na geração e adição de uma segunda tag de script para o carregamento do handler do projeto na página do cliente e a criação de um hidden input com o valor definido para a hash única do ecommerce em questão.

Com o carregamento do handler, todo o processo de captura de dados acessos é feito por ele, finalizando assim a responsabilidade do frontend.

### Estrutura do Frontend:

```bash
|-- node_modules
    |-- <Múltiplos diretórios>
|-- index.html
|-- package.json
|-- page2.html
|-- script.js
```

### Diretórios e Arquivos:
- **node_modules:** Diretório contendo arquivos das dependências utilizadas pelo projeto;
- **index.html:** Exemplo de página com o Comrade (script de captura de acesso) inserido;
- **package.json:** Arquivo de configuração utilizado para estipular e configurar dependências do seu projeto e scripts automatizados;
- **page2.html:** Segunda página, sem o Comrade inserido, utilizada para teste de captura de referrer dos acessos;
- **script.js:** O script Comrade em si com uma descrição linha-a-linha de seu funcionamento.

### O Script do Comrade:

```javascript
(function(c,o,m,r,a,d,e){
r=o.createElement('script');r.async=1;r.src='https://companion-example.com/cmp'+c.location.search;
a=o.getElementsByTagName('script')[0];a.parentNode.insertBefore(r, a);
d=o.createElement('input');d.type='hidden';d.id='eh';d.value=m;
e=o.getElementsByTagName('body')[0];e.appendChild(d);
})(window,document,'<HASH>');
```

## Handler:
O handler é um arquivo intermediário entre o frontend e o backend do projeto.

Servido pelo backend, ele é responsável pela separação dos dados úteis do projeto, gerenciamento de sessão de acesso e separação/validação dos parâmetros críticos vindos da URL.

Com todos os dados prontos, estes são mandados para o backend no formato JSON através de um websocket.

#### Estrutura do Handler:

```
|-- dist/
    |-- cmp.js
    |-- cmp.js.map
|-- node_modules
    |-- <Múltiplos diretórios>
|-- index.html
|-- index.js
|-- package.json
```

### Diretórios e Arquivos:
- **dist**: Diretório com o arquivo index.js minificado e otimizado. Contém os arquivos "cmp.js" e "cmp.js.map", utilizados pelo Companion (backend);
- **node_modules:** Diretório contendo arquivos das dependências utilizadas pelo projeto;
- **index.html:** Página com acesso direto ao Companion (backend) para facilitar a testagem do handler sem depender do Comrade (frontend);
- **index.js:** O script do handler em si. Este é o arquivo que, quando compilado, se torna os arquivos do diretório "dist";
- **package.json:** Arquivo de configuração utilizado para estipular e configurar dependências do seu projeto e scripts automatizados;

## Backend
O Companion (backend) é a parte responsável por todo o tratamento dos dados e interação com o banco de dados.

Há um websocket responsável pelo recebimento dos dados enviados pelo handler.

```bash
|-- sql
|-- src
    |-- config
    |-- controllers
    |-- db
    |-- models
    |-- repositories
    |-- static
    |-- ws
|-- .env
|-- go-webserver
|-- go.mod
|-- main.go
|-- regexes.yaml
```

### Diretórios e Arquivos:

- **sql**: Diretório com a estrutura de banco do projeto;
- **config:** Diretório com variáveis de configuração do projeto;
- **controllers:** Diretório de controladores do projeto com métodos para inserção dos acessos no banco;
- **db:** Diretório de configuração do banco de dados;
- **models:** Diretório de modelos de estruturas utilizados no projeto;
- **repositories:** Diretório onde se localizam os arquivos responsáveis pelas interações com o banco de dados. Foi criado como
módulo separado dos controles para todas as queries pudessem ser tratadas em um único local do projeto;
- **static:** Módulo para servir o arquivo do handler;
- **ws:** Módulo de implementação do websocket;
- **.env:** Arquivo com variáveis de ambiente. Não é versionado;
- **go-webserver:** Binário executável do projeto;
- **go.mod:** Arquivo com os módulos utilizados no projeto;
- **main.go:** Arquivo principal do projeto. Ponto de partida e gerenciamento de rotas;
- **regexes.yaml:** Arquivo necessário para a utilização do parser de user agents. O parser é responsável pela obtenção de dados de dispositivo, sistema operacional e navegador utilizado para o acesso, tendo como base o user agent enviado pelo handler.