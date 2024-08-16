// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NetworkInterfaceState extends com.pulumi.resources.ResourceArgs {

    public static final NetworkInterfaceState Empty = new NetworkInterfaceState();

    /**
     * Description of the ENI. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the ENI. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="instanceType")
    private @Nullable Output<String> instanceType;

    public Optional<Output<String>> instanceType() {
        return Optional.ofNullable(this.instanceType);
    }

    @Import(name="ipv4PrefixCount")
    private @Nullable Output<Integer> ipv4PrefixCount;

    public Optional<Output<Integer>> ipv4PrefixCount() {
        return Optional.ofNullable(this.ipv4PrefixCount);
    }

    @Import(name="ipv4Prefixes")
    private @Nullable Output<List<String>> ipv4Prefixes;

    public Optional<Output<List<String>>> ipv4Prefixes() {
        return Optional.ofNullable(this.ipv4Prefixes);
    }

    @Import(name="ipv6AddressCount")
    private @Nullable Output<Integer> ipv6AddressCount;

    public Optional<Output<Integer>> ipv6AddressCount() {
        return Optional.ofNullable(this.ipv6AddressCount);
    }

    @Import(name="ipv6Addresses")
    private @Nullable Output<List<String>> ipv6Addresses;

    public Optional<Output<List<String>>> ipv6Addresses() {
        return Optional.ofNullable(this.ipv6Addresses);
    }

    /**
     * (Available in 1.54.0+) The MAC address of an ENI.
     * 
     */
    @Import(name="mac")
    private @Nullable Output<String> mac;

    /**
     * @return (Available in 1.54.0+) The MAC address of an ENI.
     * 
     */
    public Optional<Output<String>> mac() {
        return Optional.ofNullable(this.mac);
    }

    /**
     * Name of the ENI. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;, &#34;.&#34;, &#34;_&#34;, and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.123.1. New field &#39;network_interface_name&#39; instead
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the ENI. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;, &#34;.&#34;, &#34;_&#34;, and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.123.1. New field &#39;network_interface_name&#39; instead
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="networkInterfaceName")
    private @Nullable Output<String> networkInterfaceName;

    public Optional<Output<String>> networkInterfaceName() {
        return Optional.ofNullable(this.networkInterfaceName);
    }

    @Import(name="networkInterfaceTrafficMode")
    private @Nullable Output<String> networkInterfaceTrafficMode;

    public Optional<Output<String>> networkInterfaceTrafficMode() {
        return Optional.ofNullable(this.networkInterfaceTrafficMode);
    }

    @Import(name="primaryIpAddress")
    private @Nullable Output<String> primaryIpAddress;

    public Optional<Output<String>> primaryIpAddress() {
        return Optional.ofNullable(this.primaryIpAddress);
    }

    /**
     * The primary private IP of the ENI.
     * 
     * @deprecated
     * Field &#39;private_ip&#39; has been deprecated from provider version 1.123.1. New field &#39;primary_ip_address&#39; instead
     * 
     */
    @Deprecated /* Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead */
    @Import(name="privateIp")
    private @Nullable Output<String> privateIp;

    /**
     * @return The primary private IP of the ENI.
     * 
     * @deprecated
     * Field &#39;private_ip&#39; has been deprecated from provider version 1.123.1. New field &#39;primary_ip_address&#39; instead
     * 
     */
    @Deprecated /* Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead */
    public Optional<Output<String>> privateIp() {
        return Optional.ofNullable(this.privateIp);
    }

    @Import(name="privateIpAddresses")
    private @Nullable Output<List<String>> privateIpAddresses;

    public Optional<Output<List<String>>> privateIpAddresses() {
        return Optional.ofNullable(this.privateIpAddresses);
    }

    /**
     * List of secondary private IPs to assign to the ENI. Don&#39;t use both private_ips and private_ips_count in the same ENI resource block.
     * 
     * @deprecated
     * Field &#39;private_ips&#39; has been deprecated from provider version 1.123.1. New field &#39;private_ip_addresses&#39; instead
     * 
     */
    @Deprecated /* Field 'private_ips' has been deprecated from provider version 1.123.1. New field 'private_ip_addresses' instead */
    @Import(name="privateIps")
    private @Nullable Output<List<String>> privateIps;

    /**
     * @return List of secondary private IPs to assign to the ENI. Don&#39;t use both private_ips and private_ips_count in the same ENI resource block.
     * 
     * @deprecated
     * Field &#39;private_ips&#39; has been deprecated from provider version 1.123.1. New field &#39;private_ip_addresses&#39; instead
     * 
     */
    @Deprecated /* Field 'private_ips' has been deprecated from provider version 1.123.1. New field 'private_ip_addresses' instead */
    public Optional<Output<List<String>>> privateIps() {
        return Optional.ofNullable(this.privateIps);
    }

    /**
     * Number of secondary private IPs to assign to the ENI. Don&#39;t use both private_ips and private_ips_count in the same ENI resource block.
     * 
     * @deprecated
     * Field &#39;private_ips_count&#39; has been deprecated from provider version 1.123.1. New field &#39;secondary_private_ip_address_count&#39; instead
     * 
     */
    @Deprecated /* Field 'private_ips_count' has been deprecated from provider version 1.123.1. New field 'secondary_private_ip_address_count' instead */
    @Import(name="privateIpsCount")
    private @Nullable Output<Integer> privateIpsCount;

    /**
     * @return Number of secondary private IPs to assign to the ENI. Don&#39;t use both private_ips and private_ips_count in the same ENI resource block.
     * 
     * @deprecated
     * Field &#39;private_ips_count&#39; has been deprecated from provider version 1.123.1. New field &#39;secondary_private_ip_address_count&#39; instead
     * 
     */
    @Deprecated /* Field 'private_ips_count' has been deprecated from provider version 1.123.1. New field 'secondary_private_ip_address_count' instead */
    public Optional<Output<Integer>> privateIpsCount() {
        return Optional.ofNullable(this.privateIpsCount);
    }

    @Import(name="queueNumber")
    private @Nullable Output<Integer> queueNumber;

    public Optional<Output<Integer>> queueNumber() {
        return Optional.ofNullable(this.queueNumber);
    }

    /**
     * The Id of resource group which the network interface belongs.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The Id of resource group which the network interface belongs.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    @Import(name="secondaryPrivateIpAddressCount")
    private @Nullable Output<Integer> secondaryPrivateIpAddressCount;

    public Optional<Output<Integer>> secondaryPrivateIpAddressCount() {
        return Optional.ofNullable(this.secondaryPrivateIpAddressCount);
    }

    @Import(name="securityGroupIds")
    private @Nullable Output<List<String>> securityGroupIds;

    public Optional<Output<List<String>>> securityGroupIds() {
        return Optional.ofNullable(this.securityGroupIds);
    }

    /**
     * A list of security group ids to associate with.
     * 
     * @deprecated
     * Field &#39;security_groups&#39; has been deprecated from provider version 1.123.1. New field &#39;security_group_ids&#39; instead
     * 
     */
    @Deprecated /* Field 'security_groups' has been deprecated from provider version 1.123.1. New field 'security_group_ids' instead */
    @Import(name="securityGroups")
    private @Nullable Output<List<String>> securityGroups;

    /**
     * @return A list of security group ids to associate with.
     * 
     * @deprecated
     * Field &#39;security_groups&#39; has been deprecated from provider version 1.123.1. New field &#39;security_group_ids&#39; instead
     * 
     */
    @Deprecated /* Field 'security_groups' has been deprecated from provider version 1.123.1. New field 'security_group_ids' instead */
    public Optional<Output<List<String>>> securityGroups() {
        return Optional.ofNullable(this.securityGroups);
    }

    @Import(name="status")
    private @Nullable Output<String> status;

    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The VSwitch to create the ENI in.
     * 
     */
    @Import(name="vswitchId")
    private @Nullable Output<String> vswitchId;

    /**
     * @return The VSwitch to create the ENI in.
     * 
     */
    public Optional<Output<String>> vswitchId() {
        return Optional.ofNullable(this.vswitchId);
    }

    private NetworkInterfaceState() {}

    private NetworkInterfaceState(NetworkInterfaceState $) {
        this.description = $.description;
        this.instanceType = $.instanceType;
        this.ipv4PrefixCount = $.ipv4PrefixCount;
        this.ipv4Prefixes = $.ipv4Prefixes;
        this.ipv6AddressCount = $.ipv6AddressCount;
        this.ipv6Addresses = $.ipv6Addresses;
        this.mac = $.mac;
        this.name = $.name;
        this.networkInterfaceName = $.networkInterfaceName;
        this.networkInterfaceTrafficMode = $.networkInterfaceTrafficMode;
        this.primaryIpAddress = $.primaryIpAddress;
        this.privateIp = $.privateIp;
        this.privateIpAddresses = $.privateIpAddresses;
        this.privateIps = $.privateIps;
        this.privateIpsCount = $.privateIpsCount;
        this.queueNumber = $.queueNumber;
        this.resourceGroupId = $.resourceGroupId;
        this.secondaryPrivateIpAddressCount = $.secondaryPrivateIpAddressCount;
        this.securityGroupIds = $.securityGroupIds;
        this.securityGroups = $.securityGroups;
        this.status = $.status;
        this.tags = $.tags;
        this.vswitchId = $.vswitchId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NetworkInterfaceState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NetworkInterfaceState $;

        public Builder() {
            $ = new NetworkInterfaceState();
        }

        public Builder(NetworkInterfaceState defaults) {
            $ = new NetworkInterfaceState(Objects.requireNonNull(defaults));
        }

        /**
         * @param description Description of the ENI. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the ENI. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        public Builder instanceType(@Nullable Output<String> instanceType) {
            $.instanceType = instanceType;
            return this;
        }

        public Builder instanceType(String instanceType) {
            return instanceType(Output.of(instanceType));
        }

        public Builder ipv4PrefixCount(@Nullable Output<Integer> ipv4PrefixCount) {
            $.ipv4PrefixCount = ipv4PrefixCount;
            return this;
        }

        public Builder ipv4PrefixCount(Integer ipv4PrefixCount) {
            return ipv4PrefixCount(Output.of(ipv4PrefixCount));
        }

        public Builder ipv4Prefixes(@Nullable Output<List<String>> ipv4Prefixes) {
            $.ipv4Prefixes = ipv4Prefixes;
            return this;
        }

        public Builder ipv4Prefixes(List<String> ipv4Prefixes) {
            return ipv4Prefixes(Output.of(ipv4Prefixes));
        }

        public Builder ipv4Prefixes(String... ipv4Prefixes) {
            return ipv4Prefixes(List.of(ipv4Prefixes));
        }

        public Builder ipv6AddressCount(@Nullable Output<Integer> ipv6AddressCount) {
            $.ipv6AddressCount = ipv6AddressCount;
            return this;
        }

        public Builder ipv6AddressCount(Integer ipv6AddressCount) {
            return ipv6AddressCount(Output.of(ipv6AddressCount));
        }

        public Builder ipv6Addresses(@Nullable Output<List<String>> ipv6Addresses) {
            $.ipv6Addresses = ipv6Addresses;
            return this;
        }

        public Builder ipv6Addresses(List<String> ipv6Addresses) {
            return ipv6Addresses(Output.of(ipv6Addresses));
        }

        public Builder ipv6Addresses(String... ipv6Addresses) {
            return ipv6Addresses(List.of(ipv6Addresses));
        }

        /**
         * @param mac (Available in 1.54.0+) The MAC address of an ENI.
         * 
         * @return builder
         * 
         */
        public Builder mac(@Nullable Output<String> mac) {
            $.mac = mac;
            return this;
        }

        /**
         * @param mac (Available in 1.54.0+) The MAC address of an ENI.
         * 
         * @return builder
         * 
         */
        public Builder mac(String mac) {
            return mac(Output.of(mac));
        }

        /**
         * @param name Name of the ENI. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;, &#34;.&#34;, &#34;_&#34;, and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;name&#39; has been deprecated from provider version 1.123.1. New field &#39;network_interface_name&#39; instead
         * 
         */
        @Deprecated /* Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the ENI. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;, &#34;.&#34;, &#34;_&#34;, and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;name&#39; has been deprecated from provider version 1.123.1. New field &#39;network_interface_name&#39; instead
         * 
         */
        @Deprecated /* Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder networkInterfaceName(@Nullable Output<String> networkInterfaceName) {
            $.networkInterfaceName = networkInterfaceName;
            return this;
        }

        public Builder networkInterfaceName(String networkInterfaceName) {
            return networkInterfaceName(Output.of(networkInterfaceName));
        }

        public Builder networkInterfaceTrafficMode(@Nullable Output<String> networkInterfaceTrafficMode) {
            $.networkInterfaceTrafficMode = networkInterfaceTrafficMode;
            return this;
        }

        public Builder networkInterfaceTrafficMode(String networkInterfaceTrafficMode) {
            return networkInterfaceTrafficMode(Output.of(networkInterfaceTrafficMode));
        }

        public Builder primaryIpAddress(@Nullable Output<String> primaryIpAddress) {
            $.primaryIpAddress = primaryIpAddress;
            return this;
        }

        public Builder primaryIpAddress(String primaryIpAddress) {
            return primaryIpAddress(Output.of(primaryIpAddress));
        }

        /**
         * @param privateIp The primary private IP of the ENI.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;private_ip&#39; has been deprecated from provider version 1.123.1. New field &#39;primary_ip_address&#39; instead
         * 
         */
        @Deprecated /* Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead */
        public Builder privateIp(@Nullable Output<String> privateIp) {
            $.privateIp = privateIp;
            return this;
        }

        /**
         * @param privateIp The primary private IP of the ENI.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;private_ip&#39; has been deprecated from provider version 1.123.1. New field &#39;primary_ip_address&#39; instead
         * 
         */
        @Deprecated /* Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead */
        public Builder privateIp(String privateIp) {
            return privateIp(Output.of(privateIp));
        }

        public Builder privateIpAddresses(@Nullable Output<List<String>> privateIpAddresses) {
            $.privateIpAddresses = privateIpAddresses;
            return this;
        }

        public Builder privateIpAddresses(List<String> privateIpAddresses) {
            return privateIpAddresses(Output.of(privateIpAddresses));
        }

        public Builder privateIpAddresses(String... privateIpAddresses) {
            return privateIpAddresses(List.of(privateIpAddresses));
        }

        /**
         * @param privateIps List of secondary private IPs to assign to the ENI. Don&#39;t use both private_ips and private_ips_count in the same ENI resource block.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;private_ips&#39; has been deprecated from provider version 1.123.1. New field &#39;private_ip_addresses&#39; instead
         * 
         */
        @Deprecated /* Field 'private_ips' has been deprecated from provider version 1.123.1. New field 'private_ip_addresses' instead */
        public Builder privateIps(@Nullable Output<List<String>> privateIps) {
            $.privateIps = privateIps;
            return this;
        }

        /**
         * @param privateIps List of secondary private IPs to assign to the ENI. Don&#39;t use both private_ips and private_ips_count in the same ENI resource block.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;private_ips&#39; has been deprecated from provider version 1.123.1. New field &#39;private_ip_addresses&#39; instead
         * 
         */
        @Deprecated /* Field 'private_ips' has been deprecated from provider version 1.123.1. New field 'private_ip_addresses' instead */
        public Builder privateIps(List<String> privateIps) {
            return privateIps(Output.of(privateIps));
        }

        /**
         * @param privateIps List of secondary private IPs to assign to the ENI. Don&#39;t use both private_ips and private_ips_count in the same ENI resource block.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;private_ips&#39; has been deprecated from provider version 1.123.1. New field &#39;private_ip_addresses&#39; instead
         * 
         */
        @Deprecated /* Field 'private_ips' has been deprecated from provider version 1.123.1. New field 'private_ip_addresses' instead */
        public Builder privateIps(String... privateIps) {
            return privateIps(List.of(privateIps));
        }

        /**
         * @param privateIpsCount Number of secondary private IPs to assign to the ENI. Don&#39;t use both private_ips and private_ips_count in the same ENI resource block.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;private_ips_count&#39; has been deprecated from provider version 1.123.1. New field &#39;secondary_private_ip_address_count&#39; instead
         * 
         */
        @Deprecated /* Field 'private_ips_count' has been deprecated from provider version 1.123.1. New field 'secondary_private_ip_address_count' instead */
        public Builder privateIpsCount(@Nullable Output<Integer> privateIpsCount) {
            $.privateIpsCount = privateIpsCount;
            return this;
        }

        /**
         * @param privateIpsCount Number of secondary private IPs to assign to the ENI. Don&#39;t use both private_ips and private_ips_count in the same ENI resource block.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;private_ips_count&#39; has been deprecated from provider version 1.123.1. New field &#39;secondary_private_ip_address_count&#39; instead
         * 
         */
        @Deprecated /* Field 'private_ips_count' has been deprecated from provider version 1.123.1. New field 'secondary_private_ip_address_count' instead */
        public Builder privateIpsCount(Integer privateIpsCount) {
            return privateIpsCount(Output.of(privateIpsCount));
        }

        public Builder queueNumber(@Nullable Output<Integer> queueNumber) {
            $.queueNumber = queueNumber;
            return this;
        }

        public Builder queueNumber(Integer queueNumber) {
            return queueNumber(Output.of(queueNumber));
        }

        /**
         * @param resourceGroupId The Id of resource group which the network interface belongs.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The Id of resource group which the network interface belongs.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        public Builder secondaryPrivateIpAddressCount(@Nullable Output<Integer> secondaryPrivateIpAddressCount) {
            $.secondaryPrivateIpAddressCount = secondaryPrivateIpAddressCount;
            return this;
        }

        public Builder secondaryPrivateIpAddressCount(Integer secondaryPrivateIpAddressCount) {
            return secondaryPrivateIpAddressCount(Output.of(secondaryPrivateIpAddressCount));
        }

        public Builder securityGroupIds(@Nullable Output<List<String>> securityGroupIds) {
            $.securityGroupIds = securityGroupIds;
            return this;
        }

        public Builder securityGroupIds(List<String> securityGroupIds) {
            return securityGroupIds(Output.of(securityGroupIds));
        }

        public Builder securityGroupIds(String... securityGroupIds) {
            return securityGroupIds(List.of(securityGroupIds));
        }

        /**
         * @param securityGroups A list of security group ids to associate with.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;security_groups&#39; has been deprecated from provider version 1.123.1. New field &#39;security_group_ids&#39; instead
         * 
         */
        @Deprecated /* Field 'security_groups' has been deprecated from provider version 1.123.1. New field 'security_group_ids' instead */
        public Builder securityGroups(@Nullable Output<List<String>> securityGroups) {
            $.securityGroups = securityGroups;
            return this;
        }

        /**
         * @param securityGroups A list of security group ids to associate with.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;security_groups&#39; has been deprecated from provider version 1.123.1. New field &#39;security_group_ids&#39; instead
         * 
         */
        @Deprecated /* Field 'security_groups' has been deprecated from provider version 1.123.1. New field 'security_group_ids' instead */
        public Builder securityGroups(List<String> securityGroups) {
            return securityGroups(Output.of(securityGroups));
        }

        /**
         * @param securityGroups A list of security group ids to associate with.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;security_groups&#39; has been deprecated from provider version 1.123.1. New field &#39;security_group_ids&#39; instead
         * 
         */
        @Deprecated /* Field 'security_groups' has been deprecated from provider version 1.123.1. New field 'security_group_ids' instead */
        public Builder securityGroups(String... securityGroups) {
            return securityGroups(List.of(securityGroups));
        }

        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param vswitchId The VSwitch to create the ENI in.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(@Nullable Output<String> vswitchId) {
            $.vswitchId = vswitchId;
            return this;
        }

        /**
         * @param vswitchId The VSwitch to create the ENI in.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(String vswitchId) {
            return vswitchId(Output.of(vswitchId));
        }

        public NetworkInterfaceState build() {
            return $;
        }
    }

}
