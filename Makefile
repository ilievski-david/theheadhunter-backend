mock-handler:
	@mockgen -destination=./mocks/handlers/handlers.go -package=handlers  github.com/ilievski-david/theheadhunter-backend/handlers Handler

mock-database:
	@mockgen -destination=./mocks/crud/crud.go -package=crud  github.com/ilievski-david/theheadhunter-backend/crud Database