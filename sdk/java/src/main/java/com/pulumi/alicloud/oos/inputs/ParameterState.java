// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.oos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ParameterState extends com.pulumi.resources.ResourceArgs {

    public static final ParameterState Empty = new ParameterState();

    /**
     * The constraints of the common parameter. This value follows the json format. By default, this parameter is null. Valid values:
     * * `AllowedValues`: The value that is allowed for the common parameter. It must be an array string.
     * * `AllowedPattern`: The pattern that is allowed for the common parameter. It must be a regular expression.
     * * `MinLength`: The minimum length of the common parameter.
     * * `MaxLength`: The maximum length of the common parameter.
     * 
     */
    @Import(name="constraints")
    private @Nullable Output<String> constraints;

    /**
     * @return The constraints of the common parameter. This value follows the json format. By default, this parameter is null. Valid values:
     * * `AllowedValues`: The value that is allowed for the common parameter. It must be an array string.
     * * `AllowedPattern`: The pattern that is allowed for the common parameter. It must be a regular expression.
     * * `MinLength`: The minimum length of the common parameter.
     * * `MaxLength`: The maximum length of the common parameter.
     * 
     */
    public Optional<Output<String>> constraints() {
        return Optional.ofNullable(this.constraints);
    }

    /**
     * The description of the common parameter. The description must be `1` to `200` characters in length.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the common parameter. The description must be `1` to `200` characters in length.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The name of the common parameter. The name must be `2` to `180` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/) and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, `ALICLOUD`, or `OOS`.
     * 
     */
    @Import(name="parameterName")
    private @Nullable Output<String> parameterName;

    /**
     * @return The name of the common parameter. The name must be `2` to `180` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/) and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, `ALICLOUD`, or `OOS`.
     * 
     */
    public Optional<Output<String>> parameterName() {
        return Optional.ofNullable(this.parameterName);
    }

    /**
     * The ID of the Resource Group.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The ID of the Resource Group.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The data type of the common parameter. Valid values: `String` and `StringList`.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return The data type of the common parameter. Valid values: `String` and `StringList`.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    /**
     * The value of the common parameter. The value must be `1` to `4096` characters in length.
     * 
     */
    @Import(name="value")
    private @Nullable Output<String> value;

    /**
     * @return The value of the common parameter. The value must be `1` to `4096` characters in length.
     * 
     */
    public Optional<Output<String>> value() {
        return Optional.ofNullable(this.value);
    }

    private ParameterState() {}

    private ParameterState(ParameterState $) {
        this.constraints = $.constraints;
        this.description = $.description;
        this.parameterName = $.parameterName;
        this.resourceGroupId = $.resourceGroupId;
        this.tags = $.tags;
        this.type = $.type;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ParameterState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ParameterState $;

        public Builder() {
            $ = new ParameterState();
        }

        public Builder(ParameterState defaults) {
            $ = new ParameterState(Objects.requireNonNull(defaults));
        }

        /**
         * @param constraints The constraints of the common parameter. This value follows the json format. By default, this parameter is null. Valid values:
         * * `AllowedValues`: The value that is allowed for the common parameter. It must be an array string.
         * * `AllowedPattern`: The pattern that is allowed for the common parameter. It must be a regular expression.
         * * `MinLength`: The minimum length of the common parameter.
         * * `MaxLength`: The maximum length of the common parameter.
         * 
         * @return builder
         * 
         */
        public Builder constraints(@Nullable Output<String> constraints) {
            $.constraints = constraints;
            return this;
        }

        /**
         * @param constraints The constraints of the common parameter. This value follows the json format. By default, this parameter is null. Valid values:
         * * `AllowedValues`: The value that is allowed for the common parameter. It must be an array string.
         * * `AllowedPattern`: The pattern that is allowed for the common parameter. It must be a regular expression.
         * * `MinLength`: The minimum length of the common parameter.
         * * `MaxLength`: The maximum length of the common parameter.
         * 
         * @return builder
         * 
         */
        public Builder constraints(String constraints) {
            return constraints(Output.of(constraints));
        }

        /**
         * @param description The description of the common parameter. The description must be `1` to `200` characters in length.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the common parameter. The description must be `1` to `200` characters in length.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param parameterName The name of the common parameter. The name must be `2` to `180` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/) and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, `ALICLOUD`, or `OOS`.
         * 
         * @return builder
         * 
         */
        public Builder parameterName(@Nullable Output<String> parameterName) {
            $.parameterName = parameterName;
            return this;
        }

        /**
         * @param parameterName The name of the common parameter. The name must be `2` to `180` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/) and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, `ALICLOUD`, or `OOS`.
         * 
         * @return builder
         * 
         */
        public Builder parameterName(String parameterName) {
            return parameterName(Output.of(parameterName));
        }

        /**
         * @param resourceGroupId The ID of the Resource Group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The ID of the Resource Group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param type The data type of the common parameter. Valid values: `String` and `StringList`.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The data type of the common parameter. Valid values: `String` and `StringList`.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param value The value of the common parameter. The value must be `1` to `4096` characters in length.
         * 
         * @return builder
         * 
         */
        public Builder value(@Nullable Output<String> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value The value of the common parameter. The value must be `1` to `4096` characters in length.
         * 
         * @return builder
         * 
         */
        public Builder value(String value) {
            return value(Output.of(value));
        }

        public ParameterState build() {
            return $;
        }
    }

}
