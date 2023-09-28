// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ots.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SearchIndexSchemaFieldSchema {
    /**
     * @return Specifies the type of the analyzer that you want to use. If fieldType is set to Text, you can configure this parameter. Otherwise, the default analyzer type single-word tokenization is used.
     * 
     */
    private @Nullable String analyzer;
    /**
     * @return Specifies whether to enable sorting and aggregation. Type: Boolean. Sorting can be enabled only for fields for which enable_sort_and_agg is set to true.
     * 
     */
    private @Nullable Boolean enableSortAndAgg;
    /**
     * @return Specifies the name of the field in the search index. The value is used as a column name. A field in a search index can be a primary key column or an attribute column.
     * 
     */
    private String fieldName;
    /**
     * @return Specifies the type of the field. Use FieldType.XXX to set the type.
     * 
     */
    private String fieldType;
    /**
     * @return Specifies whether to enable indexing for the column. Type: Boolean.
     * 
     */
    private @Nullable Boolean index;
    /**
     * @return Specifies whether the value is an array. Type: Boolean.
     * 
     */
    private @Nullable Boolean isArray;
    /**
     * @return Specifies whether to store the value of the field in the search index. Type: Boolean. If you set store to true, you can read the value of the field from the search index without querying the data table. This improves query performance.
     * 
     */
    private @Nullable Boolean store;

    private SearchIndexSchemaFieldSchema() {}
    /**
     * @return Specifies the type of the analyzer that you want to use. If fieldType is set to Text, you can configure this parameter. Otherwise, the default analyzer type single-word tokenization is used.
     * 
     */
    public Optional<String> analyzer() {
        return Optional.ofNullable(this.analyzer);
    }
    /**
     * @return Specifies whether to enable sorting and aggregation. Type: Boolean. Sorting can be enabled only for fields for which enable_sort_and_agg is set to true.
     * 
     */
    public Optional<Boolean> enableSortAndAgg() {
        return Optional.ofNullable(this.enableSortAndAgg);
    }
    /**
     * @return Specifies the name of the field in the search index. The value is used as a column name. A field in a search index can be a primary key column or an attribute column.
     * 
     */
    public String fieldName() {
        return this.fieldName;
    }
    /**
     * @return Specifies the type of the field. Use FieldType.XXX to set the type.
     * 
     */
    public String fieldType() {
        return this.fieldType;
    }
    /**
     * @return Specifies whether to enable indexing for the column. Type: Boolean.
     * 
     */
    public Optional<Boolean> index() {
        return Optional.ofNullable(this.index);
    }
    /**
     * @return Specifies whether the value is an array. Type: Boolean.
     * 
     */
    public Optional<Boolean> isArray() {
        return Optional.ofNullable(this.isArray);
    }
    /**
     * @return Specifies whether to store the value of the field in the search index. Type: Boolean. If you set store to true, you can read the value of the field from the search index without querying the data table. This improves query performance.
     * 
     */
    public Optional<Boolean> store() {
        return Optional.ofNullable(this.store);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SearchIndexSchemaFieldSchema defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String analyzer;
        private @Nullable Boolean enableSortAndAgg;
        private String fieldName;
        private String fieldType;
        private @Nullable Boolean index;
        private @Nullable Boolean isArray;
        private @Nullable Boolean store;
        public Builder() {}
        public Builder(SearchIndexSchemaFieldSchema defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.analyzer = defaults.analyzer;
    	      this.enableSortAndAgg = defaults.enableSortAndAgg;
    	      this.fieldName = defaults.fieldName;
    	      this.fieldType = defaults.fieldType;
    	      this.index = defaults.index;
    	      this.isArray = defaults.isArray;
    	      this.store = defaults.store;
        }

        @CustomType.Setter
        public Builder analyzer(@Nullable String analyzer) {
            this.analyzer = analyzer;
            return this;
        }
        @CustomType.Setter
        public Builder enableSortAndAgg(@Nullable Boolean enableSortAndAgg) {
            this.enableSortAndAgg = enableSortAndAgg;
            return this;
        }
        @CustomType.Setter
        public Builder fieldName(String fieldName) {
            this.fieldName = Objects.requireNonNull(fieldName);
            return this;
        }
        @CustomType.Setter
        public Builder fieldType(String fieldType) {
            this.fieldType = Objects.requireNonNull(fieldType);
            return this;
        }
        @CustomType.Setter
        public Builder index(@Nullable Boolean index) {
            this.index = index;
            return this;
        }
        @CustomType.Setter
        public Builder isArray(@Nullable Boolean isArray) {
            this.isArray = isArray;
            return this;
        }
        @CustomType.Setter
        public Builder store(@Nullable Boolean store) {
            this.store = store;
            return this;
        }
        public SearchIndexSchemaFieldSchema build() {
            final var o = new SearchIndexSchemaFieldSchema();
            o.analyzer = analyzer;
            o.enableSortAndAgg = enableSortAndAgg;
            o.fieldName = fieldName;
            o.fieldType = fieldType;
            o.index = index;
            o.isArray = isArray;
            o.store = store;
            return o;
        }
    }
}
