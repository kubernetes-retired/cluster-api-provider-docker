FROM debian:9.5-slim
LABEL "com.github.actions.name"="deploy image"
LABEL "com.github.actions.description"="Pushes a new image to github packages"
LABEL "com.github.actions.icon"="arrow-up"
LABEL "com.github.actions.color"="purple"

LABEL "repository"="http://github.com/kubernetes-sigs/cluster-api-provider-docker"
LABEL "maintainer"="Chuck Ha <chuckh@vmware.com>"

ADD entrypoint.sh .
ENTRYPOINT ["entrypoint.sh"]
