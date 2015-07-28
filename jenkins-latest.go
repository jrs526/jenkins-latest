package main

import "github.com/jrs526/jenkinsrss"
import "net/http"
import "encoding/xml"
import "fmt"
import "flag"
import "time"
import "strings"

func main() {
	jobPtr := flag.String("job", "ci_build_master", "Comma seperated list of jobs to check the status of")
	hostPtr := flag.String("host", "jenkins-server:8080", "Jenkins host")
	userPtr := flag.String("user", "user", "Jenkins user")
	keyPtr := flag.String("api-key", "not_an_auth", "Jenkins api key")

	flag.Parse()

	client := &http.Client{}

	for true {
		for _, job := range strings.Split(*jobPtr, ",") {
			get(client, hostPtr, userPtr, keyPtr, &job)
			time.Sleep(3 * time.Second)
		}
	}

}
func get(client *http.Client, hostPtr *string, userPtr *string, keyPtr *string, jobPtr *string) {
	req, err := http.NewRequest("GET", "http://"+*hostPtr+"/view/CI/job/"+*jobPtr+"/rssAll", nil)
	req.SetBasicAuth(*userPtr, *keyPtr)

	resp, err := client.Do(req)
	if err != nil {
		// Handle Err
	}

	defer resp.Body.Close()

	v := jenkinsrss.Feed{}

	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&v)

	fmt.Printf("%v\n", v.Entries[0].Title)
}
