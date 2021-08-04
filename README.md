# Terraform New Relic Event Action

[![Action Template](https://img.shields.io/badge/Action%20Template-Go%20Container%20Action-blue.svg?colorA=24292e&colorB=0366d6&style=flat&longCache=true&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAA4AAAAOCAYAAAAfSC3RAAAABHNCSVQICAgIfAhkiAAAAAlwSFlzAAAM6wAADOsB5dZE0gAAABl0RVh0U29mdHdhcmUAd3d3Lmlua3NjYXBlLm9yZ5vuPBoAAAERSURBVCiRhZG/SsMxFEZPfsVJ61jbxaF0cRQRcRJ9hlYn30IHN/+9iquDCOIsblIrOjqKgy5aKoJQj4O3EEtbPwhJbr6Te28CmdSKeqzeqr0YbfVIrTBKakvtOl5dtTkK+v4HfA9PEyBFCY9AGVgCBLaBp1jPAyfAJ/AAdIEG0dNAiyP7+K1qIfMdonZic6+WJoBJvQlvuwDqcXadUuqPA1NKAlexbRTAIMvMOCjTbMwl1LtI/6KWJ5Q6rT6Ht1MA58AX8Apcqqt5r2qhrgAXQC3CZ6i1+KMd9TRu3MvA3aH/fFPnBodb6oe6HM8+lYHrGdRXW8M9bMZtPXUji69lmf5Cmamq7quNLFZXD9Rq7v0Bpc1o/tp0fisAAAAASUVORK5CYII=)](https://github.com/jacobtomlinson/go-container-action)
[![Actions Status](https://github.com/jacobtomlinson/go-container-action/workflows/Build/badge.svg)](https://github.com/jacobtomlinson/go-container-action/actions)
[![Actions Status](https://github.com/jacobtomlinson/go-container-action/workflows/Integration%20Test/badge.svg)](https://github.com/jacobtomlinson/go-container-action/actions)

This is a template for creating GitHub actions and contains a small Go application which will be built into a minimal [Container Action](https://help.github.com/en/actions/automating-your-workflow-with-github-actions/creating-a-docker-container-action). Our final container from this template is ~3MB, yours may be a little bigger once you add some code, but it'll still be tiny!

In `main.go` you will find a small example of accessing Action inputs and returning Action outputs. For more information on communicating with the workflow see the [development tools for GitHub Actions](https://help.github.com/en/actions/automating-your-workflow-with-github-actions/development-tools-for-github-actions).

> 🏁 To get started, click the `Use this template` button on this repository [which will create a new repository based on this template](https://github.blog/2019-06-06-generate-new-repositories-with-repository-templates/).

## Usage

Describe how to use your action here.

### Example workflow

```yaml
      - name: Terraform Init
        working-directory: ./report/infra
        id: init
        continue-on-error: true
        run: terraform init -auto-approve -var "profile=prod"
      - name: Terraform Plan
        working-directory: ./report/infra
        id: apply
        continue-on-error: true
        run: terraform apply -auto-approve -var "profile=prod"
      - name: Terraform Apply
        working-directory: ./report/infra
        id: apply
        continue-on-error: true
        run: terraform apply -auto-approve -var "profile=prod"
      - name: Terraform Monitoring
        uses: maathor/terraform-new-relic-event-action@master
        with:
          new_relic_api_key: {{ SECRET.NEW_RELIC_API_KEY}}
          event_type_name: DeployEvent
          env: prod
          terraform_init_status: init
          terraform_apply_status: status
          terraform_tag_key: FeatureTeam
          terraform_tag_value: Report
          github_repository: {{ github.repository}}
          github_runid: {{ github.runid}}
```

### Inputs

| Input                                             | Description                                        |
|------------------------------------------------------|-----------------------------------------------|
| `new_relic_api_key`  | {{ SECRET.NEW_RELIC_API_KEY}}    |
| `event_type_name`   | custom event type name    |
| `env`   | custom event type name    |
| `terraform_init_status`   | custom event type name    |
| `terraform_apply_status`   | custom event type name    |
| `terraform_tag_key`   | tag key to put into GHA monitoring    |
| `terraform_tag_value`   | tag value to put into GHA monitoring    |
| `github_repository`   | {{ github.repository}}    |
| `github_runid`   | {{ github.runid}}    |

## Examples

> NOTE: People ❤️ cut and paste examples. Be generous with them!

This is how to use the optional input.

