// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ots.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetSecondaryIndexesIndex {
    /**
     * @return A list of defined column for index, referenced from Table&#39;s primary keys or predefined columns.
     * 
     */
    private List<String> definedColumns;
    /**
     * @return The resource ID. The value is `&lt;instance_name&gt;:&lt;table_name&gt;:&lt;indexName&gt;:&lt;indexType&gt;`.
     * 
     */
    private String id;
    /**
     * @return The index name of the OTS Table which could not be changed.
     * 
     */
    private String indexName;
    /**
     * @return The index type of the OTS Table which could not be changed.
     * 
     */
    private String indexType;
    /**
     * @return The name of OTS instance.
     * 
     */
    private String instanceName;
    /**
     * @return A list of primary keys for index, referenced from Table&#39;s primary keys or predefined columns.
     * 
     */
    private List<String> primaryKeys;
    /**
     * @return The name of OTS table.
     * 
     */
    private String tableName;

    private GetSecondaryIndexesIndex() {}
    /**
     * @return A list of defined column for index, referenced from Table&#39;s primary keys or predefined columns.
     * 
     */
    public List<String> definedColumns() {
        return this.definedColumns;
    }
    /**
     * @return The resource ID. The value is `&lt;instance_name&gt;:&lt;table_name&gt;:&lt;indexName&gt;:&lt;indexType&gt;`.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The index name of the OTS Table which could not be changed.
     * 
     */
    public String indexName() {
        return this.indexName;
    }
    /**
     * @return The index type of the OTS Table which could not be changed.
     * 
     */
    public String indexType() {
        return this.indexType;
    }
    /**
     * @return The name of OTS instance.
     * 
     */
    public String instanceName() {
        return this.instanceName;
    }
    /**
     * @return A list of primary keys for index, referenced from Table&#39;s primary keys or predefined columns.
     * 
     */
    public List<String> primaryKeys() {
        return this.primaryKeys;
    }
    /**
     * @return The name of OTS table.
     * 
     */
    public String tableName() {
        return this.tableName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSecondaryIndexesIndex defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> definedColumns;
        private String id;
        private String indexName;
        private String indexType;
        private String instanceName;
        private List<String> primaryKeys;
        private String tableName;
        public Builder() {}
        public Builder(GetSecondaryIndexesIndex defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.definedColumns = defaults.definedColumns;
    	      this.id = defaults.id;
    	      this.indexName = defaults.indexName;
    	      this.indexType = defaults.indexType;
    	      this.instanceName = defaults.instanceName;
    	      this.primaryKeys = defaults.primaryKeys;
    	      this.tableName = defaults.tableName;
        }

        @CustomType.Setter
        public Builder definedColumns(List<String> definedColumns) {
            this.definedColumns = Objects.requireNonNull(definedColumns);
            return this;
        }
        public Builder definedColumns(String... definedColumns) {
            return definedColumns(List.of(definedColumns));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder indexName(String indexName) {
            this.indexName = Objects.requireNonNull(indexName);
            return this;
        }
        @CustomType.Setter
        public Builder indexType(String indexType) {
            this.indexType = Objects.requireNonNull(indexType);
            return this;
        }
        @CustomType.Setter
        public Builder instanceName(String instanceName) {
            this.instanceName = Objects.requireNonNull(instanceName);
            return this;
        }
        @CustomType.Setter
        public Builder primaryKeys(List<String> primaryKeys) {
            this.primaryKeys = Objects.requireNonNull(primaryKeys);
            return this;
        }
        public Builder primaryKeys(String... primaryKeys) {
            return primaryKeys(List.of(primaryKeys));
        }
        @CustomType.Setter
        public Builder tableName(String tableName) {
            this.tableName = Objects.requireNonNull(tableName);
            return this;
        }
        public GetSecondaryIndexesIndex build() {
            final var o = new GetSecondaryIndexesIndex();
            o.definedColumns = definedColumns;
            o.id = id;
            o.indexName = indexName;
            o.indexType = indexType;
            o.instanceName = instanceName;
            o.primaryKeys = primaryKeys;
            o.tableName = tableName;
            return o;
        }
    }
}
