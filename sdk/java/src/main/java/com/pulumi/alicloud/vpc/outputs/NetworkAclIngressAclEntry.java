// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class NetworkAclIngressAclEntry {
    /**
     * @return The description of the network ACL.The description must be 1 to 256 characters in length and cannot start with http:// or https.
     * 
     */
    private @Nullable String description;
    /**
     * @return Name of the outbound rule entry.The name must be 1 to 128 characters in length and cannot start with http:// or https.
     * 
     */
    private @Nullable String networkAclEntryName;
    /**
     * @return Authorization policy. Value:
     * - accept: Allow.
     * - drop: Refused.
     * - accept: Allow.
     * - drop: Refused.
     * 
     */
    private @Nullable String policy;
    /**
     * @return The destination port range of the outbound rule.When the Protocol type of the outbound rule is all, icmp, or gre, the port range is - 1/-1, indicating that the port is not restricted.When the Protocol type of the outbound rule is tcp or udp, the port range is 1 to 65535, and the format is 1/200 or 80/80, indicating port 1 to port 200 or port 80.
     * 
     */
    private @Nullable String port;
    /**
     * @return The protocol type. Value:
     * - icmp: Network Control Message Protocol.
     * - gre: Generic Routing Encapsulation Protocol.
     * - tcp: Transmission Control Protocol.
     * - udp: User Datagram Protocol.
     * - all: Supports all protocols.
     * 
     * - icmp: Network Control Message Protocol.
     * - gre: Generic Routing Encapsulation Protocol.
     * - tcp: Transmission Control Protocol.
     * - udp: User Datagram Protocol.
     * - all: Supports all protocols.
     * 
     */
    private @Nullable String protocol;
    /**
     * @return Source address network segment.
     * 
     */
    private @Nullable String sourceCidrIp;

    private NetworkAclIngressAclEntry() {}
    /**
     * @return The description of the network ACL.The description must be 1 to 256 characters in length and cannot start with http:// or https.
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    /**
     * @return Name of the outbound rule entry.The name must be 1 to 128 characters in length and cannot start with http:// or https.
     * 
     */
    public Optional<String> networkAclEntryName() {
        return Optional.ofNullable(this.networkAclEntryName);
    }
    /**
     * @return Authorization policy. Value:
     * - accept: Allow.
     * - drop: Refused.
     * - accept: Allow.
     * - drop: Refused.
     * 
     */
    public Optional<String> policy() {
        return Optional.ofNullable(this.policy);
    }
    /**
     * @return The destination port range of the outbound rule.When the Protocol type of the outbound rule is all, icmp, or gre, the port range is - 1/-1, indicating that the port is not restricted.When the Protocol type of the outbound rule is tcp or udp, the port range is 1 to 65535, and the format is 1/200 or 80/80, indicating port 1 to port 200 or port 80.
     * 
     */
    public Optional<String> port() {
        return Optional.ofNullable(this.port);
    }
    /**
     * @return The protocol type. Value:
     * - icmp: Network Control Message Protocol.
     * - gre: Generic Routing Encapsulation Protocol.
     * - tcp: Transmission Control Protocol.
     * - udp: User Datagram Protocol.
     * - all: Supports all protocols.
     * 
     * - icmp: Network Control Message Protocol.
     * - gre: Generic Routing Encapsulation Protocol.
     * - tcp: Transmission Control Protocol.
     * - udp: User Datagram Protocol.
     * - all: Supports all protocols.
     * 
     */
    public Optional<String> protocol() {
        return Optional.ofNullable(this.protocol);
    }
    /**
     * @return Source address network segment.
     * 
     */
    public Optional<String> sourceCidrIp() {
        return Optional.ofNullable(this.sourceCidrIp);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(NetworkAclIngressAclEntry defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String description;
        private @Nullable String networkAclEntryName;
        private @Nullable String policy;
        private @Nullable String port;
        private @Nullable String protocol;
        private @Nullable String sourceCidrIp;
        public Builder() {}
        public Builder(NetworkAclIngressAclEntry defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.networkAclEntryName = defaults.networkAclEntryName;
    	      this.policy = defaults.policy;
    	      this.port = defaults.port;
    	      this.protocol = defaults.protocol;
    	      this.sourceCidrIp = defaults.sourceCidrIp;
        }

        @CustomType.Setter
        public Builder description(@Nullable String description) {
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder networkAclEntryName(@Nullable String networkAclEntryName) {
            this.networkAclEntryName = networkAclEntryName;
            return this;
        }
        @CustomType.Setter
        public Builder policy(@Nullable String policy) {
            this.policy = policy;
            return this;
        }
        @CustomType.Setter
        public Builder port(@Nullable String port) {
            this.port = port;
            return this;
        }
        @CustomType.Setter
        public Builder protocol(@Nullable String protocol) {
            this.protocol = protocol;
            return this;
        }
        @CustomType.Setter
        public Builder sourceCidrIp(@Nullable String sourceCidrIp) {
            this.sourceCidrIp = sourceCidrIp;
            return this;
        }
        public NetworkAclIngressAclEntry build() {
            final var o = new NetworkAclIngressAclEntry();
            o.description = description;
            o.networkAclEntryName = networkAclEntryName;
            o.policy = policy;
            o.port = port;
            o.protocol = protocol;
            o.sourceCidrIp = sourceCidrIp;
            return o;
        }
    }
}
