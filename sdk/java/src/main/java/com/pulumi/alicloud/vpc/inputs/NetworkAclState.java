// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.inputs;

import com.pulumi.alicloud.vpc.inputs.NetworkAclEgressAclEntryArgs;
import com.pulumi.alicloud.vpc.inputs.NetworkAclIngressAclEntryArgs;
import com.pulumi.alicloud.vpc.inputs.NetworkAclResourceArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NetworkAclState extends com.pulumi.resources.ResourceArgs {

    public static final NetworkAclState Empty = new NetworkAclState();

    /**
     * The creation time of the resource.
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return The creation time of the resource.
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * The description of the network ACL.  The description must be 1 to 256 characters in length and cannot start with http:// or https.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the network ACL.  The description must be 1 to 256 characters in length and cannot start with http:// or https.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Out direction rule information. See `egress_acl_entries` below.
     * 
     */
    @Import(name="egressAclEntries")
    private @Nullable Output<List<NetworkAclEgressAclEntryArgs>> egressAclEntries;

    /**
     * @return Out direction rule information. See `egress_acl_entries` below.
     * 
     */
    public Optional<Output<List<NetworkAclEgressAclEntryArgs>>> egressAclEntries() {
        return Optional.ofNullable(this.egressAclEntries);
    }

    /**
     * Inward direction rule information. See `ingress_acl_entries` below.
     * 
     */
    @Import(name="ingressAclEntries")
    private @Nullable Output<List<NetworkAclIngressAclEntryArgs>> ingressAclEntries;

    /**
     * @return Inward direction rule information. See `ingress_acl_entries` below.
     * 
     */
    public Optional<Output<List<NetworkAclIngressAclEntryArgs>>> ingressAclEntries() {
        return Optional.ofNullable(this.ingressAclEntries);
    }

    /**
     * . Field &#39;name&#39; has been deprecated from provider version 1.122.0. New field &#39;network_acl_name&#39; instead.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated since provider version 1.122.0. New field &#39;network_acl_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated since provider version 1.122.0. New field 'network_acl_name' instead. */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return . Field &#39;name&#39; has been deprecated from provider version 1.122.0. New field &#39;network_acl_name&#39; instead.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated since provider version 1.122.0. New field &#39;network_acl_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated since provider version 1.122.0. New field 'network_acl_name' instead. */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The name of the network ACL.  The name must be 1 to 128 characters in length and cannot start with http:// or https.
     * 
     */
    @Import(name="networkAclName")
    private @Nullable Output<String> networkAclName;

    /**
     * @return The name of the network ACL.  The name must be 1 to 128 characters in length and cannot start with http:// or https.
     * 
     */
    public Optional<Output<String>> networkAclName() {
        return Optional.ofNullable(this.networkAclName);
    }

    /**
     * The associated resource. See `resources` below.
     * 
     */
    @Import(name="resources")
    private @Nullable Output<List<NetworkAclResourceArgs>> resources;

    /**
     * @return The associated resource. See `resources` below.
     * 
     */
    public Optional<Output<List<NetworkAclResourceArgs>>> resources() {
        return Optional.ofNullable(this.resources);
    }

    /**
     * SOURCE NetworkAcl specified by CopyNetworkAclEntries.
     * 
     */
    @Import(name="sourceNetworkAclId")
    private @Nullable Output<String> sourceNetworkAclId;

    /**
     * @return SOURCE NetworkAcl specified by CopyNetworkAclEntries.
     * 
     */
    public Optional<Output<String>> sourceNetworkAclId() {
        return Optional.ofNullable(this.sourceNetworkAclId);
    }

    /**
     * The state of the network ACL.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The state of the network ACL.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The tags of this resource.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,Object>> tags;

    /**
     * @return The tags of this resource.
     * 
     */
    public Optional<Output<Map<String,Object>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The ID of the associated VPC.
     * 
     * The following arguments will be discarded. Please use new fields as soon as possible:
     * 
     */
    @Import(name="vpcId")
    private @Nullable Output<String> vpcId;

    /**
     * @return The ID of the associated VPC.
     * 
     * The following arguments will be discarded. Please use new fields as soon as possible:
     * 
     */
    public Optional<Output<String>> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    private NetworkAclState() {}

    private NetworkAclState(NetworkAclState $) {
        this.createTime = $.createTime;
        this.description = $.description;
        this.egressAclEntries = $.egressAclEntries;
        this.ingressAclEntries = $.ingressAclEntries;
        this.name = $.name;
        this.networkAclName = $.networkAclName;
        this.resources = $.resources;
        this.sourceNetworkAclId = $.sourceNetworkAclId;
        this.status = $.status;
        this.tags = $.tags;
        this.vpcId = $.vpcId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NetworkAclState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NetworkAclState $;

        public Builder() {
            $ = new NetworkAclState();
        }

        public Builder(NetworkAclState defaults) {
            $ = new NetworkAclState(Objects.requireNonNull(defaults));
        }

        /**
         * @param createTime The creation time of the resource.
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime The creation time of the resource.
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param description The description of the network ACL.  The description must be 1 to 256 characters in length and cannot start with http:// or https.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the network ACL.  The description must be 1 to 256 characters in length and cannot start with http:// or https.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param egressAclEntries Out direction rule information. See `egress_acl_entries` below.
         * 
         * @return builder
         * 
         */
        public Builder egressAclEntries(@Nullable Output<List<NetworkAclEgressAclEntryArgs>> egressAclEntries) {
            $.egressAclEntries = egressAclEntries;
            return this;
        }

        /**
         * @param egressAclEntries Out direction rule information. See `egress_acl_entries` below.
         * 
         * @return builder
         * 
         */
        public Builder egressAclEntries(List<NetworkAclEgressAclEntryArgs> egressAclEntries) {
            return egressAclEntries(Output.of(egressAclEntries));
        }

        /**
         * @param egressAclEntries Out direction rule information. See `egress_acl_entries` below.
         * 
         * @return builder
         * 
         */
        public Builder egressAclEntries(NetworkAclEgressAclEntryArgs... egressAclEntries) {
            return egressAclEntries(List.of(egressAclEntries));
        }

        /**
         * @param ingressAclEntries Inward direction rule information. See `ingress_acl_entries` below.
         * 
         * @return builder
         * 
         */
        public Builder ingressAclEntries(@Nullable Output<List<NetworkAclIngressAclEntryArgs>> ingressAclEntries) {
            $.ingressAclEntries = ingressAclEntries;
            return this;
        }

        /**
         * @param ingressAclEntries Inward direction rule information. See `ingress_acl_entries` below.
         * 
         * @return builder
         * 
         */
        public Builder ingressAclEntries(List<NetworkAclIngressAclEntryArgs> ingressAclEntries) {
            return ingressAclEntries(Output.of(ingressAclEntries));
        }

        /**
         * @param ingressAclEntries Inward direction rule information. See `ingress_acl_entries` below.
         * 
         * @return builder
         * 
         */
        public Builder ingressAclEntries(NetworkAclIngressAclEntryArgs... ingressAclEntries) {
            return ingressAclEntries(List.of(ingressAclEntries));
        }

        /**
         * @param name . Field &#39;name&#39; has been deprecated from provider version 1.122.0. New field &#39;network_acl_name&#39; instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;name&#39; has been deprecated since provider version 1.122.0. New field &#39;network_acl_name&#39; instead.
         * 
         */
        @Deprecated /* Field 'name' has been deprecated since provider version 1.122.0. New field 'network_acl_name' instead. */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name . Field &#39;name&#39; has been deprecated from provider version 1.122.0. New field &#39;network_acl_name&#39; instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;name&#39; has been deprecated since provider version 1.122.0. New field &#39;network_acl_name&#39; instead.
         * 
         */
        @Deprecated /* Field 'name' has been deprecated since provider version 1.122.0. New field 'network_acl_name' instead. */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param networkAclName The name of the network ACL.  The name must be 1 to 128 characters in length and cannot start with http:// or https.
         * 
         * @return builder
         * 
         */
        public Builder networkAclName(@Nullable Output<String> networkAclName) {
            $.networkAclName = networkAclName;
            return this;
        }

        /**
         * @param networkAclName The name of the network ACL.  The name must be 1 to 128 characters in length and cannot start with http:// or https.
         * 
         * @return builder
         * 
         */
        public Builder networkAclName(String networkAclName) {
            return networkAclName(Output.of(networkAclName));
        }

        /**
         * @param resources The associated resource. See `resources` below.
         * 
         * @return builder
         * 
         */
        public Builder resources(@Nullable Output<List<NetworkAclResourceArgs>> resources) {
            $.resources = resources;
            return this;
        }

        /**
         * @param resources The associated resource. See `resources` below.
         * 
         * @return builder
         * 
         */
        public Builder resources(List<NetworkAclResourceArgs> resources) {
            return resources(Output.of(resources));
        }

        /**
         * @param resources The associated resource. See `resources` below.
         * 
         * @return builder
         * 
         */
        public Builder resources(NetworkAclResourceArgs... resources) {
            return resources(List.of(resources));
        }

        /**
         * @param sourceNetworkAclId SOURCE NetworkAcl specified by CopyNetworkAclEntries.
         * 
         * @return builder
         * 
         */
        public Builder sourceNetworkAclId(@Nullable Output<String> sourceNetworkAclId) {
            $.sourceNetworkAclId = sourceNetworkAclId;
            return this;
        }

        /**
         * @param sourceNetworkAclId SOURCE NetworkAcl specified by CopyNetworkAclEntries.
         * 
         * @return builder
         * 
         */
        public Builder sourceNetworkAclId(String sourceNetworkAclId) {
            return sourceNetworkAclId(Output.of(sourceNetworkAclId));
        }

        /**
         * @param status The state of the network ACL.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The state of the network ACL.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param tags The tags of this resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,Object>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The tags of this resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,Object> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param vpcId The ID of the associated VPC.
         * 
         * The following arguments will be discarded. Please use new fields as soon as possible:
         * 
         * @return builder
         * 
         */
        public Builder vpcId(@Nullable Output<String> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vpcId The ID of the associated VPC.
         * 
         * The following arguments will be discarded. Please use new fields as soon as possible:
         * 
         * @return builder
         * 
         */
        public Builder vpcId(String vpcId) {
            return vpcId(Output.of(vpcId));
        }

        public NetworkAclState build() {
            return $;
        }
    }

}
