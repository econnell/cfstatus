# cfstatus

# Install

It's a go program, so just download the binary, decompress it, and run it:

https://github.com/econnell/cfstatus/blob/master/cfstatus.bz2?raw=true

# Usage

```cfstatus -region [region name] -stack-name [stack name]```

For example:  ```cfstatus -region us-west-2 -stack-name mystack```

The default region is us-west-2.  AWS credentials are read from ~/.aws/credentials or from standard environment variables.  For more info, go to: https://github.com/aws/aws-sdk-go#configuring-credentials

# Build from source:

Install go if you don't have it:  https://golang.org/dl/

Set up your GOPATH environment variable if you haven't already:  export GOPATH=$HOME

```
go get -u github.com/econnell/cfstatus
cd $GOPATH/src/github.com/econnell/cfstatus
go get
go install
~/bin/cfstatus -region us-west-2 -stack-name mystack
```
