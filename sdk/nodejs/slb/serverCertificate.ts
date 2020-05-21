// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A Load Balancer Server Certificate is an ssl Certificate used by the listener of the protocol https.
 *
 * For information about slb and how to use it, see [What is Server Load Balancer](https://www.alibabacloud.com/help/doc-detail/27539.htm).
 *
 * For information about Server Certificate and how to use it, see [Configure Server Certificate](https://www.alibabacloud.com/help/doc-detail/85968.htm).
 *
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * // create a server certificate
 * const foo = new alicloud.slb.ServerCertificate("foo", {
 *     privateKey: `-----BEGIN RSA PRIVATE KEY-----
 * MIICXAIBAAKBgQDO0knDrlNdiys******ErVpjsckAaOW/JDG5PCSwkaMxk=
 * -----END RSA PRIVATE KEY-----`,
 *     serverCertificate: `-----BEGIN CERTIFICATE-----
 * MIIDRjCCAq+gAwIBAgI+OuMs******XTtI90EAxEG/bJJyOm5LqoiA=
 * -----END CERTIFICATE-----`,
 * });
 * ```
 */
export class ServerCertificate extends pulumi.CustomResource {
    /**
     * Get an existing ServerCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServerCertificateState, opts?: pulumi.CustomResourceOptions): ServerCertificate {
        return new ServerCertificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:slb/serverCertificate:ServerCertificate';

    /**
     * Returns true if the given object is an instance of ServerCertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServerCertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServerCertificate.__pulumiType;
    }

    public readonly alicloudCertifacteId!: pulumi.Output<string | undefined>;
    public readonly alicloudCertifacteName!: pulumi.Output<string | undefined>;
    /**
     * an id of server certificate ssued/proxied by alibaba cloud. but it is not supported on the international site of alibaba cloud now.
     */
    public readonly alicloudCertificateId!: pulumi.Output<string | undefined>;
    /**
     * the name of the certificate specified by `alicloudCertificateId`.but it is not supported on the international site of alibaba cloud now.
     */
    public readonly alicloudCertificateName!: pulumi.Output<string | undefined>;
    /**
     * the region of the certificate specified by `alicloudCertificateId`. but it is not supported on the international site of alibaba cloud now.
     */
    public readonly alicloudCertificateRegionId!: pulumi.Output<string | undefined>;
    /**
     * Name of the Server Certificate.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * the content of privat key of the ssl certificate specified by `serverCertificate`. where `alicloudCertificateId` is null, it is required, otherwise it is ignored.
     */
    public readonly privateKey!: pulumi.Output<string | undefined>;
    /**
     * The Id of resource group which the slb server certificate belongs.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * the content of the ssl certificate. where `alicloudCertificateId` is null, it is required, otherwise it is ignored.
     */
    public readonly serverCertificate!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a ServerCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ServerCertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServerCertificateArgs | ServerCertificateState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ServerCertificateState | undefined;
            inputs["alicloudCertifacteId"] = state ? state.alicloudCertifacteId : undefined;
            inputs["alicloudCertifacteName"] = state ? state.alicloudCertifacteName : undefined;
            inputs["alicloudCertificateId"] = state ? state.alicloudCertificateId : undefined;
            inputs["alicloudCertificateName"] = state ? state.alicloudCertificateName : undefined;
            inputs["alicloudCertificateRegionId"] = state ? state.alicloudCertificateRegionId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["privateKey"] = state ? state.privateKey : undefined;
            inputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            inputs["serverCertificate"] = state ? state.serverCertificate : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ServerCertificateArgs | undefined;
            inputs["alicloudCertifacteId"] = args ? args.alicloudCertifacteId : undefined;
            inputs["alicloudCertifacteName"] = args ? args.alicloudCertifacteName : undefined;
            inputs["alicloudCertificateId"] = args ? args.alicloudCertificateId : undefined;
            inputs["alicloudCertificateName"] = args ? args.alicloudCertificateName : undefined;
            inputs["alicloudCertificateRegionId"] = args ? args.alicloudCertificateRegionId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["privateKey"] = args ? args.privateKey : undefined;
            inputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            inputs["serverCertificate"] = args ? args.serverCertificate : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ServerCertificate.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServerCertificate resources.
 */
export interface ServerCertificateState {
    readonly alicloudCertifacteId?: pulumi.Input<string>;
    readonly alicloudCertifacteName?: pulumi.Input<string>;
    /**
     * an id of server certificate ssued/proxied by alibaba cloud. but it is not supported on the international site of alibaba cloud now.
     */
    readonly alicloudCertificateId?: pulumi.Input<string>;
    /**
     * the name of the certificate specified by `alicloudCertificateId`.but it is not supported on the international site of alibaba cloud now.
     */
    readonly alicloudCertificateName?: pulumi.Input<string>;
    /**
     * the region of the certificate specified by `alicloudCertificateId`. but it is not supported on the international site of alibaba cloud now.
     */
    readonly alicloudCertificateRegionId?: pulumi.Input<string>;
    /**
     * Name of the Server Certificate.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * the content of privat key of the ssl certificate specified by `serverCertificate`. where `alicloudCertificateId` is null, it is required, otherwise it is ignored.
     */
    readonly privateKey?: pulumi.Input<string>;
    /**
     * The Id of resource group which the slb server certificate belongs.
     */
    readonly resourceGroupId?: pulumi.Input<string>;
    /**
     * the content of the ssl certificate. where `alicloudCertificateId` is null, it is required, otherwise it is ignored.
     */
    readonly serverCertificate?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a ServerCertificate resource.
 */
export interface ServerCertificateArgs {
    readonly alicloudCertifacteId?: pulumi.Input<string>;
    readonly alicloudCertifacteName?: pulumi.Input<string>;
    /**
     * an id of server certificate ssued/proxied by alibaba cloud. but it is not supported on the international site of alibaba cloud now.
     */
    readonly alicloudCertificateId?: pulumi.Input<string>;
    /**
     * the name of the certificate specified by `alicloudCertificateId`.but it is not supported on the international site of alibaba cloud now.
     */
    readonly alicloudCertificateName?: pulumi.Input<string>;
    /**
     * the region of the certificate specified by `alicloudCertificateId`. but it is not supported on the international site of alibaba cloud now.
     */
    readonly alicloudCertificateRegionId?: pulumi.Input<string>;
    /**
     * Name of the Server Certificate.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * the content of privat key of the ssl certificate specified by `serverCertificate`. where `alicloudCertificateId` is null, it is required, otherwise it is ignored.
     */
    readonly privateKey?: pulumi.Input<string>;
    /**
     * The Id of resource group which the slb server certificate belongs.
     */
    readonly resourceGroupId?: pulumi.Input<string>;
    /**
     * the content of the ssl certificate. where `alicloudCertificateId` is null, it is required, otherwise it is ignored.
     */
    readonly serverCertificate?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
