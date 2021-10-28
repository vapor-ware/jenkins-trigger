# Jenkins Trigger

A lightweight golang jenkins api client for use with vapor.io build services


## Installation

Download the binary from the Github Releases and place the bin in $PATH

## Usage

Required Params:
- Jenkins UserID (typically your github user id)
- Jenkins API Token (Minted from: https://jenkins/user/$USERID/configure)
- The server defaults to VaporIO build services

Required Arguments to `build`: the job name, including organization. This is to work around a defect in the jenkins api client library which does not include organization in internal functions. If we provide it to the build command, it will successfully resolve.

`jenkins-trigger -u $USERNAME -t $API_TOKEN -s https://some-jenkins-server build $api-job-route`


As an example using my username, a fake token, and the vapor-cloud-platform deployment job:

`jenkins-trigger -u lazypower -t 0xD34DB33f build vapor-cloud-ops/job/vapor-cloud-platform`

## Docker

jenkins-trigger is published, wrapped in a `vaporio/foundation` base image for ease of consumption. This makes the binary far more practical to be deployed in one-off cronjob fashion in k8s, or to keep your host operating system pristine.

```
docker run -e JENKINS_API_USER=example -e JENKINS_API_TOKEN=example vaporio/jenkins-trigger build vapor-cloud-ops/jobs/example-job
```

## Getting Help

Try contacting in the #devops slack channel in the VaporIO slack
