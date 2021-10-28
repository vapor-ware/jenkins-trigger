FROM vaporio/golang:1.16 as build
ADD . /build
WORKDIR /build
RUN go build main.go && ls -al dist

FROM vaporio/foundation
COPY --from=0 /build/dist/jenkins-trigger_linux_amd64/jenkins-trigger /usr/local/bin/
WORKDIR /
ENV JENKINS_API_USER=""
ENV JENKINS_API_TOKEN=""
CMD ["/usr/local/bin/jenkins-trigger"]
