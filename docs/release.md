# Cutting a new release

## Create a new release / tag

Create new tag, release and release notes:

1. Goto https://github.com/smoothglue/provider-sonarqube/releases/new
1. Click `Choose a tag`
1. Enter the new tag in the box
1. Click `Create new tag: <tag-name>`
1. Click `Generate release notes`
1. Click `Publish release`

## Generate artifacts

Create a new release branch off of main and push it to github to start the pipeline:

1. `git checkout main && git pull`
1. `git checkout -b release-<version-number>`
1. `git push --set-upstream origin <branch-name>`

Wait for pipeline to finish and the new artifacts should show up [here](https://github.com/smoothglue/provider-sonarqube/pkgs/container/provider-sonarqube).
