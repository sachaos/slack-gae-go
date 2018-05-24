# slack-gae-go

Sample GAE/Go SE application to handle slack slash command and interactive message.

## Deploy

### Prepare

1. Create Slack App.
    * https://api.slack.com/
2. Setup GCP project and development environment.
    * [Creating and Managing Projects  |  Resource Manager Documentation  |  Google Cloud](https://cloud.google.com/resource-manager/docs/creating-managing-projects?hl=en)
3. Copy to secrets.yaml and edit it.

```shell
$ cp secrets.yaml{.example,}
```

### Run deploy command

```shell
$ gcloud app deploy
```
