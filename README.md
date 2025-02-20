# Deploy

The goal of this repository is to showcase an easy deployment pipeline that just works.

The pipeline has following component:
 * Once PR is merged, docker image is built and pushed to staging Google Cloud Run instance [https://deploy-staging-555983195874.us-central1.run.app](https://deploy-staging-555983195874.us-central1.run.app)

The app is a simple server app that returns what version it is running, the version is passed as a build flag when building docker image
