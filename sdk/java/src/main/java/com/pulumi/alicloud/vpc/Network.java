// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.vpc.NetworkArgs;
import com.pulumi.alicloud.vpc.inputs.NetworkState;
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
 * ## Example Usage
 * 
 * Basic Usage
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
 *         var vpc = new Network(&#34;vpc&#34;, NetworkArgs.builder()        
 *             .cidrBlock(&#34;172.16.0.0/12&#34;)
 *             .vpcName(&#34;tf_test_foo&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Module Support
 * 
 * You can use the existing vpc module
 * to create a VPC and several VSwitches one-click.
 * 
 * ## Import
 * 
 * VPC can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:vpc/network:Network example vpc-abc123456
 * ```
 * 
 */
@ResourceType(type="alicloud:vpc/network:Network")
public class Network extends com.pulumi.resources.CustomResource {
    /**
     * The CIDR block for the VPC. The `cidr_block` is Optional and default value is `172.16.0.0/12` after v1.119.0+.
     * 
     */
    @Export(name="cidrBlock", type=String.class, parameters={})
    private Output</* @Nullable */ String> cidrBlock;

    /**
     * @return The CIDR block for the VPC. The `cidr_block` is Optional and default value is `172.16.0.0/12` after v1.119.0+.
     * 
     */
    public Output<Optional<String>> cidrBlock() {
        return Codegen.optional(this.cidrBlock);
    }
    /**
     * The VPC description. Defaults to null.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The VPC description. Defaults to null.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Specifies whether to pre-check this request only. Valid values: `true` and `false`.
     * 
     */
    @Export(name="dryRun", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> dryRun;

    /**
     * @return Specifies whether to pre-check this request only. Valid values: `true` and `false`.
     * 
     */
    public Output<Optional<Boolean>> dryRun() {
        return Codegen.optional(this.dryRun);
    }
    /**
     * Specifies whether to enable the IPv6 CIDR block. Valid values: `false` (Default): disables IPv6 CIDR blocks. `true`: enables IPv6 CIDR blocks. If the `enable_ipv6` is `true`, the system will automatically create a free version of an IPv6 gateway for your private network and assign an IPv6 network segment assigned as /56.
     * 
     * &gt; **NOTE:** Currently, the IPv4 / IPv6 dual-stack VPC function is under public testing. Only the following regions support IPv4 / IPv6 dual-stack VPC: `cn-hangzhou`, `cn-shanghai`, `cn-shenzhen`, `cn-beijing`, `cn-huhehaote`, `cn-hongkong` and `ap-southeast-1`, and need to apply for public beta qualification. To use, please [submit an application](https://help.aliyun.com/document_detail/100334.html).
     * 
     */
    @Export(name="enableIpv6", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> enableIpv6;

    /**
     * @return Specifies whether to enable the IPv6 CIDR block. Valid values: `false` (Default): disables IPv6 CIDR blocks. `true`: enables IPv6 CIDR blocks. If the `enable_ipv6` is `true`, the system will automatically create a free version of an IPv6 gateway for your private network and assign an IPv6 network segment assigned as /56.
     * 
     * &gt; **NOTE:** Currently, the IPv4 / IPv6 dual-stack VPC function is under public testing. Only the following regions support IPv4 / IPv6 dual-stack VPC: `cn-hangzhou`, `cn-shanghai`, `cn-shenzhen`, `cn-beijing`, `cn-huhehaote`, `cn-hongkong` and `ap-southeast-1`, and need to apply for public beta qualification. To use, please [submit an application](https://help.aliyun.com/document_detail/100334.html).
     * 
     */
    public Output<Optional<Boolean>> enableIpv6() {
        return Codegen.optional(this.enableIpv6);
    }
    /**
     * (Available in v1.119.0+) ) The ipv6 cidr block of VPC.
     * 
     */
    @Export(name="ipv6CidrBlock", type=String.class, parameters={})
    private Output<String> ipv6CidrBlock;

    /**
     * @return (Available in v1.119.0+) ) The ipv6 cidr block of VPC.
     * 
     */
    public Output<String> ipv6CidrBlock() {
        return this.ipv6CidrBlock;
    }
    /**
     * Field `name` has been deprecated from provider version 1.119.0. New field `vpc_name` instead.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.119.0. New field &#39;vpc_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.119.0. New field 'vpc_name' instead. */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Field `name` has been deprecated from provider version 1.119.0. New field `vpc_name` instead.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The ID of resource group which the VPC belongs.
     * 
     */
    @Export(name="resourceGroupId", type=String.class, parameters={})
    private Output<String> resourceGroupId;

    /**
     * @return The ID of resource group which the VPC belongs.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * The route table ID of the router created by default on VPC creation.
     * 
     */
    @Export(name="routeTableId", type=String.class, parameters={})
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
    @Export(name="routerId", type=String.class, parameters={})
    private Output<String> routerId;

    /**
     * @return The ID of the router created by default on VPC creation.
     * 
     */
    public Output<String> routerId() {
        return this.routerId;
    }
    /**
     * (Deprecated) It has been deprecated and replaced with `route_table_id`.
     * 
     * @deprecated
     * Attribute router_table_id has been deprecated and replaced with route_table_id.
     * 
     */
    @Deprecated /* Attribute router_table_id has been deprecated and replaced with route_table_id. */
    @Export(name="routerTableId", type=String.class, parameters={})
    private Output<String> routerTableId;

    /**
     * @return (Deprecated) It has been deprecated and replaced with `route_table_id`.
     * 
     */
    public Output<String> routerTableId() {
        return this.routerTableId;
    }
    /**
     * Field `secondary_cidr_blocks` has been deprecated from provider version 1.185.0, and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_ipv4_cidr_block&#39;. `secondary_cidr_blocks` attributes and `alicloud.vpc.Ipv4CidrBlock` resource cannot be used at the same time.
     * 
     * @deprecated
     * Field &#39;secondary_cidr_blocks&#39; has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_ipv4_cidr_block&#39;. `secondary_cidr_blocks` attributes and `alicloud_vpc_ipv4_cidr_block` resource cannot be used at the same time.
     * 
     */
    @Deprecated /* Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_ipv4_cidr_block'. `secondary_cidr_blocks` attributes and `alicloud_vpc_ipv4_cidr_block` resource cannot be used at the same time. */
    @Export(name="secondaryCidrBlocks", type=List.class, parameters={String.class})
    private Output<List<String>> secondaryCidrBlocks;

    /**
     * @return Field `secondary_cidr_blocks` has been deprecated from provider version 1.185.0, and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_ipv4_cidr_block&#39;. `secondary_cidr_blocks` attributes and `alicloud.vpc.Ipv4CidrBlock` resource cannot be used at the same time.
     * 
     */
    public Output<List<String>> secondaryCidrBlocks() {
        return this.secondaryCidrBlocks;
    }
    /**
     * The status of the VPC.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of the VPC.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Export(name="tags", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The user cidr blocks of the VPC.
     * 
     */
    @Export(name="userCidrs", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> userCidrs;

    /**
     * @return The user cidr blocks of the VPC.
     * 
     */
    public Output<Optional<List<String>>> userCidrs() {
        return Codegen.optional(this.userCidrs);
    }
    /**
     * The name of the VPC. Defaults to null.
     * 
     */
    @Export(name="vpcName", type=String.class, parameters={})
    private Output<String> vpcName;

    /**
     * @return The name of the VPC. Defaults to null.
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
