// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.fc;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.fc.TriggerArgs;
import com.pulumi.alicloud.fc.inputs.TriggerState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an Alicloud Function Compute Trigger resource. Based on trigger, execute your code in response to events in Alibaba Cloud.
 *  For information about Service and how to use it, see [What is Function Compute](https://www.alibabacloud.com/help/en/fc/developer-reference/api-fc-open-2021-04-06-createtrigger).
 * 
 * &gt; **NOTE:** The resource requires a provider field &#39;account_id&#39;. See account_id.
 * 
 * &gt; **NOTE:** Available since v1.93.0.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetRegionsArgs;
 * import com.pulumi.random.integer;
 * import com.pulumi.random.IntegerArgs;
 * import com.pulumi.alicloud.log.Project;
 * import com.pulumi.alicloud.log.ProjectArgs;
 * import com.pulumi.alicloud.log.Store;
 * import com.pulumi.alicloud.log.StoreArgs;
 * import com.pulumi.alicloud.ram.Role;
 * import com.pulumi.alicloud.ram.RoleArgs;
 * import com.pulumi.alicloud.ram.RolePolicyAttachment;
 * import com.pulumi.alicloud.ram.RolePolicyAttachmentArgs;
 * import com.pulumi.alicloud.fc.Service;
 * import com.pulumi.alicloud.fc.ServiceArgs;
 * import com.pulumi.alicloud.fc.inputs.ServiceLogConfigArgs;
 * import com.pulumi.alicloud.oss.Bucket;
 * import com.pulumi.alicloud.oss.BucketArgs;
 * import com.pulumi.alicloud.oss.BucketObject;
 * import com.pulumi.alicloud.oss.BucketObjectArgs;
 * import com.pulumi.alicloud.fc.Function;
 * import com.pulumi.alicloud.fc.FunctionArgs;
 * import com.pulumi.alicloud.fc.Trigger;
 * import com.pulumi.alicloud.fc.TriggerArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var default = AlicloudFunctions.getAccount();
 * 
 *         final var defaultGetRegions = AlicloudFunctions.getRegions(GetRegionsArgs.builder()
 *             .current(true)
 *             .build());
 * 
 *         var defaultInteger = new Integer("defaultInteger", IntegerArgs.builder()        
 *             .max(99999)
 *             .min(10000)
 *             .build());
 * 
 *         var defaultProject = new Project("defaultProject", ProjectArgs.builder()        
 *             .projectName(String.format("example-value-%s", defaultInteger.result()))
 *             .build());
 * 
 *         var defaultStore = new Store("defaultStore", StoreArgs.builder()        
 *             .projectName(defaultProject.name())
 *             .logstoreName("example-value")
 *             .build());
 * 
 *         var sourceStore = new Store("sourceStore", StoreArgs.builder()        
 *             .projectName(defaultProject.name())
 *             .logstoreName("example-source-store")
 *             .build());
 * 
 *         var defaultRole = new Role("defaultRole", RoleArgs.builder()        
 *             .name(String.format("fcservicerole-%s", defaultInteger.result()))
 *             .document("""
 *   {
 *       "Statement": [
 *         {
 *           "Action": "sts:AssumeRole",
 *           "Effect": "Allow",
 *           "Principal": {
 *             "Service": [
 *               "fc.aliyuncs.com"
 *             ]
 *           }
 *         }
 *       ],
 *       "Version": "1"
 *   }
 *             """)
 *             .description("this is a example")
 *             .force(true)
 *             .build());
 * 
 *         var defaultRolePolicyAttachment = new RolePolicyAttachment("defaultRolePolicyAttachment", RolePolicyAttachmentArgs.builder()        
 *             .roleName(defaultRole.name())
 *             .policyName("AliyunLogFullAccess")
 *             .policyType("System")
 *             .build());
 * 
 *         var defaultService = new Service("defaultService", ServiceArgs.builder()        
 *             .name(String.format("example-value-%s", defaultInteger.result()))
 *             .description("example-value")
 *             .role(defaultRole.arn())
 *             .logConfig(ServiceLogConfigArgs.builder()
 *                 .project(defaultProject.name())
 *                 .logstore(defaultStore.name())
 *                 .enableInstanceMetrics(true)
 *                 .enableRequestMetrics(true)
 *                 .build())
 *             .build());
 * 
 *         var defaultBucket = new Bucket("defaultBucket", BucketArgs.builder()        
 *             .bucket(String.format("terraform-example-%s", defaultInteger.result()))
 *             .build());
 * 
 *         // If you upload the function by OSS Bucket, you need to specify path can't upload by content.
 *         var defaultBucketObject = new BucketObject("defaultBucketObject", BucketObjectArgs.builder()        
 *             .bucket(defaultBucket.id())
 *             .key("index.py")
 *             .content("""
 * import logging 
 * def handler(event, context): 
 * logger = logging.getLogger() 
 * logger.info('hello world') 
 * return 'hello world'            """)
 *             .build());
 * 
 *         var defaultFunction = new Function("defaultFunction", FunctionArgs.builder()        
 *             .service(defaultService.name())
 *             .name("terraform-example")
 *             .description("example")
 *             .ossBucket(defaultBucket.id())
 *             .ossKey(defaultBucketObject.key())
 *             .memorySize("512")
 *             .runtime("python3.10")
 *             .handler("hello.handler")
 *             .build());
 * 
 *         var defaultTrigger = new Trigger("defaultTrigger", TriggerArgs.builder()        
 *             .service(defaultService.name())
 *             .function(defaultFunction.name())
 *             .name("terraform-example")
 *             .role(defaultRole.arn())
 *             .sourceArn(defaultProject.name().applyValue(name -> String.format("acs:log:%s:%s:project/%s", defaultGetRegions.applyValue(getRegionsResult -> getRegionsResult.regions()[0].id()),default_.id(),name)))
 *             .type("log")
 *             .config(Output.tuple(sourceStore.name(), defaultProject.name(), defaultStore.name()).applyValue(values -> {
 *                 var sourceStoreName = values.t1;
 *                 var defaultProjectName = values.t2;
 *                 var defaultStoreName = values.t3;
 *                 return """
 *     {
 *         "sourceConfig": {
 *             "logstore": "%s",
 *             "startTime": null
 *         },
 *         "jobConfig": {
 *             "maxRetryTime": 3,
 *             "triggerInterval": 60
 *         },
 *         "functionParameter": {
 *             "a": "b",
 *             "c": "d"
 *         },
 *         "logConfig": {
 *              "project": "%s",
 *             "logstore": "%s"
 *         },
 *         "targetConfig": null,
 *         "enable": true
 *     }
 *   
 * ", sourceStoreName,defaultProjectName,defaultStoreName);
 *             }))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * MNS topic trigger:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetRegionsArgs;
 * import com.pulumi.random.integer;
 * import com.pulumi.random.IntegerArgs;
 * import com.pulumi.alicloud.mns.Topic;
 * import com.pulumi.alicloud.mns.TopicArgs;
 * import com.pulumi.alicloud.ram.Role;
 * import com.pulumi.alicloud.ram.RoleArgs;
 * import com.pulumi.alicloud.ram.RolePolicyAttachment;
 * import com.pulumi.alicloud.ram.RolePolicyAttachmentArgs;
 * import com.pulumi.alicloud.fc.Service;
 * import com.pulumi.alicloud.fc.ServiceArgs;
 * import com.pulumi.alicloud.oss.Bucket;
 * import com.pulumi.alicloud.oss.BucketArgs;
 * import com.pulumi.alicloud.oss.BucketObject;
 * import com.pulumi.alicloud.oss.BucketObjectArgs;
 * import com.pulumi.alicloud.fc.Function;
 * import com.pulumi.alicloud.fc.FunctionArgs;
 * import com.pulumi.alicloud.fc.Trigger;
 * import com.pulumi.alicloud.fc.TriggerArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var default = AlicloudFunctions.getAccount();
 * 
 *         final var defaultGetRegions = AlicloudFunctions.getRegions(GetRegionsArgs.builder()
 *             .current(true)
 *             .build());
 * 
 *         var defaultInteger = new Integer("defaultInteger", IntegerArgs.builder()        
 *             .max(99999)
 *             .min(10000)
 *             .build());
 * 
 *         var defaultTopic = new Topic("defaultTopic", TopicArgs.builder()        
 *             .name(String.format("example-value-%s", defaultInteger.result()))
 *             .build());
 * 
 *         var defaultRole = new Role("defaultRole", RoleArgs.builder()        
 *             .name(String.format("fcservicerole-%s", defaultInteger.result()))
 *             .document("""
 *   {
 *       "Statement": [
 *         {
 *           "Action": "sts:AssumeRole",
 *           "Effect": "Allow",
 *           "Principal": {
 *             "Service": [
 *               "mns.aliyuncs.com"
 *             ]
 *           }
 *         }
 *       ],
 *       "Version": "1"
 *   }
 *             """)
 *             .description("this is a example")
 *             .force(true)
 *             .build());
 * 
 *         var defaultRolePolicyAttachment = new RolePolicyAttachment("defaultRolePolicyAttachment", RolePolicyAttachmentArgs.builder()        
 *             .roleName(defaultRole.name())
 *             .policyName("AliyunMNSNotificationRolePolicy")
 *             .policyType("System")
 *             .build());
 * 
 *         var defaultService = new Service("defaultService", ServiceArgs.builder()        
 *             .name(String.format("example-value-%s", defaultInteger.result()))
 *             .description("example-value")
 *             .internetAccess(false)
 *             .build());
 * 
 *         var defaultBucket = new Bucket("defaultBucket", BucketArgs.builder()        
 *             .bucket(String.format("terraform-example-%s", defaultInteger.result()))
 *             .build());
 * 
 *         // If you upload the function by OSS Bucket, you need to specify path can't upload by content.
 *         var defaultBucketObject = new BucketObject("defaultBucketObject", BucketObjectArgs.builder()        
 *             .bucket(defaultBucket.id())
 *             .key("index.py")
 *             .content("""
 * import logging 
 * def handler(event, context): 
 * logger = logging.getLogger() 
 * logger.info('hello world') 
 * return 'hello world'            """)
 *             .build());
 * 
 *         var defaultFunction = new Function("defaultFunction", FunctionArgs.builder()        
 *             .service(defaultService.name())
 *             .name(String.format("terraform-example-%s", defaultInteger.result()))
 *             .description("example")
 *             .ossBucket(defaultBucket.id())
 *             .ossKey(defaultBucketObject.key())
 *             .memorySize("512")
 *             .runtime("python3.10")
 *             .handler("hello.handler")
 *             .build());
 * 
 *         var defaultTrigger = new Trigger("defaultTrigger", TriggerArgs.builder()        
 *             .service(defaultService.name())
 *             .function(defaultFunction.name())
 *             .name("terraform-example")
 *             .role(defaultRole.arn())
 *             .sourceArn(defaultTopic.name().applyValue(name -> String.format("acs:mns:%s:%s:/topics/%s", defaultGetRegions.applyValue(getRegionsResult -> getRegionsResult.regions()[0].id()),default_.id(),name)))
 *             .type("mns_topic")
 *             .configMns("""
 *   {
 *     "filterTag":"exampleTag",
 *     "notifyContentFormat":"STREAM",
 *     "notifyStrategy":"BACKOFF_RETRY"
 *   }
 *             """)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * CDN events trigger:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.random.integer;
 * import com.pulumi.random.IntegerArgs;
 * import com.pulumi.alicloud.cdn.DomainNew;
 * import com.pulumi.alicloud.cdn.DomainNewArgs;
 * import com.pulumi.alicloud.cdn.inputs.DomainNewSourceArgs;
 * import com.pulumi.alicloud.fc.Service;
 * import com.pulumi.alicloud.fc.ServiceArgs;
 * import com.pulumi.alicloud.ram.Role;
 * import com.pulumi.alicloud.ram.RoleArgs;
 * import com.pulumi.alicloud.ram.Policy;
 * import com.pulumi.alicloud.ram.PolicyArgs;
 * import com.pulumi.alicloud.ram.RolePolicyAttachment;
 * import com.pulumi.alicloud.ram.RolePolicyAttachmentArgs;
 * import com.pulumi.alicloud.oss.Bucket;
 * import com.pulumi.alicloud.oss.BucketArgs;
 * import com.pulumi.alicloud.oss.BucketObject;
 * import com.pulumi.alicloud.oss.BucketObjectArgs;
 * import com.pulumi.alicloud.fc.Function;
 * import com.pulumi.alicloud.fc.FunctionArgs;
 * import com.pulumi.alicloud.fc.Trigger;
 * import com.pulumi.alicloud.fc.TriggerArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var default = AlicloudFunctions.getAccount();
 * 
 *         var defaultInteger = new Integer("defaultInteger", IntegerArgs.builder()        
 *             .max(99999)
 *             .min(10000)
 *             .build());
 * 
 *         var defaultDomainNew = new DomainNew("defaultDomainNew", DomainNewArgs.builder()        
 *             .domainName(String.format("example%s.tf.com", defaultInteger.result()))
 *             .cdnType("web")
 *             .scope("overseas")
 *             .sources(DomainNewSourceArgs.builder()
 *                 .content("1.1.1.1")
 *                 .type("ipaddr")
 *                 .priority(20)
 *                 .port(80)
 *                 .weight(10)
 *                 .build())
 *             .build());
 * 
 *         var defaultService = new Service("defaultService", ServiceArgs.builder()        
 *             .name(String.format("example-value-%s", defaultInteger.result()))
 *             .description("example-value")
 *             .internetAccess(false)
 *             .build());
 * 
 *         var defaultRole = new Role("defaultRole", RoleArgs.builder()        
 *             .name(String.format("fcservicerole-%s", defaultInteger.result()))
 *             .document("""
 *     {
 *       "Statement": [
 *         {
 *           "Action": "sts:AssumeRole",
 *           "Effect": "Allow",
 *           "Principal": {
 *             "Service": [
 *               "cdn.aliyuncs.com"
 *             ]
 *           }
 *         }
 *       ],
 *       "Version": "1"
 *   }
 *             """)
 *             .description("this is a example")
 *             .force(true)
 *             .build());
 * 
 *         var defaultPolicy = new Policy("defaultPolicy", PolicyArgs.builder()        
 *             .policyName(String.format("fcservicepolicy-%s", defaultInteger.result()))
 *             .policyDocument(Output.tuple(defaultService.name(), defaultService.name()).applyValue(values -> {
 *                 var defaultServiceName = values.t1;
 *                 var defaultServiceName1 = values.t2;
 *                 return """
 *     {
 *         "Version": "1",
 *         "Statement": [
 *         {
 *             "Action": [
 *             "fc:InvokeFunction"
 *             ],
 *         "Resource": [
 *             "acs:fc:*:*:services/%s/functions/*",
 *             "acs:fc:*:*:services/%s.*{@literal /}functions/*"
 *         ],
 *         "Effect": "Allow"
 *         }
 *         ]
 *     }
 * ", defaultServiceName,defaultServiceName1);
 *             }))
 *             .description("this is a example")
 *             .force(true)
 *             .build());
 * 
 *         var defaultRolePolicyAttachment = new RolePolicyAttachment("defaultRolePolicyAttachment", RolePolicyAttachmentArgs.builder()        
 *             .roleName(defaultRole.name())
 *             .policyName(defaultPolicy.name())
 *             .policyType("Custom")
 *             .build());
 * 
 *         var defaultBucket = new Bucket("defaultBucket", BucketArgs.builder()        
 *             .bucket(String.format("terraform-example-%s", defaultInteger.result()))
 *             .build());
 * 
 *         // If you upload the function by OSS Bucket, you need to specify path can't upload by content.
 *         var defaultBucketObject = new BucketObject("defaultBucketObject", BucketObjectArgs.builder()        
 *             .bucket(defaultBucket.id())
 *             .key("index.py")
 *             .content("""
 * import logging 
 * def handler(event, context): 
 * logger = logging.getLogger() 
 * logger.info('hello world') 
 * return 'hello world'            """)
 *             .build());
 * 
 *         var defaultFunction = new Function("defaultFunction", FunctionArgs.builder()        
 *             .service(defaultService.name())
 *             .name(String.format("terraform-example-%s", defaultInteger.result()))
 *             .description("example")
 *             .ossBucket(defaultBucket.id())
 *             .ossKey(defaultBucketObject.key())
 *             .memorySize("512")
 *             .runtime("python3.10")
 *             .handler("hello.handler")
 *             .build());
 * 
 *         var defaultTrigger = new Trigger("defaultTrigger", TriggerArgs.builder()        
 *             .service(defaultService.name())
 *             .function(defaultFunction.name())
 *             .name("terraform-example")
 *             .role(defaultRole.arn())
 *             .sourceArn(String.format("acs:cdn:*:%s", default_.id()))
 *             .type("cdn_events")
 *             .config(defaultDomainNew.domainName().applyValue(domainName -> """
 *       {"eventName":"LogFileCreated",
 *      "eventVersion":"1.0.0",
 *      "notes":"cdn events trigger",
 *      "filter":{
 *         "domain": ["%s"]
 *         }
 *     }
 * ", domainName)))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * EventBridge trigger:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetRegionsArgs;
 * import com.pulumi.random.integer;
 * import com.pulumi.random.IntegerArgs;
 * import com.pulumi.alicloud.eventbridge.ServiceLinkedRole;
 * import com.pulumi.alicloud.eventbridge.ServiceLinkedRoleArgs;
 * import com.pulumi.alicloud.fc.Service;
 * import com.pulumi.alicloud.fc.ServiceArgs;
 * import com.pulumi.alicloud.oss.Bucket;
 * import com.pulumi.alicloud.oss.BucketArgs;
 * import com.pulumi.alicloud.oss.BucketObject;
 * import com.pulumi.alicloud.oss.BucketObjectArgs;
 * import com.pulumi.alicloud.fc.Function;
 * import com.pulumi.alicloud.fc.FunctionArgs;
 * import com.pulumi.alicloud.fc.Trigger;
 * import com.pulumi.alicloud.fc.TriggerArgs;
 * import com.pulumi.alicloud.rocketmq.Instance;
 * import com.pulumi.alicloud.rocketmq.InstanceArgs;
 * import com.pulumi.alicloud.rocketmq.Group;
 * import com.pulumi.alicloud.rocketmq.GroupArgs;
 * import com.pulumi.alicloud.rocketmq.Topic;
 * import com.pulumi.alicloud.rocketmq.TopicArgs;
 * import com.pulumi.alicloud.amqp.Instance;
 * import com.pulumi.alicloud.amqp.InstanceArgs;
 * import com.pulumi.alicloud.amqp.VirtualHost;
 * import com.pulumi.alicloud.amqp.VirtualHostArgs;
 * import com.pulumi.alicloud.amqp.Queue;
 * import com.pulumi.alicloud.amqp.QueueArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var default = AlicloudFunctions.getAccount();
 * 
 *         final var defaultGetRegions = AlicloudFunctions.getRegions(GetRegionsArgs.builder()
 *             .current(true)
 *             .build());
 * 
 *         var defaultInteger = new Integer("defaultInteger", IntegerArgs.builder()        
 *             .max(99999)
 *             .min(10000)
 *             .build());
 * 
 *         var serviceLinkedRole = new ServiceLinkedRole("serviceLinkedRole", ServiceLinkedRoleArgs.builder()        
 *             .productName("AliyunServiceRoleForEventBridgeSendToFC")
 *             .build());
 * 
 *         var defaultService = new Service("defaultService", ServiceArgs.builder()        
 *             .name(String.format("example-value-%s", defaultInteger.result()))
 *             .description("example-value")
 *             .internetAccess(false)
 *             .build());
 * 
 *         var defaultBucket = new Bucket("defaultBucket", BucketArgs.builder()        
 *             .bucket(String.format("terraform-example-%s", defaultInteger.result()))
 *             .build());
 * 
 *         // If you upload the function by OSS Bucket, you need to specify path can't upload by content.
 *         var defaultBucketObject = new BucketObject("defaultBucketObject", BucketObjectArgs.builder()        
 *             .bucket(defaultBucket.id())
 *             .key("index.py")
 *             .content("""
 * import logging 
 * def handler(event, context): 
 * logger = logging.getLogger() 
 * logger.info('hello world') 
 * return 'hello world'            """)
 *             .build());
 * 
 *         var defaultFunction = new Function("defaultFunction", FunctionArgs.builder()        
 *             .service(defaultService.name())
 *             .name("terraform-example")
 *             .description("example")
 *             .ossBucket(defaultBucket.id())
 *             .ossKey(defaultBucketObject.key())
 *             .memorySize("512")
 *             .runtime("python3.10")
 *             .handler("hello.handler")
 *             .build());
 * 
 *         var ossTrigger = new Trigger("ossTrigger", TriggerArgs.builder()        
 *             .service(defaultService.name())
 *             .function(defaultFunction.name())
 *             .name("terraform-example-oss")
 *             .type("eventbridge")
 *             .config("""
 *     {
 *         "triggerEnable": false,
 *         "asyncInvocationType": false,
 *         "eventRuleFilterPattern": {
 *           "source":[
 *             "acs.oss"
 *             ],
 *             "type":[
 *               "oss:BucketCreated:PutBucket"
 *             ]
 *         },
 *         "eventSourceConfig": {
 *             "eventSourceType": "Default"
 *         }
 *     }
 *             """)
 *             .build());
 * 
 *         var mnsTrigger = new Trigger("mnsTrigger", TriggerArgs.builder()        
 *             .service(defaultService.name())
 *             .function(defaultFunction.name())
 *             .name("terraform-example-mns")
 *             .type("eventbridge")
 *             .config("""
 *     {
 *         "triggerEnable": false,
 *         "asyncInvocationType": false,
 *         "eventRuleFilterPattern": "{}",
 *         "eventSourceConfig": {
 *             "eventSourceType": "MNS",
 *             "eventSourceParameters": {
 *                 "sourceMNSParameters": {
 *                     "RegionId": "cn-hangzhou",
 *                     "QueueName": "mns-queue",
 *                     "IsBase64Decode": true
 *                 }
 *             }
 *         }
 *     }
 *             """)
 *             .build());
 * 
 *         var defaultInstance = new Instance("defaultInstance", InstanceArgs.builder()        
 *             .instanceName(String.format("terraform-example-%s", defaultInteger.result()))
 *             .remark("terraform-example")
 *             .build());
 * 
 *         var defaultGroup = new Group("defaultGroup", GroupArgs.builder()        
 *             .groupName("GID-example")
 *             .instanceId(defaultInstance.id())
 *             .remark("terraform-example")
 *             .build());
 * 
 *         var defaultTopic = new Topic("defaultTopic", TopicArgs.builder()        
 *             .topicName("mytopic")
 *             .instanceId(defaultInstance.id())
 *             .messageType(0)
 *             .remark("terraform-example")
 *             .build());
 * 
 *         var rocketmqTrigger = new Trigger("rocketmqTrigger", TriggerArgs.builder()        
 *             .service(defaultService.name())
 *             .function(defaultFunction.name())
 *             .name("terraform-example-rocketmq")
 *             .type("eventbridge")
 *             .config(Output.tuple(defaultInstance.id(), defaultGroup.groupName(), defaultTopic.topicName()).applyValue(values -> {
 *                 var id = values.t1;
 *                 var groupName = values.t2;
 *                 var topicName = values.t3;
 *                 return """
 *     {
 *         "triggerEnable": false,
 *         "asyncInvocationType": false,
 *         "eventRuleFilterPattern": "{}",
 *         "eventSourceConfig": {
 *             "eventSourceType": "RocketMQ",
 *             "eventSourceParameters": {
 *                 "sourceRocketMQParameters": {
 *                     "RegionId": "%s",
 *                     "InstanceId": "%s",
 *                     "GroupID": "%s",
 *                     "Topic": "%s",
 *                     "Timestamp": 1686296162,
 *                     "Tag": "example-tag",
 *                     "Offset": "CONSUME_FROM_LAST_OFFSET"
 *                 }
 *             }
 *         }
 *     }
 * ", defaultGetRegions.applyValue(getRegionsResult -> getRegionsResult.regions()[0].id()),id,groupName,topicName);
 *             }))
 *             .build());
 * 
 *         var defaultInstance2 = new Instance("defaultInstance2", InstanceArgs.builder()        
 *             .instanceName(String.format("terraform-example-%s", defaultInteger.result()))
 *             .instanceType("professional")
 *             .maxTps(1000)
 *             .queueCapacity(50)
 *             .supportEip(true)
 *             .maxEipTps(128)
 *             .paymentType("Subscription")
 *             .period(1)
 *             .build());
 * 
 *         var defaultVirtualHost = new VirtualHost("defaultVirtualHost", VirtualHostArgs.builder()        
 *             .instanceId(defaultInstance2.id())
 *             .virtualHostName("example-VirtualHost")
 *             .build());
 * 
 *         var defaultQueue = new Queue("defaultQueue", QueueArgs.builder()        
 *             .instanceId(defaultVirtualHost.instanceId())
 *             .queueName("example-queue")
 *             .virtualHostName(defaultVirtualHost.virtualHostName())
 *             .build());
 * 
 *         var rabbitmqTrigger = new Trigger("rabbitmqTrigger", TriggerArgs.builder()        
 *             .service(defaultService.name())
 *             .function(defaultFunction.name())
 *             .name("terraform-example-rabbitmq")
 *             .type("eventbridge")
 *             .config(Output.tuple(defaultInstance2.id(), defaultVirtualHost.virtualHostName(), defaultQueue.queueName()).applyValue(values -> {
 *                 var id = values.t1;
 *                 var virtualHostName = values.t2;
 *                 var queueName = values.t3;
 *                 return """
 *     {
 *         "triggerEnable": false,
 *         "asyncInvocationType": false,
 *         "eventRuleFilterPattern": "{}",
 *         "eventSourceConfig": {
 *             "eventSourceType": "RabbitMQ",
 *             "eventSourceParameters": {
 *                 "sourceRabbitMQParameters": {
 *                     "RegionId": "%s",
 *                     "InstanceId": "%s",
 *                     "VirtualHostName": "%s",
 *                     "QueueName": "%s"
 *                 }
 *             }
 *         }
 *     }
 * ", defaultGetRegions.applyValue(getRegionsResult -> getRegionsResult.regions()[0].id()),id,virtualHostName,queueName);
 *             }))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Module Support
 * 
 * You can use to the existing fc module
 * to create several triggers quickly.
 * 
 * ## Import
 * 
 * Function Compute trigger can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:fc/trigger:Trigger foo my-fc-service:hello-world:hello-trigger
 * ```
 * 
 */
@ResourceType(type="alicloud:fc/trigger:Trigger")
public class Trigger extends com.pulumi.resources.CustomResource {
    /**
     * The config of Function Compute trigger.It is valid when `type` is not &#34;mns_topic&#34;.See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
     * 
     */
    @Export(name="config", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> config;

    /**
     * @return The config of Function Compute trigger.It is valid when `type` is not &#34;mns_topic&#34;.See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
     * 
     */
    public Output<Optional<String>> config() {
        return Codegen.optional(this.config);
    }
    /**
     * The config of Function Compute trigger when the type is &#34;mns_topic&#34;.It is conflict with `config`.
     * 
     */
    @Export(name="configMns", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> configMns;

    /**
     * @return The config of Function Compute trigger when the type is &#34;mns_topic&#34;.It is conflict with `config`.
     * 
     */
    public Output<Optional<String>> configMns() {
        return Codegen.optional(this.configMns);
    }
    /**
     * The Function Compute function name.
     * 
     */
    @Export(name="function", refs={String.class}, tree="[0]")
    private Output<String> function;

    /**
     * @return The Function Compute function name.
     * 
     */
    public Output<String> function() {
        return this.function;
    }
    /**
     * The date this resource was last modified.
     * 
     */
    @Export(name="lastModified", refs={String.class}, tree="[0]")
    private Output<String> lastModified;

    /**
     * @return The date this resource was last modified.
     * 
     */
    public Output<String> lastModified() {
        return this.lastModified;
    }
    /**
     * The Function Compute trigger name. It is the only in one service and is conflict with &#34;name_prefix&#34;.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The Function Compute trigger name. It is the only in one service and is conflict with &#34;name_prefix&#34;.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Setting a prefix to get a only trigger name. It is conflict with &#34;name&#34;.
     * 
     */
    @Export(name="namePrefix", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> namePrefix;

    /**
     * @return Setting a prefix to get a only trigger name. It is conflict with &#34;name&#34;.
     * 
     */
    public Output<Optional<String>> namePrefix() {
        return Codegen.optional(this.namePrefix);
    }
    /**
     * RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is &#34;acs:ram::$account-id:role/$role-name&#34;. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
     * 
     */
    @Export(name="role", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> role;

    /**
     * @return RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is &#34;acs:ram::$account-id:role/$role-name&#34;. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
     * 
     */
    public Output<Optional<String>> role() {
        return Codegen.optional(this.role);
    }
    /**
     * The Function Compute service name.
     * 
     */
    @Export(name="service", refs={String.class}, tree="[0]")
    private Output<String> service;

    /**
     * @return The Function Compute service name.
     * 
     */
    public Output<String> service() {
        return this.service;
    }
    /**
     * Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
     * 
     */
    @Export(name="sourceArn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourceArn;

    /**
     * @return Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
     * 
     */
    public Output<Optional<String>> sourceArn() {
        return Codegen.optional(this.sourceArn);
    }
    /**
     * The Function Compute trigger ID.
     * 
     */
    @Export(name="triggerId", refs={String.class}, tree="[0]")
    private Output<String> triggerId;

    /**
     * @return The Function Compute trigger ID.
     * 
     */
    public Output<String> triggerId() {
        return this.triggerId;
    }
    /**
     * The Type of the trigger. Valid values: [&#34;oss&#34;, &#34;log&#34;, &#34;timer&#34;, &#34;http&#34;, &#34;mns_topic&#34;, &#34;cdn_events&#34;, &#34;eventbridge&#34;].
     * 
     * &gt; **NOTE:** Config does not support modification when type is mns_topic.
     * **NOTE:** type = cdn_events, available in 1.47.0+.
     * **NOTE:** type = eventbridge, available in 1.173.0+.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The Type of the trigger. Valid values: [&#34;oss&#34;, &#34;log&#34;, &#34;timer&#34;, &#34;http&#34;, &#34;mns_topic&#34;, &#34;cdn_events&#34;, &#34;eventbridge&#34;].
     * 
     * &gt; **NOTE:** Config does not support modification when type is mns_topic.
     * **NOTE:** type = cdn_events, available in 1.47.0+.
     * **NOTE:** type = eventbridge, available in 1.173.0+.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Trigger(String name) {
        this(name, TriggerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Trigger(String name, TriggerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Trigger(String name, TriggerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:fc/trigger:Trigger", name, args == null ? TriggerArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Trigger(String name, Output<String> id, @Nullable TriggerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:fc/trigger:Trigger", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Trigger get(String name, Output<String> id, @Nullable TriggerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Trigger(name, id, state, options);
    }
}
