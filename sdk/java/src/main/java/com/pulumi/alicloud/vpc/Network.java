// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.vpc.NetworkArgs;
import com.pulumi.alicloud.vpc.inputs.NetworkState;
import com.pulumi.alicloud.vpc.outputs.NetworkIpv6CidrBlock;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Vpc Vpc resource. A VPC instance creates a VPC. You can fully control your own VPC, such as selecting IP address ranges, configuring routing tables, and gateways. You can use Alibaba cloud resources such as cloud servers, apsaradb for RDS, and load balancer in your own VPC.
 * 
 * &gt; **NOTE:** Available since v1.0.0.
 * 
 * &gt; **NOTE:** This resource will auto build a router and a route table while it uses `alicloud.vpc.Network` to build a vpc resource.
 * 
 * &gt; **NOTE:** Currently, the IPv4 / IPv6 dual-stack VPC function is under public testing. Only the following regions support IPv4 / IPv6 dual-stack VPC: `cn-hangzhou`, `cn-shanghai`, `cn-shenzhen`, `cn-beijing`, `cn-huhehaote`, `cn-hongkong` and `ap-southeast-1`, and need to apply for public beta qualification. To use, please [submit an application](https://www.alibabacloud.com/help/en/vpc/getting-started/create-a-vpc-with-an-ipv6-cidr-block).
 * 
 * ## Module Support
 * 
 * You can use the existing vpc module
 * to create a VPC and several VSwitches one-click.
 * 
 * For information about Vpc Vpc and how to use it, see [What is Vpc](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/what-is-a-vpc).
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
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
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
 *         var default_ = new Network(&#34;default&#34;, NetworkArgs.builder()        
 *             .ipv6Isp(&#34;BGP&#34;)
 *             .description(&#34;test&#34;)
 *             .cidrBlock(&#34;10.0.0.0/8&#34;)
 *             .vpcName(name)
 *             .enableIpv6(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Vpc Vpc can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:vpc/network:Network example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:vpc/network:Network")
public class Network extends com.pulumi.resources.CustomResource {
    /**
     * The CIDR block for the VPC. The `cidr_block` is Optional and default value is `172.16.0.0/12` after v1.119.0+.
     * 
     */
    @Export(name="cidrBlock", refs={String.class}, tree="[0]")
    private Output<String> cidrBlock;

    /**
     * @return The CIDR block for the VPC. The `cidr_block` is Optional and default value is `172.16.0.0/12` after v1.119.0+.
     * 
     */
    public Output<String> cidrBlock() {
        return this.cidrBlock;
    }
    /**
     * The status of ClassicLink function.
     * 
     */
    @Export(name="classicLinkEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> classicLinkEnabled;

    /**
     * @return The status of ClassicLink function.
     * 
     */
    public Output<Optional<Boolean>> classicLinkEnabled() {
        return Codegen.optional(this.classicLinkEnabled);
    }
    /**
     * The creation time of the VPC.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return The creation time of the VPC.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * The VPC description. Defaults to null.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The VPC description. Defaults to null.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Whether to PreCheck this request only. Value:
     * - **true**: sends a check request and does not create a VPC. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code &#39;DryRunOperation&#39; is returned &#39;.
     * - **false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly creates a VPC.
     * 
     */
    @Export(name="dryRun", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> dryRun;

    /**
     * @return Whether to PreCheck this request only. Value:
     * - **true**: sends a check request and does not create a VPC. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code &#39;DryRunOperation&#39; is returned &#39;.
     * - **false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly creates a VPC.
     * 
     */
    public Output<Optional<Boolean>> dryRun() {
        return Codegen.optional(this.dryRun);
    }
    /**
     * Whether to enable the IPv6 network segment. Value:
     * - **false** (default): not enabled.
     * - **true**: on.
     * 
     */
    @Export(name="enableIpv6", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableIpv6;

    /**
     * @return Whether to enable the IPv6 network segment. Value:
     * - **false** (default): not enabled.
     * - **true**: on.
     * 
     */
    public Output<Optional<Boolean>> enableIpv6() {
        return Codegen.optional(this.enableIpv6);
    }
    /**
     * The IPv6 CIDR block of the VPC.
     * 
     */
    @Export(name="ipv6CidrBlock", refs={String.class}, tree="[0]")
    private Output<String> ipv6CidrBlock;

    /**
     * @return The IPv6 CIDR block of the VPC.
     * 
     */
    public Output<String> ipv6CidrBlock() {
        return this.ipv6CidrBlock;
    }
    /**
     * The IPv6 CIDR block information of the VPC.
     * 
     */
    @Export(name="ipv6CidrBlocks", refs={List.class,NetworkIpv6CidrBlock.class}, tree="[0,1]")
    private Output<List<NetworkIpv6CidrBlock>> ipv6CidrBlocks;

    /**
     * @return The IPv6 CIDR block information of the VPC.
     * 
     */
    public Output<List<NetworkIpv6CidrBlock>> ipv6CidrBlocks() {
        return this.ipv6CidrBlocks;
    }
    /**
     * The IPv6 address segment type of the VPC. Value:
     * - **BGP** (default): Alibaba Cloud BGP IPv6.
     * - **ChinaMobile**: China Mobile (single line).
     * - **ChinaUnicom**: China Unicom (single line).
     * - **ChinaTelecom**: China Telecom (single line).
     * &gt; **NOTE:**  If a single-line bandwidth whitelist is enabled, this field can be set to **ChinaTelecom** (China Telecom), **ChinaUnicom** (China Unicom), or **ChinaMobile** (China Mobile).
     * 
     */
    @Export(name="ipv6Isp", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ipv6Isp;

    /**
     * @return The IPv6 address segment type of the VPC. Value:
     * - **BGP** (default): Alibaba Cloud BGP IPv6.
     * - **ChinaMobile**: China Mobile (single line).
     * - **ChinaUnicom**: China Unicom (single line).
     * - **ChinaTelecom**: China Telecom (single line).
     * &gt; **NOTE:**  If a single-line bandwidth whitelist is enabled, this field can be set to **ChinaTelecom** (China Telecom), **ChinaUnicom** (China Unicom), or **ChinaMobile** (China Mobile).
     * 
     */
    public Output<Optional<String>> ipv6Isp() {
        return Codegen.optional(this.ipv6Isp);
    }
    /**
     * Field &#39;name&#39; has been deprecated from provider version 1.119.0. New field &#39;vpc_name&#39; instead.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.119.0. New field &#39;vpc_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.119.0. New field 'vpc_name' instead. */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Field &#39;name&#39; has been deprecated from provider version 1.119.0. New field &#39;vpc_name&#39; instead.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The ID of the resource group to which the VPC belongs.
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group to which the VPC belongs.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * The route table ID of the router created by default on VPC creation.
     * 
     */
    @Export(name="routeTableId", refs={String.class}, tree="[0]")
    private Output<String> routeTableId;

    /**
     * @return The route table ID of the router created by default on VPC creation.
     * 
     */
    public Output<String> routeTableId() {
        return this.routeTableId;
    }
    /**
     * The ID of the router created by default on VPC creation.
     * 
     */
    @Export(name="routerId", refs={String.class}, tree="[0]")
    private Output<String> routerId;

    /**
     * @return The ID of the router created by default on VPC creation.
     * 
     */
    public Output<String> routerId() {
        return this.routerId;
    }
    /**
     * (Deprecated since v1.206.0+) Field &#39;router_table_id&#39; has been deprecated from provider version 1.206.0. New field &#39;route_table_id&#39; instead.
     * 
     * @deprecated
     * Field &#39;router_table_id&#39; has been deprecated from provider version 1.206.0. New field &#39;route_table_id&#39; instead.
     * 
     */
    @Deprecated /* Field 'router_table_id' has been deprecated from provider version 1.206.0. New field 'route_table_id' instead. */
    @Export(name="routerTableId", refs={String.class}, tree="[0]")
    private Output<String> routerTableId;

    /**
     * @return (Deprecated since v1.206.0+) Field &#39;router_table_id&#39; has been deprecated from provider version 1.206.0. New field &#39;route_table_id&#39; instead.
     * 
     */
    public Output<String> routerTableId() {
        return this.routerTableId;
    }
    /**
     * Field &#39;secondary_cidr_blocks&#39; has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_ipv4_cidr_block&#39;. `secondary_cidr_blocks` attributes and `alicloud.vpc.Ipv4CidrBlock` resource cannot be used at the same time.
     * 
     * @deprecated
     * Field &#39;SecondaryCidrBlocks&#39; has been deprecated from provider version 1.206.0. Field &#39;secondary_cidr_blocks&#39; has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_ipv4_cidr_block&#39;. `secondary_cidr_blocks` attributes and `alicloud_vpc_ipv4_cidr_block` resource cannot be used at the same time.
     * 
     */
    @Deprecated /* Field 'SecondaryCidrBlocks' has been deprecated from provider version 1.206.0. Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_ipv4_cidr_block'. `secondary_cidr_blocks` attributes and `alicloud_vpc_ipv4_cidr_block` resource cannot be used at the same time. */
    @Export(name="secondaryCidrBlocks", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> secondaryCidrBlocks;

    /**
     * @return Field &#39;secondary_cidr_blocks&#39; has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_ipv4_cidr_block&#39;. `secondary_cidr_blocks` attributes and `alicloud.vpc.Ipv4CidrBlock` resource cannot be used at the same time.
     * 
     */
    public Output<List<String>> secondaryCidrBlocks() {
        return this.secondaryCidrBlocks;
    }
    /**
     * The status of the VPC. Valid values:  **Pending**: The VPC is being configured. **Available**: The VPC is available.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the VPC. Valid values:  **Pending**: The VPC is being configured. **Available**: The VPC is available.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The tags of Vpc.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return The tags of Vpc.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * A list of user CIDRs.
     * 
     */
    @Export(name="userCidrs", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> userCidrs;

    /**
     * @return A list of user CIDRs.
     * 
     */
    public Output<List<String>> userCidrs() {
        return this.userCidrs;
    }
    /**
     * The name of the VPC. Defaults to null.
     * 
     * The following arguments will be discarded. Please use new fields as soon as possible:
     * 
     */
    @Export(name="vpcName", refs={String.class}, tree="[0]")
    private Output<String> vpcName;

    /**
     * @return The name of the VPC. Defaults to null.
     * 
     * The following arguments will be discarded. Please use new fields as soon as possible:
     * 
     */
    public Output<String> vpcName() {
        return this.vpcName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Network(String name) {
        this(name, NetworkArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Network(String name, @Nullable NetworkArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Network(String name, @Nullable NetworkArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/network:Network", name, args == null ? NetworkArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Network(String name, Output<String> id, @Nullable NetworkState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/network:Network", name, state, makeResourceOptions(options, id));
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
    public static Network get(String name, Output<String> id, @Nullable NetworkState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Network(name, id, state, options);
    }
}
