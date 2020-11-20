sudo docker build -t dev_image_docker_service -f dev.Dockerfile ../..
sudo docker tag dev_image_docker_service felixgreen/dev_service_image
sudo docker push felixgreen/dev_service_image
