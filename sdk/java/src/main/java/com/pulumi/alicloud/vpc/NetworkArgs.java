// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NetworkArgs extends com.pulumi.resources.ResourceArgs {

    public static final NetworkArgs Empty = new NetworkArgs();

    /**
     * The CIDR block for the VPC. The `cidr_block` is Optional and default value is `172.16.0.0/12` after v1.119.0+.
     * 
     */
    @Import(name="cidrBlock")
    private @Nullable Output<String> cidrBlock;

    /**
     * @return The CIDR block for the VPC. The `cidr_block` is Optional and default value is `172.16.0.0/12` after v1.119.0+.
     * 
     */
    public Optional<Output<String>> cidrBlock() {
        return Optional.ofNullable(this.cidrBlock);
    }

    /**
     * The status of ClassicLink function.
     * 
     */
    @Import(name="classicLinkEnabled")
    private @Nullable Output<Boolean> classicLinkEnabled;

    /**
     * @return The status of ClassicLink function.
     * 
     */
    public Optional<Output<Boolean>> classicLinkEnabled() {
        return Optional.ofNullable(this.classicLinkEnabled);
    }

    /**
     * The VPC description. Defaults to null.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The VPC description. Defaults to null.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Whether to PreCheck this request only. Value:
     * - **true**: sends a check request and does not create a VPC. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code &#39;DryRunOperation&#39; is returned &#39;.
     * - **false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly creates a VPC.
     * 
     */
    @Import(name="dryRun")
    private @Nullable Output<Boolean> dryRun;

    /**
     * @return Whether to PreCheck this request only. Value:
     * - **true**: sends a check request and does not create a VPC. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code &#39;DryRunOperation&#39; is returned &#39;.
     * - **false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly creates a VPC.
     * 
     */
    public Optional<Output<Boolean>> dryRun() {
        return Optional.ofNullable(this.dryRun);
    }

    /**
     * Whether to enable the IPv6 network segment. Value:
     * - **false** (default): not enabled.
     * - **true**: on.
     * 
     */
    @Import(name="enableIpv6")
    private @Nullable Output<Boolean> enableIpv6;

    /**
     * @return Whether to enable the IPv6 network segment. Value:
     * - **false** (default): not enabled.
     * - **true**: on.
     * 
     */
    public Optional<Output<Boolean>> enableIpv6() {
        return Optional.ofNullable(this.enableIpv6);
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
    @Import(name="ipv6Isp")
    private @Nullable Output<String> ipv6Isp;

    /**
     * @return The IPv6 address segment type of the VPC. Value:
     * - **BGP** (default): Alibaba Cloud BGP IPv6.
     * - **ChinaMobile**: China Mobile (single line).
     * - **ChinaUnicom**: China Unicom (single line).
     * - **ChinaTelecom**: China Telecom (single line).
     * &gt; **NOTE:**  If a single-line bandwidth whitelist is enabled, this field can be set to **ChinaTelecom** (China Telecom), **ChinaUnicom** (China Unicom), or **ChinaMobile** (China Mobile).
     * 
     */
    public Optional<Output<String>> ipv6Isp() {
        return Optional.ofNullable(this.ipv6Isp);
    }

    /**
     * Field &#39;name&#39; has been deprecated from provider version 1.119.0. New field &#39;vpc_name&#39; instead.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.119.0. New field &#39;vpc_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.119.0. New field 'vpc_name' instead. */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Field &#39;name&#39; has been deprecated from provider version 1.119.0. New field &#39;vpc_name&#39; instead.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.119.0. New field &#39;vpc_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.119.0. New field 'vpc_name' instead. */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID of the resource group to which the VPC belongs.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group to which the VPC belongs.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * The route table ID of the router created by default on VPC creation.
     * 
     */
    @Import(name="routeTableId")
    private @Nullable Output<String> routeTableId;

    /**
     * @return The route table ID of the router created by default on VPC creation.
     * 
     */
    public Optional<Output<String>> routeTableId() {
        return Optional.ofNullable(this.routeTableId);
    }

    /**
     * Field &#39;router_table_id&#39; has been deprecated from provider version 1.206.0. New field &#39;route_table_id&#39; instead.
     * 
     * @deprecated
     * Field &#39;router_table_id&#39; has been deprecated from provider version 1.206.0. New field &#39;route_table_id&#39; instead.
     * 
     */
    @Deprecated /* Field 'router_table_id' has been deprecated from provider version 1.206.0. New field 'route_table_id' instead. */
    @Import(name="routerTableId")
    private @Nullable Output<String> routerTableId;

    /**
     * @return Field &#39;router_table_id&#39; has been deprecated from provider version 1.206.0. New field &#39;route_table_id&#39; instead.
     * 
     * @deprecated
     * Field &#39;router_table_id&#39; has been deprecated from provider version 1.206.0. New field &#39;route_table_id&#39; instead.
     * 
     */
    @Deprecated /* Field 'router_table_id' has been deprecated from provider version 1.206.0. New field 'route_table_id' instead. */
    public Optional<Output<String>> routerTableId() {
        return Optional.ofNullable(this.routerTableId);
    }

    /**
     * Field &#39;secondary_cidr_blocks&#39; has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_ipv4_cidr_block&#39;. `secondary_cidr_blocks` attributes and `alicloud.vpc.Ipv4CidrBlock` resource cannot be used at the same time.
     * 
     * @deprecated
     * Field &#39;SecondaryCidrBlocks&#39; has been deprecated from provider version 1.206.0. Field &#39;secondary_cidr_blocks&#39; has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_ipv4_cidr_block&#39;. `secondary_cidr_blocks` attributes and `alicloud_vpc_ipv4_cidr_block` resource cannot be used at the same time.
     * 
     */
    @Deprecated /* Field 'SecondaryCidrBlocks' has been deprecated from provider version 1.206.0. Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_ipv4_cidr_block'. `secondary_cidr_blocks` attributes and `alicloud_vpc_ipv4_cidr_block` resource cannot be used at the same time. */
    @Import(name="secondaryCidrBlocks")
    private @Nullable Output<List<String>> secondaryCidrBlocks;

    /**
     * @return Field &#39;secondary_cidr_blocks&#39; has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_ipv4_cidr_block&#39;. `secondary_cidr_blocks` attributes and `alicloud.vpc.Ipv4CidrBlock` resource cannot be used at the same time.
     * 
     * @deprecated
     * Field &#39;SecondaryCidrBlocks&#39; has been deprecated from provider version 1.206.0. Field &#39;secondary_cidr_blocks&#39; has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_ipv4_cidr_block&#39;. `secondary_cidr_blocks` attributes and `alicloud_vpc_ipv4_cidr_block` resource cannot be used at the same time.
     * 
     */
    @Deprecated /* Field 'SecondaryCidrBlocks' has been deprecated from provider version 1.206.0. Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_ipv4_cidr_block'. `secondary_cidr_blocks` attributes and `alicloud_vpc_ipv4_cidr_block` resource cannot be used at the same time. */
    public Optional<Output<List<String>>> secondaryCidrBlocks() {
        return Optional.ofNullable(this.secondaryCidrBlocks);
    }

    /**
     * The tags of Vpc.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,Object>> tags;

    /**
     * @return The tags of Vpc.
     * 
     */
    public Optional<Output<Map<String,Object>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * A list of user CIDRs.
     * 
     */
    @Import(name="userCidrs")
    private @Nullable Output<List<String>> userCidrs;

    /**
     * @return A list of user CIDRs.
     * 
     */
    public Optional<Output<List<String>>> userCidrs() {
        return Optional.ofNullable(this.userCidrs);
    }

    /**
     * The name of the VPC. Defaults to null.
     * 
     * The following arguments will be discarded. Please use new fields as soon as possible:
     * 
     */
    @Import(name="vpcName")
    private @Nullable Output<String> vpcName;

    /**
     * @return The name of the VPC. Defaults to null.
     * 
     * The following arguments will be discarded. Please use new fields as soon as possible:
     * 
     */
    public Optional<Output<String>> vpcName() {
        return Optional.ofNullable(this.vpcName);
    }

    private NetworkArgs() {}

    private NetworkArgs(NetworkArgs $) {
        this.cidrBlock = $.cidrBlock;
        this.classicLinkEnabled = $.classicLinkEnabled;
        this.description = $.description;
        this.dryRun = $.dryRun;
        this.enableIpv6 = $.enableIpv6;
        this.ipv6Isp = $.ipv6Isp;
        this.name = $.name;
        this.resourceGroupId = $.resourceGroupId;
        this.routeTableId = $.routeTableId;
        this.routerTableId = $.routerTableId;
        this.secondaryCidrBlocks = $.secondaryCidrBlocks;
        this.tags = $.tags;
        this.userCidrs = $.userCidrs;
        this.vpcName = $.vpcName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NetworkArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NetworkArgs $;

        public Builder() {
            $ = new NetworkArgs();
        }

        public Builder(NetworkArgs defaults) {
            $ = new NetworkArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param cidrBlock The CIDR block for the VPC. The `cidr_block` is Optional and default value is `172.16.0.0/12` after v1.119.0+.
         * 
         * @return builder
         * 
         */
        public Builder cidrBlock(@Nullable Output<String> cidrBlock) {
            $.cidrBlock = cidrBlock;
            return this;
        }

        /**
         * @param cidrBlock The CIDR block for the VPC. The `cidr_block` is Optional and default value is `172.16.0.0/12` after v1.119.0+.
         * 
         * @return builder
         * 
         */
        public Builder cidrBlock(String cidrBlock) {
            return cidrBlock(Output.of(cidrBlock));
        }

        /**
         * @param classicLinkEnabled The status of ClassicLink function.
         * 
         * @return builder
         * 
         */
        public Builder classicLinkEnabled(@Nullable Output<Boolean> classicLinkEnabled) {
            $.classicLinkEnabled = classicLinkEnabled;
            return this;
        }

        /**
         * @param classicLinkEnabled The status of ClassicLink function.
         * 
         * @return builder
         * 
         */
        public Builder classicLinkEnabled(Boolean classicLinkEnabled) {
            return classicLinkEnabled(Output.of(classicLinkEnabled));
        }

        /**
         * @param description The VPC description. Defaults to null.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The VPC description. Defaults to null.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param dryRun Whether to PreCheck this request only. Value:
         * - **true**: sends a check request and does not create a VPC. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code &#39;DryRunOperation&#39; is returned &#39;.
         * - **false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly creates a VPC.
         * 
         * @return builder
         * 
         */
        public Builder dryRun(@Nullable Output<Boolean> dryRun) {
            $.dryRun = dryRun;
            return this;
        }

        /**
         * @param dryRun Whether to PreCheck this request only. Value:
         * - **true**: sends a check request and does not create a VPC. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code &#39;DryRunOperation&#39; is returned &#39;.
         * - **false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly creates a VPC.
         * 
         * @return builder
         * 
         */
        public Builder dryRun(Boolean dryRun) {
            return dryRun(Output.of(dryRun));
        }

        /**
         * @param enableIpv6 Whether to enable the IPv6 network segment. Value:
         * - **false** (default): not enabled.
         * - **true**: on.
         * 
         * @return builder
         * 
         */
        public Builder enableIpv6(@Nullable Output<Boolean> enableIpv6) {
            $.enableIpv6 = enableIpv6;
            return this;
        }

        /**
         * @param enableIpv6 Whether to enable the IPv6 network segment. Value:
         * - **false** (default): not enabled.
         * - **true**: on.
         * 
         * @return builder
         * 
         */
        public Builder enableIpv6(Boolean enableIpv6) {
            return enableIpv6(Output.of(enableIpv6));
        }

        /**
         * @param ipv6Isp The IPv6 address segment type of the VPC. Value:
         * - **BGP** (default): Alibaba Cloud BGP IPv6.
         * - **ChinaMobile**: China Mobile (single line).
         * - **ChinaUnicom**: China Unicom (single line).
         * - **ChinaTelecom**: China Telecom (single line).
         * &gt; **NOTE:**  If a single-line bandwidth whitelist is enabled, this field can be set to **ChinaTelecom** (China Telecom), **ChinaUnicom** (China Unicom), or **ChinaMobile** (China Mobile).
         * 
         * @return builder
         * 
         */
        public Builder ipv6Isp(@Nullable Output<String> ipv6Isp) {
            $.ipv6Isp = ipv6Isp;
            return this;
        }

        /**
         * @param ipv6Isp The IPv6 address segment type of the VPC. Value:
         * - **BGP** (default): Alibaba Cloud BGP IPv6.
         * - **ChinaMobile**: China Mobile (single line).
         * - **ChinaUnicom**: China Unicom (single line).
         * - **ChinaTelecom**: China Telecom (single line).
         * &gt; **NOTE:**  If a single-line bandwidth whitelist is enabled, this field can be set to **ChinaTelecom** (China Telecom), **ChinaUnicom** (China Unicom), or **ChinaMobile** (China Mobile).
         * 
         * @return builder
         * 
         */
        public Builder ipv6Isp(String ipv6Isp) {
            return ipv6Isp(Output.of(ipv6Isp));
        }

        /**
         * @param name Field &#39;name&#39; has been deprecated from provider version 1.119.0. New field &#39;vpc_name&#39; instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;name&#39; has been deprecated from provider version 1.119.0. New field &#39;vpc_name&#39; instead.
         * 
         */
        @Deprecated /* Field 'name' has been deprecated from provider version 1.119.0. New field 'vpc_name' instead. */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Field &#39;name&#39; has been deprecated from provider version 1.119.0. New field &#39;vpc_name&#39; instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;name&#39; has been deprecated from provider version 1.119.0. New field &#39;vpc_name&#39; instead.
         * 
         */
        @Deprecated /* Field 'name' has been deprecated from provider version 1.119.0. New field 'vpc_name' instead. */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param resourceGroupId The ID of the resource group to which the VPC belongs.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The ID of the resource group to which the VPC belongs.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        /**
         * @param routeTableId The route table ID of the router created by default on VPC creation.
         * 
         * @return builder
         * 
         */
        public Builder routeTableId(@Nullable Output<String> routeTableId) {
            $.routeTableId = routeTableId;
            return this;
        }

        /**
         * @param routeTableId The route table ID of the router created by default on VPC creation.
         * 
         * @return builder
         * 
         */
        public Builder routeTableId(String routeTableId) {
            return routeTableId(Output.of(routeTableId));
        }

        /**
         * @param routerTableId Field &#39;router_table_id&#39; has been deprecated from provider version 1.206.0. New field &#39;route_table_id&#39; instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;router_table_id&#39; has been deprecated from provider version 1.206.0. New field &#39;route_table_id&#39; instead.
         * 
         */
        @Deprecated /* Field 'router_table_id' has been deprecated from provider version 1.206.0. New field 'route_table_id' instead. */
        public Builder routerTableId(@Nullable Output<String> routerTableId) {
            $.routerTableId = routerTableId;
            return this;
        }

        /**
         * @param routerTableId Field &#39;router_table_id&#39; has been deprecated from provider version 1.206.0. New field &#39;route_table_id&#39; instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;router_table_id&#39; has been deprecated from provider version 1.206.0. New field &#39;route_table_id&#39; instead.
         * 
         */
        @Deprecated /* Field 'router_table_id' has been deprecated from provider version 1.206.0. New field 'route_table_id' instead. */
        public Builder routerTableId(String routerTableId) {
            return routerTableId(Output.of(routerTableId));
        }

        /**
         * @param secondaryCidrBlocks Field &#39;secondary_cidr_blocks&#39; has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_ipv4_cidr_block&#39;. `secondary_cidr_blocks` attributes and `alicloud.vpc.Ipv4CidrBlock` resource cannot be used at the same time.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;SecondaryCidrBlocks&#39; has been deprecated from provider version 1.206.0. Field &#39;secondary_cidr_blocks&#39; has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_ipv4_cidr_block&#39;. `secondary_cidr_blocks` attributes and `alicloud_vpc_ipv4_cidr_block` resource cannot be used at the same time.
         * 
         */
        @Deprecated /* Field 'SecondaryCidrBlocks' has been deprecated from provider version 1.206.0. Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_ipv4_cidr_block'. `secondary_cidr_blocks` attributes and `alicloud_vpc_ipv4_cidr_block` resource cannot be used at the same time. */
        public Builder secondaryCidrBlocks(@Nullable Output<List<String>> secondaryCidrBlocks) {
            $.secondaryCidrBlocks = secondaryCidrBlocks;
            return this;
        }

        /**
         * @param secondaryCidrBlocks Field &#39;secondary_cidr_blocks&#39; has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_ipv4_cidr_block&#39;. `secondary_cidr_blocks` attributes and `alicloud.vpc.Ipv4CidrBlock` resource cannot be used at the same time.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;SecondaryCidrBlocks&#39; has been deprecated from provider version 1.206.0. Field &#39;secondary_cidr_blocks&#39; has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_ipv4_cidr_block&#39;. `secondary_cidr_blocks` attributes and `alicloud_vpc_ipv4_cidr_block` resource cannot be used at the same time.
         * 
         */
        @Deprecated /* Field 'SecondaryCidrBlocks' has been deprecated from provider version 1.206.0. Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_ipv4_cidr_block'. `secondary_cidr_blocks` attributes and `alicloud_vpc_ipv4_cidr_block` resource cannot be used at the same time. */
        public Builder secondaryCidrBlocks(List<String> secondaryCidrBlocks) {
            return secondaryCidrBlocks(Output.of(secondaryCidrBlocks));
        }

        /**
         * @param secondaryCidrBlocks Field &#39;secondary_cidr_blocks&#39; has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_ipv4_cidr_block&#39;. `secondary_cidr_blocks` attributes and `alicloud.vpc.Ipv4CidrBlock` resource cannot be used at the same time.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;SecondaryCidrBlocks&#39; has been deprecated from provider version 1.206.0. Field &#39;secondary_cidr_blocks&#39; has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_ipv4_cidr_block&#39;. `secondary_cidr_blocks` attributes and `alicloud_vpc_ipv4_cidr_block` resource cannot be used at the same time.
         * 
         */
        @Deprecated /* Field 'SecondaryCidrBlocks' has been deprecated from provider version 1.206.0. Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_ipv4_cidr_block'. `secondary_cidr_blocks` attributes and `alicloud_vpc_ipv4_cidr_block` resource cannot be used at the same time. */
        public Builder secondaryCidrBlocks(String... secondaryCidrBlocks) {
            return secondaryCidrBlocks(List.of(secondaryCidrBlocks));
        }

        /**
         * @param tags The tags of Vpc.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,Object>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The tags of Vpc.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,Object> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param userCidrs A list of user CIDRs.
         * 
         * @return builder
         * 
         */
        public Builder userCidrs(@Nullable Output<List<String>> userCidrs) {
            $.userCidrs = userCidrs;
            return this;
        }

        /**
         * @param userCidrs A list of user CIDRs.
         * 
         * @return builder
         * 
         */
        public Builder userCidrs(List<String> userCidrs) {
            return userCidrs(Output.of(userCidrs));
        }

        /**
         * @param userCidrs A list of user CIDRs.
         * 
         * @return builder
         * 
         */
        public Builder userCidrs(String... userCidrs) {
            return userCidrs(List.of(userCidrs));
        }

        /**
         * @param vpcName The name of the VPC. Defaults to null.
         * 
         * The following arguments will be discarded. Please use new fields as soon as possible:
         * 
         * @return builder
         * 
         */
        public Builder vpcName(@Nullable Output<String> vpcName) {
            $.vpcName = vpcName;
            return this;
        }

        /**
         * @param vpcName The name of the VPC. Defaults to null.
         * 
         * The following arguments will be discarded. Please use new fields as soon as possible:
         * 
         * @return builder
         * 
         */
        public Builder vpcName(String vpcName) {
            return vpcName(Output.of(vpcName));
        }

        public NetworkArgs build() {
            return $;
        }
    }

}
