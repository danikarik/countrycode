build:
	@go build -o bin/countrycode

run:
	@bin/countrycode --output=./out.txt
