// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.brain.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetIndustrialPidOrganizationsOrganization {
    /**
     * @return The ID of the Pid Organization.
     * 
     */
    private String id;
    /**
     * @return The parent organization id.
     * 
     */
    private String parentPidOrganizationId;
    /**
     * @return The organization id.
     * 
     */
    private String pidOrganizationId;
    /**
     * @return The organization level.
     * 
     */
    private Integer pidOrganizationLevel;
    /**
     * @return The organization name.
     * 
     */
    private String pidOrganizationName;

    private GetIndustrialPidOrganizationsOrganization() {}
    /**
     * @return The ID of the Pid Organization.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The parent organization id.
     * 
     */
    public String parentPidOrganizationId() {
        return this.parentPidOrganizationId;
    }
    /**
     * @return The organization id.
     * 
     */
    public String pidOrganizationId() {
        return this.pidOrganizationId;
    }
    /**
     * @return The organization level.
     * 
     */
    public Integer pidOrganizationLevel() {
        return this.pidOrganizationLevel;
    }
    /**
     * @return The organization name.
     * 
     */
    public String pidOrganizationName() {
        return this.pidOrganizationName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIndustrialPidOrganizationsOrganization defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String parentPidOrganizationId;
        private String pidOrganizationId;
        private Integer pidOrganizationLevel;
        private String pidOrganizationName;
        public Builder() {}
        public Builder(GetIndustrialPidOrganizationsOrganization defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.parentPidOrganizationId = defaults.parentPidOrganizationId;
    	      this.pidOrganizationId = defaults.pidOrganizationId;
    	      this.pidOrganizationLevel = defaults.pidOrganizationLevel;
    	      this.pidOrganizationName = defaults.pidOrganizationName;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetIndustrialPidOrganizationsOrganization", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder parentPidOrganizationId(String parentPidOrganizationId) {
            if (parentPidOrganizationId == null) {
              throw new MissingRequiredPropertyException("GetIndustrialPidOrganizationsOrganization", "parentPidOrganizationId");
            }
            this.parentPidOrganizationId = parentPidOrganizationId;
            return this;
        }
        @CustomType.Setter
        public Builder pidOrganizationId(String pidOrganizationId) {
            if (pidOrganizationId == null) {
              throw new MissingRequiredPropertyException("GetIndustrialPidOrganizationsOrganization", "pidOrganizationId");
            }
            this.pidOrganizationId = pidOrganizationId;
            return this;
        }
        @CustomType.Setter
        public Builder pidOrganizationLevel(Integer pidOrganizationLevel) {
            if (pidOrganizationLevel == null) {
              throw new MissingRequiredPropertyException("GetIndustrialPidOrganizationsOrganization", "pidOrganizationLevel");
            }
            this.pidOrganizationLevel = pidOrganizationLevel;
            return this;
        }
        @CustomType.Setter
        public Builder pidOrganizationName(String pidOrganizationName) {
            if (pidOrganizationName == null) {
              throw new MissingRequiredPropertyException("GetIndustrialPidOrganizationsOrganization", "pidOrganizationName");
            }
            this.pidOrganizationName = pidOrganizationName;
            return this;
        }
        public GetIndustrialPidOrganizationsOrganization build() {
            final var _resultValue = new GetIndustrialPidOrganizationsOrganization();
            _resultValue.id = id;
            _resultValue.parentPidOrganizationId = parentPidOrganizationId;
            _resultValue.pidOrganizationId = pidOrganizationId;
            _resultValue.pidOrganizationLevel = pidOrganizationLevel;
            _resultValue.pidOrganizationName = pidOrganizationName;
            return _resultValue;
        }
    }
}
