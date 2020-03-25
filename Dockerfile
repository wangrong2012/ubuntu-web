FROM docker-registry.saicstack.com/base/ubuntu-web:14.04
#RUN apt-get update -y && \ 
#	apt-get install software-properties-common -y && \ 
#	add-apt-repository ppa:longsleep/golang-backports -y && \
#	apt-get install golang-go curl inetutils-ping net-tools -y
	
COPY ./gin-helloworld /usr/local/bin
RUN chmod +x /usr/local/bin/gin-helloworld
ENTRYPOINT [ "/usr/local/bin/gin-helloworld" ]
