# drone-github-pr-commenter
A Drone CI plugin, for more information: https://docs.drone.io/plugins/overview/

Posts a comment to an issue/PR in Github.

Docker hub: https://hub.docker.com/repository/docker/mtyurt/drone-github-pr-commenter

## Usage

```yaml
kind: pipeline
type: docker
name: Build site for PR

trigger:
  event: pull_request

- name: Add comment to PR
  image: mtyurt/drone-github-pr-commenter
  settings:
    github_token: .....
    comment: The branch is deployed successfully! 🚀
```

* **github_token**: Some personal access token
* **comment**: Comment body
* **issue_number**: (optional) The issue/PR number to post for. If it is not provided, the PR build is running for is
    used.

### Test usage without Drone CI

```bash

docker run --rm -it \
 -e PLUGIN_GITHUB_TOKEN=$GITHUB_TOKEN \
 -e PLUGIN_COMMENT='This branch is deployed successfully! 🚀' \
 -e DRONE_REPO_NAMESPACE=<ci-repo-owner> \
 -e DRONE_REPO_NAME=<ci-repo-name> \
mtyurt/drone-github-pr-commenter:latest
```
