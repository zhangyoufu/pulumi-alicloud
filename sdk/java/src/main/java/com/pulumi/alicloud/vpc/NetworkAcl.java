// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.vpc.NetworkAclArgs;
import com.pulumi.alicloud.vpc.inputs.NetworkAclState;
import com.pulumi.alicloud.vpc.outputs.NetworkAclEgressAclEntry;
import com.pulumi.alicloud.vpc.outputs.NetworkAclIngressAclEntry;
import com.pulumi.alicloud.vpc.outputs.NetworkAclResource;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a VPC Network Acl resource.
 * &gt; **NOTE:** Currently, the resource are only available in Hongkong(cn-hongkong), India(ap-south-1), and Indonesia(ap-southeast-1) regions.
 * 
 * For information about VPC Network Acl and how to use it, see [What is Network Acl](https://www.alibabacloud.com/help/en/ens/latest/createnetworkacl).
 * 
 * &gt; **NOTE:** Available in v1.43.0+.
 * 
 * ## Import
 * 
 * VPC Network Acl can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:vpc/networkAcl:NetworkAcl example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:vpc/networkAcl:NetworkAcl")
public class NetworkAcl extends com.pulumi.resources.CustomResource {
    /**
     * The creation time of the resource.
     * 
     */
    @Export(name="createTime", type=String.class, parameters={})
    private Output<String> createTime;

    /**
     * @return The creation time of the resource.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * The description of the network ACL.The description must be 1 to 256 characters in length and cannot start with http:// or https.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the network ACL.The description must be 1 to 256 characters in length and cannot start with http:// or https.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Out direction rule information. See the following `Block EgressAclEntries`.
     * 
     */
    @Export(name="egressAclEntries", type=List.class, parameters={NetworkAclEgressAclEntry.class})
    private Output<List<NetworkAclEgressAclEntry>> egressAclEntries;

    /**
     * @return Out direction rule information. See the following `Block EgressAclEntries`.
     * 
     */
    public Output<List<NetworkAclEgressAclEntry>> egressAclEntries() {
        return this.egressAclEntries;
    }
    /**
     * Inward direction rule information. See the following `Block IngressAclEntries`.
     * 
     */
    @Export(name="ingressAclEntries", type=List.class, parameters={NetworkAclIngressAclEntry.class})
    private Output<List<NetworkAclIngressAclEntry>> ingressAclEntries;

    /**
     * @return Inward direction rule information. See the following `Block IngressAclEntries`.
     * 
     */
    public Output<List<NetworkAclIngressAclEntry>> ingressAclEntries() {
        return this.ingressAclEntries;
    }
    /**
     * Field &#39;name&#39; has been deprecated from provider version 1.122.0. New field &#39;network_acl_name&#39; instead.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.122.0. New field &#39;network_acl_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.122.0. New field 'network_acl_name' instead. */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Field &#39;name&#39; has been deprecated from provider version 1.122.0. New field &#39;network_acl_name&#39; instead.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The name of the network ACL.The name must be 1 to 128 characters in length and cannot start with http:// or https.
     * 
     */
    @Export(name="networkAclName", type=String.class, parameters={})
    private Output<String> networkAclName;

    /**
     * @return The name of the network ACL.The name must be 1 to 128 characters in length and cannot start with http:// or https.
     * 
     */
    public Output<String> networkAclName() {
        return this.networkAclName;
    }
    /**
     * The associated resource. See the following `Block Resources`.
     * 
     */
    @Export(name="resources", type=List.class, parameters={NetworkAclResource.class})
    private Output<List<NetworkAclResource>> resources;

    /**
     * @return The associated resource. See the following `Block Resources`.
     * 
     */
    public Output<List<NetworkAclResource>> resources() {
        return this.resources;
    }
    /**
     * The state of the network ACL.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The state of the network ACL.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The tags of this resource.
     * 
     */
    @Export(name="tags", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return The tags of this resource.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The ID of the associated VPC.
     * 
     * The following arguments will be discarded. Please use new fields as soon as possible:
     * 
     */
    @Export(name="vpcId", type=String.class, parameters={})
    private Output<String> vpcId;

    /**
     * @return The ID of the associated VPC.
     * 
     * The following arguments will be discarded. Please use new fields as soon as possible:
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NetworkAcl(String name) {
        this(name, NetworkAclArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NetworkAcl(String name, NetworkAclArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NetworkAcl(String name, NetworkAclArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/networkAcl:NetworkAcl", name, args == null ? NetworkAclArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private NetworkAcl(String name, Output<String> id, @Nullable NetworkAclState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/networkAcl:NetworkAcl", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static NetworkAcl get(String name, Output<String> id, @Nullable NetworkAclState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NetworkAcl(name, id, state, options);
    }
}
