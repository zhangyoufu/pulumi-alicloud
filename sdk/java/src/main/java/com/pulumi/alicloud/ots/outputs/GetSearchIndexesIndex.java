// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ots.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetSearchIndexesIndex {
    private Integer createTime;
    private Integer currentSyncTimestamp;
    private String id;
    private String indexName;
    private String instanceName;
    private Integer meteringLastUpdateTime;
    private Integer reservedReadCu;
    private Integer rowCount;
    private String schema;
    private Integer storageSize;
    private String syncPhase;
    private String tableName;
    private Integer timeToLive;

    private GetSearchIndexesIndex() {}
    public Integer createTime() {
        return this.createTime;
    }
    public Integer currentSyncTimestamp() {
        return this.currentSyncTimestamp;
    }
    public String id() {
        return this.id;
    }
    public String indexName() {
        return this.indexName;
    }
    public String instanceName() {
        return this.instanceName;
    }
    public Integer meteringLastUpdateTime() {
        return this.meteringLastUpdateTime;
    }
    public Integer reservedReadCu() {
        return this.reservedReadCu;
    }
    public Integer rowCount() {
        return this.rowCount;
    }
    public String schema() {
        return this.schema;
    }
    public Integer storageSize() {
        return this.storageSize;
    }
    public String syncPhase() {
        return this.syncPhase;
    }
    public String tableName() {
        return this.tableName;
    }
    public Integer timeToLive() {
        return this.timeToLive;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSearchIndexesIndex defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer createTime;
        private Integer currentSyncTimestamp;
        private String id;
        private String indexName;
        private String instanceName;
        private Integer meteringLastUpdateTime;
        private Integer reservedReadCu;
        private Integer rowCount;
        private String schema;
        private Integer storageSize;
        private String syncPhase;
        private String tableName;
        private Integer timeToLive;
        public Builder() {}
        public Builder(GetSearchIndexesIndex defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createTime = defaults.createTime;
    	      this.currentSyncTimestamp = defaults.currentSyncTimestamp;
    	      this.id = defaults.id;
    	      this.indexName = defaults.indexName;
    	      this.instanceName = defaults.instanceName;
    	      this.meteringLastUpdateTime = defaults.meteringLastUpdateTime;
    	      this.reservedReadCu = defaults.reservedReadCu;
    	      this.rowCount = defaults.rowCount;
    	      this.schema = defaults.schema;
    	      this.storageSize = defaults.storageSize;
    	      this.syncPhase = defaults.syncPhase;
    	      this.tableName = defaults.tableName;
    	      this.timeToLive = defaults.timeToLive;
        }

        @CustomType.Setter
        public Builder createTime(Integer createTime) {
            this.createTime = Objects.requireNonNull(createTime);
            return this;
        }
        @CustomType.Setter
        public Builder currentSyncTimestamp(Integer currentSyncTimestamp) {
            this.currentSyncTimestamp = Objects.requireNonNull(currentSyncTimestamp);
            return this;
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
        public Builder instanceName(String instanceName) {
            this.instanceName = Objects.requireNonNull(instanceName);
            return this;
        }
        @CustomType.Setter
        public Builder meteringLastUpdateTime(Integer meteringLastUpdateTime) {
            this.meteringLastUpdateTime = Objects.requireNonNull(meteringLastUpdateTime);
            return this;
        }
        @CustomType.Setter
        public Builder reservedReadCu(Integer reservedReadCu) {
            this.reservedReadCu = Objects.requireNonNull(reservedReadCu);
            return this;
        }
        @CustomType.Setter
        public Builder rowCount(Integer rowCount) {
            this.rowCount = Objects.requireNonNull(rowCount);
            return this;
        }
        @CustomType.Setter
        public Builder schema(String schema) {
            this.schema = Objects.requireNonNull(schema);
            return this;
        }
        @CustomType.Setter
        public Builder storageSize(Integer storageSize) {
            this.storageSize = Objects.requireNonNull(storageSize);
            return this;
        }
        @CustomType.Setter
        public Builder syncPhase(String syncPhase) {
            this.syncPhase = Objects.requireNonNull(syncPhase);
            return this;
        }
        @CustomType.Setter
        public Builder tableName(String tableName) {
            this.tableName = Objects.requireNonNull(tableName);
            return this;
        }
        @CustomType.Setter
        public Builder timeToLive(Integer timeToLive) {
            this.timeToLive = Objects.requireNonNull(timeToLive);
            return this;
        }
        public GetSearchIndexesIndex build() {
            final var o = new GetSearchIndexesIndex();
            o.createTime = createTime;
            o.currentSyncTimestamp = currentSyncTimestamp;
            o.id = id;
            o.indexName = indexName;
            o.instanceName = instanceName;
            o.meteringLastUpdateTime = meteringLastUpdateTime;
            o.reservedReadCu = reservedReadCu;
            o.rowCount = rowCount;
            o.schema = schema;
            o.storageSize = storageSize;
            o.syncPhase = syncPhase;
            o.tableName = tableName;
            o.timeToLive = timeToLive;
            return o;
        }
    }
}
