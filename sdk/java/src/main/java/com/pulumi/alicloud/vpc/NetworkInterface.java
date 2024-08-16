// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.vpc.NetworkInterfaceArgs;
import com.pulumi.alicloud.vpc.inputs.NetworkInterfaceState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * &gt; **DEPRECATED:** This resource has been renamed to alicloud.ecs.EcsNetworkInterface from version 1.123.1.
 * 
 * Provides an ECS Elastic Network Interface resource.
 * 
 * For information about Elastic Network Interface and how to use it, see [Elastic Network Interface](https://www.alibabacloud.com/help/doc-detail/58496.html).
 * 
 * &gt; **NOTE** Only one of private_ips or private_ips_count can be specified when assign private IPs.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.ecs.SecurityGroup;
 * import com.pulumi.alicloud.ecs.SecurityGroupArgs;
 * import com.pulumi.alicloud.vpc.NetworkInterface;
 * import com.pulumi.alicloud.vpc.NetworkInterfaceArgs;
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
 *         final var name = config.get("name").orElse("networkInterfaceName");
 *         var vpc = new Network("vpc", NetworkArgs.builder()
 *             .vpcName(name)
 *             .cidrBlock("192.168.0.0/24")
 *             .build());
 * 
 *         final var default = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation("VSwitch")
 *             .build());
 * 
 *         var vswitch = new Switch("vswitch", SwitchArgs.builder()
 *             .name(name)
 *             .cidrBlock("192.168.0.0/24")
 *             .zoneId(default_.zones()[0].id())
 *             .vpcId(vpc.id())
 *             .build());
 * 
 *         var group = new SecurityGroup("group", SecurityGroupArgs.builder()
 *             .name(name)
 *             .vpcId(vpc.id())
 *             .build());
 * 
 *         var defaultNetworkInterface = new NetworkInterface("defaultNetworkInterface", NetworkInterfaceArgs.builder()
 *             .networkInterfaceName(name)
 *             .vswitchId(vswitch.id())
 *             .securityGroupIds(group.id())
 *             .privateIp("192.168.0.2")
 *             .privateIpsCount(3)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * ENI can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:vpc/networkInterface:NetworkInterface eni eni-abc1234567890000
 * ```
 * 
 */
@ResourceType(type="alicloud:vpc/networkInterface:NetworkInterface")
public class NetworkInterface extends com.pulumi.resources.CustomResource {
    /**
     * Description of the ENI. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the ENI. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    @Export(name="instanceType", refs={String.class}, tree="[0]")
    private Output<String> instanceType;

    public Output<String> instanceType() {
        return this.instanceType;
    }
    @Export(name="ipv4PrefixCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> ipv4PrefixCount;

    public Output<Integer> ipv4PrefixCount() {
        return this.ipv4PrefixCount;
    }
    @Export(name="ipv4Prefixes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> ipv4Prefixes;

    public Output<List<String>> ipv4Prefixes() {
        return this.ipv4Prefixes;
    }
    @Export(name="ipv6AddressCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> ipv6AddressCount;

    public Output<Integer> ipv6AddressCount() {
        return this.ipv6AddressCount;
    }
    @Export(name="ipv6Addresses", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> ipv6Addresses;

    public Output<List<String>> ipv6Addresses() {
        return this.ipv6Addresses;
    }
    /**
     * (Available in 1.54.0+) The MAC address of an ENI.
     * 
     */
    @Export(name="mac", refs={String.class}, tree="[0]")
    private Output<String> mac;

    /**
     * @return (Available in 1.54.0+) The MAC address of an ENI.
     * 
     */
    public Output<String> mac() {
        return this.mac;
    }
    /**
     * Name of the ENI. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;, &#34;.&#34;, &#34;_&#34;, and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.123.1. New field &#39;network_interface_name&#39; instead
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the ENI. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;, &#34;.&#34;, &#34;_&#34;, and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    @Export(name="networkInterfaceName", refs={String.class}, tree="[0]")
    private Output<String> networkInterfaceName;

    public Output<String> networkInterfaceName() {
        return this.networkInterfaceName;
    }
    @Export(name="networkInterfaceTrafficMode", refs={String.class}, tree="[0]")
    private Output<String> networkInterfaceTrafficMode;

    public Output<String> networkInterfaceTrafficMode() {
        return this.networkInterfaceTrafficMode;
    }
    @Export(name="primaryIpAddress", refs={String.class}, tree="[0]")
    private Output<String> primaryIpAddress;

    public Output<String> primaryIpAddress() {
        return this.primaryIpAddress;
    }
    /**
     * The primary private IP of the ENI.
     * 
     * @deprecated
     * Field &#39;private_ip&#39; has been deprecated from provider version 1.123.1. New field &#39;primary_ip_address&#39; instead
     * 
     */
    @Deprecated /* Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead */
    @Export(name="privateIp", refs={String.class}, tree="[0]")
    private Output<String> privateIp;

    /**
     * @return The primary private IP of the ENI.
     * 
     */
    public Output<String> privateIp() {
        return this.privateIp;
    }
    @Export(name="privateIpAddresses", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> privateIpAddresses;

    public Output<List<String>> privateIpAddresses() {
        return this.privateIpAddresses;
    }
    /**
     * List of secondary private IPs to assign to the ENI. Don&#39;t use both private_ips and private_ips_count in the same ENI resource block.
     * 
     * @deprecated
     * Field &#39;private_ips&#39; has been deprecated from provider version 1.123.1. New field &#39;private_ip_addresses&#39; instead
     * 
     */
    @Deprecated /* Field 'private_ips' has been deprecated from provider version 1.123.1. New field 'private_ip_addresses' instead */
    @Export(name="privateIps", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> privateIps;

    /**
     * @return List of secondary private IPs to assign to the ENI. Don&#39;t use both private_ips and private_ips_count in the same ENI resource block.
     * 
     */
    public Output<List<String>> privateIps() {
        return this.privateIps;
    }
    /**
     * Number of secondary private IPs to assign to the ENI. Don&#39;t use both private_ips and private_ips_count in the same ENI resource block.
     * 
     * @deprecated
     * Field &#39;private_ips_count&#39; has been deprecated from provider version 1.123.1. New field &#39;secondary_private_ip_address_count&#39; instead
     * 
     */
    @Deprecated /* Field 'private_ips_count' has been deprecated from provider version 1.123.1. New field 'secondary_private_ip_address_count' instead */
    @Export(name="privateIpsCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> privateIpsCount;

    /**
     * @return Number of secondary private IPs to assign to the ENI. Don&#39;t use both private_ips and private_ips_count in the same ENI resource block.
     * 
     */
    public Output<Integer> privateIpsCount() {
        return this.privateIpsCount;
    }
    @Export(name="queueNumber", refs={Integer.class}, tree="[0]")
    private Output<Integer> queueNumber;

    public Output<Integer> queueNumber() {
        return this.queueNumber;
    }
    /**
     * The Id of resource group which the network interface belongs.
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> resourceGroupId;

    /**
     * @return The Id of resource group which the network interface belongs.
     * 
     */
    public Output<Optional<String>> resourceGroupId() {
        return Codegen.optional(this.resourceGroupId);
    }
    @Export(name="secondaryPrivateIpAddressCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> secondaryPrivateIpAddressCount;

    public Output<Integer> secondaryPrivateIpAddressCount() {
        return this.secondaryPrivateIpAddressCount;
    }
    @Export(name="securityGroupIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> securityGroupIds;

    public Output<List<String>> securityGroupIds() {
        return this.securityGroupIds;
    }
    /**
     * A list of security group ids to associate with.
     * 
     * @deprecated
     * Field &#39;security_groups&#39; has been deprecated from provider version 1.123.1. New field &#39;security_group_ids&#39; instead
     * 
     */
    @Deprecated /* Field 'security_groups' has been deprecated from provider version 1.123.1. New field 'security_group_ids' instead */
    @Export(name="securityGroups", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> securityGroups;

    /**
     * @return A list of security group ids to associate with.
     * 
     */
    public Output<List<String>> securityGroups() {
        return this.securityGroups;
    }
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    public Output<String> status() {
        return this.status;
    }
    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The VSwitch to create the ENI in.
     * 
     */
    @Export(name="vswitchId", refs={String.class}, tree="[0]")
    private Output<String> vswitchId;

    /**
     * @return The VSwitch to create the ENI in.
     * 
     */
    public Output<String> vswitchId() {
        return this.vswitchId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NetworkInterface(java.lang.String name) {
        this(name, NetworkInterfaceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NetworkInterface(java.lang.String name, NetworkInterfaceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NetworkInterface(java.lang.String name, NetworkInterfaceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/networkInterface:NetworkInterface", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private NetworkInterface(java.lang.String name, Output<java.lang.String> id, @Nullable NetworkInterfaceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/networkInterface:NetworkInterface", name, state, makeResourceOptions(options, id), false);
    }

    private static NetworkInterfaceArgs makeArgs(NetworkInterfaceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? NetworkInterfaceArgs.Empty : args;
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
    public static NetworkInterface get(java.lang.String name, Output<java.lang.String> id, @Nullable NetworkInterfaceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NetworkInterface(name, id, state, options);
    }
}
