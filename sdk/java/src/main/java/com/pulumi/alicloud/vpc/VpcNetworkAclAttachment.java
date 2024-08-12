// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.vpc.VpcNetworkAclAttachmentArgs;
import com.pulumi.alicloud.vpc.inputs.VpcNetworkAclAttachmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a VPC Network Acl Attachment resource. Resources associated with network Acl.
 * 
 * For information about VPC Network Acl Attachment and how to use it, see [What is Network Acl Attachment](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/associatenetworkacl).
 * 
 * &gt; **NOTE:** Available since v1.193.0.
 * 
 * ## Import
 * 
 * VPC Network Acl Attachment can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:vpc/vpcNetworkAclAttachment:VpcNetworkAclAttachment example &lt;network_acl_id&gt;:&lt;resource_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:vpc/vpcNetworkAclAttachment:VpcNetworkAclAttachment")
public class VpcNetworkAclAttachment extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the network ACL.
     * 
     */
    @Export(name="networkAclId", refs={String.class}, tree="[0]")
    private Output<String> networkAclId;

    /**
     * @return The ID of the network ACL.
     * 
     */
    public Output<String> networkAclId() {
        return this.networkAclId;
    }
    /**
     * The ID of the associated resource.
     * 
     */
    @Export(name="resourceId", refs={String.class}, tree="[0]")
    private Output<String> resourceId;

    /**
     * @return The ID of the associated resource.
     * 
     */
    public Output<String> resourceId() {
        return this.resourceId;
    }
    /**
     * The type of the associated resource. Valid values: `VSwitch`.
     * 
     */
    @Export(name="resourceType", refs={String.class}, tree="[0]")
    private Output<String> resourceType;

    /**
     * @return The type of the associated resource. Valid values: `VSwitch`.
     * 
     */
    public Output<String> resourceType() {
        return this.resourceType;
    }
    /**
     * The status of the Network Acl Attachment.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the Network Acl Attachment.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VpcNetworkAclAttachment(java.lang.String name) {
        this(name, VpcNetworkAclAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VpcNetworkAclAttachment(java.lang.String name, VpcNetworkAclAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VpcNetworkAclAttachment(java.lang.String name, VpcNetworkAclAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/vpcNetworkAclAttachment:VpcNetworkAclAttachment", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private VpcNetworkAclAttachment(java.lang.String name, Output<java.lang.String> id, @Nullable VpcNetworkAclAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/vpcNetworkAclAttachment:VpcNetworkAclAttachment", name, state, makeResourceOptions(options, id), false);
    }

    private static VpcNetworkAclAttachmentArgs makeArgs(VpcNetworkAclAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? VpcNetworkAclAttachmentArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static VpcNetworkAclAttachment get(java.lang.String name, Output<java.lang.String> id, @Nullable VpcNetworkAclAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VpcNetworkAclAttachment(name, id, state, options);
    }
}
