// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.oss.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetInstancesInstance {
    /**
     * @return The cluster type of the instance. Possible values: `SSD`, `HYBRID`.
     * 
     */
    private String clusterType;
    /**
     * @return The create time of the instance.
     * 
     */
    private String createTime;
    /**
     * @return The description of the instance.
     * 
     */
    private String description;
    /**
     * @return (Removed since v1.221.0) The instance quota which indicating the maximum number of tables.
     * 
     */
    private Integer entityQuota;
    /**
     * @return ID of the instance.
     * 
     */
    private String id;
    /**
     * @return Instance name.
     * 
     */
    private String name;
    /**
     * @return (Removed since v1.221.0) The network type of the instance. Possible values: `NORMAL`, `VPC`, `VPC_CONSOLE`.
     * 
     */
    private String network;
    /**
     * @return (Available since v1.221.0) The set of request sources that are allowed access. Possible values: `TRUST_PROXY`.
     * 
     */
    private List<String> networkSourceAcls;
    /**
     * @return (Available since v1.221.0) The set of network types that are allowed access. Possible values: `CLASSIC`, `VPC`, `INTERNET`.
     * 
     */
    private List<String> networkTypeAcls;
    /**
     * @return (Available since v1.221.0) instance policy, json string.
     * 
     */
    private String policy;
    /**
     * @return (Available since v1.221.0) instance policy version.
     * 
     */
    private Integer policyVersion;
    /**
     * @return (Available since v1.221.0) The resource group the instance belongs to.
     * 
     */
    private String resourceGroupId;
    /**
     * @return Instance status. Possible values: `Running`, `Disabled`, `Deleting`.
     * 
     */
    private String status;
    /**
     * @return (Available since v1.221.0) The instance quota which indicating the maximum number of tables.
     * 
     */
    private Integer tableQuota;
    /**
     * @return A map of tags assigned to the instance. It must be in the format:
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.ots.OtsFunctions;
     * import com.pulumi.alicloud.ots.inputs.GetInstancesArgs;
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
     *         final var instancesDs = OtsFunctions.getInstances(GetInstancesArgs.builder()
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
    private Map<String,Object> tags;
    /**
     * @return The user id of the instance.
     * 
     */
    private String userId;

    private GetInstancesInstance() {}
    /**
     * @return The cluster type of the instance. Possible values: `SSD`, `HYBRID`.
     * 
     */
    public String clusterType() {
        return this.clusterType;
    }
    /**
     * @return The create time of the instance.
     * 
     */
    public String createTime() {
        return this.createTime;
    }
    /**
     * @return The description of the instance.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return (Removed since v1.221.0) The instance quota which indicating the maximum number of tables.
     * 
     */
    public Integer entityQuota() {
        return this.entityQuota;
    }
    /**
     * @return ID of the instance.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Instance name.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return (Removed since v1.221.0) The network type of the instance. Possible values: `NORMAL`, `VPC`, `VPC_CONSOLE`.
     * 
     */
    public String network() {
        return this.network;
    }
    /**
     * @return (Available since v1.221.0) The set of request sources that are allowed access. Possible values: `TRUST_PROXY`.
     * 
     */
    public List<String> networkSourceAcls() {
        return this.networkSourceAcls;
    }
    /**
     * @return (Available since v1.221.0) The set of network types that are allowed access. Possible values: `CLASSIC`, `VPC`, `INTERNET`.
     * 
     */
    public List<String> networkTypeAcls() {
        return this.networkTypeAcls;
    }
    /**
     * @return (Available since v1.221.0) instance policy, json string.
     * 
     */
    public String policy() {
        return this.policy;
    }
    /**
     * @return (Available since v1.221.0) instance policy version.
     * 
     */
    public Integer policyVersion() {
        return this.policyVersion;
    }
    /**
     * @return (Available since v1.221.0) The resource group the instance belongs to.
     * 
     */
    public String resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * @return Instance status. Possible values: `Running`, `Disabled`, `Deleting`.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return (Available since v1.221.0) The instance quota which indicating the maximum number of tables.
     * 
     */
    public Integer tableQuota() {
        return this.tableQuota;
    }
    /**
     * @return A map of tags assigned to the instance. It must be in the format:
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.ots.OtsFunctions;
     * import com.pulumi.alicloud.ots.inputs.GetInstancesArgs;
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
     *         final var instancesDs = OtsFunctions.getInstances(GetInstancesArgs.builder()
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
        return this.tags;
    }
    /**
     * @return The user id of the instance.
     * 
     */
    public String userId() {
        return this.userId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetInstancesInstance defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String clusterType;
        private String createTime;
        private String description;
        private Integer entityQuota;
        private String id;
        private String name;
        private String network;
        private List<String> networkSourceAcls;
        private List<String> networkTypeAcls;
        private String policy;
        private Integer policyVersion;
        private String resourceGroupId;
        private String status;
        private Integer tableQuota;
        private Map<String,Object> tags;
        private String userId;
        public Builder() {}
        public Builder(GetInstancesInstance defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clusterType = defaults.clusterType;
    	      this.createTime = defaults.createTime;
    	      this.description = defaults.description;
    	      this.entityQuota = defaults.entityQuota;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.network = defaults.network;
    	      this.networkSourceAcls = defaults.networkSourceAcls;
    	      this.networkTypeAcls = defaults.networkTypeAcls;
    	      this.policy = defaults.policy;
    	      this.policyVersion = defaults.policyVersion;
    	      this.resourceGroupId = defaults.resourceGroupId;
    	      this.status = defaults.status;
    	      this.tableQuota = defaults.tableQuota;
    	      this.tags = defaults.tags;
    	      this.userId = defaults.userId;
        }

        @CustomType.Setter
        public Builder clusterType(String clusterType) {
            if (clusterType == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "clusterType");
            }
            this.clusterType = clusterType;
            return this;
        }
        @CustomType.Setter
        public Builder createTime(String createTime) {
            if (createTime == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "createTime");
            }
            this.createTime = createTime;
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
        public Builder entityQuota(Integer entityQuota) {
            if (entityQuota == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "entityQuota");
            }
            this.entityQuota = entityQuota;
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
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder network(String network) {
            if (network == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "network");
            }
            this.network = network;
            return this;
        }
        @CustomType.Setter
        public Builder networkSourceAcls(List<String> networkSourceAcls) {
            if (networkSourceAcls == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "networkSourceAcls");
            }
            this.networkSourceAcls = networkSourceAcls;
            return this;
        }
        public Builder networkSourceAcls(String... networkSourceAcls) {
            return networkSourceAcls(List.of(networkSourceAcls));
        }
        @CustomType.Setter
        public Builder networkTypeAcls(List<String> networkTypeAcls) {
            if (networkTypeAcls == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "networkTypeAcls");
            }
            this.networkTypeAcls = networkTypeAcls;
            return this;
        }
        public Builder networkTypeAcls(String... networkTypeAcls) {
            return networkTypeAcls(List.of(networkTypeAcls));
        }
        @CustomType.Setter
        public Builder policy(String policy) {
            if (policy == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "policy");
            }
            this.policy = policy;
            return this;
        }
        @CustomType.Setter
        public Builder policyVersion(Integer policyVersion) {
            if (policyVersion == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "policyVersion");
            }
            this.policyVersion = policyVersion;
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
        public Builder status(String status) {
            if (status == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "status");
            }
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder tableQuota(Integer tableQuota) {
            if (tableQuota == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "tableQuota");
            }
            this.tableQuota = tableQuota;
            return this;
        }
        @CustomType.Setter
        public Builder tags(Map<String,Object> tags) {
            if (tags == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "tags");
            }
            this.tags = tags;
            return this;
        }
        @CustomType.Setter
        public Builder userId(String userId) {
            if (userId == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "userId");
            }
            this.userId = userId;
            return this;
        }
        public GetInstancesInstance build() {
            final var _resultValue = new GetInstancesInstance();
            _resultValue.clusterType = clusterType;
            _resultValue.createTime = createTime;
            _resultValue.description = description;
            _resultValue.entityQuota = entityQuota;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.network = network;
            _resultValue.networkSourceAcls = networkSourceAcls;
            _resultValue.networkTypeAcls = networkTypeAcls;
            _resultValue.policy = policy;
            _resultValue.policyVersion = policyVersion;
            _resultValue.resourceGroupId = resourceGroupId;
            _resultValue.status = status;
            _resultValue.tableQuota = tableQuota;
            _resultValue.tags = tags;
            _resultValue.userId = userId;
            return _resultValue;
        }
    }
}
