// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export class OpenApiImageCache extends pulumi.CustomResource {
    /**
     * Get an existing OpenApiImageCache resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OpenApiImageCacheState, opts?: pulumi.CustomResourceOptions): OpenApiImageCache {
        return new OpenApiImageCache(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:eci/openApiImageCache:OpenApiImageCache';

    /**
     * Returns true if the given object is an instance of OpenApiImageCache.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OpenApiImageCache {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OpenApiImageCache.__pulumiType;
    }

    public /*out*/ readonly containerGroupId!: pulumi.Output<string>;
    public readonly eipInstanceId!: pulumi.Output<string | undefined>;
    public readonly imageCacheName!: pulumi.Output<string>;
    public readonly imageCacheSize!: pulumi.Output<number | undefined>;
    public readonly imageRegistryCredentials!: pulumi.Output<outputs.eci.OpenApiImageCacheImageRegistryCredential[] | undefined>;
    public readonly images!: pulumi.Output<string[]>;
    public readonly resourceGroupId!: pulumi.Output<string | undefined>;
    public readonly retentionDays!: pulumi.Output<number | undefined>;
    public readonly securityGroupId!: pulumi.Output<string>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    public readonly vswitchId!: pulumi.Output<string>;
    public readonly zoneId!: pulumi.Output<string | undefined>;

    /**
     * Create a OpenApiImageCache resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OpenApiImageCacheArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OpenApiImageCacheArgs | OpenApiImageCacheState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OpenApiImageCacheState | undefined;
            inputs["containerGroupId"] = state ? state.containerGroupId : undefined;
            inputs["eipInstanceId"] = state ? state.eipInstanceId : undefined;
            inputs["imageCacheName"] = state ? state.imageCacheName : undefined;
            inputs["imageCacheSize"] = state ? state.imageCacheSize : undefined;
            inputs["imageRegistryCredentials"] = state ? state.imageRegistryCredentials : undefined;
            inputs["images"] = state ? state.images : undefined;
            inputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            inputs["retentionDays"] = state ? state.retentionDays : undefined;
            inputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["vswitchId"] = state ? state.vswitchId : undefined;
            inputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as OpenApiImageCacheArgs | undefined;
            if ((!args || args.imageCacheName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'imageCacheName'");
            }
            if ((!args || args.images === undefined) && !opts.urn) {
                throw new Error("Missing required property 'images'");
            }
            if ((!args || args.securityGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroupId'");
            }
            if ((!args || args.vswitchId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vswitchId'");
            }
            inputs["eipInstanceId"] = args ? args.eipInstanceId : undefined;
            inputs["imageCacheName"] = args ? args.imageCacheName : undefined;
            inputs["imageCacheSize"] = args ? args.imageCacheSize : undefined;
            inputs["imageRegistryCredentials"] = args ? args.imageRegistryCredentials : undefined;
            inputs["images"] = args ? args.images : undefined;
            inputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            inputs["retentionDays"] = args ? args.retentionDays : undefined;
            inputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            inputs["vswitchId"] = args ? args.vswitchId : undefined;
            inputs["zoneId"] = args ? args.zoneId : undefined;
            inputs["containerGroupId"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(OpenApiImageCache.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OpenApiImageCache resources.
 */
export interface OpenApiImageCacheState {
    readonly containerGroupId?: pulumi.Input<string>;
    readonly eipInstanceId?: pulumi.Input<string>;
    readonly imageCacheName?: pulumi.Input<string>;
    readonly imageCacheSize?: pulumi.Input<number>;
    readonly imageRegistryCredentials?: pulumi.Input<pulumi.Input<inputs.eci.OpenApiImageCacheImageRegistryCredential>[]>;
    readonly images?: pulumi.Input<pulumi.Input<string>[]>;
    readonly resourceGroupId?: pulumi.Input<string>;
    readonly retentionDays?: pulumi.Input<number>;
    readonly securityGroupId?: pulumi.Input<string>;
    readonly status?: pulumi.Input<string>;
    readonly vswitchId?: pulumi.Input<string>;
    readonly zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OpenApiImageCache resource.
 */
export interface OpenApiImageCacheArgs {
    readonly eipInstanceId?: pulumi.Input<string>;
    readonly imageCacheName: pulumi.Input<string>;
    readonly imageCacheSize?: pulumi.Input<number>;
    readonly imageRegistryCredentials?: pulumi.Input<pulumi.Input<inputs.eci.OpenApiImageCacheImageRegistryCredential>[]>;
    readonly images: pulumi.Input<pulumi.Input<string>[]>;
    readonly resourceGroupId?: pulumi.Input<string>;
    readonly retentionDays?: pulumi.Input<number>;
    readonly securityGroupId: pulumi.Input<string>;
    readonly vswitchId: pulumi.Input<string>;
    readonly zoneId?: pulumi.Input<string>;
}
