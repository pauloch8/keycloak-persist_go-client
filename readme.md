# Keycloak persistindo dados e cliente em Go.lang

Subir o container

`docker-compose up`

## Configuração da aplicação no Keycloak

### Manualmente

Adicionar o realm `demo`

Adicionar o client `app` com a Root URL *http://localhost:8081*, e o type access `confidential`. Copiar o segredo da aba Credentials e criar o arquivo .env com o segredo setado na variável `CLIENTSECRET`

Adicionar um usuario `myuser` com Email Verified `ON`, e setar o password na aba Credentials, com Temporary `Off`

### Importando JSON

Entre no [Admin Console](http://keycloak:8080/admin/) e clique em Import no menu lateral esquerdo

Selecione o arquivo realm-export.json

## Container app

O container app vai parar quando subir e não conseguir se conectar com o Keycloak. Portanto, enquanto o keycloak não tiver subido ele vai cair.

Quando o Keycloak estiver pronto, dê start no container do app novamente

## Arquivo .env

O arquivo deve conter as seguintes variáveis:

```
CLIENTSECRET=XXXXX
KEYCLOAK_HOST=xxxxx
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