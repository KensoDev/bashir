# bashir

### Readme driven development

This repo is using README driven development. The features described here might
not be available if you download/compile/run the code.

I am targetting 0.5.0 as the first "real" version that you can run and count on
the features being there.

The point of README driven development is to have the documentation extensively
written **before** writing code while making sure the API makes sense.

Another goal is to have a sort of running design

During the development, feel free to open issues and PRs if you find somehting
useful and want to participate.

## What is Bashir

Bashir is bash automation. It has a DSL that allows you to define commands that
will run in sequence or in parallel.

Each command has args, params, environment variables, files and more.

One really common workflow for us is running `ansible-playbook --check --diff`,
checking the output and running `ansible-playbook --diff`. BUT, we also have a
habbit of notifying others if we run it so we also need to slack the team.

Bashir takes care of that for you, you can slack multiple teams, mention people
and much more.

The development of Bashir is following very strict minimal approach and I like
to verify if something is useful before putting it into the product.

## Configuration

Create a YAML file called `bashir.yml` and paste this content.

```
---
command_args:
  - environment
  - ansible_pass_file

commands:
  - name: Migrate service X
    command: ansible-playbook ... {.env}
    virtualenv: ansible-deploy
    report_to:
      - KensoDev

  - name: Migrate service X
    command: ansible-playbook ... {.Env}
    virtualenv: ansible-deploy
    report_to:
      - KensoDev

```
