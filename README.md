# 🦕 Deploydocus 

The goal of this repository is to showcase an easy deployment pipeline that just works.

The pipeline has following component:
 * Once PR is merged, docker image is built and pushed to staging Google Cloud Run instance [https://deploydocus-555983195874.us-central1.run.app](https://deploydocus-555983195874.us-central1.run.app). This workflow is defined in [deploy-staging.yml](https://github.com/yb172/deploydocus/blob/main/.github/workflows/deploy-staging.yml)

The app is a simple server app that returns what version it is running, the version is passed as a build flag when building docker image

## Google cloud resources

In order for this thing to work, following things are needed to be configured on Google cloud:
 * For [`google-github-actions/auth@v2`](https://github.com/google-github-actions/auth) (see setup instructions [here](https://github.com/google-github-actions/auth?tab=readme-ov-file#preferred-direct-workload-identity-federation))
   * IAM: Workload identity pool (`deploydocus`)
   * IAM: Workload identity pool provider (`deploydocus-provider`, used in actions config: `projects/555983195874/locations/global/workloadIdentityPools/deploydocus/providers/deploydocus-provider`)
   * IAM: Service account for deployment (`deploydocus-github@windy-orb-218701.iam.gserviceaccount.com`) with bindings
     * `roles/artifactregistry.writer`
     * `roles/iam.serviceAccountTokenCreator`
     * `roles/run.admin`
 * Artifact registry: `deploydocus`
 * Cloud run
   * `deploydocus-staging` (updated on each PR merge)

...and probably something else
