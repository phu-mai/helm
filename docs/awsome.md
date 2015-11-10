# The Helm Guide to Writing Awesome Charts

A Helm Chart provides a recipe for installing and running a containerized application inside of Kubernetes. This guide explains how to write an outstanding Chart.

## Creating a New Chart

You can create a new chart using the `helm create` command. This will put the chart in your workspace, which is the perfect place for trying it out. You may choose to use the `helm edit` command to editor your chart, or you may be more comfortable editing directly with your favorite editor.

## A Brief Anatomy of a Chart

Inside of your newly created chart, there are three pieces:

- The `Chart.yaml` file contains descriptive data about a chart.
- The `manifests/` directory contains one or more Kubernetes manifest files.
- A `README.md` file, which is a Markdown-formatted file that explains how to configure your chart.

This guide walks you through the process of creating those parts.

## Naming Your Chart

A good name succinctly expresses what your chart provides. If your chart installs an application in the "normal" way or is a basic configuration of a package, you should name it with the name of the project. For example, a singe MongoDB chart may simple be called `mongo`.

If your chart displays some more advanced features, or creates a cluster of servers, you should give it a more specific name: `redis-cluster` creates and configures a cluster or Redis servers.

Please take a look at the existing charts before naming your own.

_A Note on example charts:_ As Helm was getting started, we created a few example charts, whose purpose was to illustrate how to write a chart. We are trying to only add new _example_ charts when they illustrate something new and helpful for chart developers.

## The Chart.yaml File

A simple Chart.yaml file should look like this:

```
name: NAME OF CHART
home: URL TO IMAGE HOME
version: VERSION
description: SHORT DESCRIPTION
dependencies:
  - name: CHART
    version: SEMVER FILTER
maintainers:
  - YOUR NAME <YOUR EMAIL>
details:
  ONE PARAGRAPH DESCRIPTION
```

- name: This should be the same name as your chart directory.
- home: The URL to the project where the container image came from. This is to assist people in finding out how their image was built.
- version: A SemVer 2 version strings.
- description: A short (several word) description of this chart.
- dependencies: A set of name/filter pairs
	- name: The name of the chart this chart depends upon
	- version: A filter indicating what version of the chart is required. Example: `~1.2` (greater than or equal to 1.2.0, and less than 1.3.0)
- maintainers: A set of maintainer names, together with an email
- details: A single paragraph describing the chart

Except for `dependencies`, all fields are required.

## The README.md File

The README file performs one important function: It tells the user how to use your chart. It is automatically displayed when a chart is installed.

Your README.md file _may_:

- Tell the user how to configure the chart (if necessary)
- How to work with the chart once it's in Kubernetes (if this is not self-evident)
- Explain any environment variables, secrets, annotations, or special labels
- Provide URLs to external sources that are helpful

Your REDME.md file _should not_:

- Add a license constraint (all charts in `deis/charts` are Apache 2)
- Describe the project in detail

In a nutshell, the README is intended as help text for your user: It gets the user from fetching your chart to using your chart.

## Manifest Files

All Kubernetes manifest files should be in YAML format. We know some people prefer JSON, but we've decided for a number of reasons to adopt YAML as the single format for Helm manifests.

A good simple manifest typically includes:

- A replication controller (`foo-rc.yaml`)
- A service definition for each backend service this chart relies upon (`foo-db-svc.yaml`)
- A secrets file for each secret the RC uses (`foo-password-sec.yaml`)

### Labels

All Helm charts should have a `name` label and a `heritage: helm` label in their metadata sections. These provide a base-level consistency across all Helm charts. (`heritage: helm` makes it easy to search a Kubernetes cluster for all components installed via Helm.)

We suggest using the following labels where appropriate:

- `app`
- `tier`
- `provider`: Indicates that this chart provides a specific backend service (e.g. MariaDB or Percona might include `provider: mysql` since they provide MySQL compatible service)

### Environment Variables

We recommend documenting these in the README. Something like this is great:

```
- DB_PASSWORD: The password used to connect to the MySQL database
```

### Pods vs. Replication Controllers

From the Kubernetes documentation:

> [W]e recommend that you use a replication controller even if your application requires only a single pod

Because replication controllers come with better lifecycle management, we suggest using RCs instead of using pods. This will give your users all of the features of a pod, but with added assurances. That said, when you use an RC to run a service that _cannot be scaled_, you should indicate this to your users either by adding the `scalable: false` label to the RC, or by saying so directly in the README.

### Secrets

We recommend providing default secrets whenever a secret is needed, and also explaining the secret in the README.

### Jobs, DaemonSets, and other experimental features

When using a definition from the experimental set, please note this in the `Chart.yaml` description with the notation `(experimental)`.

Example:

```yaml
name: sweep
home: Job to sweep the cluster once. (experimental)
# ...
details:
  This sweeps out the dusty corners of the cluster. It uses the experimental Jobs type.
```

## Testing

The Helm chart reviewers expect that you have tested your chart, both for compatibility with Helm, and also with Kubernetes. We run some basic tests on your charts on submission, but these perform only rudimentary checks. So please test before submitting.

## Publishing a Pull Request

The suggested workflow for publishing a pull request goes like this:

1. From GitHub, clone the `github.com/deis/charts` repository
2. Add your fork to Helm: `helm repo add technosophos git@github.com:technosophos/charts.git`
3. Copy your chart from your workspace to your new charts repo (see `helm publish`)
4. Commit and push to your charts repo
5. Using GitHub, file a PR (Pull Request) against the `deis/charts` repository
6. Follow along in the `deis/charts` issue queue

## Review

The core Helm charts team reviews all charts to see if they comply with our chart best practices. **All charts must be reviewed and marked LGTM by two members of the Helm charts team.** Once that is done, your chart will be merged into the repository.

We absolutely _love_ contributions, so don't fret about this part of the process. When we ask for additional changes, it's only because we (like you) want the Helm community to have the best experience possible.

See you in the issue queue!


