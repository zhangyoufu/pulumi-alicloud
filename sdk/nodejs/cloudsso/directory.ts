// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a Cloud SSO Directory resource.
 *
 * For information about Cloud SSO Directory and how to use it, see [What is Directory](https://www.alibabacloud.com/help/doc-detail/263624.html).
 *
 * > **NOTE:** Available in v1.135.0+.
 *
 * > **NOTE:** Cloud SSO Only Support `cn-shanghai` And `us-west-1` Region
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.cloudsso.Directory("example", {
 *     directoryName: "example_value",
 * });
 * ```
 *
 * ## Import
 *
 * Cloud SSO Directory can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:cloudsso/directory:Directory example <id>
 * ```
 */
export class Directory extends pulumi.CustomResource {
    /**
     * Get an existing Directory resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DirectoryState, opts?: pulumi.CustomResourceOptions): Directory {
        return new Directory(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cloudsso/directory:Directory';

    /**
     * Returns true if the given object is an instance of Directory.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Directory {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Directory.__pulumiType;
    }

    /**
     * The name of the CloudSSO directory. The length is 2-64 characters, and it can contain lowercase letters, numbers, and dashes (-). It cannot start or end with a dash and cannot have two consecutive dashes. Need to be globally unique, and capitalization is not supported. Cannot start with `d-`.
     */
    public readonly directoryName!: pulumi.Output<string | undefined>;
    /**
     * The mfa authentication status. Valid values: `Enabled` or `Disabled`. Default to `Disabled`.
     */
    public readonly mfaAuthenticationStatus!: pulumi.Output<string>;
    /**
     * The saml identity provider configuration.
     */
    public readonly samlIdentityProviderConfiguration!: pulumi.Output<outputs.cloudsso.DirectorySamlIdentityProviderConfiguration>;
    /**
     * The scim synchronization status. Valid values: `Enabled` or `Disabled`. Default to `Disabled`.
     */
    public readonly scimSynchronizationStatus!: pulumi.Output<string>;

    /**
     * Create a Directory resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DirectoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DirectoryArgs | DirectoryState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DirectoryState | undefined;
            inputs["directoryName"] = state ? state.directoryName : undefined;
            inputs["mfaAuthenticationStatus"] = state ? state.mfaAuthenticationStatus : undefined;
            inputs["samlIdentityProviderConfiguration"] = state ? state.samlIdentityProviderConfiguration : undefined;
            inputs["scimSynchronizationStatus"] = state ? state.scimSynchronizationStatus : undefined;
        } else {
            const args = argsOrState as DirectoryArgs | undefined;
            inputs["directoryName"] = args ? args.directoryName : undefined;
            inputs["mfaAuthenticationStatus"] = args ? args.mfaAuthenticationStatus : undefined;
            inputs["samlIdentityProviderConfiguration"] = args ? args.samlIdentityProviderConfiguration : undefined;
            inputs["scimSynchronizationStatus"] = args ? args.scimSynchronizationStatus : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Directory.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Directory resources.
 */
export interface DirectoryState {
    /**
     * The name of the CloudSSO directory. The length is 2-64 characters, and it can contain lowercase letters, numbers, and dashes (-). It cannot start or end with a dash and cannot have two consecutive dashes. Need to be globally unique, and capitalization is not supported. Cannot start with `d-`.
     */
    readonly directoryName?: pulumi.Input<string>;
    /**
     * The mfa authentication status. Valid values: `Enabled` or `Disabled`. Default to `Disabled`.
     */
    readonly mfaAuthenticationStatus?: pulumi.Input<string>;
    /**
     * The saml identity provider configuration.
     */
    readonly samlIdentityProviderConfiguration?: pulumi.Input<inputs.cloudsso.DirectorySamlIdentityProviderConfiguration>;
    /**
     * The scim synchronization status. Valid values: `Enabled` or `Disabled`. Default to `Disabled`.
     */
    readonly scimSynchronizationStatus?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Directory resource.
 */
export interface DirectoryArgs {
    /**
     * The name of the CloudSSO directory. The length is 2-64 characters, and it can contain lowercase letters, numbers, and dashes (-). It cannot start or end with a dash and cannot have two consecutive dashes. Need to be globally unique, and capitalization is not supported. Cannot start with `d-`.
     */
    readonly directoryName?: pulumi.Input<string>;
    /**
     * The mfa authentication status. Valid values: `Enabled` or `Disabled`. Default to `Disabled`.
     */
    readonly mfaAuthenticationStatus?: pulumi.Input<string>;
    /**
     * The saml identity provider configuration.
     */
    readonly samlIdentityProviderConfiguration?: pulumi.Input<inputs.cloudsso.DirectorySamlIdentityProviderConfiguration>;
    /**
     * The scim synchronization status. Valid values: `Enabled` or `Disabled`. Default to `Disabled`.
     */
    readonly scimSynchronizationStatus?: pulumi.Input<string>;
}
