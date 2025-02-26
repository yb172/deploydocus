# 🦕 Deploydocus 

The goal of this repository is to showcase an easy deployment pipeline that just works.

 * Developers work as usual (the only thing: commits must follow [Conventional commits spec](https://www.conventionalcommits.org/en/v1.0.0/))
 * [`staging.deploydocus.yb172.dev`](http://staging.deploydocus.yb172.dev) is updated with image built from head every time merged PR has label "deploy to staging"
 * [`demo.deploydocus.yb172.dev`](http://demo.deploydocus.yb172.dev) is updated with image built from head every time release PR auto-created by [release please](https://github.com/googleapis/release-please-action) is merged
 * [`deploydocus.yb172.dev`](http://deploydocus.yb172.dev) is updated with `<version>` when `./deploy-prod <version>` command is executed and created PR is merged

The app is a simple server app that returns what version it is running, the version is passed as a build flag when building docker image

```css
_________________________________________________________________________________________________________________________
/~~~~~~~\__/~~~~~~~~\_/~~~~~~~\__/~~\________/~~~~~~\__/~~\__/~~\_/~~~~~~~\___/~~~~~~\___/~~~~~~\__/~~\__/~~\__/~~~~~~\__
/~~\__/~~\_/~~\_______/~~\__/~~\_/~~\_______/~~\__/~~\__/~~\/~~\__/~~\__/~~\_/~~\__/~~\_/~~\__/~~\_/~~\__/~~\_/~~\_______
/~~\__/~~\_/~~~~~~\___/~~~~~~~\__/~~\_______/~~\__/~~\___/~~~~\___/~~\__/~~\_/~~\__/~~\_/~~\_______/~~\__/~~\__/~~~~~~\__
/~~\__/~~\_/~~\_______/~~\_______/~~\_______/~~\__/~~\____/~~\____/~~\__/~~\_/~~\__/~~\_/~~\__/~~\_/~~\__/~~\_______/~~\_
/~~~~~~~\__/~~~~~~~~\_/~~\_______/~~~~~~~~\__/~~~~~~\_____/~~\____/~~~~~~~\___/~~~~~~\___/~~~~~~\___/~~~~~~\___/~~~~~~\__
_________________________________________________________________________________________________________________________
env: <env>, v: v1.1.1, build date: <date>
```

## Google cloud resources

In order for this thing to work, following things are needed to be configured on Google cloud:
 * For [`google-github-actions/auth@v2`](https://github.com/google-github-actions/auth) (see setup instructions [here](https://github.com/google-github-actions/auth?tab=readme-ov-file#preferred-direct-workload-identity-federation))
   * IAM: Workload identity pool (`deploydocus`)
   * IAM: Workload identity pool provider (`deploydocus-provider`, used in actions config, passed as a secret `WORKLOAD_IDENTITY_PROVIDER`)
   * IAM: Service account for deployment, passed as a secret `SERVICE_ACCOUNT` with bindings
     * `roles/artifactregistry.writer`
     * `roles/iam.serviceAccountTokenCreator`
     * `roles/run.admin`
 * Artifact registry: `deploydocus`
 * Cloud run
   * `deploydocus-staging` (updated on each PR merge, if PR was labeled with 'deploy to staging')
   * `deploydocus-demo` (updated once new release is created)
   * `deploydocus` (updated when workflow that promotes demo image to prod is manually triggered and completed)

...and probably something else
