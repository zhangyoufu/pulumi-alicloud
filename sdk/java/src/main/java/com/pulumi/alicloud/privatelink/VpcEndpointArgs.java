// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.privatelink;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VpcEndpointArgs extends com.pulumi.resources.ResourceArgs {

    public static final VpcEndpointArgs Empty = new VpcEndpointArgs();

    /**
     * Specifies whether to perform only a dry run, without performing the actual request. Valid values:
     * - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
     * - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
     * 
     */
    @Import(name="dryRun")
    private @Nullable Output<Boolean> dryRun;

    /**
     * @return Specifies whether to perform only a dry run, without performing the actual request. Valid values:
     * - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
     * - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
     * 
     */
    public Optional<Output<Boolean>> dryRun() {
        return Optional.ofNullable(this.dryRun);
    }

    /**
     * The description of the endpoint.
     * 
     */
    @Import(name="endpointDescription")
    private @Nullable Output<String> endpointDescription;

    /**
     * @return The description of the endpoint.
     * 
     */
    public Optional<Output<String>> endpointDescription() {
        return Optional.ofNullable(this.endpointDescription);
    }

    /**
     * The endpoint type. Only the value: Interface, indicating the Interface endpoint. You can add the service resource types of Application Load Balancer (ALB), Classic Load Balancer (CLB), and Network Load Balancer (NLB).
     * 
     */
    @Import(name="endpointType")
    private @Nullable Output<String> endpointType;

    /**
     * @return The endpoint type. Only the value: Interface, indicating the Interface endpoint. You can add the service resource types of Application Load Balancer (ALB), Classic Load Balancer (CLB), and Network Load Balancer (NLB).
     * 
     */
    public Optional<Output<String>> endpointType() {
        return Optional.ofNullable(this.endpointType);
    }

    /**
     * RAM access policies.
     * 
     */
    @Import(name="policyDocument")
    private @Nullable Output<String> policyDocument;

    /**
     * @return RAM access policies.
     * 
     */
    public Optional<Output<String>> policyDocument() {
        return Optional.ofNullable(this.policyDocument);
    }

    /**
     * Specifies whether to enable user authentication. This parameter is available in Security Token Service (STS) mode. Valid values:
     * - **true**: enables user authentication. After user authentication is enabled, only the user who creates the endpoint can modify or delete the endpoint in STS mode.
     * - **false (default)**: disables user authentication.
     * 
     */
    @Import(name="protectedEnabled")
    private @Nullable Output<Boolean> protectedEnabled;

    /**
     * @return Specifies whether to enable user authentication. This parameter is available in Security Token Service (STS) mode. Valid values:
     * - **true**: enables user authentication. After user authentication is enabled, only the user who creates the endpoint can modify or delete the endpoint in STS mode.
     * - **false (default)**: disables user authentication.
     * 
     */
    public Optional<Output<Boolean>> protectedEnabled() {
        return Optional.ofNullable(this.protectedEnabled);
    }

    /**
     * The resource group ID.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The resource group ID.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * The ID of the security group that is associated with the endpoint ENI. The security group can be used to control data transfer between the VPC and the endpoint ENI.The endpoint can be associated with up to 10 security groups.
     * 
     */
    @Import(name="securityGroupIds", required=true)
    private Output<List<String>> securityGroupIds;

    /**
     * @return The ID of the security group that is associated with the endpoint ENI. The security group can be used to control data transfer between the VPC and the endpoint ENI.The endpoint can be associated with up to 10 security groups.
     * 
     */
    public Output<List<String>> securityGroupIds() {
        return this.securityGroupIds;
    }

    /**
     * The ID of the endpoint service with which the endpoint is associated.
     * 
     */
    @Import(name="serviceId")
    private @Nullable Output<String> serviceId;

    /**
     * @return The ID of the endpoint service with which the endpoint is associated.
     * 
     */
    public Optional<Output<String>> serviceId() {
        return Optional.ofNullable(this.serviceId);
    }

    /**
     * The name of the endpoint service with which the endpoint is associated.
     * 
     */
    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    /**
     * @return The name of the endpoint service with which the endpoint is associated.
     * 
     */
    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    /**
     * The list of tags.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,Object>> tags;

    /**
     * @return The list of tags.
     * 
     */
    public Optional<Output<Map<String,Object>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The name of the endpoint.
     * 
     */
    @Import(name="vpcEndpointName")
    private @Nullable Output<String> vpcEndpointName;

    /**
     * @return The name of the endpoint.
     * 
     */
    public Optional<Output<String>> vpcEndpointName() {
        return Optional.ofNullable(this.vpcEndpointName);
    }

    /**
     * The ID of the VPC to which the endpoint belongs.
     * 
     */
    @Import(name="vpcId", required=true)
    private Output<String> vpcId;

    /**
     * @return The ID of the VPC to which the endpoint belongs.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }

    /**
     * The number of private IP addresses that are assigned to an elastic network interface (ENI) in each zone. Only 1 is returned.
     * 
     */
    @Import(name="zonePrivateIpAddressCount")
    private @Nullable Output<Integer> zonePrivateIpAddressCount;

    /**
     * @return The number of private IP addresses that are assigned to an elastic network interface (ENI) in each zone. Only 1 is returned.
     * 
     */
    public Optional<Output<Integer>> zonePrivateIpAddressCount() {
        return Optional.ofNullable(this.zonePrivateIpAddressCount);
    }

    private VpcEndpointArgs() {}

    private VpcEndpointArgs(VpcEndpointArgs $) {
        this.dryRun = $.dryRun;
        this.endpointDescription = $.endpointDescription;
        this.endpointType = $.endpointType;
        this.policyDocument = $.policyDocument;
        this.protectedEnabled = $.protectedEnabled;
        this.resourceGroupId = $.resourceGroupId;
        this.securityGroupIds = $.securityGroupIds;
        this.serviceId = $.serviceId;
        this.serviceName = $.serviceName;
        this.tags = $.tags;
        this.vpcEndpointName = $.vpcEndpointName;
        this.vpcId = $.vpcId;
        this.zonePrivateIpAddressCount = $.zonePrivateIpAddressCount;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VpcEndpointArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VpcEndpointArgs $;

        public Builder() {
            $ = new VpcEndpointArgs();
        }

        public Builder(VpcEndpointArgs defaults) {
            $ = new VpcEndpointArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dryRun Specifies whether to perform only a dry run, without performing the actual request. Valid values:
         * - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
         * - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
         * 
         * @return builder
         * 
         */
        public Builder dryRun(@Nullable Output<Boolean> dryRun) {
            $.dryRun = dryRun;
            return this;
        }

        /**
         * @param dryRun Specifies whether to perform only a dry run, without performing the actual request. Valid values:
         * - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
         * - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
         * 
         * @return builder
         * 
         */
        public Builder dryRun(Boolean dryRun) {
            return dryRun(Output.of(dryRun));
        }

        /**
         * @param endpointDescription The description of the endpoint.
         * 
         * @return builder
         * 
         */
        public Builder endpointDescription(@Nullable Output<String> endpointDescription) {
            $.endpointDescription = endpointDescription;
            return this;
        }

        /**
         * @param endpointDescription The description of the endpoint.
         * 
         * @return builder
         * 
         */
        public Builder endpointDescription(String endpointDescription) {
            return endpointDescription(Output.of(endpointDescription));
        }

        /**
         * @param endpointType The endpoint type. Only the value: Interface, indicating the Interface endpoint. You can add the service resource types of Application Load Balancer (ALB), Classic Load Balancer (CLB), and Network Load Balancer (NLB).
         * 
         * @return builder
         * 
         */
        public Builder endpointType(@Nullable Output<String> endpointType) {
            $.endpointType = endpointType;
            return this;
        }

        /**
         * @param endpointType The endpoint type. Only the value: Interface, indicating the Interface endpoint. You can add the service resource types of Application Load Balancer (ALB), Classic Load Balancer (CLB), and Network Load Balancer (NLB).
         * 
         * @return builder
         * 
         */
        public Builder endpointType(String endpointType) {
            return endpointType(Output.of(endpointType));
        }

        /**
         * @param policyDocument RAM access policies.
         * 
         * @return builder
         * 
         */
        public Builder policyDocument(@Nullable Output<String> policyDocument) {
            $.policyDocument = policyDocument;
            return this;
        }

        /**
         * @param policyDocument RAM access policies.
         * 
         * @return builder
         * 
         */
        public Builder policyDocument(String policyDocument) {
            return policyDocument(Output.of(policyDocument));
        }

        /**
         * @param protectedEnabled Specifies whether to enable user authentication. This parameter is available in Security Token Service (STS) mode. Valid values:
         * - **true**: enables user authentication. After user authentication is enabled, only the user who creates the endpoint can modify or delete the endpoint in STS mode.
         * - **false (default)**: disables user authentication.
         * 
         * @return builder
         * 
         */
        public Builder protectedEnabled(@Nullable Output<Boolean> protectedEnabled) {
            $.protectedEnabled = protectedEnabled;
            return this;
        }

        /**
         * @param protectedEnabled Specifies whether to enable user authentication. This parameter is available in Security Token Service (STS) mode. Valid values:
         * - **true**: enables user authentication. After user authentication is enabled, only the user who creates the endpoint can modify or delete the endpoint in STS mode.
         * - **false (default)**: disables user authentication.
         * 
         * @return builder
         * 
         */
        public Builder protectedEnabled(Boolean protectedEnabled) {
            return protectedEnabled(Output.of(protectedEnabled));
        }

        /**
         * @param resourceGroupId The resource group ID.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The resource group ID.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        /**
         * @param securityGroupIds The ID of the security group that is associated with the endpoint ENI. The security group can be used to control data transfer between the VPC and the endpoint ENI.The endpoint can be associated with up to 10 security groups.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupIds(Output<List<String>> securityGroupIds) {
            $.securityGroupIds = securityGroupIds;
            return this;
        }

        /**
         * @param securityGroupIds The ID of the security group that is associated with the endpoint ENI. The security group can be used to control data transfer between the VPC and the endpoint ENI.The endpoint can be associated with up to 10 security groups.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupIds(List<String> securityGroupIds) {
            return securityGroupIds(Output.of(securityGroupIds));
        }

        /**
         * @param securityGroupIds The ID of the security group that is associated with the endpoint ENI. The security group can be used to control data transfer between the VPC and the endpoint ENI.The endpoint can be associated with up to 10 security groups.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupIds(String... securityGroupIds) {
            return securityGroupIds(List.of(securityGroupIds));
        }

        /**
         * @param serviceId The ID of the endpoint service with which the endpoint is associated.
         * 
         * @return builder
         * 
         */
        public Builder serviceId(@Nullable Output<String> serviceId) {
            $.serviceId = serviceId;
            return this;
        }

        /**
         * @param serviceId The ID of the endpoint service with which the endpoint is associated.
         * 
         * @return builder
         * 
         */
        public Builder serviceId(String serviceId) {
            return serviceId(Output.of(serviceId));
        }

        /**
         * @param serviceName The name of the endpoint service with which the endpoint is associated.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(@Nullable Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The name of the endpoint service with which the endpoint is associated.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        /**
         * @param tags The list of tags.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,Object>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The list of tags.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,Object> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param vpcEndpointName The name of the endpoint.
         * 
         * @return builder
         * 
         */
        public Builder vpcEndpointName(@Nullable Output<String> vpcEndpointName) {
            $.vpcEndpointName = vpcEndpointName;
            return this;
        }

        /**
         * @param vpcEndpointName The name of the endpoint.
         * 
         * @return builder
         * 
         */
        public Builder vpcEndpointName(String vpcEndpointName) {
            return vpcEndpointName(Output.of(vpcEndpointName));
        }

        /**
         * @param vpcId The ID of the VPC to which the endpoint belongs.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(Output<String> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vpcId The ID of the VPC to which the endpoint belongs.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(String vpcId) {
            return vpcId(Output.of(vpcId));
        }

        /**
         * @param zonePrivateIpAddressCount The number of private IP addresses that are assigned to an elastic network interface (ENI) in each zone. Only 1 is returned.
         * 
         * @return builder
         * 
         */
        public Builder zonePrivateIpAddressCount(@Nullable Output<Integer> zonePrivateIpAddressCount) {
            $.zonePrivateIpAddressCount = zonePrivateIpAddressCount;
            return this;
        }

        /**
         * @param zonePrivateIpAddressCount The number of private IP addresses that are assigned to an elastic network interface (ENI) in each zone. Only 1 is returned.
         * 
         * @return builder
         * 
         */
        public Builder zonePrivateIpAddressCount(Integer zonePrivateIpAddressCount) {
            return zonePrivateIpAddressCount(Output.of(zonePrivateIpAddressCount));
        }

        public VpcEndpointArgs build() {
            if ($.securityGroupIds == null) {
                throw new MissingRequiredPropertyException("VpcEndpointArgs", "securityGroupIds");
            }
            if ($.vpcId == null) {
                throw new MissingRequiredPropertyException("VpcEndpointArgs", "vpcId");
            }
            return $;
        }
    }

}
