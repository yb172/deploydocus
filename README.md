# 🦕 Deploydocus 

The goal of this repository is to showcase an easy deployment pipeline that just works.

 * Developers work as usual (the only thing: commits must follow [Conventional commits spec](https://www.conventionalcommits.org/en/v1.0.0/))
 * [`staging.deploydocus.yb172.dev`](http://staging.deploydocus.yb172.dev) is updated with image built from head every time merged PR has label "deploy to staging"
 * [`demo.deploydocus.yb172.dev`](http://demo.deploydocus.yb172.dev) is updated with image built from head every time release PR auto-created by [release please](https://github.com/googleapis/release-please-action) is merged
 * [`deploydocus.yb172.dev`](http://deploydocus.yb172.dev) is updated with `<version>` when `./release/deploy-prod.sh <version>` command is executed and created PR is merged

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

## Cherrypicks

Oh no, there is an issue in production and we're halfway in migrating forms validation which requires extensive testing. No worries, there is a cherrypick process just for that kind of situations:

```sh
./release/cherrypick.sh <version>
```

This command will create a new branch and checkout code as of `<version>`.

 * Do the fix, commit it with `fix:` prefix (to bump patch only), and push the `cherrypick-draft-<version>`.

 * Open PR to merge the draft into `release-cherrypick-<version>` branch (use link provided in the output of `cherrypick.sh` command).
 * Once PR is merged, you should see release-please create a release PR. Review and merge it.
 * Once this PR is merged a new release of `<version major>.<version minor>.<version patch+1>` will be created and deployed to demo environment.

> [!IMPORTANT]  
> PR must be opened to merge into `release-cherrypick-<version>`, not `main`

 * Once tested, it could be deployed to production using `./release/deploy-prod.sh <vesion with fix>`
 * Create a PR to merge fixes and `CHANGELOG` updates to `main` by opening PR from `release-cherrypick-<version>` to `main` (use link provided in the output of `cherrypick.sh` command)
 * Done!

## GitHub config

For this setup to work you need to create Personal Access Token (PAT) and save it as a repository secret named `RELEASE_PLEASE_TOKEN` (used in `release-please*.yml` configs).

Also `staging`, `demo` and `prod` environments are needed with `WORKLOAD_IDENTITY_PROVIDER` and `SERVICE_ACCOUNT` secrets

It's also nice to create `release-cherrypick-*` protection rule to require a pull request before merging with approvals. That way only approved PRs will be merged to produce docker image and go to demo env.

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
