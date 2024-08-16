// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GatewayEndpointState extends com.pulumi.resources.ResourceArgs {

    public static final GatewayEndpointState Empty = new GatewayEndpointState();

    /**
     * The creation time of the gateway endpoint.
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return The creation time of the gateway endpoint.
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * The description of the gateway endpoint.
     * 
     */
    @Import(name="gatewayEndpointDescrption")
    private @Nullable Output<String> gatewayEndpointDescrption;

    /**
     * @return The description of the gateway endpoint.
     * 
     */
    public Optional<Output<String>> gatewayEndpointDescrption() {
        return Optional.ofNullable(this.gatewayEndpointDescrption);
    }

    /**
     * The name of the gateway endpoint.
     * 
     */
    @Import(name="gatewayEndpointName")
    private @Nullable Output<String> gatewayEndpointName;

    /**
     * @return The name of the gateway endpoint.
     * 
     */
    public Optional<Output<String>> gatewayEndpointName() {
        return Optional.ofNullable(this.gatewayEndpointName);
    }

    /**
     * Access control policies for cloud services. This parameter is required when the cloud service is oss. For details about the syntax and structure of access policies, see [syntax and structure of permission Policies](https://help.aliyun.com/document_detail/93739.html).
     * 
     */
    @Import(name="policyDocument")
    private @Nullable Output<String> policyDocument;

    /**
     * @return Access control policies for cloud services. This parameter is required when the cloud service is oss. For details about the syntax and structure of access policies, see [syntax and structure of permission Policies](https://help.aliyun.com/document_detail/93739.html).
     * 
     */
    public Optional<Output<String>> policyDocument() {
        return Optional.ofNullable(this.policyDocument);
    }

    /**
     * The ID of the resource group to which the instance belongs.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group to which the instance belongs.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * The name of endpoint service.
     * 
     */
    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    /**
     * @return The name of endpoint service.
     * 
     */
    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    /**
     * The status of VPC gateway endpoint.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of VPC gateway endpoint.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The tags of the resource.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return The tags of the resource.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The ID of the VPC.
     * 
     */
    @Import(name="vpcId")
    private @Nullable Output<String> vpcId;

    /**
     * @return The ID of the VPC.
     * 
     */
    public Optional<Output<String>> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    private GatewayEndpointState() {}

    private GatewayEndpointState(GatewayEndpointState $) {
        this.createTime = $.createTime;
        this.gatewayEndpointDescrption = $.gatewayEndpointDescrption;
        this.gatewayEndpointName = $.gatewayEndpointName;
        this.policyDocument = $.policyDocument;
        this.resourceGroupId = $.resourceGroupId;
        this.serviceName = $.serviceName;
        this.status = $.status;
        this.tags = $.tags;
        this.vpcId = $.vpcId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GatewayEndpointState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GatewayEndpointState $;

        public Builder() {
            $ = new GatewayEndpointState();
        }

        public Builder(GatewayEndpointState defaults) {
            $ = new GatewayEndpointState(Objects.requireNonNull(defaults));
        }

        /**
         * @param createTime The creation time of the gateway endpoint.
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime The creation time of the gateway endpoint.
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param gatewayEndpointDescrption The description of the gateway endpoint.
         * 
         * @return builder
         * 
         */
        public Builder gatewayEndpointDescrption(@Nullable Output<String> gatewayEndpointDescrption) {
            $.gatewayEndpointDescrption = gatewayEndpointDescrption;
            return this;
        }

        /**
         * @param gatewayEndpointDescrption The description of the gateway endpoint.
         * 
         * @return builder
         * 
         */
        public Builder gatewayEndpointDescrption(String gatewayEndpointDescrption) {
            return gatewayEndpointDescrption(Output.of(gatewayEndpointDescrption));
        }

        /**
         * @param gatewayEndpointName The name of the gateway endpoint.
         * 
         * @return builder
         * 
         */
        public Builder gatewayEndpointName(@Nullable Output<String> gatewayEndpointName) {
            $.gatewayEndpointName = gatewayEndpointName;
            return this;
        }

        /**
         * @param gatewayEndpointName The name of the gateway endpoint.
         * 
         * @return builder
         * 
         */
        public Builder gatewayEndpointName(String gatewayEndpointName) {
            return gatewayEndpointName(Output.of(gatewayEndpointName));
        }

        /**
         * @param policyDocument Access control policies for cloud services. This parameter is required when the cloud service is oss. For details about the syntax and structure of access policies, see [syntax and structure of permission Policies](https://help.aliyun.com/document_detail/93739.html).
         * 
         * @return builder
         * 
         */
        public Builder policyDocument(@Nullable Output<String> policyDocument) {
            $.policyDocument = policyDocument;
            return this;
        }

        /**
         * @param policyDocument Access control policies for cloud services. This parameter is required when the cloud service is oss. For details about the syntax and structure of access policies, see [syntax and structure of permission Policies](https://help.aliyun.com/document_detail/93739.html).
         * 
         * @return builder
         * 
         */
        public Builder policyDocument(String policyDocument) {
            return policyDocument(Output.of(policyDocument));
        }

        /**
         * @param resourceGroupId The ID of the resource group to which the instance belongs.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The ID of the resource group to which the instance belongs.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        /**
         * @param serviceName The name of endpoint service.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(@Nullable Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The name of endpoint service.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        /**
         * @param status The status of VPC gateway endpoint.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of VPC gateway endpoint.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param tags The tags of the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The tags of the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param vpcId The ID of the VPC.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(@Nullable Output<String> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vpcId The ID of the VPC.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(String vpcId) {
            return vpcId(Output.of(vpcId));
        }

        public GatewayEndpointState build() {
            return $;
        }
    }

}
