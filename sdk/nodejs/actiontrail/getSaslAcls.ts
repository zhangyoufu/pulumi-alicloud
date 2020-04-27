// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list of ALIKAFKA Sasl acls in an Alibaba Cloud account according to the specified filters.
 * 
 * > **NOTE:** Available in 1.66.0+
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const saslAclsDs = pulumi.output(alicloud.actiontrail.getSaslAcls({
 *     aclResourceName: "testTopic",
 *     aclResourceType: "Topic",
 *     instanceId: "xxx",
 *     outputFile: "saslAcls.txt",
 *     username: "username",
 * }, { async: true }));
 * 
 * export const firstSaslAclUsername = saslAclsDs.acls[0].username;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/alikafka_sasl_acls.html.markdown.
 */
export function getSaslAcls(args: GetSaslAclsArgs, opts?: pulumi.InvokeOptions): Promise<GetSaslAclsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:actiontrail/getSaslAcls:getSaslAcls", {
        "aclResourceName": args.aclResourceName,
        "aclResourceType": args.aclResourceType,
        "instanceId": args.instanceId,
        "outputFile": args.outputFile,
        "username": args.username,
    }, opts);
}

/**
 * A collection of arguments for invoking getSaslAcls.
 */
export interface GetSaslAclsArgs {
    /**
     * Get results for the specified resource name. 
     */
    readonly aclResourceName: string;
    /**
     * Get results for the specified resource type. 
     */
    readonly aclResourceType: string;
    /**
     * ID of the ALIKAFKA Instance that owns the sasl acls.
     */
    readonly instanceId: string;
    readonly outputFile?: string;
    /**
     * Get results for the specified username. 
     */
    readonly username: string;
}

/**
 * A collection of values returned by getSaslAcls.
 */
export interface GetSaslAclsResult {
    /**
     * The resource name of the sasl acl.
     */
    readonly aclResourceName: string;
    /**
     * The resource type of the sasl acl.
     */
    readonly aclResourceType: string;
    /**
     * A list of sasl acls. Each element contains the following attributes:
     */
    readonly acls: outputs.actiontrail.GetSaslAclsAcl[];
    readonly instanceId: string;
    readonly outputFile?: string;
    /**
     * The username of the sasl acl.
     */
    readonly username: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
