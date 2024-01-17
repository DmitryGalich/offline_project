## Project

* whole project <https://www.youtube.com/playlist?list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE>

## Design

* <https://www.youtube.com/watch?v=mr3pywHOz7I&t=2780s>

## Docker

* docker and docker-compose
  * <https://www.youtube.com/watch?v=Q0OwEKtncPc>
  * <https://www.youtube.com/watch?v=ZB8JBWriDVo>
* open in vscode <https://www.youtube.com/watch?v=mi8kpAgHYFo&t=38s>
* attach to running container <https://www.youtube.com/watch?v=8gUtN5j4QnY>

* migrations and docker-compose <https://medium.com/pengenpaham/postgres-database-migration-using-golang-migrate-docker-compose-and-makefile-159ef50670cf>

## PostgreSQL

* basics <https://www.youtube.com/watch?v=i5-1HNf3W_Y&list=LL&index=15&t=2454s>

* open docker in terminal
```docker exec -it offline_project_db bash```

* go to postgres mode
```psql -U postgres```

* list of databases
```\l```

* connect to database
```\c offline_project_db```

## Flutter

* Docker
  * <https://blog.codemagic.io/how-to-dockerize-flutter-apps/>
  * <https://github.com/sbis04/flutter_docker>

## Authorization

* flutter and go authorization <https://www.youtube.com/watch?v=q7WvkKio7ek>

# GO

* Project structure <https://github.com/golang-standards/project-layout/blob/master/README_ru.md>
* Awesome Go <https://github.com/avelino/awesome-go>
* JWT POSTGRES GO <https://www.youtube.com/watch?v=ma7rUS_vW9M>

for vscode
  gotests
  gomodifytags
  impl
  goplay
  dlv
  staticcheck
  gopls
  goimports

* For protobuf generating
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
protoc -I proto/ proto/sso/sso.proto --go_out=gen/ --go_opt=paths=source_relative --go-grpc_out=gen/ --go-grpc_opt=paths=source_relative

* For env configs parsing
go get -u github.com/ilyakaznacheev/cleanenv

* For grpc
go get -u google.golang.org/grpc

* For crypto
go get golang.org/x/crypto

* BOOKS
<https://github.com/diptomondal007/GoLangBooks.git>

## Rust

* oauth2 for rust
  * <https://docs.rs/oauth2/latest/oauth2/>
  * <https://www.youtube.com/watch?v=S8Usi3zsLIs>

## Architecture

* Kong
* GraphQL

* Patterns
<https://www.youtube.com/watch?v=_RCiOo4Dv8w>

* Avito
<https://www.youtube.com/watch?v=eI1QQUrFUZI&t=1678s>

* Uber architecture
<https://www.youtube.com/watch?v=R_agd5qZ26Y>

* NGINX microservices
<https://www.grizzlypeaksoftware.com/articles?id=4vlfDp2ZanpAh3bSuUzPeZ>

<https://habr.com/ru/articles/779276/>

* NGINX and keycloack
<https://kevalnagda.github.io/configure-nginx-and-keycloak-to-enable-sso-for-proxied-applications>
<https://github.com/flavienbwk/nginx-keycloak?tab=readme-ov-file><<https://yasithev.medium.com/>

* NGINX Big video
<https://www.youtube.com/watch?v=NwijBVfiK_o>

* NGINX nginx-jwt-module
<https://github.com/max-lt/nginx-jwt-module>

* DOCKERFILE
<https://habr.com/ru/companies/ruvds/articles/439980/>

* NGINX and 0Auth 2.0
<nginx-as-an-api-gateway-with-oauth-2-0-authorization-on-aws-4d7dbfe2a85b>

<https://habr.com/ru/articles/650911/>

* Лекция об авторизации и авторизации
<https://www.youtube.com/watch?v=dLt0ktv6Ca0>

* Keycloack and kong

<https://ncarlier.gitbooks.io/oss-api-management/content/howto-kong_with_keycloak.html>
<https://www.jerney.io/secure-apis-kong-keycloak-1/>
<https://callistaenterprise.se/blogg/teknik/2023/04/20/kong-api-gateway-part1/>

* Zero trust
<https://medium.com/@krishnarajanarunachalam/securing-microservices-in-hyperscalers-a-zero-trust-approach-2f5a43aeb0fa>
