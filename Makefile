.PHONY: f-test

f-test:
	docker-compose run --rm frontend npm run test
