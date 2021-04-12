# Terraform Provider

**Note: Issues are currently closed on this repo as we are not quite ready for feedback. Thanks!**

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.11 (to build the provider plugin)

## Building The Provider

Clone repository to: `$GOPATH/src/github.com/twilio/terraform-provider-twilio`

```sh
$ mkdir -p $GOPATH/src/github.com/twilio; cd $GOPATH/src/github.com/twilio
$ git clone git@github.com:twilio/terraform-provider-twilio
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/twilio/terraform-provider-twilio
$ make build
```

**Note**: For contributions created from forks, the repository should still be cloned under the `$GOPATH/src/github.com/twilio/terraform-provider-twilio` directory to allow the provided `make` commands to properly run, build, and test this project.

## Using the provider
If you're building the provider, follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin) After placing it into your plugins directory,  run `terraform init` to initialize it.

Further [usage documentation is provided in usage.md](usage.md).

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.11+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-twilio
...
```

In order to run the full suite of Acceptance tests, run `make testacc`. Provide your Account SID and Auth Token as environment variables to properly configure the test suite.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$  make testacc ACCOUNT_SID=YOUR_ACCOUNT_SID AUTH_TOKEN=YOUR_AUTH_TOKEN
```

An example test file can be found [here](https://github.com/twilio/terraform-provider-twilio/blob/master/twilio/resource_taskrouter_workspace_test.go).

### Inspirations and References ###
We found these resources to be helpful in guiding Twilio's Terraform Provider implementation and in illustrating Terraform's [documentation](https://www.terraform.io/docs/extend/writing-custom-providers.html): 
- [Github](https://github.com/terraform-providers/terraform-provider-github)
- [DigitalOcean](https://github.com/terraform-providers/terraform-provider-digitalocean)


### Debugging
First:
```sh
$ export TF_LOG=TRACE
```
then refer to the [Terraform Debugging Documentation](https://www.terraform.io/docs/internals/debugging.html).