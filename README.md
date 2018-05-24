# slack-gae-go

Sample GAE/Go SE application to handle slack slash command and interactive message.

## Demo

![deploycommand mov](https://user-images.githubusercontent.com/6121271/40494986-1d006294-5fb1-11e8-9722-fbf5eaf3fc91.gif)

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
