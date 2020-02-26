# Sobre o Projeto

Esse projeto é um teste de Lincoln Coutinho para candidatura para vaga de Engenheiro de Software na LuizaLabs.

# Iniciando o projeto
Para criar todo ambiente desse projeto basta rodar o comando:
	```bash
	 docker-compose up -d --build && docker logs -f app
	 ```
OBS: Na primeira vez que esse processo é feito demora aproximadamente 5 minutos (Pois será necessário instalar as dependências do projeto), porém nas proximas vezes será bem mais rápido (Em torno de 2 minutos), quando tudo estiver operante a execução desse comando irá mostrar a mensagem ```[GIN-debug] Listening and serving HTTP on :8000``` na tela.


# Rotas
	Rota de autenticação e autorização
	Method: POST
	URL: localhost:8000/refresh
	Payload da requisição:
	```json
	 {
         "username": "admin",
         "password": "luizalabs"
     }
	 ```
    Retorno esperado:
    ```json
    {
        "success": true,
        "data": {
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODI2Nzc3MDQsInJlc291cmNlIjoicHJvZHVjdCJ9.gnesT44uozO78F4JlPMkFQ3KIA8xxO5L3qBjWNUrYwU"
        },
        "total": 0,
        "error": false
    }
     ```
	Headers:
    Content-Type:application/json
	Resource:product
	OBS: No header Resource é necessário espeficicar qual tipo de recurso que você gostaria de ter acesso, pois o token que irá ser gerado só terá acesso ao recurso solicitado, por exemplo:
	Os valores possíveis desse header são:
	product
	customer
	product,customer
	Caso o recurso solicitado seja apenas "product", o token gerado a partir desse recurso não conseguirá fazer nenhum tipo de operação de escrita no recurso "customer", vice-e-versa.
	
	Rota para listar os produtos
	Method: GET
	URL: localhost:8000/product?page={page}
	Retorno esperado:
	```json
	{
        "success": true,
        "data": [
            {
                "id": "5e55b454c8de4f208e95cedc",
                "price": 110,
                "image": "http://teste.com",
                "brand": "Brand",
                "title": "Teste Produto1",
                "reviewScore": 1
            },
            {
                "id": "5e55b455c8de4f208e95cede",
                "price": 110,
                "image": "http://teste.com",
                "brand": "Brand",
                "title": "Teste Produto1",
                "reviewScore": 1
            }],
              "total": 10,
              "error": false
          }
	 ```
	Headers:
    Content-Type:application/json
    
	 A paginação traz 10 resultados por página.
	 
	Rota para visualizar um produto
	Method: GET
	URL: http://localhost:8000/product/{id}
	Retorno esperado:
	```json
    {
        "success": true,
        "data": {
            "id": "5e55b455c8de4f208e95cede",
            "price": 110,
            "image": "http://teste.com",
            "brand": "Brand",
            "title": "Teste Produto1",
            "reviewScore": 1
        },
        "total": 1,
        "error": false
    }
	 ```
	Headers:
    Content-Type:application/json
    
    	
	Rota para criar produtos
	Method: POST
	URL: localhost:8000/product
    Payload da requisição:
    ```json
     {
         "price": 110.0,
         "image": "http://teste.com",
         "brand": "Brand",
         "title": "Teste Produto1",
         "reviewScore": 1
     }
     ```
    Todos os itens estão configurados como obrigatório.
	Retorno esperado:
	```json
	{
        "success": true,
        "data": {
            "price": 110,
            "image": "http://teste.com",
            "brand": "Brand",
            "title": "Teste Produto1",
            "reviewScore": 1
        },
        "total": 1,
        "error": false
    }
	 ```
	Headers:
    Content-Type:application/json	
    Authorization:{token gerado com recurso product ou product,customer}
     
	Rota para editar um produto
	Method: PUT
	URL: http://localhost:8000/product/{id}
    Payload da requisição:
    ```json
     {
         "price": 110.0,
         "image": "http://teste.com",
         "brand": "Brand",
         "title": "Teste Produto2",
         "reviewScore": 1
     }
     ```
	Retorno esperado:
	```json
    {
        "success": true,
        "data": {
            "id": "5e55b455c8de4f208e95cede",
            "price": 110,
            "image": "http://teste.com",
            "brand": "Brand",
            "title": "Teste Produto2",
            "reviewScore": 1
        },
        "total": 1,
        "error": false
    }
	 ```
	Headers:
    Content-Type:application/json	
    Authorization:{token gerado com recurso product ou product,customer}

	Rota para excluir um produto
	Method: DELETE
	URL: http://localhost:8000/product/{id}
	Retorno esperado:
	```json
    {
        "success": true,
        "data": "Success",
        "total": 1,
        "error": false
    }
	 ```
	Headers:
    Content-Type:application/json	
    Authorization:{token gerado com recurso product ou product,customer}

	Rota para listar clientes
	Method: GET
	URL: localhost:8000/customer
	Retorno esperado:
	```json
	{
        "success": true,
        "data": [
            {
                "id": "5e55c2f0c8de4f208e95d765",
                "name": "Teste",
                "email": "teste@teste.com"
            },
            {
                "id": "5e55c2f0c8de4f208e95d745",
                "name": "Teste 1",
                "email": "teste1@teste1.com"
            }
        ],
        "total": 1,
        "error": false
    }
	 ```
	Headers:
    Content-Type:application/json
	 
	Rota para visualizar um cliente
	Method: GET
	URL: http://localhost:8000/customer/{id}
	Retorno esperado:
	```json
    {
        "success": true,
        "data": {
            "id": "5e55c2f0c8de4f208e95d765",
            "name": "Teste",
            "email": "teste@teste.com"
        },
        "total": 1,
        "error": false
    }
	 ```
	Headers:
    Content-Type:application/json
    
    	
	Rota para criar clientes
	Method: POST
	URL: localhost:8000/customer
    Payload da requisição:
    ```json
     {
         "name": "Teste",
         "email": "teste@teste.com"
     }
     ```
    Todos os itens estão configurados como obrigatório.
	Retorno esperado:
	```json
	{
        "success": true,
        "data": {
            "name": "Teste",
            "email": "teste@teste.com"
        },
        "total": 1,
        "error": false
    }
	 ```
	Headers:
    Content-Type:application/json	
    Authorization:{token gerado com recurso customer ou product,customer}
     
	Rota para editar um cliente
	Method: PUT
	URL: http://localhost:8000/customer/{id}
    Payload da requisição:
    ```json
     {
         "name": "Teste 1",
         "email": "teste1@teste.com"
     }
     ```
	Retorno esperado:
	```json
    {
        "success": true,
        "data": {
            "name": "Teste 1",
            "email": "teste1@teste.com"
        },
        "total": 1,
        "error": false
    }
	 ```
	Headers:
    Content-Type:application/json	
    Authorization:{token gerado com recurso customer ou product,customer}

	Rota para excluir um cliente
	Method: DELETE
	URL: http://localhost:8000/customer/{id}
	Retorno esperado:
	```json
    {
        "success": true,
        "data": "Success",
        "total": 1,
        "error": false
    }
	 ```
	Headers:
    Content-Type:application/json	
    Authorization:{token gerado com recurso customer ou product,customer}

	Rota para adicionar um produto favorita para um cliente
	Method: POST
	URL: localhost:8000/customer/favorites/{id}
    Payload da requisição:
    ```json
     [
         "5e55aed70c43745f3d9c0be0"
     ]
     ```
    No payload são enviados os ids dos produtos.
	Retorno esperado:
	```json
	{
        "success": true,
        "data": "1 favorite products added for this customer",
        "total": 1,
        "error": false
    }
	 ```
	Headers:
    Content-Type:application/json	
    Authorization:{token gerado com recurso customer ou product,customer}
        
