// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.apigateway.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetPluginsPlugin {
    /**
     * @return The CreateTime of the resource.
     * 
     */
    private String createTime;
    /**
     * @return The description of the plug-in, which cannot exceed 200 characters.
     * 
     */
    private String description;
    /**
     * @return The ID of the Plugin.
     * 
     */
    private String id;
    /**
     * @return The ModifiedTime of the resource.
     * 
     */
    private String modifiedTime;
    /**
     * @return The definition statement of the plug-in. Plug-in definition statements in the JSON and YAML formats are supported.
     * 
     */
    private String pluginData;
    /**
     * @return The first ID of the resource.
     * 
     */
    private String pluginId;
    /**
     * @return The name of the plug-in that you want to create.
     * 
     */
    private String pluginName;
    /**
     * @return The type of the plug-in.
     * 
     */
    private String pluginType;
    /**
     * @return The tag of the resource.
     * 
     */
    private Map<String,String> tags;

    private GetPluginsPlugin() {}
    /**
     * @return The CreateTime of the resource.
     * 
     */
    public String createTime() {
        return this.createTime;
    }
    /**
     * @return The description of the plug-in, which cannot exceed 200 characters.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The ID of the Plugin.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The ModifiedTime of the resource.
     * 
     */
    public String modifiedTime() {
        return this.modifiedTime;
    }
    /**
     * @return The definition statement of the plug-in. Plug-in definition statements in the JSON and YAML formats are supported.
     * 
     */
    public String pluginData() {
        return this.pluginData;
    }
    /**
     * @return The first ID of the resource.
     * 
     */
    public String pluginId() {
        return this.pluginId;
    }
    /**
     * @return The name of the plug-in that you want to create.
     * 
     */
    public String pluginName() {
        return this.pluginName;
    }
    /**
     * @return The type of the plug-in.
     * 
     */
    public String pluginType() {
        return this.pluginType;
    }
    /**
     * @return The tag of the resource.
     * 
     */
    public Map<String,String> tags() {
        return this.tags;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPluginsPlugin defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String createTime;
        private String description;
        private String id;
        private String modifiedTime;
        private String pluginData;
        private String pluginId;
        private String pluginName;
        private String pluginType;
        private Map<String,String> tags;
        public Builder() {}
        public Builder(GetPluginsPlugin defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createTime = defaults.createTime;
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.modifiedTime = defaults.modifiedTime;
    	      this.pluginData = defaults.pluginData;
    	      this.pluginId = defaults.pluginId;
    	      this.pluginName = defaults.pluginName;
    	      this.pluginType = defaults.pluginType;
    	      this.tags = defaults.tags;
        }

        @CustomType.Setter
        public Builder createTime(String createTime) {
            if (createTime == null) {
              throw new MissingRequiredPropertyException("GetPluginsPlugin", "createTime");
            }
            this.createTime = createTime;
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetPluginsPlugin", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetPluginsPlugin", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder modifiedTime(String modifiedTime) {
            if (modifiedTime == null) {
              throw new MissingRequiredPropertyException("GetPluginsPlugin", "modifiedTime");
            }
            this.modifiedTime = modifiedTime;
            return this;
        }
        @CustomType.Setter
        public Builder pluginData(String pluginData) {
            if (pluginData == null) {
              throw new MissingRequiredPropertyException("GetPluginsPlugin", "pluginData");
            }
            this.pluginData = pluginData;
            return this;
        }
        @CustomType.Setter
        public Builder pluginId(String pluginId) {
            if (pluginId == null) {
              throw new MissingRequiredPropertyException("GetPluginsPlugin", "pluginId");
            }
            this.pluginId = pluginId;
            return this;
        }
        @CustomType.Setter
        public Builder pluginName(String pluginName) {
            if (pluginName == null) {
              throw new MissingRequiredPropertyException("GetPluginsPlugin", "pluginName");
            }
            this.pluginName = pluginName;
            return this;
        }
        @CustomType.Setter
        public Builder pluginType(String pluginType) {
            if (pluginType == null) {
              throw new MissingRequiredPropertyException("GetPluginsPlugin", "pluginType");
            }
            this.pluginType = pluginType;
            return this;
        }
        @CustomType.Setter
        public Builder tags(Map<String,String> tags) {
            if (tags == null) {
              throw new MissingRequiredPropertyException("GetPluginsPlugin", "tags");
            }
            this.tags = tags;
            return this;
        }
        public GetPluginsPlugin build() {
            final var _resultValue = new GetPluginsPlugin();
            _resultValue.createTime = createTime;
            _resultValue.description = description;
            _resultValue.id = id;
            _resultValue.modifiedTime = modifiedTime;
            _resultValue.pluginData = pluginData;
            _resultValue.pluginId = pluginId;
            _resultValue.pluginName = pluginName;
            _resultValue.pluginType = pluginType;
            _resultValue.tags = tags;
            return _resultValue;
        }
    }
}
