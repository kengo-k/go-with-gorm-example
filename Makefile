build-image:
	docker build -t sample-go-with-gorm:v1.0.0 .

connect:
	docker run -it --rm --name sample-go-with-gorm -v "${PWD}/:/app" sample-go-with-gorm:v1.0.0