# cfstatus

# Usage

```cfstatus -region [region name] -stack-name [stack name]```

For example:  ```cfstatus -region us-west-2 -stack-name mystack```

The default region is us-west-2.  AWS credentials are read from ~/.aws/credentials or from standard environment variables.  For more info, go to: https://github.com/aws/aws-sdk-go#configuring-credentials

# Build from source:

Install go:  https://golang.org/dl/

Set up your GOPATH environment variable:  export GOPATH=$HOME

1. mkdir -p ~/src/github.com/econnell
2. cd ~/src/github.com/econnell
3. git clone https://github.com/econnell/cfstatus.git
4. go get -u github.com/aws/aws-sdk-go
5. go get -u github.com/vaughan0/go-ini
6. go install
7. ~/bin/cfstatus -region us-west-2 -stack-name mystack
