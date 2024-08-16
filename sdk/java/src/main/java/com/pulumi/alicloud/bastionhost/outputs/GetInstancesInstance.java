// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.bastionhost.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetInstancesInstance {
    /**
     * @return The instance&#39;s remark.
     * 
     */
    private String description;
    /**
     * @return The instance&#39;s id.
     * 
     */
    private String id;
    /**
     * @return The instance&#39;s status.
     * 
     */
    private String instanceStatus;
    private String licenseCode;
    /**
     * @return The instance&#39;s private domain name.
     * 
     */
    private String privateDomain;
    /**
     * @return The instance&#39;s public domain name.
     * 
     */
    private String publicDomain;
    /**
     * @return The instance&#39;s public network access configuration.
     * 
     */
    private Boolean publicNetworkAccess;
    /**
     * @return The instance&#39;s security group configuration.
     * 
     */
    private List<String> securityGroupIds;
    /**
     * @return A map of tags assigned to the bastionhost instance. It must be in the format:
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.bastionhost.BastionhostFunctions;
     * import com.pulumi.alicloud.bastionhost.inputs.GetInstancesArgs;
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
     *         final var instance = BastionhostFunctions.getInstances(GetInstancesArgs.builder()
     *             .tags(Map.of("tagKey1", "tagValue1"))
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    private @Nullable Map<String,String> tags;
    /**
     * @return The instance&#39;s vSwitch ID.
     * 
     */
    private String userVswitchId;

    private GetInstancesInstance() {}
    /**
     * @return The instance&#39;s remark.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The instance&#39;s id.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The instance&#39;s status.
     * 
     */
    public String instanceStatus() {
        return this.instanceStatus;
    }
    public String licenseCode() {
        return this.licenseCode;
    }
    /**
     * @return The instance&#39;s private domain name.
     * 
     */
    public String privateDomain() {
        return this.privateDomain;
    }
    /**
     * @return The instance&#39;s public domain name.
     * 
     */
    public String publicDomain() {
        return this.publicDomain;
    }
    /**
     * @return The instance&#39;s public network access configuration.
     * 
     */
    public Boolean publicNetworkAccess() {
        return this.publicNetworkAccess;
    }
    /**
     * @return The instance&#39;s security group configuration.
     * 
     */
    public List<String> securityGroupIds() {
        return this.securityGroupIds;
    }
    /**
     * @return A map of tags assigned to the bastionhost instance. It must be in the format:
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.bastionhost.BastionhostFunctions;
     * import com.pulumi.alicloud.bastionhost.inputs.GetInstancesArgs;
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
     *         final var instance = BastionhostFunctions.getInstances(GetInstancesArgs.builder()
     *             .tags(Map.of("tagKey1", "tagValue1"))
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public Map<String,String> tags() {
        return this.tags == null ? Map.of() : this.tags;
    }
    /**
     * @return The instance&#39;s vSwitch ID.
     * 
     */
    public String userVswitchId() {
        return this.userVswitchId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetInstancesInstance defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private String id;
        private String instanceStatus;
        private String licenseCode;
        private String privateDomain;
        private String publicDomain;
        private Boolean publicNetworkAccess;
        private List<String> securityGroupIds;
        private @Nullable Map<String,String> tags;
        private String userVswitchId;
        public Builder() {}
        public Builder(GetInstancesInstance defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.instanceStatus = defaults.instanceStatus;
    	      this.licenseCode = defaults.licenseCode;
    	      this.privateDomain = defaults.privateDomain;
    	      this.publicDomain = defaults.publicDomain;
    	      this.publicNetworkAccess = defaults.publicNetworkAccess;
    	      this.securityGroupIds = defaults.securityGroupIds;
    	      this.tags = defaults.tags;
    	      this.userVswitchId = defaults.userVswitchId;
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
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder instanceStatus(String instanceStatus) {
            if (instanceStatus == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "instanceStatus");
            }
            this.instanceStatus = instanceStatus;
            return this;
        }
        @CustomType.Setter
        public Builder licenseCode(String licenseCode) {
            if (licenseCode == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "licenseCode");
            }
            this.licenseCode = licenseCode;
            return this;
        }
        @CustomType.Setter
        public Builder privateDomain(String privateDomain) {
            if (privateDomain == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "privateDomain");
            }
            this.privateDomain = privateDomain;
            return this;
        }
        @CustomType.Setter
        public Builder publicDomain(String publicDomain) {
            if (publicDomain == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "publicDomain");
            }
            this.publicDomain = publicDomain;
            return this;
        }
        @CustomType.Setter
        public Builder publicNetworkAccess(Boolean publicNetworkAccess) {
            if (publicNetworkAccess == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "publicNetworkAccess");
            }
            this.publicNetworkAccess = publicNetworkAccess;
            return this;
        }
        @CustomType.Setter
        public Builder securityGroupIds(List<String> securityGroupIds) {
            if (securityGroupIds == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "securityGroupIds");
            }
            this.securityGroupIds = securityGroupIds;
            return this;
        }
        public Builder securityGroupIds(String... securityGroupIds) {
            return securityGroupIds(List.of(securityGroupIds));
        }
        @CustomType.Setter
        public Builder tags(@Nullable Map<String,String> tags) {

            this.tags = tags;
            return this;
        }
        @CustomType.Setter
        public Builder userVswitchId(String userVswitchId) {
            if (userVswitchId == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstance", "userVswitchId");
            }
            this.userVswitchId = userVswitchId;
            return this;
        }
        public GetInstancesInstance build() {
            final var _resultValue = new GetInstancesInstance();
            _resultValue.description = description;
            _resultValue.id = id;
            _resultValue.instanceStatus = instanceStatus;
            _resultValue.licenseCode = licenseCode;
            _resultValue.privateDomain = privateDomain;
            _resultValue.publicDomain = publicDomain;
            _resultValue.publicNetworkAccess = publicNetworkAccess;
            _resultValue.securityGroupIds = securityGroupIds;
            _resultValue.tags = tags;
            _resultValue.userVswitchId = userVswitchId;
            return _resultValue;
        }
    }
}
