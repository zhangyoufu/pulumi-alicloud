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
 * 
 * For information about VPC Network Acl and how to use it, see [What is Network Acl](https://www.alibabacloud.com/help/en/ens/latest/createnetworkacl).
 * 
 * &gt; **NOTE:** Available since v1.43.0.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.vpc.NetworkAcl;
 * import com.pulumi.alicloud.vpc.NetworkAclArgs;
 * import com.pulumi.alicloud.vpc.inputs.NetworkAclIngressAclEntryArgs;
 * import com.pulumi.alicloud.vpc.inputs.NetworkAclEgressAclEntryArgs;
 * import com.pulumi.alicloud.vpc.inputs.NetworkAclResourceArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var config = ctx.config();
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;tf-example&#34;);
 *         final var default = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(&#34;VSwitch&#34;)
 *             .build());
 * 
 *         var exampleNetwork = new Network(&#34;exampleNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(name)
 *             .cidrBlock(&#34;10.4.0.0/16&#34;)
 *             .build());
 * 
 *         var exampleSwitch = new Switch(&#34;exampleSwitch&#34;, SwitchArgs.builder()        
 *             .vswitchName(name)
 *             .cidrBlock(&#34;10.4.0.0/24&#34;)
 *             .vpcId(exampleNetwork.id())
 *             .zoneId(default_.zones()[0].id())
 *             .build());
 * 
 *         var exampleNetworkAcl = new NetworkAcl(&#34;exampleNetworkAcl&#34;, NetworkAclArgs.builder()        
 *             .vpcId(exampleNetwork.id())
 *             .networkAclName(name)
 *             .description(name)
 *             .ingressAclEntries(NetworkAclIngressAclEntryArgs.builder()
 *                 .description(String.format(&#34;%s-ingress&#34;, name))
 *                 .networkAclEntryName(String.format(&#34;%s-ingress&#34;, name))
 *                 .sourceCidrIp(&#34;10.0.0.0/24&#34;)
 *                 .policy(&#34;accept&#34;)
 *                 .port(&#34;20/80&#34;)
 *                 .protocol(&#34;tcp&#34;)
 *                 .build())
 *             .egressAclEntries(NetworkAclEgressAclEntryArgs.builder()
 *                 .description(String.format(&#34;%s-egress&#34;, name))
 *                 .networkAclEntryName(String.format(&#34;%s-egress&#34;, name))
 *                 .destinationCidrIp(&#34;10.0.0.0/24&#34;)
 *                 .policy(&#34;accept&#34;)
 *                 .port(&#34;20/80&#34;)
 *                 .protocol(&#34;tcp&#34;)
 *                 .build())
 *             .resources(NetworkAclResourceArgs.builder()
 *                 .resourceId(exampleSwitch.id())
 *                 .resourceType(&#34;VSwitch&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * VPC Network Acl can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:vpc/networkAcl:NetworkAcl example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:vpc/networkAcl:NetworkAcl")
public class NetworkAcl extends com.pulumi.resources.CustomResource {
    /**
     * The creation time of the resource.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
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
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the network ACL.The description must be 1 to 256 characters in length and cannot start with http:// or https.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Out direction rule information. See `egress_acl_entries` below.
     * 
     */
    @Export(name="egressAclEntries", refs={List.class,NetworkAclEgressAclEntry.class}, tree="[0,1]")
    private Output<List<NetworkAclEgressAclEntry>> egressAclEntries;

    /**
     * @return Out direction rule information. See `egress_acl_entries` below.
     * 
     */
    public Output<List<NetworkAclEgressAclEntry>> egressAclEntries() {
        return this.egressAclEntries;
    }
    /**
     * Inward direction rule information. See `ingress_acl_entries` below.
     * 
     */
    @Export(name="ingressAclEntries", refs={List.class,NetworkAclIngressAclEntry.class}, tree="[0,1]")
    private Output<List<NetworkAclIngressAclEntry>> ingressAclEntries;

    /**
     * @return Inward direction rule information. See `ingress_acl_entries` below.
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
    @Export(name="name", refs={String.class}, tree="[0]")
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
    @Export(name="networkAclName", refs={String.class}, tree="[0]")
    private Output<String> networkAclName;

    /**
     * @return The name of the network ACL.The name must be 1 to 128 characters in length and cannot start with http:// or https.
     * 
     */
    public Output<String> networkAclName() {
        return this.networkAclName;
    }
    /**
     * The associated resource. See `resources` below.
     * 
     */
    @Export(name="resources", refs={List.class,NetworkAclResource.class}, tree="[0,1]")
    private Output<List<NetworkAclResource>> resources;

    /**
     * @return The associated resource. See `resources` below.
     * 
     */
    public Output<List<NetworkAclResource>> resources() {
        return this.resources;
    }
    /**
     * The status of the associated resource.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the associated resource.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The tags of this resource.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
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
    @Export(name="vpcId", refs={String.class}, tree="[0]")
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
