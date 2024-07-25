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
     * The CIDR block of the VPC.
     * - You can specify one of the following CIDR blocks or their subsets as the primary IPv4 CIDR block of the VPC: 192.168.0.0/16, 172.16.0.0/12, and 10.0.0.0/8. These CIDR blocks are standard private CIDR blocks as defined by Request for Comments (RFC) documents. The subnet mask must be 8 to 28 bits in length.
     * - You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, 169.254.0.0/16, and their subnets as the primary IPv4 CIDR block of the VPC.
     * 
     */
    @Import(name="cidrBlock")
    private @Nullable Output<String> cidrBlock;

    /**
     * @return The CIDR block of the VPC.
     * - You can specify one of the following CIDR blocks or their subsets as the primary IPv4 CIDR block of the VPC: 192.168.0.0/16, 172.16.0.0/12, and 10.0.0.0/8. These CIDR blocks are standard private CIDR blocks as defined by Request for Comments (RFC) documents. The subnet mask must be 8 to 28 bits in length.
     * - You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, 169.254.0.0/16, and their subnets as the primary IPv4 CIDR block of the VPC.
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
     * The new description of the VPC. The description must be 1 to 256 characters in length, and cannot start with `http://` or `https://`.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The new description of the VPC. The description must be 1 to 256 characters in length, and cannot start with `http://` or `https://`.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Specifies whether to perform a dry run. Valid values:
     * 
     */
    @Import(name="dryRun")
    private @Nullable Output<Boolean> dryRun;

    /**
     * @return Specifies whether to perform a dry run. Valid values:
     * 
     */
    public Optional<Output<Boolean>> dryRun() {
        return Optional.ofNullable(this.dryRun);
    }

    /**
     * The name of the VPC. The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
     * 
     */
    @Import(name="enableIpv6")
    private @Nullable Output<Boolean> enableIpv6;

    /**
     * @return The name of the VPC. The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
     * 
     */
    public Optional<Output<Boolean>> enableIpv6() {
        return Optional.ofNullable(this.enableIpv6);
    }

    /**
     * The ID of the IP Address Manager (IPAM) pool that contains IPv4 addresses.
     * 
     */
    @Import(name="ipv4IpamPoolId")
    private @Nullable Output<String> ipv4IpamPoolId;

    /**
     * @return The ID of the IP Address Manager (IPAM) pool that contains IPv4 addresses.
     * 
     */
    public Optional<Output<String>> ipv4IpamPoolId() {
        return Optional.ofNullable(this.ipv4IpamPoolId);
    }

    /**
     * The IPv6 CIDR block of the default VPC.
     * 
     * &gt; **NOTE:**  When `EnableIpv6` is set to `true`, this parameter is required.
     * 
     */
    @Import(name="ipv6CidrBlock")
    private @Nullable Output<String> ipv6CidrBlock;

    /**
     * @return The IPv6 CIDR block of the default VPC.
     * 
     * &gt; **NOTE:**  When `EnableIpv6` is set to `true`, this parameter is required.
     * 
     */
    public Optional<Output<String>> ipv6CidrBlock() {
        return Optional.ofNullable(this.ipv6CidrBlock);
    }

    /**
     * The IPv6 address segment type of the VPC. Value:
     * - `BGP` (default): Alibaba Cloud BGP IPv6.
     * - `ChinaMobile`: China Mobile (single line).
     * - `ChinaUnicom`: China Unicom (single line).
     * - `ChinaTelecom`: China Telecom (single line).
     * 
     * &gt; **NOTE:**  If a single-line bandwidth whitelist is enabled, this field can be set to `ChinaTelecom` (China Telecom), `ChinaUnicom` (China Unicom), or `ChinaMobile` (China Mobile).
     * 
     */
    @Import(name="ipv6Isp")
    private @Nullable Output<String> ipv6Isp;

    /**
     * @return The IPv6 address segment type of the VPC. Value:
     * - `BGP` (default): Alibaba Cloud BGP IPv6.
     * - `ChinaMobile`: China Mobile (single line).
     * - `ChinaUnicom`: China Unicom (single line).
     * - `ChinaTelecom`: China Telecom (single line).
     * 
     * &gt; **NOTE:**  If a single-line bandwidth whitelist is enabled, this field can be set to `ChinaTelecom` (China Telecom), `ChinaUnicom` (China Unicom), or `ChinaMobile` (China Mobile).
     * 
     */
    public Optional<Output<String>> ipv6Isp() {
        return Optional.ofNullable(this.ipv6Isp);
    }

    /**
     * Specifies whether to create the default VPC in the specified region. Valid values:
     * 
     */
    @Import(name="isDefault")
    private @Nullable Output<Boolean> isDefault;

    /**
     * @return Specifies whether to create the default VPC in the specified region. Valid values:
     * 
     */
    public Optional<Output<Boolean>> isDefault() {
        return Optional.ofNullable(this.isDefault);
    }

    /**
     * . Field &#39;name&#39; has been deprecated from provider version 1.119.0. New field &#39;vpc_name&#39; instead.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated since provider version 1.119.0. New field &#39;vpc_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated since provider version 1.119.0. New field 'vpc_name' instead. */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return . Field &#39;name&#39; has been deprecated from provider version 1.119.0. New field &#39;vpc_name&#39; instead.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated since provider version 1.119.0. New field &#39;vpc_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated since provider version 1.119.0. New field 'vpc_name' instead. */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID of the resource group to which you want to move the resource.
     * 
     * &gt; **NOTE:**   You can use resource groups to facilitate resource grouping and permission management for an Alibaba Cloud. For more information, see [What is resource management?](https://www.alibabacloud.com/help/en/doc-detail/94475.html)
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group to which you want to move the resource.
     * 
     * &gt; **NOTE:**   You can use resource groups to facilitate resource grouping and permission management for an Alibaba Cloud. For more information, see [What is resource management?](https://www.alibabacloud.com/help/en/doc-detail/94475.html)
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * Field &#39;secondary_cidr_blocks&#39; has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_ipv4_cidr_block&#39;. `secondary_cidr_blocks` attributes and `alicloud.vpc.Ipv4CidrBlock` resource cannot be used at the same time.
     * 
     * @deprecated
     * Field &#39;secondary_cidr_blocks&#39; has been deprecated from provider version 1.185.0. Field &#39;secondary_cidr_blocks&#39; has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_ipv4_cidr_block&#39;. `secondary_cidr_blocks` attributes and `alicloud.vpc.Ipv4CidrBlock` resource cannot be used at the same time.
     * 
     */
    @Deprecated /* Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0. Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_ipv4_cidr_block'. `secondary_cidr_blocks` attributes and `alicloud.vpc.Ipv4CidrBlock` resource cannot be used at the same time. */
    @Import(name="secondaryCidrBlocks")
    private @Nullable Output<List<String>> secondaryCidrBlocks;

    /**
     * @return Field &#39;secondary_cidr_blocks&#39; has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_ipv4_cidr_block&#39;. `secondary_cidr_blocks` attributes and `alicloud.vpc.Ipv4CidrBlock` resource cannot be used at the same time.
     * 
     * @deprecated
     * Field &#39;secondary_cidr_blocks&#39; has been deprecated from provider version 1.185.0. Field &#39;secondary_cidr_blocks&#39; has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_ipv4_cidr_block&#39;. `secondary_cidr_blocks` attributes and `alicloud.vpc.Ipv4CidrBlock` resource cannot be used at the same time.
     * 
     */
    @Deprecated /* Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0. Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_ipv4_cidr_block'. `secondary_cidr_blocks` attributes and `alicloud.vpc.Ipv4CidrBlock` resource cannot be used at the same time. */
    public Optional<Output<List<String>>> secondaryCidrBlocks() {
        return Optional.ofNullable(this.secondaryCidrBlocks);
    }

    /**
     * The description of the route table. The description must be 1 to 256 characters in length, and cannot start with `http://` or `https://`.
     * 
     */
    @Import(name="systemRouteTableDescription")
    private @Nullable Output<String> systemRouteTableDescription;

    /**
     * @return The description of the route table. The description must be 1 to 256 characters in length, and cannot start with `http://` or `https://`.
     * 
     */
    public Optional<Output<String>> systemRouteTableDescription() {
        return Optional.ofNullable(this.systemRouteTableDescription);
    }

    /**
     * The name of the route table. The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
     * 
     */
    @Import(name="systemRouteTableName")
    private @Nullable Output<String> systemRouteTableName;

    /**
     * @return The name of the route table. The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
     * 
     */
    public Optional<Output<String>> systemRouteTableName() {
        return Optional.ofNullable(this.systemRouteTableName);
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
     * The new name of the VPC. The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
     * 
     * The following arguments will be discarded. Please use new fields as soon as possible:
     * 
     */
    @Import(name="vpcName")
    private @Nullable Output<String> vpcName;

    /**
     * @return The new name of the VPC. The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
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
        this.ipv4IpamPoolId = $.ipv4IpamPoolId;
        this.ipv6CidrBlock = $.ipv6CidrBlock;
        this.ipv6Isp = $.ipv6Isp;
        this.isDefault = $.isDefault;
        this.name = $.name;
        this.resourceGroupId = $.resourceGroupId;
        this.secondaryCidrBlocks = $.secondaryCidrBlocks;
        this.systemRouteTableDescription = $.systemRouteTableDescription;
        this.systemRouteTableName = $.systemRouteTableName;
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
         * @param cidrBlock The CIDR block of the VPC.
         * - You can specify one of the following CIDR blocks or their subsets as the primary IPv4 CIDR block of the VPC: 192.168.0.0/16, 172.16.0.0/12, and 10.0.0.0/8. These CIDR blocks are standard private CIDR blocks as defined by Request for Comments (RFC) documents. The subnet mask must be 8 to 28 bits in length.
         * - You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, 169.254.0.0/16, and their subnets as the primary IPv4 CIDR block of the VPC.
         * 
         * @return builder
         * 
         */
        public Builder cidrBlock(@Nullable Output<String> cidrBlock) {
            $.cidrBlock = cidrBlock;
            return this;
        }

        /**
         * @param cidrBlock The CIDR block of the VPC.
         * - You can specify one of the following CIDR blocks or their subsets as the primary IPv4 CIDR block of the VPC: 192.168.0.0/16, 172.16.0.0/12, and 10.0.0.0/8. These CIDR blocks are standard private CIDR blocks as defined by Request for Comments (RFC) documents. The subnet mask must be 8 to 28 bits in length.
         * - You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, 169.254.0.0/16, and their subnets as the primary IPv4 CIDR block of the VPC.
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
         * @param description The new description of the VPC. The description must be 1 to 256 characters in length, and cannot start with `http://` or `https://`.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The new description of the VPC. The description must be 1 to 256 characters in length, and cannot start with `http://` or `https://`.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param dryRun Specifies whether to perform a dry run. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder dryRun(@Nullable Output<Boolean> dryRun) {
            $.dryRun = dryRun;
            return this;
        }

        /**
         * @param dryRun Specifies whether to perform a dry run. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder dryRun(Boolean dryRun) {
            return dryRun(Output.of(dryRun));
        }

        /**
         * @param enableIpv6 The name of the VPC. The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
         * 
         * @return builder
         * 
         */
        public Builder enableIpv6(@Nullable Output<Boolean> enableIpv6) {
            $.enableIpv6 = enableIpv6;
            return this;
        }

        /**
         * @param enableIpv6 The name of the VPC. The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
         * 
         * @return builder
         * 
         */
        public Builder enableIpv6(Boolean enableIpv6) {
            return enableIpv6(Output.of(enableIpv6));
        }

        /**
         * @param ipv4IpamPoolId The ID of the IP Address Manager (IPAM) pool that contains IPv4 addresses.
         * 
         * @return builder
         * 
         */
        public Builder ipv4IpamPoolId(@Nullable Output<String> ipv4IpamPoolId) {
            $.ipv4IpamPoolId = ipv4IpamPoolId;
            return this;
        }

        /**
         * @param ipv4IpamPoolId The ID of the IP Address Manager (IPAM) pool that contains IPv4 addresses.
         * 
         * @return builder
         * 
         */
        public Builder ipv4IpamPoolId(String ipv4IpamPoolId) {
            return ipv4IpamPoolId(Output.of(ipv4IpamPoolId));
        }

        /**
         * @param ipv6CidrBlock The IPv6 CIDR block of the default VPC.
         * 
         * &gt; **NOTE:**  When `EnableIpv6` is set to `true`, this parameter is required.
         * 
         * @return builder
         * 
         */
        public Builder ipv6CidrBlock(@Nullable Output<String> ipv6CidrBlock) {
            $.ipv6CidrBlock = ipv6CidrBlock;
            return this;
        }

        /**
         * @param ipv6CidrBlock The IPv6 CIDR block of the default VPC.
         * 
         * &gt; **NOTE:**  When `EnableIpv6` is set to `true`, this parameter is required.
         * 
         * @return builder
         * 
         */
        public Builder ipv6CidrBlock(String ipv6CidrBlock) {
            return ipv6CidrBlock(Output.of(ipv6CidrBlock));
        }

        /**
         * @param ipv6Isp The IPv6 address segment type of the VPC. Value:
         * - `BGP` (default): Alibaba Cloud BGP IPv6.
         * - `ChinaMobile`: China Mobile (single line).
         * - `ChinaUnicom`: China Unicom (single line).
         * - `ChinaTelecom`: China Telecom (single line).
         * 
         * &gt; **NOTE:**  If a single-line bandwidth whitelist is enabled, this field can be set to `ChinaTelecom` (China Telecom), `ChinaUnicom` (China Unicom), or `ChinaMobile` (China Mobile).
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
         * - `BGP` (default): Alibaba Cloud BGP IPv6.
         * - `ChinaMobile`: China Mobile (single line).
         * - `ChinaUnicom`: China Unicom (single line).
         * - `ChinaTelecom`: China Telecom (single line).
         * 
         * &gt; **NOTE:**  If a single-line bandwidth whitelist is enabled, this field can be set to `ChinaTelecom` (China Telecom), `ChinaUnicom` (China Unicom), or `ChinaMobile` (China Mobile).
         * 
         * @return builder
         * 
         */
        public Builder ipv6Isp(String ipv6Isp) {
            return ipv6Isp(Output.of(ipv6Isp));
        }

        /**
         * @param isDefault Specifies whether to create the default VPC in the specified region. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder isDefault(@Nullable Output<Boolean> isDefault) {
            $.isDefault = isDefault;
            return this;
        }

        /**
         * @param isDefault Specifies whether to create the default VPC in the specified region. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder isDefault(Boolean isDefault) {
            return isDefault(Output.of(isDefault));
        }

        /**
         * @param name . Field &#39;name&#39; has been deprecated from provider version 1.119.0. New field &#39;vpc_name&#39; instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;name&#39; has been deprecated since provider version 1.119.0. New field &#39;vpc_name&#39; instead.
         * 
         */
        @Deprecated /* Field 'name' has been deprecated since provider version 1.119.0. New field 'vpc_name' instead. */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name . Field &#39;name&#39; has been deprecated from provider version 1.119.0. New field &#39;vpc_name&#39; instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;name&#39; has been deprecated since provider version 1.119.0. New field &#39;vpc_name&#39; instead.
         * 
         */
        @Deprecated /* Field 'name' has been deprecated since provider version 1.119.0. New field 'vpc_name' instead. */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param resourceGroupId The ID of the resource group to which you want to move the resource.
         * 
         * &gt; **NOTE:**   You can use resource groups to facilitate resource grouping and permission management for an Alibaba Cloud. For more information, see [What is resource management?](https://www.alibabacloud.com/help/en/doc-detail/94475.html)
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The ID of the resource group to which you want to move the resource.
         * 
         * &gt; **NOTE:**   You can use resource groups to facilitate resource grouping and permission management for an Alibaba Cloud. For more information, see [What is resource management?](https://www.alibabacloud.com/help/en/doc-detail/94475.html)
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        /**
         * @param secondaryCidrBlocks Field &#39;secondary_cidr_blocks&#39; has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_ipv4_cidr_block&#39;. `secondary_cidr_blocks` attributes and `alicloud.vpc.Ipv4CidrBlock` resource cannot be used at the same time.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;secondary_cidr_blocks&#39; has been deprecated from provider version 1.185.0. Field &#39;secondary_cidr_blocks&#39; has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_ipv4_cidr_block&#39;. `secondary_cidr_blocks` attributes and `alicloud.vpc.Ipv4CidrBlock` resource cannot be used at the same time.
         * 
         */
        @Deprecated /* Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0. Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_ipv4_cidr_block'. `secondary_cidr_blocks` attributes and `alicloud.vpc.Ipv4CidrBlock` resource cannot be used at the same time. */
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
         * Field &#39;secondary_cidr_blocks&#39; has been deprecated from provider version 1.185.0. Field &#39;secondary_cidr_blocks&#39; has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_ipv4_cidr_block&#39;. `secondary_cidr_blocks` attributes and `alicloud.vpc.Ipv4CidrBlock` resource cannot be used at the same time.
         * 
         */
        @Deprecated /* Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0. Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_ipv4_cidr_block'. `secondary_cidr_blocks` attributes and `alicloud.vpc.Ipv4CidrBlock` resource cannot be used at the same time. */
        public Builder secondaryCidrBlocks(List<String> secondaryCidrBlocks) {
            return secondaryCidrBlocks(Output.of(secondaryCidrBlocks));
        }

        /**
         * @param secondaryCidrBlocks Field &#39;secondary_cidr_blocks&#39; has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_ipv4_cidr_block&#39;. `secondary_cidr_blocks` attributes and `alicloud.vpc.Ipv4CidrBlock` resource cannot be used at the same time.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;secondary_cidr_blocks&#39; has been deprecated from provider version 1.185.0. Field &#39;secondary_cidr_blocks&#39; has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_ipv4_cidr_block&#39;. `secondary_cidr_blocks` attributes and `alicloud.vpc.Ipv4CidrBlock` resource cannot be used at the same time.
         * 
         */
        @Deprecated /* Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0. Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_ipv4_cidr_block'. `secondary_cidr_blocks` attributes and `alicloud.vpc.Ipv4CidrBlock` resource cannot be used at the same time. */
        public Builder secondaryCidrBlocks(String... secondaryCidrBlocks) {
            return secondaryCidrBlocks(List.of(secondaryCidrBlocks));
        }

        /**
         * @param systemRouteTableDescription The description of the route table. The description must be 1 to 256 characters in length, and cannot start with `http://` or `https://`.
         * 
         * @return builder
         * 
         */
        public Builder systemRouteTableDescription(@Nullable Output<String> systemRouteTableDescription) {
            $.systemRouteTableDescription = systemRouteTableDescription;
            return this;
        }

        /**
         * @param systemRouteTableDescription The description of the route table. The description must be 1 to 256 characters in length, and cannot start with `http://` or `https://`.
         * 
         * @return builder
         * 
         */
        public Builder systemRouteTableDescription(String systemRouteTableDescription) {
            return systemRouteTableDescription(Output.of(systemRouteTableDescription));
        }

        /**
         * @param systemRouteTableName The name of the route table. The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
         * 
         * @return builder
         * 
         */
        public Builder systemRouteTableName(@Nullable Output<String> systemRouteTableName) {
            $.systemRouteTableName = systemRouteTableName;
            return this;
        }

        /**
         * @param systemRouteTableName The name of the route table. The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
         * 
         * @return builder
         * 
         */
        public Builder systemRouteTableName(String systemRouteTableName) {
            return systemRouteTableName(Output.of(systemRouteTableName));
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
         * @param vpcName The new name of the VPC. The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
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
         * @param vpcName The new name of the VPC. The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
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
