// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetEcsNetworkInterfacesArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetEcsNetworkInterfacesArgs Empty = new GetEcsNetworkInterfacesArgs();

    /**
     * A list of Network Interface IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of Network Interface IDs.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * The instance id.
     * 
     */
    @Import(name="instanceId")
    private @Nullable Output<String> instanceId;

    /**
     * @return The instance id.
     * 
     */
    public Optional<Output<String>> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    /**
     * Field `name` has been deprecated from provider version 1.123.1. New field `network_interface_name` instead
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.123.1. New field &#39;network_interface_name&#39; instead
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Field `name` has been deprecated from provider version 1.123.1. New field `network_interface_name` instead
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.123.1. New field &#39;network_interface_name&#39; instead
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * A regex string to filter results by Network Interface name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable Output<String> nameRegex;

    /**
     * @return A regex string to filter results by Network Interface name.
     * 
     */
    public Optional<Output<String>> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    /**
     * The network interface name.
     * 
     */
    @Import(name="networkInterfaceName")
    private @Nullable Output<String> networkInterfaceName;

    /**
     * @return The network interface name.
     * 
     */
    public Optional<Output<String>> networkInterfaceName() {
        return Optional.ofNullable(this.networkInterfaceName);
    }

    /**
     * File name where to save data source results (after running `pulumi preview`).
     * 
     */
    @Import(name="outputFile")
    private @Nullable Output<String> outputFile;

    /**
     * @return File name where to save data source results (after running `pulumi preview`).
     * 
     */
    public Optional<Output<String>> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * The primary private IP address of the ENI.
     * 
     */
    @Import(name="primaryIpAddress")
    private @Nullable Output<String> primaryIpAddress;

    /**
     * @return The primary private IP address of the ENI.
     * 
     */
    public Optional<Output<String>> primaryIpAddress() {
        return Optional.ofNullable(this.primaryIpAddress);
    }

    /**
     * Field `private_ip` has been deprecated from provider version 1.123.1. New field `primary_ip_address` instead
     * 
     * @deprecated
     * Field &#39;private_ip&#39; has been deprecated from provider version 1.123.1. New field &#39;primary_ip_address&#39; instead
     * 
     */
    @Deprecated /* Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead */
    @Import(name="privateIp")
    private @Nullable Output<String> privateIp;

    /**
     * @return Field `private_ip` has been deprecated from provider version 1.123.1. New field `primary_ip_address` instead
     * 
     * @deprecated
     * Field &#39;private_ip&#39; has been deprecated from provider version 1.123.1. New field &#39;primary_ip_address&#39; instead
     * 
     */
    @Deprecated /* Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead */
    public Optional<Output<String>> privateIp() {
        return Optional.ofNullable(this.privateIp);
    }

    /**
     * The resource group id.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The resource group id.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * The security group id.
     * 
     */
    @Import(name="securityGroupId")
    private @Nullable Output<String> securityGroupId;

    /**
     * @return The security group id.
     * 
     */
    public Optional<Output<String>> securityGroupId() {
        return Optional.ofNullable(this.securityGroupId);
    }

    /**
     * Whether the user of the elastic network card is a cloud product or a virtual vendor.
     * 
     */
    @Import(name="serviceManaged")
    private @Nullable Output<Boolean> serviceManaged;

    /**
     * @return Whether the user of the elastic network card is a cloud product or a virtual vendor.
     * 
     */
    public Optional<Output<Boolean>> serviceManaged() {
        return Optional.ofNullable(this.serviceManaged);
    }

    /**
     * The status of ENI. Valid Values: `Attaching`, `Available`, `CreateFailed`, `Creating`, `Deleting`, `Detaching`, `InUse`, `Linked`, `Linking`, `Unlinking`.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of ENI. Valid Values: `Attaching`, `Available`, `CreateFailed`, `Creating`, `Deleting`, `Detaching`, `InUse`, `Linked`, `Linking`, `Unlinking`.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * A map of tags assigned to ENIs.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags assigned to ENIs.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The type of ENI. Valid Values: `Primary`, `Secondary`.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return The type of ENI. Valid Values: `Primary`, `Secondary`.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    /**
     * The vpc id.
     * 
     */
    @Import(name="vpcId")
    private @Nullable Output<String> vpcId;

    /**
     * @return The vpc id.
     * 
     */
    public Optional<Output<String>> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    /**
     * The vswitch id.
     * 
     */
    @Import(name="vswitchId")
    private @Nullable Output<String> vswitchId;

    /**
     * @return The vswitch id.
     * 
     */
    public Optional<Output<String>> vswitchId() {
        return Optional.ofNullable(this.vswitchId);
    }

    private GetEcsNetworkInterfacesArgs() {}

    private GetEcsNetworkInterfacesArgs(GetEcsNetworkInterfacesArgs $) {
        this.ids = $.ids;
        this.instanceId = $.instanceId;
        this.name = $.name;
        this.nameRegex = $.nameRegex;
        this.networkInterfaceName = $.networkInterfaceName;
        this.outputFile = $.outputFile;
        this.primaryIpAddress = $.primaryIpAddress;
        this.privateIp = $.privateIp;
        this.resourceGroupId = $.resourceGroupId;
        this.securityGroupId = $.securityGroupId;
        this.serviceManaged = $.serviceManaged;
        this.status = $.status;
        this.tags = $.tags;
        this.type = $.type;
        this.vpcId = $.vpcId;
        this.vswitchId = $.vswitchId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetEcsNetworkInterfacesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetEcsNetworkInterfacesArgs $;

        public Builder() {
            $ = new GetEcsNetworkInterfacesArgs();
        }

        public Builder(GetEcsNetworkInterfacesArgs defaults) {
            $ = new GetEcsNetworkInterfacesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of Network Interface IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Network Interface IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids A list of Network Interface IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param instanceId The instance id.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(@Nullable Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId The instance id.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param name Field `name` has been deprecated from provider version 1.123.1. New field `network_interface_name` instead
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
         * @param name Field `name` has been deprecated from provider version 1.123.1. New field `network_interface_name` instead
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

        /**
         * @param nameRegex A regex string to filter results by Network Interface name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable Output<String> nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by Network Interface name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(String nameRegex) {
            return nameRegex(Output.of(nameRegex));
        }

        /**
         * @param networkInterfaceName The network interface name.
         * 
         * @return builder
         * 
         */
        public Builder networkInterfaceName(@Nullable Output<String> networkInterfaceName) {
            $.networkInterfaceName = networkInterfaceName;
            return this;
        }

        /**
         * @param networkInterfaceName The network interface name.
         * 
         * @return builder
         * 
         */
        public Builder networkInterfaceName(String networkInterfaceName) {
            return networkInterfaceName(Output.of(networkInterfaceName));
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(@Nullable Output<String> outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(String outputFile) {
            return outputFile(Output.of(outputFile));
        }

        /**
         * @param primaryIpAddress The primary private IP address of the ENI.
         * 
         * @return builder
         * 
         */
        public Builder primaryIpAddress(@Nullable Output<String> primaryIpAddress) {
            $.primaryIpAddress = primaryIpAddress;
            return this;
        }

        /**
         * @param primaryIpAddress The primary private IP address of the ENI.
         * 
         * @return builder
         * 
         */
        public Builder primaryIpAddress(String primaryIpAddress) {
            return primaryIpAddress(Output.of(primaryIpAddress));
        }

        /**
         * @param privateIp Field `private_ip` has been deprecated from provider version 1.123.1. New field `primary_ip_address` instead
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
         * @param privateIp Field `private_ip` has been deprecated from provider version 1.123.1. New field `primary_ip_address` instead
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

        /**
         * @param resourceGroupId The resource group id.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The resource group id.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        /**
         * @param securityGroupId The security group id.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupId(@Nullable Output<String> securityGroupId) {
            $.securityGroupId = securityGroupId;
            return this;
        }

        /**
         * @param securityGroupId The security group id.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupId(String securityGroupId) {
            return securityGroupId(Output.of(securityGroupId));
        }

        /**
         * @param serviceManaged Whether the user of the elastic network card is a cloud product or a virtual vendor.
         * 
         * @return builder
         * 
         */
        public Builder serviceManaged(@Nullable Output<Boolean> serviceManaged) {
            $.serviceManaged = serviceManaged;
            return this;
        }

        /**
         * @param serviceManaged Whether the user of the elastic network card is a cloud product or a virtual vendor.
         * 
         * @return builder
         * 
         */
        public Builder serviceManaged(Boolean serviceManaged) {
            return serviceManaged(Output.of(serviceManaged));
        }

        /**
         * @param status The status of ENI. Valid Values: `Attaching`, `Available`, `CreateFailed`, `Creating`, `Deleting`, `Detaching`, `InUse`, `Linked`, `Linking`, `Unlinking`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of ENI. Valid Values: `Attaching`, `Available`, `CreateFailed`, `Creating`, `Deleting`, `Detaching`, `InUse`, `Linked`, `Linking`, `Unlinking`.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param tags A map of tags assigned to ENIs.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags assigned to ENIs.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param type The type of ENI. Valid Values: `Primary`, `Secondary`.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The type of ENI. Valid Values: `Primary`, `Secondary`.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param vpcId The vpc id.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(@Nullable Output<String> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vpcId The vpc id.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(String vpcId) {
            return vpcId(Output.of(vpcId));
        }

        /**
         * @param vswitchId The vswitch id.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(@Nullable Output<String> vswitchId) {
            $.vswitchId = vswitchId;
            return this;
        }

        /**
         * @param vswitchId The vswitch id.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(String vswitchId) {
            return vswitchId(Output.of(vswitchId));
        }

        public GetEcsNetworkInterfacesArgs build() {
            return $;
        }
    }

}
