FROM vaporio/foundation
# Depends on the goreleaser build context
ADD ./jenkins-trigger /usr/local/bin/jenkins-trigger
WORKDIR /
ENV JENKINS_API_USER=""
ENV JENKINS_API_TOKEN=""
CMD ["/usr/local/bin/jenkins-trigger"]
