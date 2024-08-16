// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.rocketmq.outputs;

import com.pulumi.alicloud.rocketmq.outputs.GetGroupsGroup;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetGroupsResult {
    private @Nullable String groupIdRegex;
    /**
     * @return Specify the protocol applicable to the created Group ID.
     * 
     */
    private @Nullable String groupType;
    /**
     * @return A list of groups. Each element contains the following attributes:
     * 
     */
    private List<GetGroupsGroup> groups;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return A list of group names.
     * 
     */
    private List<String> ids;
    private String instanceId;
    private @Nullable String nameRegex;
    private List<String> names;
    private @Nullable String outputFile;
    /**
     * @return A map of tags assigned to the Ons group.
     * 
     */
    private @Nullable Map<String,String> tags;

    private GetGroupsResult() {}
    public Optional<String> groupIdRegex() {
        return Optional.ofNullable(this.groupIdRegex);
    }
    /**
     * @return Specify the protocol applicable to the created Group ID.
     * 
     */
    public Optional<String> groupType() {
        return Optional.ofNullable(this.groupType);
    }
    /**
     * @return A list of groups. Each element contains the following attributes:
     * 
     */
    public List<GetGroupsGroup> groups() {
        return this.groups;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return A list of group names.
     * 
     */
    public List<String> ids() {
        return this.ids;
    }
    public String instanceId() {
        return this.instanceId;
    }
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    public List<String> names() {
        return this.names;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    /**
     * @return A map of tags assigned to the Ons group.
     * 
     */
    public Map<String,String> tags() {
        return this.tags == null ? Map.of() : this.tags;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetGroupsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String groupIdRegex;
        private @Nullable String groupType;
        private List<GetGroupsGroup> groups;
        private String id;
        private List<String> ids;
        private String instanceId;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String outputFile;
        private @Nullable Map<String,String> tags;
        public Builder() {}
        public Builder(GetGroupsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.groupIdRegex = defaults.groupIdRegex;
    	      this.groupType = defaults.groupType;
    	      this.groups = defaults.groups;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.instanceId = defaults.instanceId;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.outputFile = defaults.outputFile;
    	      this.tags = defaults.tags;
        }

        @CustomType.Setter
        public Builder groupIdRegex(@Nullable String groupIdRegex) {

            this.groupIdRegex = groupIdRegex;
            return this;
        }
        @CustomType.Setter
        public Builder groupType(@Nullable String groupType) {

            this.groupType = groupType;
            return this;
        }
        @CustomType.Setter
        public Builder groups(List<GetGroupsGroup> groups) {
            if (groups == null) {
              throw new MissingRequiredPropertyException("GetGroupsResult", "groups");
            }
            this.groups = groups;
            return this;
        }
        public Builder groups(GetGroupsGroup... groups) {
            return groups(List.of(groups));
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetGroupsResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder ids(List<String> ids) {
            if (ids == null) {
              throw new MissingRequiredPropertyException("GetGroupsResult", "ids");
            }
            this.ids = ids;
            return this;
        }
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }
        @CustomType.Setter
        public Builder instanceId(String instanceId) {
            if (instanceId == null) {
              throw new MissingRequiredPropertyException("GetGroupsResult", "instanceId");
            }
            this.instanceId = instanceId;
            return this;
        }
        @CustomType.Setter
        public Builder nameRegex(@Nullable String nameRegex) {

            this.nameRegex = nameRegex;
            return this;
        }
        @CustomType.Setter
        public Builder names(List<String> names) {
            if (names == null) {
              throw new MissingRequiredPropertyException("GetGroupsResult", "names");
            }
            this.names = names;
            return this;
        }
        public Builder names(String... names) {
            return names(List.of(names));
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {

            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder tags(@Nullable Map<String,String> tags) {

            this.tags = tags;
            return this;
        }
        public GetGroupsResult build() {
            final var _resultValue = new GetGroupsResult();
            _resultValue.groupIdRegex = groupIdRegex;
            _resultValue.groupType = groupType;
            _resultValue.groups = groups;
            _resultValue.id = id;
            _resultValue.ids = ids;
            _resultValue.instanceId = instanceId;
            _resultValue.nameRegex = nameRegex;
            _resultValue.names = names;
            _resultValue.outputFile = outputFile;
            _resultValue.tags = tags;
            return _resultValue;
        }
    }
}
