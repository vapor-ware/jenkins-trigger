run:
	JENKINS_API_USER=lazypower JENKINS_API_TOKEN=${JENKINS_TOKEN}	go run main.go build vapor-cloud-ops/job/ansible-microk8s/
