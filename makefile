.PHONY: intodb
intodb:
	docker container exec -it postgres bash

.PHONY: intogo
intogo:
	docker container exec -it todo_app-web-1 bash

.PHONY: upd
upd:
	docker compose up -d

#Dockerfileの更新を反映させる
.PHONY: dockerfile
updb:
	docker compose up -d --build

#docker-compose.ymlの更新を反映させる
.PHONY: composeyml
upd:
	docker-compose up -d