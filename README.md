# go-fipe
Serviço desenvolvido para consumir dados da tabela FIPE no site https://veiculos.fipe.org.br

# Instruções

- Verifique se a versão mínima é ``1.17``, não foi testado em versões anteriores mas provável que funcione sem problemas.
- Se quiser altere a porta em ``.env`` a padrão é ``8080``
- Execute o comando abaixo na raiz do projeto:
```shell
docker-compose up
```
ou
```shell
docker-compose up -d
```

# Postman

Existe um arquivo na raiz do projeto chamado ``go-fipe.postman-collection.json``, carregue ele no seu postman para poder ver os endpoints disponíveis.