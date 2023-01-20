// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetTransitRouterMulticastDomainPeerMembersMember {
    /**
     * @return The IP address of the multicast group to which the multicast member belongs. Value range: **224.0.0.1** to **239.255.255.254**.If the multicast group you specified does not exist in the current multicast domain, the system will automatically create a new multicast group for you in the current multicast domain.
     * 
     */
    private String groupIpAddress;
    /**
     * @return The `key` of the resource supplied above.The value is formulated as `&lt;transit_router_multicast_domain_id&gt;:&lt;group_ip_address&gt;:&lt;peer_transit_router_multicast_domain_id&gt;`.
     * 
     */
    private String id;
    /**
     * @return The multicast domain ID of the peer transit router.
     * 
     */
    private String peerTransitRouterMulticastDomainId;
    /**
     * @return The status of the resource
     * 
     */
    private String status;
    /**
     * @return The ID of the multicast domain to which the multicast member belongs.
     * 
     */
    private String transitRouterMulticastDomainId;

    private GetTransitRouterMulticastDomainPeerMembersMember() {}
    /**
     * @return The IP address of the multicast group to which the multicast member belongs. Value range: **224.0.0.1** to **239.255.255.254**.If the multicast group you specified does not exist in the current multicast domain, the system will automatically create a new multicast group for you in the current multicast domain.
     * 
     */
    public String groupIpAddress() {
        return this.groupIpAddress;
    }
    /**
     * @return The `key` of the resource supplied above.The value is formulated as `&lt;transit_router_multicast_domain_id&gt;:&lt;group_ip_address&gt;:&lt;peer_transit_router_multicast_domain_id&gt;`.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The multicast domain ID of the peer transit router.
     * 
     */
    public String peerTransitRouterMulticastDomainId() {
        return this.peerTransitRouterMulticastDomainId;
    }
    /**
     * @return The status of the resource
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return The ID of the multicast domain to which the multicast member belongs.
     * 
     */
    public String transitRouterMulticastDomainId() {
        return this.transitRouterMulticastDomainId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetTransitRouterMulticastDomainPeerMembersMember defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String groupIpAddress;
        private String id;
        private String peerTransitRouterMulticastDomainId;
        private String status;
        private String transitRouterMulticastDomainId;
        public Builder() {}
        public Builder(GetTransitRouterMulticastDomainPeerMembersMember defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.groupIpAddress = defaults.groupIpAddress;
    	      this.id = defaults.id;
    	      this.peerTransitRouterMulticastDomainId = defaults.peerTransitRouterMulticastDomainId;
    	      this.status = defaults.status;
    	      this.transitRouterMulticastDomainId = defaults.transitRouterMulticastDomainId;
        }

        @CustomType.Setter
        public Builder groupIpAddress(String groupIpAddress) {
            this.groupIpAddress = Objects.requireNonNull(groupIpAddress);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder peerTransitRouterMulticastDomainId(String peerTransitRouterMulticastDomainId) {
            this.peerTransitRouterMulticastDomainId = Objects.requireNonNull(peerTransitRouterMulticastDomainId);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder transitRouterMulticastDomainId(String transitRouterMulticastDomainId) {
            this.transitRouterMulticastDomainId = Objects.requireNonNull(transitRouterMulticastDomainId);
            return this;
        }
        public GetTransitRouterMulticastDomainPeerMembersMember build() {
            final var o = new GetTransitRouterMulticastDomainPeerMembersMember();
            o.groupIpAddress = groupIpAddress;
            o.id = id;
            o.peerTransitRouterMulticastDomainId = peerTransitRouterMulticastDomainId;
            o.status = status;
            o.transitRouterMulticastDomainId = transitRouterMulticastDomainId;
            return o;
        }
    }
}
