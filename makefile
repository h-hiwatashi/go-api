.PHONY: intodb
intodb:
	docker container exec -it go_api bash

.PHONY: intogo
intogo:
	docker container exec -it go_clean_arch_mysql bash

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
	docker compose up -d