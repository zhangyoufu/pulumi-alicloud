// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.nas.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetAccessGroupsGroup {
    /**
     * @return The name of access group.
     * 
     */
    private String accessGroupName;
    /**
     * @return Filter results by a specific AccessGroupType.
     * 
     */
    private String accessGroupType;
    /**
     * @return Filter results by a specific Description.
     * 
     */
    private String description;
    /**
     * @return This ID of this AccessGroup. It is formatted to ``&lt;access_group_id&gt;:&lt;file_system_type&gt;``. Before version 1.95.0, the value is `access_group_name`.
     * 
     */
    private String id;
    /**
     * @return MountTargetCount block of the AccessGroup
     * 
     */
    private Integer mountTargetCount;
    /**
     * @return RuleCount of the AccessGroup.
     * 
     */
    private Integer ruleCount;
    /**
     * @return Field `type` has been deprecated from version 1.95.0. Use `access_group_type` instead.
     * 
     */
    private String type;

    private GetAccessGroupsGroup() {}
    /**
     * @return The name of access group.
     * 
     */
    public String accessGroupName() {
        return this.accessGroupName;
    }
    /**
     * @return Filter results by a specific AccessGroupType.
     * 
     */
    public String accessGroupType() {
        return this.accessGroupType;
    }
    /**
     * @return Filter results by a specific Description.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return This ID of this AccessGroup. It is formatted to ``&lt;access_group_id&gt;:&lt;file_system_type&gt;``. Before version 1.95.0, the value is `access_group_name`.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return MountTargetCount block of the AccessGroup
     * 
     */
    public Integer mountTargetCount() {
        return this.mountTargetCount;
    }
    /**
     * @return RuleCount of the AccessGroup.
     * 
     */
    public Integer ruleCount() {
        return this.ruleCount;
    }
    /**
     * @return Field `type` has been deprecated from version 1.95.0. Use `access_group_type` instead.
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAccessGroupsGroup defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String accessGroupName;
        private String accessGroupType;
        private String description;
        private String id;
        private Integer mountTargetCount;
        private Integer ruleCount;
        private String type;
        public Builder() {}
        public Builder(GetAccessGroupsGroup defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessGroupName = defaults.accessGroupName;
    	      this.accessGroupType = defaults.accessGroupType;
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.mountTargetCount = defaults.mountTargetCount;
    	      this.ruleCount = defaults.ruleCount;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder accessGroupName(String accessGroupName) {
            if (accessGroupName == null) {
              throw new MissingRequiredPropertyException("GetAccessGroupsGroup", "accessGroupName");
            }
            this.accessGroupName = accessGroupName;
            return this;
        }
        @CustomType.Setter
        public Builder accessGroupType(String accessGroupType) {
            if (accessGroupType == null) {
              throw new MissingRequiredPropertyException("GetAccessGroupsGroup", "accessGroupType");
            }
            this.accessGroupType = accessGroupType;
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetAccessGroupsGroup", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetAccessGroupsGroup", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder mountTargetCount(Integer mountTargetCount) {
            if (mountTargetCount == null) {
              throw new MissingRequiredPropertyException("GetAccessGroupsGroup", "mountTargetCount");
            }
            this.mountTargetCount = mountTargetCount;
            return this;
        }
        @CustomType.Setter
        public Builder ruleCount(Integer ruleCount) {
            if (ruleCount == null) {
              throw new MissingRequiredPropertyException("GetAccessGroupsGroup", "ruleCount");
            }
            this.ruleCount = ruleCount;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("GetAccessGroupsGroup", "type");
            }
            this.type = type;
            return this;
        }
        public GetAccessGroupsGroup build() {
            final var _resultValue = new GetAccessGroupsGroup();
            _resultValue.accessGroupName = accessGroupName;
            _resultValue.accessGroupType = accessGroupType;
            _resultValue.description = description;
            _resultValue.id = id;
            _resultValue.mountTargetCount = mountTargetCount;
            _resultValue.ruleCount = ruleCount;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}
