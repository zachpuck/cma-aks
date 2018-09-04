# Protocol Documentation
<a name="top"/>

## Table of Contents

- [api.proto](#api.proto)
    - [AzureCredentials](#cmaaks.AzureCredentials)
    - [ClusterDetailItem](#cmaaks.ClusterDetailItem)
    - [ClusterItem](#cmaaks.ClusterItem)
    - [CreateClusterAKSSpec](#cmaaks.CreateClusterAKSSpec)
    - [CreateClusterAKSSpec.AKSInstanceGroup](#cmaaks.CreateClusterAKSSpec.AKSInstanceGroup)
    - [CreateClusterMsg](#cmaaks.CreateClusterMsg)
    - [CreateClusterProviderSpec](#cmaaks.CreateClusterProviderSpec)
    - [CreateClusterReply](#cmaaks.CreateClusterReply)
    - [DeleteClusterMsg](#cmaaks.DeleteClusterMsg)
    - [DeleteClusterReply](#cmaaks.DeleteClusterReply)
    - [GetClusterListMsg](#cmaaks.GetClusterListMsg)
    - [GetClusterListReply](#cmaaks.GetClusterListReply)
    - [GetClusterMsg](#cmaaks.GetClusterMsg)
    - [GetClusterReply](#cmaaks.GetClusterReply)
    - [GetVersionMsg](#cmaaks.GetVersionMsg)
    - [GetVersionReply](#cmaaks.GetVersionReply)
    - [GetVersionReply.VersionInformation](#cmaaks.GetVersionReply.VersionInformation)
  
  
  
    - [Cluster](#cmaaks.Cluster)
  

- [api.proto](#api.proto)
    - [AzureCredentials](#cmaaks.AzureCredentials)
    - [ClusterDetailItem](#cmaaks.ClusterDetailItem)
    - [ClusterItem](#cmaaks.ClusterItem)
    - [CreateClusterAKSSpec](#cmaaks.CreateClusterAKSSpec)
    - [CreateClusterAKSSpec.AKSInstanceGroup](#cmaaks.CreateClusterAKSSpec.AKSInstanceGroup)
    - [CreateClusterMsg](#cmaaks.CreateClusterMsg)
    - [CreateClusterProviderSpec](#cmaaks.CreateClusterProviderSpec)
    - [CreateClusterReply](#cmaaks.CreateClusterReply)
    - [DeleteClusterMsg](#cmaaks.DeleteClusterMsg)
    - [DeleteClusterReply](#cmaaks.DeleteClusterReply)
    - [GetClusterListMsg](#cmaaks.GetClusterListMsg)
    - [GetClusterListReply](#cmaaks.GetClusterListReply)
    - [GetClusterMsg](#cmaaks.GetClusterMsg)
    - [GetClusterReply](#cmaaks.GetClusterReply)
    - [GetVersionMsg](#cmaaks.GetVersionMsg)
    - [GetVersionReply](#cmaaks.GetVersionReply)
    - [GetVersionReply.VersionInformation](#cmaaks.GetVersionReply.VersionInformation)
  
  
  
    - [Cluster](#cmaaks.Cluster)
  

- [Scalar Value Types](#scalar-value-types)



<a name="api.proto"/>
<p align="right"><a href="#top">Top</a></p>

## api.proto



<a name="cmaaks.AzureCredentials"/>

### AzureCredentials
The credentials to use for creating the cluster


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| app_id | [string](#string) |  | The AppId for API Access |
| tenant | [string](#string) |  | The Tenant for API access |
| password | [string](#string) |  | The Password for API access |






<a name="cmaaks.ClusterDetailItem"/>

### ClusterDetailItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | ID of the cluster |
| name | [string](#string) |  | Name of the cluster |
| status | [string](#string) |  | What is the status of the cluster |
| kubeconfig | [string](#string) |  | What is the kubeconfig to connect to the cluster |






<a name="cmaaks.ClusterItem"/>

### ClusterItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | ID of the cluster |
| name | [string](#string) |  | Name of the cluster |
| status | [string](#string) |  | What is the status of the cluster |






<a name="cmaaks.CreateClusterAKSSpec"/>

### CreateClusterAKSSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| location | [string](#string) |  | The Azure Data Center |
| credentials | [AzureCredentials](#cmaaks.AzureCredentials) |  | Credentials to build the cluster |
| instance_groups | [CreateClusterAKSSpec.AKSInstanceGroup](#cmaaks.CreateClusterAKSSpec.AKSInstanceGroup) | repeated | Instance groups |






<a name="cmaaks.CreateClusterAKSSpec.AKSInstanceGroup"/>

### CreateClusterAKSSpec.AKSInstanceGroup
Instance groups define a type and number of instances


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the group |
| type | [string](#string) |  | Instance type (Standard_D2_v2, etc.) |
| min_quantity | [int32](#int32) |  | Minimum number of instances (defaults to zero) |
| max_quantity | [int32](#int32) |  | Maximum number of instances (defaults to zero) |






<a name="cmaaks.CreateClusterMsg"/>

### CreateClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the cluster to be provisioned |
| provider | [CreateClusterProviderSpec](#cmaaks.CreateClusterProviderSpec) |  | The provider specification |






<a name="cmaaks.CreateClusterProviderSpec"/>

### CreateClusterProviderSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the provider - like aks |
| k8s_version | [string](#string) |  | The version of Kubernetes |
| azure | [CreateClusterAKSSpec](#cmaaks.CreateClusterAKSSpec) |  | The AKS specification |
| high_availability | [bool](#bool) |  | Whether or not the cluster is HA |
| network_fabric | [string](#string) |  | The fabric to be used |






<a name="cmaaks.CreateClusterReply"/>

### CreateClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Whether or not the cluster was provisioned by this request |
| cluster | [ClusterItem](#cmaaks.ClusterItem) |  | The details of the cluster request response |






<a name="cmaaks.DeleteClusterMsg"/>

### DeleteClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the cluster&#39;s name to destroy |
| credentials | [AzureCredentials](#cmaaks.AzureCredentials) |  | Credentials to delete the cluster |






<a name="cmaaks.DeleteClusterReply"/>

### DeleteClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Could the cluster be destroyed |
| status | [string](#string) |  | Status of the request |






<a name="cmaaks.GetClusterListMsg"/>

### GetClusterListMsg







<a name="cmaaks.GetClusterListReply"/>

### GetClusterListReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Is the cluster in the system |
| clusters | [ClusterItem](#cmaaks.ClusterItem) | repeated | List of clusters |






<a name="cmaaks.GetClusterMsg"/>

### GetClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the cluster to be looked up |
| credentials | [AzureCredentials](#cmaaks.AzureCredentials) |  | Credentials to query for the cluster |






<a name="cmaaks.GetClusterReply"/>

### GetClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Is the cluster in the system |
| cluster | [ClusterDetailItem](#cmaaks.ClusterDetailItem) |  |  |






<a name="cmaaks.GetVersionMsg"/>

### GetVersionMsg
Get version of API Server






<a name="cmaaks.GetVersionReply"/>

### GetVersionReply
Reply for version request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | If operation was OK |
| version_information | [GetVersionReply.VersionInformation](#cmaaks.GetVersionReply.VersionInformation) |  | Version Information |






<a name="cmaaks.GetVersionReply.VersionInformation"/>

### GetVersionReply.VersionInformation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| git_version | [string](#string) |  | The tag on the git repository |
| git_commit | [string](#string) |  | The hash of the git commit |
| git_tree_state | [string](#string) |  | Whether or not the tree was clean when built |
| build_date | [string](#string) |  | Date of build |
| go_version | [string](#string) |  | Version of go used to compile |
| compiler | [string](#string) |  | Compiler used |
| platform | [string](#string) |  | Platform it was compiled for / running on |





 

 

 


<a name="cmaaks.Cluster"/>

### Cluster


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateCluster | [CreateClusterMsg](#cmaaks.CreateClusterMsg) | [CreateClusterReply](#cmaaks.CreateClusterMsg) | Will provision a cluster |
| GetCluster | [GetClusterMsg](#cmaaks.GetClusterMsg) | [GetClusterReply](#cmaaks.GetClusterMsg) | Will retrieve the status of a cluster and its kubeconfig for connectivity |
| DeleteCluster | [DeleteClusterMsg](#cmaaks.DeleteClusterMsg) | [DeleteClusterReply](#cmaaks.DeleteClusterMsg) | Will delete a cluster |
| GetClusterList | [GetClusterListMsg](#cmaaks.GetClusterListMsg) | [GetClusterListReply](#cmaaks.GetClusterListMsg) | Will retrieve a list of clusters |
| GetVersionInformation | [GetVersionMsg](#cmaaks.GetVersionMsg) | [GetVersionReply](#cmaaks.GetVersionMsg) | Will return version information about api server |

 



<a name="api.proto"/>
<p align="right"><a href="#top">Top</a></p>

## api.proto



<a name="cmaaks.AzureCredentials"/>

### AzureCredentials
The credentials to use for creating the cluster


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| app_id | [string](#string) |  | The AppId for API Access |
| tenant | [string](#string) |  | The Tenant for API access |
| password | [string](#string) |  | The Password for API access |






<a name="cmaaks.ClusterDetailItem"/>

### ClusterDetailItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | ID of the cluster |
| name | [string](#string) |  | Name of the cluster |
| status | [string](#string) |  | What is the status of the cluster |
| kubeconfig | [string](#string) |  | What is the kubeconfig to connect to the cluster |






<a name="cmaaks.ClusterItem"/>

### ClusterItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | ID of the cluster |
| name | [string](#string) |  | Name of the cluster |
| status | [string](#string) |  | What is the status of the cluster |






<a name="cmaaks.CreateClusterAKSSpec"/>

### CreateClusterAKSSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| location | [string](#string) |  | The Azure Data Center |
| credentials | [AzureCredentials](#cmaaks.AzureCredentials) |  | Credentials to build the cluster |
| instance_groups | [CreateClusterAKSSpec.AKSInstanceGroup](#cmaaks.CreateClusterAKSSpec.AKSInstanceGroup) | repeated | Instance groups |






<a name="cmaaks.CreateClusterAKSSpec.AKSInstanceGroup"/>

### CreateClusterAKSSpec.AKSInstanceGroup
Instance groups define a type and number of instances


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the group |
| type | [string](#string) |  | Instance type (Standard_D2_v2, etc.) |
| min_quantity | [int32](#int32) |  | Minimum number of instances (defaults to zero) |
| max_quantity | [int32](#int32) |  | Maximum number of instances (defaults to zero) |






<a name="cmaaks.CreateClusterMsg"/>

### CreateClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the cluster to be provisioned |
| provider | [CreateClusterProviderSpec](#cmaaks.CreateClusterProviderSpec) |  | The provider specification |






<a name="cmaaks.CreateClusterProviderSpec"/>

### CreateClusterProviderSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the provider - like aks |
| k8s_version | [string](#string) |  | The version of Kubernetes |
| azure | [CreateClusterAKSSpec](#cmaaks.CreateClusterAKSSpec) |  | The AKS specification |
| high_availability | [bool](#bool) |  | Whether or not the cluster is HA |
| network_fabric | [string](#string) |  | The fabric to be used |






<a name="cmaaks.CreateClusterReply"/>

### CreateClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Whether or not the cluster was provisioned by this request |
| cluster | [ClusterItem](#cmaaks.ClusterItem) |  | The details of the cluster request response |






<a name="cmaaks.DeleteClusterMsg"/>

### DeleteClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the cluster&#39;s name to destroy |
| credentials | [AzureCredentials](#cmaaks.AzureCredentials) |  | Credentials to delete the cluster |






<a name="cmaaks.DeleteClusterReply"/>

### DeleteClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Could the cluster be destroyed |
| status | [string](#string) |  | Status of the request |






<a name="cmaaks.GetClusterListMsg"/>

### GetClusterListMsg







<a name="cmaaks.GetClusterListReply"/>

### GetClusterListReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Is the cluster in the system |
| clusters | [ClusterItem](#cmaaks.ClusterItem) | repeated | List of clusters |






<a name="cmaaks.GetClusterMsg"/>

### GetClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the cluster to be looked up |
| credentials | [AzureCredentials](#cmaaks.AzureCredentials) |  | Credentials to query for the cluster |






<a name="cmaaks.GetClusterReply"/>

### GetClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Is the cluster in the system |
| cluster | [ClusterDetailItem](#cmaaks.ClusterDetailItem) |  |  |






<a name="cmaaks.GetVersionMsg"/>

### GetVersionMsg
Get version of API Server






<a name="cmaaks.GetVersionReply"/>

### GetVersionReply
Reply for version request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | If operation was OK |
| version_information | [GetVersionReply.VersionInformation](#cmaaks.GetVersionReply.VersionInformation) |  | Version Information |






<a name="cmaaks.GetVersionReply.VersionInformation"/>

### GetVersionReply.VersionInformation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| git_version | [string](#string) |  | The tag on the git repository |
| git_commit | [string](#string) |  | The hash of the git commit |
| git_tree_state | [string](#string) |  | Whether or not the tree was clean when built |
| build_date | [string](#string) |  | Date of build |
| go_version | [string](#string) |  | Version of go used to compile |
| compiler | [string](#string) |  | Compiler used |
| platform | [string](#string) |  | Platform it was compiled for / running on |





 

 

 


<a name="cmaaks.Cluster"/>

### Cluster


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateCluster | [CreateClusterMsg](#cmaaks.CreateClusterMsg) | [CreateClusterReply](#cmaaks.CreateClusterMsg) | Will provision a cluster |
| GetCluster | [GetClusterMsg](#cmaaks.GetClusterMsg) | [GetClusterReply](#cmaaks.GetClusterMsg) | Will retrieve the status of a cluster and its kubeconfig for connectivity |
| DeleteCluster | [DeleteClusterMsg](#cmaaks.DeleteClusterMsg) | [DeleteClusterReply](#cmaaks.DeleteClusterMsg) | Will delete a cluster |
| GetClusterList | [GetClusterListMsg](#cmaaks.GetClusterListMsg) | [GetClusterListReply](#cmaaks.GetClusterListMsg) | Will retrieve a list of clusters |
| GetVersionInformation | [GetVersionMsg](#cmaaks.GetVersionMsg) | [GetVersionReply](#cmaaks.GetVersionMsg) | Will return version information about api server |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |
