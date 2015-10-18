# jira2lametric

A small wrapper to push your just created JIRA issues to your [LaMetric](http://lametric.com) written in golang.

## Setup

### Create a private lametric app

Go to https://developer.lametric.com and create a new indicator app.

1. Select an icon (just choose what you like).
2. Enter a default text (i.e. JIRA).
3. Set Push for communication type.
4. Click on Next.
5. Enter an app name (i.e. JIRA).
6. Enter a short description.
7. Ensure that "Private App" is selected.
8. Save and publish the app.

On the detail page of your just published app you find the necessary configuration options (Push URL and Access Token).


### Deploy jira2lametric application

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

This repository is prepared for heroku. Of course, you can use it on your own server. Just install it with:

```
$ go get github.com/MMore/jira2lametric
```


### Set Application Configuration

Set the following environment variables:

- PORT (HTTP Port where the application will be served. Heroku will set it automatically.)
- LAMETRIC_PUSH_URL (Your personal lametric device push url. You find it in your app settings.)
- LAMETRIC_TOKEN (Your personal token for your lametric app. You find it in your app settings.)


### JIRA Configuration

Tested with JIRA v6.4.
Go to JIRA Administration and add a new webhook with the URL where your deployed app is running. The hook should fire for created issues only.


## Contributing
This is an open source project and your contribution is very much appreciated.

1. Check for open issues or open a fresh issue to start a discussion around a feature idea or a bug.
2. Fork the repository on Github and make your changes on the **develop** branch (or branch off of it).
3. Send a pull request (with the **develop** branch as the target).


## Changelog
See [CHANGELOG.md](changelog.md)


## License
jira2lametric is available under the GPL v3 license. See the [LICENSE](LICENSE) file for more info.
