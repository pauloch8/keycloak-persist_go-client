# Keycloak persistindo dados e cliente em Go.lang

Subir o container

`docker-compose up`

Adicionar o realm `demo`

Adicionar o client `app` com a Root URL *http://localhost:8081*, e o type access `confidential`. Copiar o segredo da aba Credentials e criar o arquivo .env com o segredo setado na variável `CLIENTSECRET`

Adicionar um usuario `myuser` com Email Verified `ON`, e setar o password na aba Credentials, com Temporary `Off`

## Arquivo .env

O arquivo deve conter as seguintes variáveis:

```
CLIENTSECRET=XXXXX
POSTGRESQL_USER=xxxxx
POSTGRESQL_PASS=xxxxx
POSTGRESQL_DB=xxxxx
```

## Host Name

O keycloak está rodando com o hostname `keycloak`, portanto, para funcionar precisa-se alterar o arquivo hosts no cliente
- `sudo nano /etc/hosts` no Linux
- `code C:\Windows\System32\drivers\etc\hosts` no Windows

Incluir o host:

```
127.0.1.1       keycloak
```