sudo docker build -t main_image_docke_service -f dockerfiles/main.Dockerfile ..
sudo docker tag main_image_docke_service felixgreen/cinema_interfaceservice_image
sudo docker push felixgreen/cinema_interfaceservice_image