## Rodando testes da aplicação
	Essa aplicação tem testes de BDD, e testes de integração, para executar os testes basta rodar o comando:
	```bash
	 docker exec -it test go test github.com/lcoutinho/luizalabs-client-api github.com/lcoutinho/luizalabs-client-api/services -coverprofile cover.out
	 ```
	Toda lógica de negócio está dentro da pasta github.com/lcoutinho/luizalabs-client-api/services, sendo que todos cenários propostos no desafio estão cobertos por testes
## Tecnologias utilizadas
________________________________
  	- Golang
	- Gin-gonic
	- Go Json Schema Validator
	- JWT Autenticator
	- MongoDB
	- TLS 1.2 CA
________________________________

## Estratégias de autenticação e autorização
	Essa API conta com duas estratégias de autenticação e autorização: 
	    TLS_CERTIFICATE_JWT : Os tokens são gerados também com o processo de certificados dinâmicos de TLS, ou seja, a cada vez que a API irá gerar um novo token, um novo certificado de TLS também é gerado e criptografado nesse token JWT como secret.
	    SIMPLE_JWT : Processo de utilização simples no JWT, com a secret fixa na aplicação.
    No projeto já vem com o padrão utilizando TLS_CERTIFICATE_JWT, para modificar basta alterar o arquivo ./config/config.go, mudar o valor da constante AUTH_STRATEGY
