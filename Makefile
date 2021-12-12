.PHONY: gen migrate

gen:
ifdef NAME
	migrate create -ext sql -dir migrations -seq $(NAME)
else
	echo "Set NAME for migration"
endif

migrate:
	set -o allexport; source .env; set +o allexport; \
	export $(grep -v '^#' .env | xargs -d '\n'); \
	migrate -path migrations -database $$DATABASE_URL up

migrate_down:
	set -o allexport; source .env; set +o allexport; \
	export $(grep -v '^#' .env | xargs -d '\n'); \
	migrate -path migrations -database $$DATABASE_URL down 1






