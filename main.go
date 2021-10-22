package main

import (
	"github.com/bndr/gojenkins"
	"github.com/urfave/cli/v2"
	"context"
	"fmt"
	"os"
	"log"
)


func main() {
	app := &cli.App{
		Name: "trigger-build",
		Usage: "Interface with build.vio.sh",
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name: "user",
				Aliases: []string{"u"},
				Required: true,
				EnvVars: []string{"JENKINS_API_USER"},
			},
			&cli.StringFlag{
				Name: "token",
				Aliases: []string{"t"},
				Required: true,
				EnvVars: []string{"JENKINS_API_TOKEN"},
			},
			&cli.StringFlag{
				Name: "server",
				Aliases: []string{"s"},
				EnvVars: []string{"JENKINS_API_HOST"},
				Value: "https://build.vio.sh",
			},
		},
		Commands: []*cli.Command{
			{
				Name: "build",
				Aliases: []string{"b"},
				Usage: "Run a build in Jenkins",
				Action: func(c *cli.Context) error {
					jobName := c.Args().Get(0)
					if jobName == "" {
						panic("Missing positional argument job, eg: vapor-cloud-ops/job/ansible-microk8s/")
					}
					enqueue(c.String("user"), c.String("token") , c.String("server"), c.Args().Get(0))
					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}


// Enqueue a job in the configured Jenkins server
// param: user - the userid for jenkins. At vapor this is your Github User ID
// param: token - an API token minted from jenkins for access
// param: server - the https url to the jenkins server
// param jobname - A formatted job name per the organization structure - eg: vapor-cloud-ops/job/ansible-microk8s
func enqueue(user, token, server, jobname string) {
	ctx := context.Background()

	// Init the client
	jenkins := gojenkins.CreateJenkins(nil, server, user, token)
	_, err := jenkins.Init(ctx)

	// Error checking on the client init
	if err != nil {
		panic(err)
	}

	// Format of job listing is org/job/jobname. The library does a shit job of documenting itself
	job, err := jenkins.GetJob(ctx, jobname)

	// Error checking on job query
	if err != nil {
		fmt.Println("Error getting job:")
		panic(err)
	}

	// Enqueue the job
	_, err = job.InvokeSimple(ctx, nil)
	if err != nil {
		fmt.Println("Error during InvokeSimple:")
		panic(err)
	}
	// The remaining functions for job interaction in the library dont work
	// because the author didn't consider organizations in the hierarchy of api calls
	log.Printf("Queued up job %s", job.Raw.URL)

}

