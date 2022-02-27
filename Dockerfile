# 1. using golang: alpine because golang tools 
#are not available in plain alpine docker images
# 2. Copy all the files and add them into 
#the app directory into the image

FROM golang:alpine as builder
COPY . /app   
WORKDIR /var/app/
CMD /var/app/
