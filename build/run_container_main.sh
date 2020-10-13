docker build -t teamproject-docker-image -f dockerfiles/main.Dockerfile ..
docker run -dp 8080:8080 teamproject-docker-image 
