// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetRouteTablesTable {
    /**
     * @return The description of the route table instance.
     * 
     */
    private String description;
    /**
     * @return ID of the Route Table.
     * 
     */
    private String id;
    /**
     * @return Name of the route table.
     * 
     */
    private String name;
    /**
     * @return The Id of resource group which route tables belongs.
     * 
     */
    private String resourceGroupId;
    /**
     * @return The route table id.
     * 
     */
    private String routeTableId;
    /**
     * @return The route table name.
     * 
     */
    private String routeTableName;
    /**
     * @return The type of route table.
     * 
     */
    private String routeTableType;
    /**
     * @return The router ID.
     * 
     */
    private String routerId;
    /**
     * @return The route type of route table. Valid values: `VRouter` and `VBR`.
     * 
     */
    private String routerType;
    /**
     * @return The status of resource. Valid values: `Available` and `Pending`.
     * 
     */
    private String status;
    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    private Map<String,String> tags;
    /**
     * @return Vpc id of the route table.
     * 
     */
    private String vpcId;
    /**
     * @return A list of vswitch id.
     * 
     */
    private List<String> vswitchIds;

    private GetRouteTablesTable() {}
    /**
     * @return The description of the route table instance.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return ID of the Route Table.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Name of the route table.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The Id of resource group which route tables belongs.
     * 
     */
    public String resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * @return The route table id.
     * 
     */
    public String routeTableId() {
        return this.routeTableId;
    }
    /**
     * @return The route table name.
     * 
     */
    public String routeTableName() {
        return this.routeTableName;
    }
    /**
     * @return The type of route table.
     * 
     */
    public String routeTableType() {
        return this.routeTableType;
    }
    /**
     * @return The router ID.
     * 
     */
    public String routerId() {
        return this.routerId;
    }
    /**
     * @return The route type of route table. Valid values: `VRouter` and `VBR`.
     * 
     */
    public String routerType() {
        return this.routerType;
    }
    /**
     * @return The status of resource. Valid values: `Available` and `Pending`.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Map<String,String> tags() {
        return this.tags;
    }
    /**
     * @return Vpc id of the route table.
     * 
     */
    public String vpcId() {
        return this.vpcId;
    }
    /**
     * @return A list of vswitch id.
     * 
     */
    public List<String> vswitchIds() {
        return this.vswitchIds;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRouteTablesTable defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private String id;
        private String name;
        private String resourceGroupId;
        private String routeTableId;
        private String routeTableName;
        private String routeTableType;
        private String routerId;
        private String routerType;
        private String status;
        private Map<String,String> tags;
        private String vpcId;
        private List<String> vswitchIds;
        public Builder() {}
        public Builder(GetRouteTablesTable defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.resourceGroupId = defaults.resourceGroupId;
    	      this.routeTableId = defaults.routeTableId;
    	      this.routeTableName = defaults.routeTableName;
    	      this.routeTableType = defaults.routeTableType;
    	      this.routerId = defaults.routerId;
    	      this.routerType = defaults.routerType;
    	      this.status = defaults.status;
    	      this.tags = defaults.tags;
    	      this.vpcId = defaults.vpcId;
    	      this.vswitchIds = defaults.vswitchIds;
        }

        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetRouteTablesTable", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetRouteTablesTable", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetRouteTablesTable", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder resourceGroupId(String resourceGroupId) {
            if (resourceGroupId == null) {
              throw new MissingRequiredPropertyException("GetRouteTablesTable", "resourceGroupId");
            }
            this.resourceGroupId = resourceGroupId;
            return this;
        }
        @CustomType.Setter
        public Builder routeTableId(String routeTableId) {
            if (routeTableId == null) {
              throw new MissingRequiredPropertyException("GetRouteTablesTable", "routeTableId");
            }
            this.routeTableId = routeTableId;
            return this;
        }
        @CustomType.Setter
        public Builder routeTableName(String routeTableName) {
            if (routeTableName == null) {
              throw new MissingRequiredPropertyException("GetRouteTablesTable", "routeTableName");
            }
            this.routeTableName = routeTableName;
            return this;
        }
        @CustomType.Setter
        public Builder routeTableType(String routeTableType) {
            if (routeTableType == null) {
              throw new MissingRequiredPropertyException("GetRouteTablesTable", "routeTableType");
            }
            this.routeTableType = routeTableType;
            return this;
        }
        @CustomType.Setter
        public Builder routerId(String routerId) {
            if (routerId == null) {
              throw new MissingRequiredPropertyException("GetRouteTablesTable", "routerId");
            }
            this.routerId = routerId;
            return this;
        }
        @CustomType.Setter
        public Builder routerType(String routerType) {
            if (routerType == null) {
              throw new MissingRequiredPropertyException("GetRouteTablesTable", "routerType");
            }
            this.routerType = routerType;
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            if (status == null) {
              throw new MissingRequiredPropertyException("GetRouteTablesTable", "status");
            }
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder tags(Map<String,String> tags) {
            if (tags == null) {
              throw new MissingRequiredPropertyException("GetRouteTablesTable", "tags");
            }
            this.tags = tags;
            return this;
        }
        @CustomType.Setter
        public Builder vpcId(String vpcId) {
            if (vpcId == null) {
              throw new MissingRequiredPropertyException("GetRouteTablesTable", "vpcId");
            }
            this.vpcId = vpcId;
            return this;
        }
        @CustomType.Setter
        public Builder vswitchIds(List<String> vswitchIds) {
            if (vswitchIds == null) {
              throw new MissingRequiredPropertyException("GetRouteTablesTable", "vswitchIds");
            }
            this.vswitchIds = vswitchIds;
            return this;
        }
        public Builder vswitchIds(String... vswitchIds) {
            return vswitchIds(List.of(vswitchIds));
        }
        public GetRouteTablesTable build() {
            final var _resultValue = new GetRouteTablesTable();
            _resultValue.description = description;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.resourceGroupId = resourceGroupId;
            _resultValue.routeTableId = routeTableId;
            _resultValue.routeTableName = routeTableName;
            _resultValue.routeTableType = routeTableType;
            _resultValue.routerId = routerId;
            _resultValue.routerType = routerType;
            _resultValue.status = status;
            _resultValue.tags = tags;
            _resultValue.vpcId = vpcId;
            _resultValue.vswitchIds = vswitchIds;
            return _resultValue;
        }
    }
}
