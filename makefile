.PHONY: intogo
intogo:
	docker container exec -it go_api bash

.PHONY: intodb
intodb:
	docker container exec -it go_api_mysql bash

.PHONY: upd
upd:
	docker compose up -d

#Dockerfileの更新を反映させる
.PHONY: updb
updb:
	docker compose up -d --build

#docker-compose.ymlの更新を反映させる
# .PHONY: composeyml
# upd:
# 	docker compose up -d


.PHONY: gofmt
gofmt:
	gofmt -l -s -w .


# curl http://localhost:8080/hello -X POST -w '%{http_code}\n'