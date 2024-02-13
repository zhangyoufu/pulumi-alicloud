// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.FC
{
    /// <summary>
    /// Provides an Alicloud Function Compute Trigger resource. Based on trigger, execute your code in response to events in Alibaba Cloud.
    ///  For information about Service and how to use it, see [What is Function Compute](https://www.alibabacloud.com/help/en/fc/developer-reference/api-fc-open-2021-04-06-createtrigger).
    /// 
    /// &gt; **NOTE:** The resource requires a provider field 'account_id'. See account_id.
    /// 
    /// &gt; **NOTE:** Available since v1.93.0.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// using Random = Pulumi.Random;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var defaultAccount = AliCloud.GetAccount.Invoke();
    /// 
    ///     var defaultRegions = AliCloud.GetRegions.Invoke(new()
    ///     {
    ///         Current = true,
    ///     });
    /// 
    ///     var defaultRandomInteger = new Random.RandomInteger("defaultRandomInteger", new()
    ///     {
    ///         Max = 99999,
    ///         Min = 10000,
    ///     });
    /// 
    ///     var defaultProject = new AliCloud.Log.Project("defaultProject");
    /// 
    ///     var defaultStore = new AliCloud.Log.Store("defaultStore", new()
    ///     {
    ///         Project = defaultProject.Name,
    ///     });
    /// 
    ///     var sourceStore = new AliCloud.Log.Store("sourceStore", new()
    ///     {
    ///         Project = defaultProject.Name,
    ///     });
    /// 
    ///     var defaultRole = new AliCloud.Ram.Role("defaultRole", new()
    ///     {
    ///         Document = @"  {
    ///       ""Statement"": [
    ///         {
    ///           ""Action"": ""sts:AssumeRole"",
    ///           ""Effect"": ""Allow"",
    ///           ""Principal"": {
    ///             ""Service"": [
    ///               ""fc.aliyuncs.com""
    ///             ]
    ///           }
    ///         }
    ///       ],
    ///       ""Version"": ""1""
    ///   }
    /// ",
    ///         Description = "this is a example",
    ///         Force = true,
    ///     });
    /// 
    ///     var defaultRolePolicyAttachment = new AliCloud.Ram.RolePolicyAttachment("defaultRolePolicyAttachment", new()
    ///     {
    ///         RoleName = defaultRole.Name,
    ///         PolicyName = "AliyunLogFullAccess",
    ///         PolicyType = "System",
    ///     });
    /// 
    ///     var defaultService = new AliCloud.FC.Service("defaultService", new()
    ///     {
    ///         Description = "example-value",
    ///         Role = defaultRole.Arn,
    ///         LogConfig = new AliCloud.FC.Inputs.ServiceLogConfigArgs
    ///         {
    ///             Project = defaultProject.Name,
    ///             Logstore = defaultStore.Name,
    ///             EnableInstanceMetrics = true,
    ///             EnableRequestMetrics = true,
    ///         },
    ///     });
    /// 
    ///     var defaultBucket = new AliCloud.Oss.Bucket("defaultBucket", new()
    ///     {
    ///         BucketName = defaultRandomInteger.Result.Apply(result =&gt; $"terraform-example-{result}"),
    ///     });
    /// 
    ///     // If you upload the function by OSS Bucket, you need to specify path can't upload by content.
    ///     var defaultBucketObject = new AliCloud.Oss.BucketObject("defaultBucketObject", new()
    ///     {
    ///         Bucket = defaultBucket.Id,
    ///         Key = "index.py",
    ///         Content = @"import logging 
    /// def handler(event, context): 
    /// logger = logging.getLogger() 
    /// logger.info('hello world') 
    /// return 'hello world'",
    ///     });
    /// 
    ///     var defaultFunction = new AliCloud.FC.Function("defaultFunction", new()
    ///     {
    ///         Service = defaultService.Name,
    ///         Description = "example",
    ///         OssBucket = defaultBucket.Id,
    ///         OssKey = defaultBucketObject.Key,
    ///         MemorySize = 512,
    ///         Runtime = "python3.10",
    ///         Handler = "hello.handler",
    ///     });
    /// 
    ///     var defaultTrigger = new AliCloud.FC.Trigger("defaultTrigger", new()
    ///     {
    ///         Service = defaultService.Name,
    ///         Function = defaultFunction.Name,
    ///         Role = defaultRole.Arn,
    ///         SourceArn = Output.Tuple(defaultRegions, defaultAccount, defaultProject.Name).Apply(values =&gt;
    ///         {
    ///             var defaultRegions = values.Item1;
    ///             var defaultAccount = values.Item2;
    ///             var name = values.Item3;
    ///             return $"acs:log:{defaultRegions.Apply(getRegionsResult =&gt; getRegionsResult.Regions[0]?.Id)}:{defaultAccount.Apply(getAccountResult =&gt; getAccountResult.Id)}:project/{name}";
    ///         }),
    ///         Type = "log",
    ///         Config = Output.Tuple(sourceStore.Name, defaultProject.Name, defaultStore.Name).Apply(values =&gt;
    ///         {
    ///             var sourceStoreName = values.Item1;
    ///             var defaultProjectName = values.Item2;
    ///             var defaultStoreName = values.Item3;
    ///             return @$"    {{
    ///         ""sourceConfig"": {{
    ///             ""logstore"": ""{sourceStoreName}"",
    ///             ""startTime"": null
    ///         }},
    ///         ""jobConfig"": {{
    ///             ""maxRetryTime"": 3,
    ///             ""triggerInterval"": 60
    ///         }},
    ///         ""functionParameter"": {{
    ///             ""a"": ""b"",
    ///             ""c"": ""d""
    ///         }},
    ///         ""logConfig"": {{
    ///              ""project"": ""{defaultProjectName}"",
    ///             ""logstore"": ""{defaultStoreName}""
    ///         }},
    ///         ""targetConfig"": null,
    ///         ""enable"": true
    ///     }}
    ///   
    /// ";
    ///         }),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// MNS topic trigger:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// using Random = Pulumi.Random;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var defaultAccount = AliCloud.GetAccount.Invoke();
    /// 
    ///     var defaultRegions = AliCloud.GetRegions.Invoke(new()
    ///     {
    ///         Current = true,
    ///     });
    /// 
    ///     var defaultRandomInteger = new Random.RandomInteger("defaultRandomInteger", new()
    ///     {
    ///         Max = 99999,
    ///         Min = 10000,
    ///     });
    /// 
    ///     var defaultTopic = new AliCloud.Mns.Topic("defaultTopic");
    /// 
    ///     var defaultRole = new AliCloud.Ram.Role("defaultRole", new()
    ///     {
    ///         Document = @"  {
    ///       ""Statement"": [
    ///         {
    ///           ""Action"": ""sts:AssumeRole"",
    ///           ""Effect"": ""Allow"",
    ///           ""Principal"": {
    ///             ""Service"": [
    ///               ""mns.aliyuncs.com""
    ///             ]
    ///           }
    ///         }
    ///       ],
    ///       ""Version"": ""1""
    ///   }
    /// ",
    ///         Description = "this is a example",
    ///         Force = true,
    ///     });
    /// 
    ///     var defaultRolePolicyAttachment = new AliCloud.Ram.RolePolicyAttachment("defaultRolePolicyAttachment", new()
    ///     {
    ///         RoleName = defaultRole.Name,
    ///         PolicyName = "AliyunMNSNotificationRolePolicy",
    ///         PolicyType = "System",
    ///     });
    /// 
    ///     var defaultService = new AliCloud.FC.Service("defaultService", new()
    ///     {
    ///         Description = "example-value",
    ///         InternetAccess = false,
    ///     });
    /// 
    ///     var defaultBucket = new AliCloud.Oss.Bucket("defaultBucket", new()
    ///     {
    ///         BucketName = defaultRandomInteger.Result.Apply(result =&gt; $"terraform-example-{result}"),
    ///     });
    /// 
    ///     // If you upload the function by OSS Bucket, you need to specify path can't upload by content.
    ///     var defaultBucketObject = new AliCloud.Oss.BucketObject("defaultBucketObject", new()
    ///     {
    ///         Bucket = defaultBucket.Id,
    ///         Key = "index.py",
    ///         Content = @"import logging 
    /// def handler(event, context): 
    /// logger = logging.getLogger() 
    /// logger.info('hello world') 
    /// return 'hello world'",
    ///     });
    /// 
    ///     var defaultFunction = new AliCloud.FC.Function("defaultFunction", new()
    ///     {
    ///         Service = defaultService.Name,
    ///         Description = "example",
    ///         OssBucket = defaultBucket.Id,
    ///         OssKey = defaultBucketObject.Key,
    ///         MemorySize = 512,
    ///         Runtime = "python3.10",
    ///         Handler = "hello.handler",
    ///     });
    /// 
    ///     var defaultTrigger = new AliCloud.FC.Trigger("defaultTrigger", new()
    ///     {
    ///         Service = defaultService.Name,
    ///         Function = defaultFunction.Name,
    ///         Role = defaultRole.Arn,
    ///         SourceArn = Output.Tuple(defaultRegions, defaultAccount, defaultTopic.Name).Apply(values =&gt;
    ///         {
    ///             var defaultRegions = values.Item1;
    ///             var defaultAccount = values.Item2;
    ///             var name = values.Item3;
    ///             return $"acs:mns:{defaultRegions.Apply(getRegionsResult =&gt; getRegionsResult.Regions[0]?.Id)}:{defaultAccount.Apply(getAccountResult =&gt; getAccountResult.Id)}:/topics/{name}";
    ///         }),
    ///         Type = "mns_topic",
    ///         ConfigMns = @"  {
    ///     ""filterTag"":""exampleTag"",
    ///     ""notifyContentFormat"":""STREAM"",
    ///     ""notifyStrategy"":""BACKOFF_RETRY""
    ///   }
    /// ",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// CDN events trigger:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// using Random = Pulumi.Random;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var defaultAccount = AliCloud.GetAccount.Invoke();
    /// 
    ///     var defaultRandomInteger = new Random.RandomInteger("defaultRandomInteger", new()
    ///     {
    ///         Max = 99999,
    ///         Min = 10000,
    ///     });
    /// 
    ///     var defaultDomainNew = new AliCloud.Cdn.DomainNew("defaultDomainNew", new()
    ///     {
    ///         DomainName = defaultRandomInteger.Result.Apply(result =&gt; $"example{result}.tf.com"),
    ///         CdnType = "web",
    ///         Scope = "overseas",
    ///         Sources = new[]
    ///         {
    ///             new AliCloud.Cdn.Inputs.DomainNewSourceArgs
    ///             {
    ///                 Content = "1.1.1.1",
    ///                 Type = "ipaddr",
    ///                 Priority = 20,
    ///                 Port = 80,
    ///                 Weight = 10,
    ///             },
    ///         },
    ///     });
    /// 
    ///     var defaultService = new AliCloud.FC.Service("defaultService", new()
    ///     {
    ///         Description = "example-value",
    ///         InternetAccess = false,
    ///     });
    /// 
    ///     var defaultRole = new AliCloud.Ram.Role("defaultRole", new()
    ///     {
    ///         Document = @"    {
    ///       ""Statement"": [
    ///         {
    ///           ""Action"": ""sts:AssumeRole"",
    ///           ""Effect"": ""Allow"",
    ///           ""Principal"": {
    ///             ""Service"": [
    ///               ""cdn.aliyuncs.com""
    ///             ]
    ///           }
    ///         }
    ///       ],
    ///       ""Version"": ""1""
    ///   }
    /// ",
    ///         Description = "this is a example",
    ///         Force = true,
    ///     });
    /// 
    ///     var defaultPolicy = new AliCloud.Ram.Policy("defaultPolicy", new()
    ///     {
    ///         PolicyName = defaultRandomInteger.Result.Apply(result =&gt; $"fcservicepolicy-{result}"),
    ///         PolicyDocument = Output.Tuple(defaultService.Name, defaultService.Name).Apply(values =&gt;
    ///         {
    ///             var defaultServiceName = values.Item1;
    ///             var defaultServiceName1 = values.Item2;
    ///             return @$"    {{
    ///         ""Version"": ""1"",
    ///         ""Statement"": [
    ///         {{
    ///             ""Action"": [
    ///             ""fc:InvokeFunction""
    ///             ],
    ///         ""Resource"": [
    ///             ""acs:fc:*:*:services/{defaultServiceName}/functions/*"",
    ///             ""acs:fc:*:*:services/{defaultServiceName1}.*/functions/*""
    ///         ],
    ///         ""Effect"": ""Allow""
    ///         }}
    ///         ]
    ///     }}
    /// ";
    ///         }),
    ///         Description = "this is a example",
    ///         Force = true,
    ///     });
    /// 
    ///     var defaultRolePolicyAttachment = new AliCloud.Ram.RolePolicyAttachment("defaultRolePolicyAttachment", new()
    ///     {
    ///         RoleName = defaultRole.Name,
    ///         PolicyName = defaultPolicy.Name,
    ///         PolicyType = "Custom",
    ///     });
    /// 
    ///     var defaultBucket = new AliCloud.Oss.Bucket("defaultBucket", new()
    ///     {
    ///         BucketName = defaultRandomInteger.Result.Apply(result =&gt; $"terraform-example-{result}"),
    ///     });
    /// 
    ///     // If you upload the function by OSS Bucket, you need to specify path can't upload by content.
    ///     var defaultBucketObject = new AliCloud.Oss.BucketObject("defaultBucketObject", new()
    ///     {
    ///         Bucket = defaultBucket.Id,
    ///         Key = "index.py",
    ///         Content = @"import logging 
    /// def handler(event, context): 
    /// logger = logging.getLogger() 
    /// logger.info('hello world') 
    /// return 'hello world'",
    ///     });
    /// 
    ///     var defaultFunction = new AliCloud.FC.Function("defaultFunction", new()
    ///     {
    ///         Service = defaultService.Name,
    ///         Description = "example",
    ///         OssBucket = defaultBucket.Id,
    ///         OssKey = defaultBucketObject.Key,
    ///         MemorySize = 512,
    ///         Runtime = "python3.10",
    ///         Handler = "hello.handler",
    ///     });
    /// 
    ///     var defaultTrigger = new AliCloud.FC.Trigger("defaultTrigger", new()
    ///     {
    ///         Service = defaultService.Name,
    ///         Function = defaultFunction.Name,
    ///         Role = defaultRole.Arn,
    ///         SourceArn = $"acs:cdn:*:{defaultAccount.Apply(getAccountResult =&gt; getAccountResult.Id)}",
    ///         Type = "cdn_events",
    ///         Config = defaultDomainNew.DomainName.Apply(domainName =&gt; @$"      {{""eventName"":""LogFileCreated"",
    ///      ""eventVersion"":""1.0.0"",
    ///      ""notes"":""cdn events trigger"",
    ///      ""filter"":{{
    ///         ""domain"": [""{domainName}""]
    ///         }}
    ///     }}
    /// "),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// EventBridge trigger:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// using Random = Pulumi.Random;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var defaultAccount = AliCloud.GetAccount.Invoke();
    /// 
    ///     var defaultRegions = AliCloud.GetRegions.Invoke(new()
    ///     {
    ///         Current = true,
    ///     });
    /// 
    ///     var defaultRandomInteger = new Random.RandomInteger("defaultRandomInteger", new()
    ///     {
    ///         Max = 99999,
    ///         Min = 10000,
    ///     });
    /// 
    ///     var serviceLinkedRole = new AliCloud.EventBridge.ServiceLinkedRole("serviceLinkedRole", new()
    ///     {
    ///         ProductName = "AliyunServiceRoleForEventBridgeSendToFC",
    ///     });
    /// 
    ///     var defaultService = new AliCloud.FC.Service("defaultService", new()
    ///     {
    ///         Description = "example-value",
    ///         InternetAccess = false,
    ///     });
    /// 
    ///     var defaultBucket = new AliCloud.Oss.Bucket("defaultBucket", new()
    ///     {
    ///         BucketName = defaultRandomInteger.Result.Apply(result =&gt; $"terraform-example-{result}"),
    ///     });
    /// 
    ///     // If you upload the function by OSS Bucket, you need to specify path can't upload by content.
    ///     var defaultBucketObject = new AliCloud.Oss.BucketObject("defaultBucketObject", new()
    ///     {
    ///         Bucket = defaultBucket.Id,
    ///         Key = "index.py",
    ///         Content = @"import logging 
    /// def handler(event, context): 
    /// logger = logging.getLogger() 
    /// logger.info('hello world') 
    /// return 'hello world'",
    ///     });
    /// 
    ///     var defaultFunction = new AliCloud.FC.Function("defaultFunction", new()
    ///     {
    ///         Service = defaultService.Name,
    ///         Description = "example",
    ///         OssBucket = defaultBucket.Id,
    ///         OssKey = defaultBucketObject.Key,
    ///         MemorySize = 512,
    ///         Runtime = "python3.10",
    ///         Handler = "hello.handler",
    ///     });
    /// 
    ///     var ossTrigger = new AliCloud.FC.Trigger("ossTrigger", new()
    ///     {
    ///         Service = defaultService.Name,
    ///         Function = defaultFunction.Name,
    ///         Type = "eventbridge",
    ///         Config = @"    {
    ///         ""triggerEnable"": false,
    ///         ""asyncInvocationType"": false,
    ///         ""eventRuleFilterPattern"": {
    ///           ""source"":[
    ///             ""acs.oss""
    ///             ],
    ///             ""type"":[
    ///               ""oss:BucketCreated:PutBucket""
    ///             ]
    ///         },
    ///         ""eventSourceConfig"": {
    ///             ""eventSourceType"": ""Default""
    ///         }
    ///     }
    /// ",
    ///     });
    /// 
    ///     var mnsTrigger = new AliCloud.FC.Trigger("mnsTrigger", new()
    ///     {
    ///         Service = defaultService.Name,
    ///         Function = defaultFunction.Name,
    ///         Type = "eventbridge",
    ///         Config = @"    {
    ///         ""triggerEnable"": false,
    ///         ""asyncInvocationType"": false,
    ///         ""eventRuleFilterPattern"": ""{}"",
    ///         ""eventSourceConfig"": {
    ///             ""eventSourceType"": ""MNS"",
    ///             ""eventSourceParameters"": {
    ///                 ""sourceMNSParameters"": {
    ///                     ""RegionId"": ""cn-hangzhou"",
    ///                     ""QueueName"": ""mns-queue"",
    ///                     ""IsBase64Decode"": true
    ///                 }
    ///             }
    ///         }
    ///     }
    /// ",
    ///     });
    /// 
    ///     var defaultInstance = new AliCloud.RocketMQ.Instance("defaultInstance", new()
    ///     {
    ///         InstanceName = defaultRandomInteger.Result.Apply(result =&gt; $"terraform-example-{result}"),
    ///         Remark = "terraform-example",
    ///     });
    /// 
    ///     var defaultGroup = new AliCloud.RocketMQ.Group("defaultGroup", new()
    ///     {
    ///         GroupName = "GID-example",
    ///         InstanceId = defaultInstance.Id,
    ///         Remark = "terraform-example",
    ///     });
    /// 
    ///     var defaultTopic = new AliCloud.RocketMQ.Topic("defaultTopic", new()
    ///     {
    ///         TopicName = "mytopic",
    ///         InstanceId = defaultInstance.Id,
    ///         MessageType = 0,
    ///         Remark = "terraform-example",
    ///     });
    /// 
    ///     var rocketmqTrigger = new AliCloud.FC.Trigger("rocketmqTrigger", new()
    ///     {
    ///         Service = defaultService.Name,
    ///         Function = defaultFunction.Name,
    ///         Type = "eventbridge",
    ///         Config = Output.Tuple(defaultRegions, defaultInstance.Id, defaultGroup.GroupName, defaultTopic.TopicName).Apply(values =&gt;
    ///         {
    ///             var defaultRegions = values.Item1;
    ///             var id = values.Item2;
    ///             var groupName = values.Item3;
    ///             var topicName = values.Item4;
    ///             return @$"    {{
    ///         ""triggerEnable"": false,
    ///         ""asyncInvocationType"": false,
    ///         ""eventRuleFilterPattern"": ""{{}}"",
    ///         ""eventSourceConfig"": {{
    ///             ""eventSourceType"": ""RocketMQ"",
    ///             ""eventSourceParameters"": {{
    ///                 ""sourceRocketMQParameters"": {{
    ///                     ""RegionId"": ""{defaultRegions.Apply(getRegionsResult =&gt; getRegionsResult.Regions[0]?.Id)}"",
    ///                     ""InstanceId"": ""{id}"",
    ///                     ""GroupID"": ""{groupName}"",
    ///                     ""Topic"": ""{topicName}"",
    ///                     ""Timestamp"": 1686296162,
    ///                     ""Tag"": ""example-tag"",
    ///                     ""Offset"": ""CONSUME_FROM_LAST_OFFSET""
    ///                 }}
    ///             }}
    ///         }}
    ///     }}
    /// ";
    ///         }),
    ///     });
    /// 
    ///     var defaultAmqp_instanceInstance = new AliCloud.Amqp.Instance("defaultAmqp/instanceInstance", new()
    ///     {
    ///         InstanceName = defaultRandomInteger.Result.Apply(result =&gt; $"terraform-example-{result}"),
    ///         InstanceType = "professional",
    ///         MaxTps = "1000",
    ///         QueueCapacity = "50",
    ///         SupportEip = true,
    ///         MaxEipTps = "128",
    ///         PaymentType = "Subscription",
    ///         Period = 1,
    ///     });
    /// 
    ///     var defaultVirtualHost = new AliCloud.Amqp.VirtualHost("defaultVirtualHost", new()
    ///     {
    ///         InstanceId = defaultAmqp / instanceInstance.Id,
    ///         VirtualHostName = "example-VirtualHost",
    ///     });
    /// 
    ///     var defaultQueue = new AliCloud.Amqp.Queue("defaultQueue", new()
    ///     {
    ///         InstanceId = defaultVirtualHost.InstanceId,
    ///         QueueName = "example-queue",
    ///         VirtualHostName = defaultVirtualHost.VirtualHostName,
    ///     });
    /// 
    ///     var rabbitmqTrigger = new AliCloud.FC.Trigger("rabbitmqTrigger", new()
    ///     {
    ///         Service = defaultService.Name,
    ///         Function = defaultFunction.Name,
    ///         Type = "eventbridge",
    ///         Config = Output.Tuple(defaultRegions, defaultVirtualHost.VirtualHostName, defaultQueue.QueueName).Apply(values =&gt;
    ///         {
    ///             var defaultRegions = values.Item1;
    ///             var virtualHostName = values.Item2;
    ///             var queueName = values.Item3;
    ///             return @$"    {{
    ///         ""triggerEnable"": false,
    ///         ""asyncInvocationType"": false,
    ///         ""eventRuleFilterPattern"": ""{{}}"",
    ///         ""eventSourceConfig"": {{
    ///             ""eventSourceType"": ""RabbitMQ"",
    ///             ""eventSourceParameters"": {{
    ///                 ""sourceRabbitMQParameters"": {{
    ///                     ""RegionId"": ""{defaultRegions.Apply(getRegionsResult =&gt; getRegionsResult.Regions[0]?.Id)}"",
    ///                     ""InstanceId"": ""{defaultAmqp / instanceInstance.Id}"",
    ///                     ""VirtualHostName"": ""{virtualHostName}"",
    ///                     ""QueueName"": ""{queueName}""
    ///                 }}
    ///             }}
    ///         }}
    ///     }}
    /// ";
    ///         }),
    ///     });
    /// 
    /// });
    /// ```
    /// ## Module Support
    /// 
    /// You can use to the existing fc module
    /// to create several triggers quickly.
    /// 
    /// ## Import
    /// 
    /// Function Compute trigger can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:fc/trigger:Trigger foo my-fc-service:hello-world:hello-trigger
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:fc/trigger:Trigger")]
    public partial class Trigger : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The config of Function Compute trigger.It is valid when `type` is not "mns_topic".See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
        /// </summary>
        [Output("config")]
        public Output<string?> Config { get; private set; } = null!;

        /// <summary>
        /// The config of Function Compute trigger when the type is "mns_topic".It is conflict with `config`.
        /// </summary>
        [Output("configMns")]
        public Output<string?> ConfigMns { get; private set; } = null!;

        /// <summary>
        /// The Function Compute function name.
        /// </summary>
        [Output("function")]
        public Output<string> Function { get; private set; } = null!;

        /// <summary>
        /// The date this resource was last modified.
        /// </summary>
        [Output("lastModified")]
        public Output<string> LastModified { get; private set; } = null!;

        /// <summary>
        /// The Function Compute trigger name. It is the only in one service and is conflict with "name_prefix".
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Setting a prefix to get a only trigger name. It is conflict with "name".
        /// </summary>
        [Output("namePrefix")]
        public Output<string?> NamePrefix { get; private set; } = null!;

        /// <summary>
        /// RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
        /// </summary>
        [Output("role")]
        public Output<string?> Role { get; private set; } = null!;

        /// <summary>
        /// The Function Compute service name.
        /// </summary>
        [Output("service")]
        public Output<string> Service { get; private set; } = null!;

        /// <summary>
        /// Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
        /// </summary>
        [Output("sourceArn")]
        public Output<string?> SourceArn { get; private set; } = null!;

        /// <summary>
        /// The Function Compute trigger ID.
        /// </summary>
        [Output("triggerId")]
        public Output<string> TriggerId { get; private set; } = null!;

        /// <summary>
        /// The Type of the trigger. Valid values: ["oss", "log", "timer", "http", "mns_topic", "cdn_events", "eventbridge"].
        /// 
        /// &gt; **NOTE:** Config does not support modification when type is mns_topic.
        /// &gt; **NOTE:** type = cdn_events, available in 1.47.0+.
        /// &gt; **NOTE:** type = eventbridge, available in 1.173.0+.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Trigger resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Trigger(string name, TriggerArgs args, CustomResourceOptions? options = null)
            : base("alicloud:fc/trigger:Trigger", name, args ?? new TriggerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Trigger(string name, Input<string> id, TriggerState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:fc/trigger:Trigger", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Trigger resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Trigger Get(string name, Input<string> id, TriggerState? state = null, CustomResourceOptions? options = null)
        {
            return new Trigger(name, id, state, options);
        }
    }

    public sealed class TriggerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The config of Function Compute trigger.It is valid when `type` is not "mns_topic".See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
        /// </summary>
        [Input("config")]
        public Input<string>? Config { get; set; }

        /// <summary>
        /// The config of Function Compute trigger when the type is "mns_topic".It is conflict with `config`.
        /// </summary>
        [Input("configMns")]
        public Input<string>? ConfigMns { get; set; }

        /// <summary>
        /// The Function Compute function name.
        /// </summary>
        [Input("function", required: true)]
        public Input<string> Function { get; set; } = null!;

        /// <summary>
        /// The Function Compute trigger name. It is the only in one service and is conflict with "name_prefix".
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Setting a prefix to get a only trigger name. It is conflict with "name".
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// The Function Compute service name.
        /// </summary>
        [Input("service", required: true)]
        public Input<string> Service { get; set; } = null!;

        /// <summary>
        /// Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
        /// </summary>
        [Input("sourceArn")]
        public Input<string>? SourceArn { get; set; }

        /// <summary>
        /// The Type of the trigger. Valid values: ["oss", "log", "timer", "http", "mns_topic", "cdn_events", "eventbridge"].
        /// 
        /// &gt; **NOTE:** Config does not support modification when type is mns_topic.
        /// &gt; **NOTE:** type = cdn_events, available in 1.47.0+.
        /// &gt; **NOTE:** type = eventbridge, available in 1.173.0+.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public TriggerArgs()
        {
        }
        public static new TriggerArgs Empty => new TriggerArgs();
    }

    public sealed class TriggerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The config of Function Compute trigger.It is valid when `type` is not "mns_topic".See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
        /// </summary>
        [Input("config")]
        public Input<string>? Config { get; set; }

        /// <summary>
        /// The config of Function Compute trigger when the type is "mns_topic".It is conflict with `config`.
        /// </summary>
        [Input("configMns")]
        public Input<string>? ConfigMns { get; set; }

        /// <summary>
        /// The Function Compute function name.
        /// </summary>
        [Input("function")]
        public Input<string>? Function { get; set; }

        /// <summary>
        /// The date this resource was last modified.
        /// </summary>
        [Input("lastModified")]
        public Input<string>? LastModified { get; set; }

        /// <summary>
        /// The Function Compute trigger name. It is the only in one service and is conflict with "name_prefix".
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Setting a prefix to get a only trigger name. It is conflict with "name".
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// The Function Compute service name.
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        /// <summary>
        /// Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
        /// </summary>
        [Input("sourceArn")]
        public Input<string>? SourceArn { get; set; }

        /// <summary>
        /// The Function Compute trigger ID.
        /// </summary>
        [Input("triggerId")]
        public Input<string>? TriggerId { get; set; }

        /// <summary>
        /// The Type of the trigger. Valid values: ["oss", "log", "timer", "http", "mns_topic", "cdn_events", "eventbridge"].
        /// 
        /// &gt; **NOTE:** Config does not support modification when type is mns_topic.
        /// &gt; **NOTE:** type = cdn_events, available in 1.47.0+.
        /// &gt; **NOTE:** type = eventbridge, available in 1.173.0+.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public TriggerState()
        {
        }
        public static new TriggerState Empty => new TriggerState();
    }
}
