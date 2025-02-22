# 🦕 Deploydocus

The goal of this repository is to showcase an easy deployment pipeline that just works.

The pipeline has following component:
 * Once PR is merged, docker image is built and pushed to staging Google Cloud Run instance [https://deploydocus-555983195874.us-central1.run.app](https://deploydocus-555983195874.us-central1.run.app). This workflow is defined in [deploy-staging.yml](https://github.com/yb172/deploydocus/blob/main/.github/workflows/deploy-staging.yml)

The app is a simple server app that returns what version it is running, the version is passed as a build flag when building docker image
