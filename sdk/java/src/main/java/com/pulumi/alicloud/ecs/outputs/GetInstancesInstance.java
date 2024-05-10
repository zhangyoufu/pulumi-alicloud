// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs.outputs;

import com.pulumi.alicloud.ecs.outputs.GetInstancesInstanceDiskDeviceMapping;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetInstancesInstance {
    /**
     * @return Availability zone where instances are located.
     * 
     */
    private String availabilityZone;
    /**
     * @return Instance creation time.
     * 
     */
    private String creationTime;
    /**
     * @return Instance description.
     * 
     */
    private String description;
    /**
     * @return Description of the attached disks.
     * 
     */
    private List<GetInstancesInstanceDiskDeviceMapping> diskDeviceMappings;
    /**
     * @return EIP address the VPC instance is using.
     * 
     */
    private String eip;
    /**
     * @return ID of the instance.
     * 
     */
    private String id;
    /**
     * @return The image ID of some ECS instance used.
     * 
     */
    private String imageId;
    /**
     * @return Instance charge type.
     * 
     */
    private String instanceChargeType;
    /**
     * @return Instance type.
     * 
     */
    private String instanceType;
    /**
     * @return Instance network charge type.
     * 
     */
    private String internetChargeType;
    /**
     * @return Max output bandwidth for internet.
     * 
     */
    private Integer internetMaxBandwidthOut;
    /**
     * @return Key pair the instance is using.
     * 
     */
    private String keyName;
    /**
     * @return Instance name.
     * 
     */
    private String name;
    /**
     * @return Instance private IP address.
     * 
     */
    private String privateIp;
    /**
     * @return Instance public IP address.
     * 
     */
    private String publicIp;
    /**
     * @return The RAM role name which the instance attaches.
     * 
     */
    private String ramRoleName;
    /**
     * @return Region ID the instance belongs to.
     * 
     */
    private String regionId;
    /**
     * @return The ID of resource group which the instance belongs.
     * 
     */
    private String resourceGroupId;
    /**
     * @return List of security group IDs the instance belongs to.
     * 
     */
    private List<String> securityGroups;
    /**
     * @return Spot strategy the instance is using.
     * 
     */
    private String spotStrategy;
    /**
     * @return Instance status. Valid values: &#34;Creating&#34;, &#34;Starting&#34;, &#34;Running&#34;, &#34;Stopping&#34; and &#34;Stopped&#34;. If undefined, all statuses are considered.
     * 
     */
    private String status;
    /**
     * @return A map of tags assigned to the ECS instances. It must be in the format:
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.ecs.EcsFunctions;
     * import com.pulumi.alicloud.ecs.inputs.GetInstancesArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var taggedInstances = EcsFunctions.getInstances(GetInstancesArgs.builder()
     *             .tags(Map.ofEntries(
     *                 Map.entry("tagKey1", "tagValue1"),
     *                 Map.entry("tagKey2", "tagValue2")
     *             ))
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    private @Nullable Map<String,Object> tags;
    /**
     * @return ID of the VPC linked to the instances.
     * 
     */
    private String vpcId;
    /**
     * @return ID of the VSwitch linked to the instances.
     * 
     */
    private String vswitchId;

    private GetInstancesInstance() {}
    /**
     * @return Availability zone where instances are located.
     * 
     */
    public String availabilityZone() {
        return this.availabilityZone;
    }
    /**
     * @return Instance creation time.
     * 
     */
    public String creationTime() {
        return this.creationTime;
    }
    /**
     * @return Instance description.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return Description of the attached disks.
     * 
     */
    public List<GetInstancesInstanceDiskDeviceMapping> diskDeviceMappings() {
        return this.diskDeviceMappings;
    }
    /**
     * @return EIP address the VPC instance is using.
     * 
     */
    public String eip() {
        return this.eip;
    }
    /**
     * @return ID of the instance.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The image ID of some ECS instance used.
     * 
     */
    public String imageId() {
        return this.imageId;
    }
    /**
     * @return Instance charge type.
     * 
     */
    public String instanceChargeType() {
        return this.instanceChargeType;
    }
    /**
     * @return Instance type.
     * 
     */
    public String instanceType() {
        return this.instanceType;
    }
    /**
     * @return Instance network charge type.
     * 
     */
    public String internetChargeType() {
        return this.internetChargeType;
    }
    /**
     * @return Max output bandwidth for internet.
     * 
     */
    public Integer internetMaxBandwidthOut() {
        return this.internetMaxBandwidthOut;
    }
    /**
     * @return Key pair the instance is using.
     * 
     */
    public String keyName() {
        return this.keyName;
    }
    /**
     * @return Instance name.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Instance private IP address.
     * 
     */
    public String privateIp() {
        return this.privateIp;
    }
    /**
     * @return Instance public IP address.
     * 
     */
    public String publicIp() {
        return this.publicIp;
    }
    /**
     * @return The RAM role name which the instance attaches.
     * 
     */
    public String ramRoleName() {
        return this.ramRoleName;
    }
    /**
     * @return Region ID the instance belongs to.
     * 
     */
    public String regionId() {
        return this.regionId;
    }
    /**
     * @return The ID of resource group which the instance belongs.
     * 
     */
    public String resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * @return List of security group IDs the instance belongs to.
     * 
     */
    public List<String> securityGroups() {
        return this.securityGroups;
    }
    /**
     * @return Spot strategy the instance is using.
     * 
     */
    public String spotStrategy() {
        return this.spotStrategy;
    }
    /**
     * @return Instance status. Valid values: &#34;Creating&#34;, &#34;Starting&#34;, &#34;Running&#34;, &#34;Stopping&#34; and &#34;Stopped&#34;. If undefined, all statuses are considered.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return A map of tags assigned to the ECS instances. It must be in the format:
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.ecs.EcsFunctions;
     * import com.pulumi.alicloud.ecs.inputs.GetInstancesArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var taggedInstances = EcsFunctions.getInstances(GetInstancesArgs.builder()
     *             .tags(Map.ofEntries(
     *                 Map.entry("tagKey1", "tagValue1"),
     *                 Map.entry("tagKey2", "tagValue2")
     *             ))
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public Map<String,Object> tags() {
        return this.tags == null ? Map.of() : this.tags;
    }
    /**
     * @return ID of the VPC linked to the instances.
     * 
     */
    public String vpcId() {
        return this.vpcId;
    }
    /**
     * @return ID of the VSwitch linked to the instances.
     * 
     */
    public String vswitchId() {
        return this.vswitchId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetInstancesInstance defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String availabilityZone;
        private String creationTime;
        private String description;
        private List<GetInstancesInstanceDiskDeviceMapping> diskDeviceMappings;
        private String eip;
        private String id;
        private String imageId;
        private String instanceChargeType;
        private String instanceType;
        private String internetChargeType;
        private Integer internetMaxBandwidthOut;
        private String keyName;
        private String name;
        private String privateIp;
        private String publicIp;
        private String ramRoleName;
        private String regionId;
        private String resourceGroupId;
        private List<String> securityGroups;
        private String spotStrategy;
        private String status;
        private @Nullable Map<String,Object> tags;
        private String vpcId;
        private String vswitchId;
        public Builder() {}
        public Builder(GetInstancesInstance defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.availabilityZone = defaults.availabilityZone;
    	      this.creationTime = defaults.creationTime;
    	      this.description = defaults.description;
    	      this.diskDeviceMappings = defaults.diskDeviceMappings;
    	      this.eip = defaults.eip;
    	      this.id = defaults.id;
    	      this.imageId = defaults.imageId;
    	      this.instanceChargeType = defaults.instanceChargeType;
    	      this.instanceType = defaults.instanceType;
    	      this.internetChargeType = defaults.internetChargeType;
    	      this.internetMaxBandwidthOut = defaults.internetMaxBandwidthOut;
    	      this.keyName = defaults.keyName;
    	      this.name = defaults.name;
    	      this.privateIp = defaults.privateIp;
    	      this.publicIp = defaults.publicIp;
    	      this.ramRoleName = defaults.ramRoleName;
    	      this.regionId = defaults.regionId;
    	      this.resourceGroupId = defaults.resourceGroupId;
    	      this.securityGroups = defaults.securityGroups;
    	      this.spotStrategy = defaults.spotStrategy;
    	      this.status = defaults.status;
    	      this.tags = defaults.tags;
    	      this.vpcId = defaults.vpcId;
    	      this.vswitchId = defaults.vswitchId;
        }

        @CustomType.Setter
        public Builder availabilityZone(String availabilityZone) {
            if (availabilityZone == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "availabilityZone");
            }
            this.availabilityZone = availabilityZone;
            return this;
        }
        @CustomType.Setter
        public Builder creationTime(String creationTime) {
            if (creationTime == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "creationTime");
            }
            this.creationTime = creationTime;
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder diskDeviceMappings(List<GetInstancesInstanceDiskDeviceMapping> diskDeviceMappings) {
            if (diskDeviceMappings == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "diskDeviceMappings");
            }
            this.diskDeviceMappings = diskDeviceMappings;
            return this;
        }
        public Builder diskDeviceMappings(GetInstancesInstanceDiskDeviceMapping... diskDeviceMappings) {
            return diskDeviceMappings(List.of(diskDeviceMappings));
        }
        @CustomType.Setter
        public Builder eip(String eip) {
            if (eip == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "eip");
            }
            this.eip = eip;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder imageId(String imageId) {
            if (imageId == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "imageId");
            }
            this.imageId = imageId;
            return this;
        }
        @CustomType.Setter
        public Builder instanceChargeType(String instanceChargeType) {
            if (instanceChargeType == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "instanceChargeType");
            }
            this.instanceChargeType = instanceChargeType;
            return this;
        }
        @CustomType.Setter
        public Builder instanceType(String instanceType) {
            if (instanceType == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "instanceType");
            }
            this.instanceType = instanceType;
            return this;
        }
        @CustomType.Setter
        public Builder internetChargeType(String internetChargeType) {
            if (internetChargeType == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "internetChargeType");
            }
            this.internetChargeType = internetChargeType;
            return this;
        }
        @CustomType.Setter
        public Builder internetMaxBandwidthOut(Integer internetMaxBandwidthOut) {
            if (internetMaxBandwidthOut == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "internetMaxBandwidthOut");
            }
            this.internetMaxBandwidthOut = internetMaxBandwidthOut;
            return this;
        }
        @CustomType.Setter
        public Builder keyName(String keyName) {
            if (keyName == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "keyName");
            }
            this.keyName = keyName;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder privateIp(String privateIp) {
            if (privateIp == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "privateIp");
            }
            this.privateIp = privateIp;
            return this;
        }
        @CustomType.Setter
        public Builder publicIp(String publicIp) {
            if (publicIp == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "publicIp");
            }
            this.publicIp = publicIp;
            return this;
        }
        @CustomType.Setter
        public Builder ramRoleName(String ramRoleName) {
            if (ramRoleName == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "ramRoleName");
            }
            this.ramRoleName = ramRoleName;
            return this;
        }
        @CustomType.Setter
        public Builder regionId(String regionId) {
            if (regionId == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "regionId");
            }
            this.regionId = regionId;
            return this;
        }
        @CustomType.Setter
        public Builder resourceGroupId(String resourceGroupId) {
            if (resourceGroupId == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "resourceGroupId");
            }
            this.resourceGroupId = resourceGroupId;
            return this;
        }
        @CustomType.Setter
        public Builder securityGroups(List<String> securityGroups) {
            if (securityGroups == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "securityGroups");
            }
            this.securityGroups = securityGroups;
            return this;
        }
        public Builder securityGroups(String... securityGroups) {
            return securityGroups(List.of(securityGroups));
        }
        @CustomType.Setter
        public Builder spotStrategy(String spotStrategy) {
            if (spotStrategy == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "spotStrategy");
            }
            this.spotStrategy = spotStrategy;
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            if (status == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "status");
            }
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder tags(@Nullable Map<String,Object> tags) {

            this.tags = tags;
            return this;
        }
        @CustomType.Setter
        public Builder vpcId(String vpcId) {
            if (vpcId == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "vpcId");
            }
            this.vpcId = vpcId;
            return this;
        }
        @CustomType.Setter
        public Builder vswitchId(String vswitchId) {
            if (vswitchId == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "vswitchId");
            }
            this.vswitchId = vswitchId;
            return this;
        }
        public GetInstancesInstance build() {
            final var _resultValue = new GetInstancesInstance();
            _resultValue.availabilityZone = availabilityZone;
            _resultValue.creationTime = creationTime;
            _resultValue.description = description;
            _resultValue.diskDeviceMappings = diskDeviceMappings;
            _resultValue.eip = eip;
            _resultValue.id = id;
            _resultValue.imageId = imageId;
            _resultValue.instanceChargeType = instanceChargeType;
            _resultValue.instanceType = instanceType;
            _resultValue.internetChargeType = internetChargeType;
            _resultValue.internetMaxBandwidthOut = internetMaxBandwidthOut;
            _resultValue.keyName = keyName;
            _resultValue.name = name;
            _resultValue.privateIp = privateIp;
            _resultValue.publicIp = publicIp;
            _resultValue.ramRoleName = ramRoleName;
            _resultValue.regionId = regionId;
            _resultValue.resourceGroupId = resourceGroupId;
            _resultValue.securityGroups = securityGroups;
            _resultValue.spotStrategy = spotStrategy;
            _resultValue.status = status;
            _resultValue.tags = tags;
            _resultValue.vpcId = vpcId;
            _resultValue.vswitchId = vswitchId;
            return _resultValue;
        }
    }
}
