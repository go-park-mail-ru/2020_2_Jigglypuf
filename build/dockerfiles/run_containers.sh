sudo docker build -t main_image_docker_service -f main.Dockerfile ../..
sudo docker tag main_image_docker_service felixgreen/cinema_interfaceservice_image
#sudo docker push felixgreen/cinema_interfaceservice_image

sudo docker build -t auth_service -f auth.Dockerfile ../..
sudo docker tag auth_service felixgreen/cinema-auth-service
#sudo docker push felixgreen/cinema-auth-service

sudo docker build -t profile_service -f profile.Dockerfile ../..
sudo docker tag profile_service felixgreen/cinema-profile-service
#sudo docker push felixgreen/cinema-profile-service