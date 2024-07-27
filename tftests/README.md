# Terraform Tests

We highly recommend making terraform tests to make sure that your Observe Apps don’t break unexpectedly while they are changed over time. 

Today, a “terraform test” is a set of terraform inputs that, when you run a `terraform apply` and `destroy`, works seamlessly without returning any errors. This basic level of testing ensures that, for a given set of inputs…

1. All OPAL is valid
2. No resource names conflict
3. No other issues block a terraform apply from executing successfully

If, for example, I were to change the name of a field in one datastream that was also a dependency in a subsequent datastream, a basic terraform test would catch this before the change would be made available for use by customers. 

In the future, you will be able to define custom assertions for your terraform tests to make sure that the results of a `terraform apply` match your own pre-defined expectations. 

**Note**: Tests are run in parallel, so adding more tests doesn’t necessarily increase your CI time. You'll generally want to have a test for every reasonable combination of terraform module inputs (e.g, for every reasonably expected combination of feature flags or variable inputs, etc.)

<!--
## Required Github Secrets

Before your `terraform-observe-*` repo can run terraform tests in its github actions, you must set the following Github secrets as either organization or repository secrets:

1. `TERRAFORM_MODULES_TEST_OBSERVE_CUSTOMER` - the customer account id for the observe instance you wish to use for CI testing
2. `TERRAFORM_MODULES_TEST_OBSERVE_DOMAIN` - the observe domain for the observe instance you use for CI testing (defaults to `staging-observe.com`, can be `observeinc.com`)
3. `TERRAFORM_MODULES_TEST_OBSERVE_USER_EMAIL` - the user email your CI tests should use to authenticate with your observe instance
4. `TERRAFORM_MODULES_TEST_OBSERVE_USER_PASSWORD` - the user password your CI tests should use to authenticate with your observe instance
-->

## How to make a terraform test

1. Add a new subdirectory under `tftests/`. The name of this subdirectory will effectively be the name of your test, so name it in a way that indicates the goals of the test. 
2. Create a `main.tf` in your test directory. This should contain a source value of `"../.."` so that the test uses your Observe App’s terraform module, along with any other variables, resources, data sources that your module requires. 
3. Create a `versions.tf` in your test directory that defines the required providers and versions your test uses (see the example below for inspiration)
4. Create an empty `variables.tf` in your test directory (so that pre-commit doesn’t complain)
5. Create an empty `outputs.tf` in your test directory (so that pre-commit doesn’t complain) 

If your App has nested terraform modules, the same workflow applies. The tests for your root-level module will be at `tftests/TEST_NAME/` and the tests for your nested modules will be at `NESTED_MODULE_PATH/tftests/TEST_NAME/`, and the pieces required to make a test for a nested module are the same. 

### Random Inputs

You may want to introduce some amount of randomness in your test definitions. For example, if you have multiple tests, you may need to ensure they have separate values under the `name_format` variable. One way to do this would be to add a `random_*` resource from the `random` provider, e.g, `resource "random_pet" "test" {}`, at which point you could set your `name_format` value to something like `name_format = "test_${random_pet.test.id}/%s"` (Since these tests get destroyed quickly after they are applied, you likely won’t get to see the name of your random_pet in this case 😞, but it can help avoid conflicts between your tests nonetheless.) 

### An Example Test:

A common test example is a `default` test in which you might define the bare minimum variable inputs required for your App to run, so as to ensure that everything applies smoothly with whatever are the default variable inputs. This test would…

1. Be defined at `tftests/default/`
2. Have a `tftests/default/main.tf` whose content might be
    
    ```terraform
    resource "random_pet" "test" {}
    
    data "observe_workspace" "default" {
      name = "Default"
    }
    
    module "default" {
      source            = "../.."
      workspace         = data.observe_workspace.default
      name_format       = "test_${random_pet.test.id}/%s"
    }
    ```
    
3. Have a `tftests/default/versions.tf` whose content might be
    
    ```terraform
    terraform {
      required_providers {
        observe = {
          source  = "terraform.observeinc.com/observeinc/observe"
          version = ">= 0.4.15"
        }
        random = {
          version = ">= 3"
        }
      }
      required_version = ">= 0.13"
    }
    ```
    
4. Have an empty `tftests/default/variables.tf`
5. Have an empty `tftests/default/outputs.tf` 

Other useful tests might be `all_options` where you would define every possible non-default option for your test module, or any specific combination of input options that are especially interesting for your Observe App.

## Running your tests locally

Tests are run with a bash make file. They require the following environment variables…

- `OBSERVE_CUSTOMER` - the customer account id for the Observe instance that your observe provider will use in your test
- `OBSERVE_DOMAIN` - defaults to `observe-staging.com`, can be set to `observeinc.com`
- `OBSERVE_USER_EMAIL`
- `OBSERVE_USER_PASSWORD`

To run your test, go to the root directory of your Observe App repo (e.g, `./terraform-observe-example/`), which is also the location of your Makefile, and run `make test tftests/TEST_NAME/` or for tests of nested modules, `make test NESTED_MODULE_PATH/tftests/TEST_NAME/` (tabbing through the directories is your friend)
