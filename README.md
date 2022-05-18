# apptrack-server
Apptrack-server is a GraphQL server built with Go. This server functions as the backend for [apptrack](https://github.com/petergeorgas/apptrack), which as the name suggests, is an application tracking web app. 
Originally, this repository was going to be the gist of the project, as I just wanted to get more comfortable with GraphQL and Go. Once the API was implemented, I decided that it would be interesting to make this project full-stack, so the apptrack website was created.

This server is currently deployed to Google Cloud Platform in a Cloud Run instance, that is what the `Dockerfile` in the root of the repository is for. Soon, the `playground` endpoint of the API will be closed, as I will not be needing it anymore.

## Upcoming Features
* Create/add tests. 
