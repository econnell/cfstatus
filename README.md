# cfstatus

# Usage

```cfstatus -region [region name] -stack-name [stack name]```

For example:  ```cfstatus -region us-west-2 -stack-name mystack```

The default region is us-west-2.  AWS credentials are read from ~/.aws/credentials or from standard environment variables.  For more info, go to: https://github.com/aws/aws-sdk-go#configuring-credentials

# Build from source:

1. Set up your GOPATH environment variable:  export GOPATH=$HOME
2. go get -u github.com/aws/aws-sdk-go
3. go get github.com/vaughan0/go-ini
4. go install
5. ~/bin/cfstatus
