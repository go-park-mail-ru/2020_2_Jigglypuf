docker build -t teamproject-docker-image -f dockerfiles/main.Dockerfile ..
docker run -v ~/projects/go-docker:/app -dp 8080:8080 teamproject-docker-image 
