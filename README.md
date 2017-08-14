# bashir

![logo](http://assets.avi.io/small-logo-icon.png)

## What is this project?

This project was created to automate running commands in a safe way. Specifically for me, every time we release software there's a series of commands needed to run and reported back to the command creators.

These include ansible commands, local commands in a python virtual environment and more. So, this means we have a command manifest that we need to run in sequenece.

For example:

1. Scale up the cluster (Runs a command that will wait for the cluster to be active)
2. Update the counts of the docker containers you want to run
3. Run ansible on the cluster (This does multiple things like migrating data between services and more)
4. Scale down the other cluster

Bashir was created to minimize the copy pasting from release notes, essentially having a single configuration file that engineers edit and will be run by the release manager BUT it can really automate any number of commands that you usually run in sequence.

For the initial version, Bashir focuses on docker containers, running commands on docker containers and outputting the result.

## Alternatives

I researched multiple alternatives for this and came up with a couple

1. RunDeck
2. Spinnaker

Both of these allow you to have some sort of a manifest of arbitrary commands that you can run on your cluster. Both of these are **way** richer than Bashir. I wanted to focus on the simplicity and extensions in Bashir so the first class citizen for me is the reporting back to Slack and saving a log of the commands.

I also looked into working directly with the docker client but this didn't pan out because in the next version of bashir docker will be supported but also running other commands. Each command will have a `type`, one of these will be `docker`.

## Limitations

1. File has to be executable `chmod +x` and if it's a shell file, you will need
   a proper header such as `#!/bin/sh`

## Configuration

Configuration in Bashir is a YAML file and it has multiple sections

`slack`

Slack configures the webhook URL that the command will be reported to. The defaults for reporting a command is to say `:pushpin: starting {command name} cc/ @{report_to}, @{report_to}`.

You can easily create webhook URLs for the channel you want. For us it's a `#release` channel that tracks all of the commands we run in a release.

`defaults`

Defaults configures items that will be attached to every command. Imagine it's a docker container so here you will configure env vars and arguments that will get attached by default to all containers.

`commands`

Configures the commands:

* `name` Name of the command
* `command` the actual command to run inside the container
* `args` arguments to call on the command being ran
* `envvars` ENV vars to attach to the container
* `image_name` docker container image name to run
* `report_to` Who should we report the command output to?


```
---
slack:
  webhook_url: https://test.com
  channel: "#channelname"
  icon: ":iconname:"
  botname: botname

defaults:
  envvars:
    - TEST_A
    - TEST_B
  args:
    - -x
    - test

commands:
  - name: Migrate service X
    description: |
      This is the command decription
      and it can include lots of information
    image_name: kensodev/bashir
    command: /run.sh
    envvars:
      - TEST_C
    args:
      - --check
      - -i
      - inventory/aws.py
      - -b
      - someting/aws.py
    report_to:
      - KensoDev

  - name: Migrate service X
    image_name: kensodev/bashir
    command: /run.sh
    envvars:
      - TEST_C
    args:
      - --check
      - -i
      - inventory/aws.py
      - -b
      - someting/aws.py
    report_to:
      - KensoDev

```

## Project status

Project is actively developed and being worked on. If you have comments/suggestions please feel free to let me know

## Credits

* [@kayteh](https://github.com/kayteh) - Pairing and solving the env vars and args merging nicely into a single command and debugging the docker run command.

