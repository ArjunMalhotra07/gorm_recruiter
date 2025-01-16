This projects requires :-->

1. Running the email service which exposes port :50051

2. running a container mysql instance using docker command:
```
docker-compose -f <docker_fileName>.yml up
```
It maps the containers 3306 port to our Host 3307 port


3. Then, switch on prometheus server by running the following command after going to the prometheus folder in the Downloads directory:
```
./prometheus --config.file=prometheus.yml
```
Visit prometheus dashboard at localhhost:9090

4. finally run this project by 
```
go run main.go
```