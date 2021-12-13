FROM docker.io/vaporio/foundation
# Depends on the goreleaser build context
ADD ./jenkins-trigger /usr/local/bin/jenkins-trigger
WORKDIR /
RUN apt-get update && apt-get install -y ca-certificates
ENV JENKINS_API_USER=""
ENV JENKINS_API_TOKEN=""
CMD ["/usr/local/bin/jenkins-trigger"]
