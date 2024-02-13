// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.vpc.Ipv6AddressArgs;
import com.pulumi.alicloud.vpc.inputs.Ipv6AddressState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a VPC Ipv6 Address resource.
 * 
 * For information about VPC Ipv6 Address and how to use it, see [What is Ipv6 Address](https://next.api.alibabacloud.com/document/Vpc/2016-04-28/AllocateIpv6Address).
 * 
 * &gt; **NOTE:** Available since v1.216.0.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.resourcemanager.ResourcemanagerFunctions;
 * import com.pulumi.alicloud.resourcemanager.inputs.GetResourceGroupsArgs;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.vpc.Ipv6Address;
 * import com.pulumi.alicloud.vpc.Ipv6AddressArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;terraform-example&#34;);
 *         final var defaultResourceGroups = ResourcemanagerFunctions.getResourceGroups();
 * 
 *         final var defaultZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(&#34;VSwitch&#34;)
 *             .build());
 * 
 *         var vpc = new Network(&#34;vpc&#34;, NetworkArgs.builder()        
 *             .ipv6Isp(&#34;BGP&#34;)
 *             .cidrBlock(&#34;172.168.0.0/16&#34;)
 *             .enableIpv6(true)
 *             .vpcName(name)
 *             .build());
 * 
 *         var vswich = new Switch(&#34;vswich&#34;, SwitchArgs.builder()        
 *             .vpcId(vpc.id())
 *             .cidrBlock(&#34;172.168.0.0/24&#34;)
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .vswitchName(name)
 *             .ipv6CidrBlockMask(&#34;1&#34;)
 *             .build());
 * 
 *         var defaultIpv6Address = new Ipv6Address(&#34;defaultIpv6Address&#34;, Ipv6AddressArgs.builder()        
 *             .resourceGroupId(defaultResourceGroups.applyValue(getResourceGroupsResult -&gt; getResourceGroupsResult.ids()[0]))
 *             .vswitchId(vswich.id())
 *             .ipv6AddressDescription(&#34;create_description&#34;)
 *             .ipv6AddressName(name)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * VPC Ipv6 Address can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:vpc/ipv6Address:Ipv6Address example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:vpc/ipv6Address:Ipv6Address")
public class Ipv6Address extends com.pulumi.resources.CustomResource {
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
     * IPv6 address.
     * 
     */
    @Export(name="ipv6Address", refs={String.class}, tree="[0]")
    private Output<String> ipv6Address;

    /**
     * @return IPv6 address.
     * 
     */
    public Output<String> ipv6Address() {
        return this.ipv6Address;
    }
    /**
     * The description of the IPv6 Address. The description must be 2 to 256 characters in length. It cannot start with http:// or https://.
     * 
     */
    @Export(name="ipv6AddressDescription", refs={String.class}, tree="[0]")
    private Output<String> ipv6AddressDescription;

    /**
     * @return The description of the IPv6 Address. The description must be 2 to 256 characters in length. It cannot start with http:// or https://.
     * 
     */
    public Output<String> ipv6AddressDescription() {
        return this.ipv6AddressDescription;
    }
    /**
     * The name of the IPv6 Address. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with http:// or https://.
     * 
     */
    @Export(name="ipv6AddressName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ipv6AddressName;

    /**
     * @return The name of the IPv6 Address. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with http:// or https://.
     * 
     */
    public Output<Optional<String>> ipv6AddressName() {
        return Codegen.optional(this.ipv6AddressName);
    }
    /**
     * The ID of the resource group to which the instance belongs.
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group to which the instance belongs.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * The status of the resource.  Available, Pending and Deleting.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the resource.  Available, Pending and Deleting.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The tags for the resource.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return The tags for the resource.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The VSwitchId of the IPv6 address.
     * 
     */
    @Export(name="vswitchId", refs={String.class}, tree="[0]")
    private Output<String> vswitchId;

    /**
     * @return The VSwitchId of the IPv6 address.
     * 
     */
    public Output<String> vswitchId() {
        return this.vswitchId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Ipv6Address(String name) {
        this(name, Ipv6AddressArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Ipv6Address(String name, Ipv6AddressArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Ipv6Address(String name, Ipv6AddressArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/ipv6Address:Ipv6Address", name, args == null ? Ipv6AddressArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Ipv6Address(String name, Output<String> id, @Nullable Ipv6AddressState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/ipv6Address:Ipv6Address", name, state, makeResourceOptions(options, id));
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
    public static Ipv6Address get(String name, Output<String> id, @Nullable Ipv6AddressState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Ipv6Address(name, id, state, options);
    }
}
