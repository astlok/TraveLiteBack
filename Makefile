run:
	docker build -t travelite:image .
	docker run -d -p 8080:8080 travelite:image
del:
	docker rm -f $(shell docker ps -aq)
del_all_images:
	docker rmi -f $(shell docker images -a -q)