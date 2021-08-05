# Terraform New Relic Event Action

## Usage

This github action take in parameters:
- terraform init status
- terraform apply status

It writes into your New Relic account a new customEvent that contains all the parameter given in parameters and status.

### Example workflow

```yaml
      - name: Terraform Init
        working-directory: ./report/infra
        id: init
        continue-on-error: true
        run: terraform init -auto-approve -var "profile=prod"
      - name: Terraform Plan
        working-directory: ./report/infra
        id: plan
        continue-on-error: true
        run: |
          terraform plan -var "profile=prod" -out tfplan.out
          changes_list=$(terraform show -json tfplan.out | jq '.resource_changes[].change.actions')
          echo "::set-output name=changes::$changes_list"
      - name: Terraform Apply
        working-directory: ./report/infra
        id: apply
        continue-on-error: true
        run: terraform apply -auto-approve -var "profile=prod"
      - name: Terraform Monitoring
        uses: maathor/terraform-new-relic-event-action@master
        with:
          new_relic_licence_key: ${{ secrets.NEW_RELIC_API_KEY}}
          event_type_name: DeployEvent
          env: prod
          terraform_init_status: ${{ steps.init.outcome }}
          terraform_operation_list: ${{ steps.plan.outputs.changes}}
          terraform_apply_status: ${{ steps.apply.outcome }}
          terraform_tag_key: FeatureTeam
          terraform_tag_value: Report
          github_repository: ${{ github.repository}}
          github_run_id: ${{ github.run_id}}
```

### Inputs

| Input                                             | Description                                        |
|------------------------------------------------------|-----------------------------------------------|
| `new_relic_licence_key`  | {{ SECRET.NEW_RELIC_API_KEY}}    |
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

