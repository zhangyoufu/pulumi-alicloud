// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs.outputs;

import com.pulumi.alicloud.ecs.outputs.GetNetworkInterfacesInterfaceAssociatedPublicIp;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetNetworkInterfacesInterface {
    private List<GetNetworkInterfacesInterfaceAssociatedPublicIp> associatedPublicIps;
    /**
     * @return Creation time of the ENI.
     * 
     */
    private String creationTime;
    /**
     * @return Description of the ENI.
     * 
     */
    private String description;
    /**
     * @return ID of the ENI.
     * 
     */
    private String id;
    /**
     * @return ID of the instance that the ENI is attached to.
     * 
     */
    private String instanceId;
    private List<String> ipv6Sets;
    /**
     * @return MAC address of the ENI.
     * 
     */
    private String mac;
    /**
     * @return Name of the ENI.
     * 
     */
    private String name;
    private String networkInterfaceId;
    private String networkInterfaceName;
    private String networkInterfaceTrafficMode;
    private String ownerId;
    private String primaryIpAddress;
    /**
     * @return Primary private IP of the ENI.
     * 
     */
    private String privateIp;
    private List<String> privateIpAddresses;
    /**
     * @return A list of secondary private IP address that is assigned to the ENI.
     * 
     */
    private List<String> privateIps;
    private Integer queueNumber;
    /**
     * @return The Id of resource group.
     * 
     */
    private String resourceGroupId;
    private List<String> securityGroupIds;
    /**
     * @return A list of security group that the ENI belongs to.
     * 
     */
    private List<String> securityGroups;
    private Integer serviceId;
    private Boolean serviceManaged;
    /**
     * @return Current status of the ENI.
     * 
     */
    private String status;
    /**
     * @return A map of tags assigned to the ENI.
     * 
     */
    private Map<String,String> tags;
    private String type;
    /**
     * @return ID of the VPC that the ENI belongs to.
     * 
     */
    private String vpcId;
    /**
     * @return ID of the vSwitch that the ENI is linked to.
     * 
     */
    private String vswitchId;
    /**
     * @return ID of the availability zone that the ENI belongs to.
     * 
     */
    private String zoneId;

    private GetNetworkInterfacesInterface() {}
    public List<GetNetworkInterfacesInterfaceAssociatedPublicIp> associatedPublicIps() {
        return this.associatedPublicIps;
    }
    /**
     * @return Creation time of the ENI.
     * 
     */
    public String creationTime() {
        return this.creationTime;
    }
    /**
     * @return Description of the ENI.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return ID of the ENI.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return ID of the instance that the ENI is attached to.
     * 
     */
    public String instanceId() {
        return this.instanceId;
    }
    public List<String> ipv6Sets() {
        return this.ipv6Sets;
    }
    /**
     * @return MAC address of the ENI.
     * 
     */
    public String mac() {
        return this.mac;
    }
    /**
     * @return Name of the ENI.
     * 
     */
    public String name() {
        return this.name;
    }
    public String networkInterfaceId() {
        return this.networkInterfaceId;
    }
    public String networkInterfaceName() {
        return this.networkInterfaceName;
    }
    public String networkInterfaceTrafficMode() {
        return this.networkInterfaceTrafficMode;
    }
    public String ownerId() {
        return this.ownerId;
    }
    public String primaryIpAddress() {
        return this.primaryIpAddress;
    }
    /**
     * @return Primary private IP of the ENI.
     * 
     */
    public String privateIp() {
        return this.privateIp;
    }
    public List<String> privateIpAddresses() {
        return this.privateIpAddresses;
    }
    /**
     * @return A list of secondary private IP address that is assigned to the ENI.
     * 
     */
    public List<String> privateIps() {
        return this.privateIps;
    }
    public Integer queueNumber() {
        return this.queueNumber;
    }
    /**
     * @return The Id of resource group.
     * 
     */
    public String resourceGroupId() {
        return this.resourceGroupId;
    }
    public List<String> securityGroupIds() {
        return this.securityGroupIds;
    }
    /**
     * @return A list of security group that the ENI belongs to.
     * 
     */
    public List<String> securityGroups() {
        return this.securityGroups;
    }
    public Integer serviceId() {
        return this.serviceId;
    }
    public Boolean serviceManaged() {
        return this.serviceManaged;
    }
    /**
     * @return Current status of the ENI.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return A map of tags assigned to the ENI.
     * 
     */
    public Map<String,String> tags() {
        return this.tags;
    }
    public String type() {
        return this.type;
    }
    /**
     * @return ID of the VPC that the ENI belongs to.
     * 
     */
    public String vpcId() {
        return this.vpcId;
    }
    /**
     * @return ID of the vSwitch that the ENI is linked to.
     * 
     */
    public String vswitchId() {
        return this.vswitchId;
    }
    /**
     * @return ID of the availability zone that the ENI belongs to.
     * 
     */
    public String zoneId() {
        return this.zoneId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetNetworkInterfacesInterface defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetNetworkInterfacesInterfaceAssociatedPublicIp> associatedPublicIps;
        private String creationTime;
        private String description;
        private String id;
        private String instanceId;
        private List<String> ipv6Sets;
        private String mac;
        private String name;
        private String networkInterfaceId;
        private String networkInterfaceName;
        private String networkInterfaceTrafficMode;
        private String ownerId;
        private String primaryIpAddress;
        private String privateIp;
        private List<String> privateIpAddresses;
        private List<String> privateIps;
        private Integer queueNumber;
        private String resourceGroupId;
        private List<String> securityGroupIds;
        private List<String> securityGroups;
        private Integer serviceId;
        private Boolean serviceManaged;
        private String status;
        private Map<String,String> tags;
        private String type;
        private String vpcId;
        private String vswitchId;
        private String zoneId;
        public Builder() {}
        public Builder(GetNetworkInterfacesInterface defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.associatedPublicIps = defaults.associatedPublicIps;
    	      this.creationTime = defaults.creationTime;
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.instanceId = defaults.instanceId;
    	      this.ipv6Sets = defaults.ipv6Sets;
    	      this.mac = defaults.mac;
    	      this.name = defaults.name;
    	      this.networkInterfaceId = defaults.networkInterfaceId;
    	      this.networkInterfaceName = defaults.networkInterfaceName;
    	      this.networkInterfaceTrafficMode = defaults.networkInterfaceTrafficMode;
    	      this.ownerId = defaults.ownerId;
    	      this.primaryIpAddress = defaults.primaryIpAddress;
    	      this.privateIp = defaults.privateIp;
    	      this.privateIpAddresses = defaults.privateIpAddresses;
    	      this.privateIps = defaults.privateIps;
    	      this.queueNumber = defaults.queueNumber;
    	      this.resourceGroupId = defaults.resourceGroupId;
    	      this.securityGroupIds = defaults.securityGroupIds;
    	      this.securityGroups = defaults.securityGroups;
    	      this.serviceId = defaults.serviceId;
    	      this.serviceManaged = defaults.serviceManaged;
    	      this.status = defaults.status;
    	      this.tags = defaults.tags;
    	      this.type = defaults.type;
    	      this.vpcId = defaults.vpcId;
    	      this.vswitchId = defaults.vswitchId;
    	      this.zoneId = defaults.zoneId;
        }

        @CustomType.Setter
        public Builder associatedPublicIps(List<GetNetworkInterfacesInterfaceAssociatedPublicIp> associatedPublicIps) {
            if (associatedPublicIps == null) {
              throw new MissingRequiredPropertyException("GetNetworkInterfacesInterface", "associatedPublicIps");
            }
            this.associatedPublicIps = associatedPublicIps;
            return this;
        }
        public Builder associatedPublicIps(GetNetworkInterfacesInterfaceAssociatedPublicIp... associatedPublicIps) {
            return associatedPublicIps(List.of(associatedPublicIps));
        }
        @CustomType.Setter
        public Builder creationTime(String creationTime) {
            if (creationTime == null) {
              throw new MissingRequiredPropertyException("GetNetworkInterfacesInterface", "creationTime");
            }
            this.creationTime = creationTime;
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetNetworkInterfacesInterface", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetNetworkInterfacesInterface", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder instanceId(String instanceId) {
            if (instanceId == null) {
              throw new MissingRequiredPropertyException("GetNetworkInterfacesInterface", "instanceId");
            }
            this.instanceId = instanceId;
            return this;
        }
        @CustomType.Setter
        public Builder ipv6Sets(List<String> ipv6Sets) {
            if (ipv6Sets == null) {
              throw new MissingRequiredPropertyException("GetNetworkInterfacesInterface", "ipv6Sets");
            }
            this.ipv6Sets = ipv6Sets;
            return this;
        }
        public Builder ipv6Sets(String... ipv6Sets) {
            return ipv6Sets(List.of(ipv6Sets));
        }
        @CustomType.Setter
        public Builder mac(String mac) {
            if (mac == null) {
              throw new MissingRequiredPropertyException("GetNetworkInterfacesInterface", "mac");
            }
            this.mac = mac;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetNetworkInterfacesInterface", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder networkInterfaceId(String networkInterfaceId) {
            if (networkInterfaceId == null) {
              throw new MissingRequiredPropertyException("GetNetworkInterfacesInterface", "networkInterfaceId");
            }
            this.networkInterfaceId = networkInterfaceId;
            return this;
        }
        @CustomType.Setter
        public Builder networkInterfaceName(String networkInterfaceName) {
            if (networkInterfaceName == null) {
              throw new MissingRequiredPropertyException("GetNetworkInterfacesInterface", "networkInterfaceName");
            }
            this.networkInterfaceName = networkInterfaceName;
            return this;
        }
        @CustomType.Setter
        public Builder networkInterfaceTrafficMode(String networkInterfaceTrafficMode) {
            if (networkInterfaceTrafficMode == null) {
              throw new MissingRequiredPropertyException("GetNetworkInterfacesInterface", "networkInterfaceTrafficMode");
            }
            this.networkInterfaceTrafficMode = networkInterfaceTrafficMode;
            return this;
        }
        @CustomType.Setter
        public Builder ownerId(String ownerId) {
            if (ownerId == null) {
              throw new MissingRequiredPropertyException("GetNetworkInterfacesInterface", "ownerId");
            }
            this.ownerId = ownerId;
            return this;
        }
        @CustomType.Setter
        public Builder primaryIpAddress(String primaryIpAddress) {
            if (primaryIpAddress == null) {
              throw new MissingRequiredPropertyException("GetNetworkInterfacesInterface", "primaryIpAddress");
            }
            this.primaryIpAddress = primaryIpAddress;
            return this;
        }
        @CustomType.Setter
        public Builder privateIp(String privateIp) {
            if (privateIp == null) {
              throw new MissingRequiredPropertyException("GetNetworkInterfacesInterface", "privateIp");
            }
            this.privateIp = privateIp;
            return this;
        }
        @CustomType.Setter
        public Builder privateIpAddresses(List<String> privateIpAddresses) {
            if (privateIpAddresses == null) {
              throw new MissingRequiredPropertyException("GetNetworkInterfacesInterface", "privateIpAddresses");
            }
            this.privateIpAddresses = privateIpAddresses;
            return this;
        }
        public Builder privateIpAddresses(String... privateIpAddresses) {
            return privateIpAddresses(List.of(privateIpAddresses));
        }
        @CustomType.Setter
        public Builder privateIps(List<String> privateIps) {
            if (privateIps == null) {
              throw new MissingRequiredPropertyException("GetNetworkInterfacesInterface", "privateIps");
            }
            this.privateIps = privateIps;
            return this;
        }
        public Builder privateIps(String... privateIps) {
            return privateIps(List.of(privateIps));
        }
        @CustomType.Setter
        public Builder queueNumber(Integer queueNumber) {
            if (queueNumber == null) {
              throw new MissingRequiredPropertyException("GetNetworkInterfacesInterface", "queueNumber");
            }
            this.queueNumber = queueNumber;
            return this;
        }
        @CustomType.Setter
        public Builder resourceGroupId(String resourceGroupId) {
            if (resourceGroupId == null) {
              throw new MissingRequiredPropertyException("GetNetworkInterfacesInterface", "resourceGroupId");
            }
            this.resourceGroupId = resourceGroupId;
            return this;
        }
        @CustomType.Setter
        public Builder securityGroupIds(List<String> securityGroupIds) {
            if (securityGroupIds == null) {
              throw new MissingRequiredPropertyException("GetNetworkInterfacesInterface", "securityGroupIds");
            }
            this.securityGroupIds = securityGroupIds;
            return this;
        }
        public Builder securityGroupIds(String... securityGroupIds) {
            return securityGroupIds(List.of(securityGroupIds));
        }
        @CustomType.Setter
        public Builder securityGroups(List<String> securityGroups) {
            if (securityGroups == null) {
              throw new MissingRequiredPropertyException("GetNetworkInterfacesInterface", "securityGroups");
            }
            this.securityGroups = securityGroups;
            return this;
        }
        public Builder securityGroups(String... securityGroups) {
            return securityGroups(List.of(securityGroups));
        }
        @CustomType.Setter
        public Builder serviceId(Integer serviceId) {
            if (serviceId == null) {
              throw new MissingRequiredPropertyException("GetNetworkInterfacesInterface", "serviceId");
            }
            this.serviceId = serviceId;
            return this;
        }
        @CustomType.Setter
        public Builder serviceManaged(Boolean serviceManaged) {
            if (serviceManaged == null) {
              throw new MissingRequiredPropertyException("GetNetworkInterfacesInterface", "serviceManaged");
            }
            this.serviceManaged = serviceManaged;
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            if (status == null) {
              throw new MissingRequiredPropertyException("GetNetworkInterfacesInterface", "status");
            }
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder tags(Map<String,String> tags) {
            if (tags == null) {
              throw new MissingRequiredPropertyException("GetNetworkInterfacesInterface", "tags");
            }
            this.tags = tags;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("GetNetworkInterfacesInterface", "type");
            }
            this.type = type;
            return this;
        }
        @CustomType.Setter
        public Builder vpcId(String vpcId) {
            if (vpcId == null) {
              throw new MissingRequiredPropertyException("GetNetworkInterfacesInterface", "vpcId");
            }
            this.vpcId = vpcId;
            return this;
        }
        @CustomType.Setter
        public Builder vswitchId(String vswitchId) {
            if (vswitchId == null) {
              throw new MissingRequiredPropertyException("GetNetworkInterfacesInterface", "vswitchId");
            }
            this.vswitchId = vswitchId;
            return this;
        }
        @CustomType.Setter
        public Builder zoneId(String zoneId) {
            if (zoneId == null) {
              throw new MissingRequiredPropertyException("GetNetworkInterfacesInterface", "zoneId");
            }
            this.zoneId = zoneId;
            return this;
        }
        public GetNetworkInterfacesInterface build() {
            final var _resultValue = new GetNetworkInterfacesInterface();
            _resultValue.associatedPublicIps = associatedPublicIps;
            _resultValue.creationTime = creationTime;
            _resultValue.description = description;
            _resultValue.id = id;
            _resultValue.instanceId = instanceId;
            _resultValue.ipv6Sets = ipv6Sets;
            _resultValue.mac = mac;
            _resultValue.name = name;
            _resultValue.networkInterfaceId = networkInterfaceId;
            _resultValue.networkInterfaceName = networkInterfaceName;
            _resultValue.networkInterfaceTrafficMode = networkInterfaceTrafficMode;
            _resultValue.ownerId = ownerId;
            _resultValue.primaryIpAddress = primaryIpAddress;
            _resultValue.privateIp = privateIp;
            _resultValue.privateIpAddresses = privateIpAddresses;
            _resultValue.privateIps = privateIps;
            _resultValue.queueNumber = queueNumber;
            _resultValue.resourceGroupId = resourceGroupId;
            _resultValue.securityGroupIds = securityGroupIds;
            _resultValue.securityGroups = securityGroups;
            _resultValue.serviceId = serviceId;
            _resultValue.serviceManaged = serviceManaged;
            _resultValue.status = status;
            _resultValue.tags = tags;
            _resultValue.type = type;
            _resultValue.vpcId = vpcId;
            _resultValue.vswitchId = vswitchId;
            _resultValue.zoneId = zoneId;
            return _resultValue;
        }
    }
}
