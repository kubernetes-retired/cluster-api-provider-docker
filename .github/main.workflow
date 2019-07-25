workflow "Release" {
  on = "push"
  resolves = ["push images"]
}

action "is-tag" {
  uses = "actions/bin/filter@master"
  args = "tag"
}

action "goreleaser" {
  uses = "docker://goreleaser/goreleaser"
  secrets = ["GORELEASER_GITHUB_TOKEN"]
  args = "release"
  needs = ["is-tag"]
}

action "docker build" {
  uses = "actions/docker/cli@master"
  args = "build -t capd-manager ."
  needs = ["goreleaser"]
}

action "tag images" {
  uses = "actions/docker/tag@master"
  args = "capd-manager gcr.io/kubernetes1-226021/capd-manager"
  needs = ["docker build"]
}

action "Setup Google Cloud" {
  uses = "actions/gcloud/auth@master"
  secrets = ["GCLOUD_AUTH"]
  needs = ["tag images"]
}

action "Set Credential Helper for Docker" {
  uses = "actions/gcloud/cli@master"
  args = ["auth", "configure-docker", "--quiet"]
  needs = ["Setup Google Cloud"]
}

action "push images" {
  uses = "actions/docker/cli@master"
  runs = "sh -c"
  env = {
    IMAGE_NAME = "gcr.io/kubernetes1-226021/capd-manager"
  }
  args = ["source $HOME/.profile && docker push $IMAGE_NAME:latest && docker push $IMAGE_NAME:$IMAGE_REF && docker push $IMAGE_NAME:$IMAGE_SHA && docker push $IMAGE_NAME:$IMAGE_VERSION"]
  needs = ["Set Credential Helper for Docker"]
}
