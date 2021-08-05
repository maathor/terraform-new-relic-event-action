# Terraform New Relic Event Action

## Usage

This github action take in parameters:
- terraform init status
- terraform apply status
- number of create/update/delete operations
- GHA link

It writes into your New Relic account a new customEvent that contains all the parameter given in parameters and status.

## Explaining JQ

Why do we have some jq operations piped in the terraform show command ?

To make sure the github action does not contain any critical information!
We will not collect any data, and we prove it;)

### Example workflow

```yaml
      - name: Terraform Init
        working-directory: ./report/infra
        id: init
        continue-on-error: true
        run: terraform init
      - name: Terraform Plan
        working-directory: ../report/infra
        id: plan
        continue-on-error: true
        run: |
          terraform plan -out tfplan.out
          terraform show -json tfplan.out | jq '.resource_changes[].change.actions[]' > changes.log
          change_path=$(readlink -f changes.log)
          echo "::set-output name=changes_path::$change_path"
      - name: Terraform Apply
        working-directory: ./report/infra
        id: apply
        continue-on-error: true
        run: terraform apply -auto-approve
      - name: Terraform Monitoring
        uses: maathor/terraform-new-relic-event-action@master
        with:
          new_relic_licence_key: ${{ secrets.NEW_RELIC_API_KEY}}
          event_type_name: DeployEvent
          env: prod
          terraform_init_status: ${{ steps.init.outcome }}
          terraform_operation_list_path: ${{ steps.plan.outputs.changes_path}}
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
| `terraform_operation_list_path`   | output of `terraform show -json tfplan.out <pipe> jq .resource_changes[].change.actions[]`    |
| `terraform_apply_status`   | custom event type name    |
| `terraform_tag_key`   | tag key to put into GHA monitoring    |
| `terraform_tag_value`   | tag value to put into GHA monitoring    |
| `github_repository`   | {{ github.repository}}    |
| `github_runid`   | {{ github.runid}}    |

### Outputs

| Output                                             | Description                                        |
|------------------------------------------------------|-----------------------------------------------|
| `terraform_update`  | terraform update resource number    |
| `terraform_create`  | terraform create resource number    |
| `terraform_delete`  | terraform delete resource number    |
| `terraform_noop`  | terraform no operation resource number    |