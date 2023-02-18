# kubebenchctl
kubebenchctl is a CLI tool written in Go that utilizes the kube-bench open-source tool to diagnose Kubernetes clusters and provide CIS benchmark scores. With kubebenchctl, you can easily and automatically diagnose your Kubernetes cluster and identify security issues and areas for improvement.

## Installation
To install kubebenchctl, you can use Go's go get command:

```bash
go get github.com/your-username/kubebenchctl
```

Alternatively, you can download the binary release from the project's GitHub releases page and add it to your $PATH.

## Usage
Once you have installed kubebenchctl, you can run it by executing the following command:

```bash
kubebenchctl diagnose
```

This will run the kube-bench tool and provide a report of your cluster's CIS benchmark scores. You can also specify flags to configure the tool's behavior, such as --cluster to specify the cluster name or --node to diagnose a specific node.

For more information about the tool's usage and options, you can run kubebenchctl help.

## Contributing
kubebenchctl is an open-source project, and we welcome contributions from the community. To contribute to the project, you can follow these steps:

Fork the project to your GitHub account.
Clone your forked repository to your local machine.
Create a new feature branch.
Make your changes and commit them with descriptive messages.
Push your feature branch to your forked repository.
Open a pull request to merge your changes into the main repository.
For more information about contributing to the project, please refer to the CONTRIBUTING.md file.

## License
This project is licensed under the MIT License. See the LICENSE file for more information.

## Credits
This project is inspired by the kube-bench tool, which is an open-source tool for diagnosing Kubernetes clusters against the CIS benchmarks. Special thanks to the kube-bench team for their great work.