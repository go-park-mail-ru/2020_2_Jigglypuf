FROM ubuntu:20.04 AS release
RUN apt-get -y update && apt-get install -y locales gnupg2
RUN locale-gen en_US.UTF-8
RUN update-locale LANG=en_US.UTF-8
ENV PGVER 12
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update -y && apt-get install -y postgresql postgresql-contrib

USER postgres
RUN    /etc/init.d/postgresql start &&\
    psql --command "CREATE USER cinema_backend_interface WITH SUPERUSER PASSWORD 'some_password';" &&\
    createdb -E UTF8 cinema_backend_interface &&\
    /etc/init.d/postgresql stop

USER root
CMD service postgresql start