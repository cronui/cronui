dev-fe:
	cd ui && npm run dev

dev-be: ent
	air

ent:
	go generate ./ent


.PHONY: ent dev-be dev-fe
