// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.log.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class OssShipperParquetConfig {
    private String name;
    private String type;

    private OssShipperParquetConfig() {}
    public String name() {
        return this.name;
    }
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(OssShipperParquetConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String name;
        private String type;
        public Builder() {}
        public Builder(OssShipperParquetConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("OssShipperParquetConfig", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("OssShipperParquetConfig", "type");
            }
            this.type = type;
            return this;
        }
        public OssShipperParquetConfig build() {
            final var _resultValue = new OssShipperParquetConfig();
            _resultValue.name = name;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}
