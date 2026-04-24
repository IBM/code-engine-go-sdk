[![Build Status](https://app.travis-ci.com/IBM/code-engine-go-sdk.svg?token=FXK1AJgZc9KQsRMKbgpi&branch=main)](https://app.travis-ci.com/github/IBM/code-engine-go-sdk)

[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)

# IBM Cloud Code Engine Go SDK 6.0.0

Go client library to interact with the [Code Engine API](https://cloud.ibm.com/apidocs/codeengine).

## Table of Contents
<!--
  The TOC below is generated using the `markdown-toc` node package.

      https://github.com/jonschlinkert/markdown-toc

  You should regenerate the TOC after making changes to this file.

      npx markdown-toc -i README.md
  -->

<!-- toc -->

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Breaking Changes in `codeenginev2` (March 2026)](#breaking-changes-in-codeenginev2-march-2026)
- [Installation](#installation)
  * [`go get` command](#go-get-command)
  * [Go modules](#go-modules)
- [Using the SDK](#using-the-sdk)
- [Questions](#questions)
- [Issues](#issues)
- [Open source @ IBM](#open-source--ibm)
- [Contributing](#contributing)
- [License](#license)

<!-- tocstop -->

## Overview

The IBM Cloud Code Engine Go SDK allows developers to programmatically interact with the following IBM Cloud services:

Service Name | Package name
--- | ---
[Code Engine](https://cloud.ibm.com/apidocs/codeengine/v2) | codeenginev2
[Code Engine](https://cloud.ibm.com/apidocs/codeengine/v1) | ibmcloudcodeenginev1

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration

- An [IBM Cloud][ibm-cloud-onboarding] account.
- An IAM API key to allow the SDK to access your account. Create one [here](https://cloud.ibm.com/iam/apikeys).
- Go version 1.24 or above.

## Breaking Changes in `codeenginev2` (April 2026)

As part of the introduction of `persistent_data_stores` as a Volume Mount type for **apps** and **jobs**, the optional `name` property of entries in `run_volume_mounts` is no longer supported.

If you used the optional `name` property, remove it from the prototype of the **app** or **job**.

## Breaking Changes in `codeenginev2` (March 2026)

For consistency, the March 2026 update introduces **pluralized list APIs**, **new outbound destination types**, and updated **constructor/patch models**. These changes require updates to existing client code.

- **Method renames (pluralization)**
    Update all list calls and their context variants:

    ```go
    // before
    res, resp, err := ce.ListAllowedOutboundDestination(opts)
    res, resp, err := ce.ListAllowedOutboundDestinationWithContext(ctx, opts)

    // after
    res, resp, err := ce.ListAllowedOutboundDestinations(opts)
    res, resp, err := ce.ListAllowedOutboundDestinationsWithContext(ctx, opts)
    ```

    ```go
    // before
    res, resp, err := ce.ListPersistentDataStore(opts)
    res, resp, err := ce.ListPersistentDataStoreWithContext(ctx, opts)

    // after
    res, resp, err := ce.ListPersistentDataStores(opts)
    res, resp, err := ce.ListPersistentDataStoresWithContext(ctx, opts)
    ```

- **Options types and constructors renamed**
    Replace option structs and their helpers:

    ```go
    // before
    opts := &codeenginev2.ListAllowedOutboundDestinationOptions{ ProjectID: core.StringPtr(pid) }
    opts := ce.NewListAllowedOutboundDestinationOptions(pid)

    // after
    opts := &codeenginev2.ListAllowedOutboundDestinationsOptions{ ProjectID: core.StringPtr(pid) }
    opts := ce.NewListAllowedOutboundDestinationsOptions(pid)
    ```

    ```go
    // before
    opts := &codeenginev2.ListPersistentDataStoreOptions{ ProjectID: core.StringPtr(pid) }
    opts := ce.NewListPersistentDataStoreOptions(pid)

    // after
    opts := &codeenginev2.ListPersistentDataStoresOptions{ ProjectID: core.StringPtr(pid) }
    opts := ce.NewListPersistentDataStoresOptions(pid)
    ```

- **Pager types and factories renamed**
    Switch to the new pager names and factories:

    ```go
    // before
    pager, err := ce.NewAllowedOutboundDestinationPager(opts)

    // after
    pager, err := ce.NewAllowedOutboundDestinationsPager(opts)
    ```

    ```go
    // before
    pager, err := ce.NewPersistentDataStorePager(opts)

    // after
    pager, err := ce.NewPersistentDataStoresPager(opts)
    ```

- **Constructor signature changed for CIDR prototype**
    `NewAllowedOutboundDestinationPrototypeCidrBlockDataPrototype` parameter order changed:

    ```go
    // before: (type, cidrBlock, name)
    proto, err := ce.NewAllowedOutboundDestinationPrototypeCidrBlockDataPrototype(
      "cidr_block", "10.0.0.0/24", "allow-vpc-egress",
    )

    // after: (type, name, cidrBlock)
    proto, err := ce.NewAllowedOutboundDestinationPrototypeCidrBlockDataPrototype(
      "cidr_block", "allow-vpc-egress", "10.0.0.0/24",
    )
    ```

- **Allowed outbound destination patch payloads changed**
    Do **not** send `type` in patch payloads anymore. Use the specific fields instead:

    ```go
    // CIDR patch (remove Type; only patch fields that apply)
    patch := &codeenginev2.AllowedOutboundDestinationPatchCidrBlockDataPatch{
      CidrBlock: core.StringPtr("10.0.1.0/24"),
    }

    // Private Path service gateway patch (new)
    patch := &codeenginev2.AllowedOutboundDestinationPatchPrivatePathServiceGatewayDataPatch{
      IsolationPolicy: core.StringPtr(codeenginev2.AllowedOutboundDestinationPatchPrivatePathServiceGatewayDataPatch_IsolationPolicy_Dedicated),
    }
    ```

- **Allowed outbound destination prototypes now require `name`**
    The base prototype `AllowedOutboundDestinationPrototype` adds a required `Name`. Ensure you set it for all create flows (already required for the CIDR-specific prototype, but now also applicable at the base type).
    For Private Path service gateway, use the new prototype:

    ```go
    // New: Private Path prototype
    proto, err := ce.NewAllowedOutboundDestinationPrototypePrivatePathServiceGatewayDataPrototype(
      "private_path_service_gateway",
      "pps-to-service-x",
      "<private-path-service-gateway-crn>",
    )
    // Optional: set isolation policy
    proto.IsolationPolicy = core.StringPtr(
      codeenginev2.AllowedOutboundDestinationPrototypePrivatePathServiceGatewayDataPrototype_IsolationPolicy_Shared,
    )
    ```

- **`Probe.Type` is now required**
    When providing probes (e.g., in app/job templates), you must set the protocol:

    ```go
    // before
    probe := &codeenginev2.Probe{ InitialDelay: core.Int64Ptr(5) }

    // after (Type required)
    probe := &codeenginev2.Probe{
      InitialDelay: core.Int64Ptr(5),
      Type: core.StringPtr(codeenginev2.Probe_Type_Http), // or Probe_Type_Tcp
    }
    ```

- **`Secret.Format` is now required on `Secret`**
    If you construct `Secret` models in requests (e.g., replace flows), ensure `Format` is set:

    ```go
    secret := &codeenginev2.Secret{
      EntityTag: core.StringPtr(etag),
      Format:    core.StringPtr(codeenginev2.Secret_Format_Generic), // example; set appropriate format
      // ...
    }
    ```

- **Behavioral changes in response models (nullable → required)**
    Several response fields are now marked `validate:"required"` (e.g., `ComputedEnvVariables`, `RunServiceAccount`, `RunBuildParams`, `SourceType`, `StrategySize`, `StrategyType`, `FunctionRuntimes`, etc.).
    While this mainly affects SDK-side validation of **requests**, if your code programmatically (re)uses these response structs as inputs to API calls, you must set the now-required fields before sending.

> **Action checklist:**
>
> - [ ] Rename the list methods, option types, and pagers as shown above.
> - [ ] Adjust constructor argument order for `NewAllowedOutboundDestinationPrototypeCidrBlockDataPrototype`.
> - [ ] Remove `type` from allowed-outbound-destination patch payloads; use specific patch models/fields.
> - [ ] Ensure `name` is provided when creating allowed outbound destinations.
> - [ ] If using probes in requests, set `Probe.Type`.
> - [ ] If constructing `Secret` in requests, set `Secret.Format`.
> - [ ] Include the new `private_path_service_gateway` type in any client-side branching/validation.

## Installation

The current version of this SDK: 6.0.0

There are a few different ways to download and install the Code Engine Go SDK project for use by your
Go application:

### `go get` command

Use this command to download and install the SDK to allow your Go application to
use it:

```sh
go get -u github.com/IBM/code-engine-go-sdk
```

### Go modules

If your application is using Go modules, you can add a suitable import to your
Go application, like this:

```go
import (
  "github.com/IBM/code-engine-go-sdk/codeenginev2"
)
```

then run `go mod tidy` to download and install the new dependency and update your Go application's `go.mod` file.

## Using the SDK

Examples are available [here](./example/v2/README.md).

For general SDK usage information, please see [this link](https://github.com/IBM/ibm-cloud-sdk-common/blob/master/README.md)

## Questions

If you are having difficulties using this SDK or have a question about the IBM Cloud services,
please ask a question at
[Stack Overflow](http://stackoverflow.com/questions/ask?tags=ibm-cloud).

## Issues

If you encounter an issue with the project, you are welcome to submit a [bug report](https://github.com/IBM/code-engine-go-sdk/issues).
Before that, please search for similar issues. It's possible that someone has already reported the problem.

## Open source @ IBM

Find more open source projects on the [IBM Github Page](http://ibm.github.io/)

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).

## License

This SDK project is released under the Apache 2.0 license.
The license's full text can be found in [LICENSE](LICENSE).
