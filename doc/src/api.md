# API

## Objects

WebGPU objects are referenced through blind handles, which are homogenous to pointers. The following objects are defined and detailed below:


 - [`WGPUAdapter`](#WGPUAdapter)
 - [`WGPUBindGroup`](#WGPUBindGroup)
 - [`WGPUBindGroupLayout`](#WGPUBindGroupLayout)
 - [`WGPUBuffer`](#WGPUBuffer)
 - [`WGPUCommandBuffer`](#WGPUCommandBuffer)
 - [`WGPUCommandEncoder`](#WGPUCommandEncoder)
 - [`WGPUComputePassEncoder`](#WGPUComputePassEncoder)
 - [`WGPUComputePipeline`](#WGPUComputePipeline)
 - [`WGPUDevice`](#WGPUDevice)
 - [`WGPUInstance`](#WGPUInstance)
 - [`WGPUPipelineLayout`](#WGPUPipelineLayout)
 - [`WGPUQuerySet`](#WGPUQuerySet)
 - [`WGPUQueue`](#WGPUQueue)
 - [`WGPURenderBundle`](#WGPURenderBundle)
 - [`WGPURenderBundleEncoder`](#WGPURenderBundleEncoder)
 - [`WGPURenderPassEncoder`](#WGPURenderPassEncoder)
 - [`WGPURenderPipeline`](#WGPURenderPipeline)
 - [`WGPUSampler`](#WGPUSampler)
 - [`WGPUShaderModule`](#WGPUShaderModule)
 - [`WGPUSurface`](#WGPUSurface)
 - [`WGPUTexture`](#WGPUTexture)
 - [`WGPUTextureView`](#WGPUTextureView)



### Adapter { #WGPUAdapter }


This is a test.




#### `wgpuAdapterEnumerateFeatures` { #wgpuAdapterEnumerateFeatures }

**Prototype:** `size_t wgpuAdapterEnumerateFeatures(WGPUAdapter adapter, WGPUFeatureName * features)`


TODO


**Arguments:**

 - TODO



**Returns:** size_t 
TODO





#### `wgpuAdapterGetLimits` { #wgpuAdapterGetLimits }

**Prototype:** `WGPUBool wgpuAdapterGetLimits(WGPUAdapter adapter, WGPUSupportedLimits * limits)`


Get the limits that the adapter supports. This is used to drive the
choice of required limits in {device_descriptor.required_limits}.


**Arguments:**

 - TODO



**Returns:** WGPUBool 
True iff limits were successfully retrieved.





#### `wgpuAdapterGetProperties` { #wgpuAdapterGetProperties }

**Prototype:** `void wgpuAdapterGetProperties(WGPUAdapter adapter, WGPUAdapterProperties * properties)`


TODO


**Arguments:**

 - TODO




#### `wgpuAdapterHasFeature` { #wgpuAdapterHasFeature }

**Prototype:** `WGPUBool wgpuAdapterHasFeature(WGPUAdapter adapter, WGPUFeatureName feature)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPUBool 
TODO





#### `wgpuAdapterRequestAdapterInfo` { #wgpuAdapterRequestAdapterInfo }

**Prototype:** `void wgpuAdapterRequestAdapterInfo(WGPUAdapter adapter, WGPUAdapterRequestAdapterInfoCallback callback, WGPU_NULLABLE void * userdata)`


TODO


**Arguments:**

 - TODO




#### `wgpuAdapterRequestDevice` { #wgpuAdapterRequestDevice }

**Prototype:** `void wgpuAdapterRequestDevice(WGPUAdapter adapter, WGPU_NULLABLE WGPUDeviceDescriptor const * descriptor, WGPUAdapterRequestDeviceCallback callback, WGPU_NULLABLE void * userdata)`


TODO


**Arguments:**

 - TODO




 - `void wgpuAdapterReference(WGPUAdapter adapter)`
 - `void wgpuAdapterRelease(WGPUAdapter adapter)`




### BindGroup { #WGPUBindGroup }


TODO




#### `wgpuBindGroupSetLabel` { #wgpuBindGroupSetLabel }

**Prototype:** `void wgpuBindGroupSetLabel(WGPUBindGroup bindGroup, char const * label)`


TODO


**Arguments:**

 - TODO




 - `void wgpuBindGroupReference(WGPUBindGroup bindGroup)`
 - `void wgpuBindGroupRelease(WGPUBindGroup bindGroup)`




### BindGroupLayout { #WGPUBindGroupLayout }


TODO




#### `wgpuBindGroupLayoutSetLabel` { #wgpuBindGroupLayoutSetLabel }

**Prototype:** `void wgpuBindGroupLayoutSetLabel(WGPUBindGroupLayout bindGroupLayout, char const * label)`


TODO


**Arguments:**

 - TODO




 - `void wgpuBindGroupLayoutReference(WGPUBindGroupLayout bindGroupLayout)`
 - `void wgpuBindGroupLayoutRelease(WGPUBindGroupLayout bindGroupLayout)`




### Buffer { #WGPUBuffer }


TODO




#### `wgpuBufferDestroy` { #wgpuBufferDestroy }

**Prototype:** `void wgpuBufferDestroy(WGPUBuffer buffer)`


TODO



**Arguments:**

 - TODO




#### `wgpuBufferGetConstMappedRange` { #wgpuBufferGetConstMappedRange }

**Prototype:** `void const * wgpuBufferGetConstMappedRange(WGPUBuffer buffer, size_t offset, size_t size)`


TODO


**Arguments:**

 - TODO



**Returns:** void const * 
TODO





#### `wgpuBufferGetMapState` { #wgpuBufferGetMapState }

**Prototype:** `WGPUBufferMapState wgpuBufferGetMapState(WGPUBuffer buffer)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPUBufferMapState 
TODO





#### `wgpuBufferGetMappedRange` { #wgpuBufferGetMappedRange }

**Prototype:** `void * wgpuBufferGetMappedRange(WGPUBuffer buffer, size_t offset, size_t size)`


TODO


**Arguments:**

 - TODO



**Returns:** void * 
TODO





#### `wgpuBufferGetSize` { #wgpuBufferGetSize }

**Prototype:** `uint64_t wgpuBufferGetSize(WGPUBuffer buffer)`


TODO


**Arguments:**

 - TODO



**Returns:** uint64_t 
TODO





#### `wgpuBufferGetUsage` { #wgpuBufferGetUsage }

**Prototype:** `WGPUBufferUsageFlags wgpuBufferGetUsage(WGPUBuffer buffer)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPUBufferUsageFlags 
TODO





#### `wgpuBufferMapAsync` { #wgpuBufferMapAsync }

**Prototype:** `void wgpuBufferMapAsync(WGPUBuffer buffer, WGPUMapModeFlags mode, size_t offset, size_t size, WGPUBufferMapAsyncCallback callback, WGPU_NULLABLE void * userdata)`


TODO


**Arguments:**

 - TODO




#### `wgpuBufferSetLabel` { #wgpuBufferSetLabel }

**Prototype:** `void wgpuBufferSetLabel(WGPUBuffer buffer, char const * label)`


TODO


**Arguments:**

 - TODO




#### `wgpuBufferUnmap` { #wgpuBufferUnmap }

**Prototype:** `void wgpuBufferUnmap(WGPUBuffer buffer)`


TODO


**Arguments:**

 - TODO




 - `void wgpuBufferReference(WGPUBuffer buffer)`
 - `void wgpuBufferRelease(WGPUBuffer buffer)`




### CommandBuffer { #WGPUCommandBuffer }


TODO




#### `wgpuCommandBufferSetLabel` { #wgpuCommandBufferSetLabel }

**Prototype:** `void wgpuCommandBufferSetLabel(WGPUCommandBuffer commandBuffer, char const * label)`


TODO


**Arguments:**

 - TODO




 - `void wgpuCommandBufferReference(WGPUCommandBuffer commandBuffer)`
 - `void wgpuCommandBufferRelease(WGPUCommandBuffer commandBuffer)`




### CommandEncoder { #WGPUCommandEncoder }


TODO




#### `wgpuCommandEncoderBeginComputePass` { #wgpuCommandEncoderBeginComputePass }

**Prototype:** `WGPUComputePassEncoder wgpuCommandEncoderBeginComputePass(WGPUCommandEncoder commandEncoder, WGPU_NULLABLE WGPUComputePassDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPUComputePassEncoder 
TODO





#### `wgpuCommandEncoderBeginRenderPass` { #wgpuCommandEncoderBeginRenderPass }

**Prototype:** `WGPURenderPassEncoder wgpuCommandEncoderBeginRenderPass(WGPUCommandEncoder commandEncoder, WGPURenderPassDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPURenderPassEncoder 
TODO





#### `wgpuCommandEncoderClearBuffer` { #wgpuCommandEncoderClearBuffer }

**Prototype:** `void wgpuCommandEncoderClearBuffer(WGPUCommandEncoder commandEncoder, WGPUBuffer buffer, uint64_t offset, uint64_t size)`


TODO


**Arguments:**

 - TODO




#### `wgpuCommandEncoderCopyBufferToBuffer` { #wgpuCommandEncoderCopyBufferToBuffer }

**Prototype:** `void wgpuCommandEncoderCopyBufferToBuffer(WGPUCommandEncoder commandEncoder, WGPUBuffer source, uint64_t sourceOffset, WGPUBuffer destination, uint64_t destinationOffset, uint64_t size)`


TODO


**Arguments:**

 - TODO




#### `wgpuCommandEncoderCopyBufferToTexture` { #wgpuCommandEncoderCopyBufferToTexture }

**Prototype:** `void wgpuCommandEncoderCopyBufferToTexture(WGPUCommandEncoder commandEncoder, WGPUImageCopyBuffer const * source, WGPUImageCopyTexture const * destination, WGPUExtent3D const * copySize)`


TODO


**Arguments:**

 - TODO




#### `wgpuCommandEncoderCopyTextureToBuffer` { #wgpuCommandEncoderCopyTextureToBuffer }

**Prototype:** `void wgpuCommandEncoderCopyTextureToBuffer(WGPUCommandEncoder commandEncoder, WGPUImageCopyTexture const * source, WGPUImageCopyBuffer const * destination, WGPUExtent3D const * copySize)`


TODO


**Arguments:**

 - TODO




#### `wgpuCommandEncoderCopyTextureToTexture` { #wgpuCommandEncoderCopyTextureToTexture }

**Prototype:** `void wgpuCommandEncoderCopyTextureToTexture(WGPUCommandEncoder commandEncoder, WGPUImageCopyTexture const * source, WGPUImageCopyTexture const * destination, WGPUExtent3D const * copySize)`


TODO


**Arguments:**

 - TODO




#### `wgpuCommandEncoderFinish` { #wgpuCommandEncoderFinish }

**Prototype:** `WGPUCommandBuffer wgpuCommandEncoderFinish(WGPUCommandEncoder commandEncoder, WGPU_NULLABLE WGPUCommandBufferDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPUCommandBuffer 
TODO





#### `wgpuCommandEncoderInsertDebugMarker` { #wgpuCommandEncoderInsertDebugMarker }

**Prototype:** `void wgpuCommandEncoderInsertDebugMarker(WGPUCommandEncoder commandEncoder, char const * markerLabel)`


TODO


**Arguments:**

 - TODO




#### `wgpuCommandEncoderPopDebugGroup` { #wgpuCommandEncoderPopDebugGroup }

**Prototype:** `void wgpuCommandEncoderPopDebugGroup(WGPUCommandEncoder commandEncoder)`


TODO


**Arguments:**

 - TODO




#### `wgpuCommandEncoderPushDebugGroup` { #wgpuCommandEncoderPushDebugGroup }

**Prototype:** `void wgpuCommandEncoderPushDebugGroup(WGPUCommandEncoder commandEncoder, char const * groupLabel)`


TODO


**Arguments:**

 - TODO




#### `wgpuCommandEncoderResolveQuerySet` { #wgpuCommandEncoderResolveQuerySet }

**Prototype:** `void wgpuCommandEncoderResolveQuerySet(WGPUCommandEncoder commandEncoder, WGPUQuerySet querySet, uint32_t firstQuery, uint32_t queryCount, WGPUBuffer destination, uint64_t destinationOffset)`


TODO


**Arguments:**

 - TODO




#### `wgpuCommandEncoderSetLabel` { #wgpuCommandEncoderSetLabel }

**Prototype:** `void wgpuCommandEncoderSetLabel(WGPUCommandEncoder commandEncoder, char const * label)`


TODO


**Arguments:**

 - TODO




#### `wgpuCommandEncoderWriteTimestamp` { #wgpuCommandEncoderWriteTimestamp }

**Prototype:** `void wgpuCommandEncoderWriteTimestamp(WGPUCommandEncoder commandEncoder, WGPUQuerySet querySet, uint32_t queryIndex)`


TODO


**Arguments:**

 - TODO




 - `void wgpuCommandEncoderReference(WGPUCommandEncoder commandEncoder)`
 - `void wgpuCommandEncoderRelease(WGPUCommandEncoder commandEncoder)`




### ComputePassEncoder { #WGPUComputePassEncoder }


TODO




#### `wgpuComputePassEncoderDispatchWorkgroups` { #wgpuComputePassEncoderDispatchWorkgroups }

**Prototype:** `void wgpuComputePassEncoderDispatchWorkgroups(WGPUComputePassEncoder computePassEncoder, uint32_t workgroupCountX, uint32_t workgroupCountY, uint32_t workgroupCountZ)`


TODO


**Arguments:**

 - TODO




#### `wgpuComputePassEncoderDispatchWorkgroupsIndirect` { #wgpuComputePassEncoderDispatchWorkgroupsIndirect }

**Prototype:** `void wgpuComputePassEncoderDispatchWorkgroupsIndirect(WGPUComputePassEncoder computePassEncoder, WGPUBuffer indirectBuffer, uint64_t indirectOffset)`


TODO


**Arguments:**

 - TODO




#### `wgpuComputePassEncoderEnd` { #wgpuComputePassEncoderEnd }

**Prototype:** `void wgpuComputePassEncoderEnd(WGPUComputePassEncoder computePassEncoder)`


TODO


**Arguments:**

 - TODO




#### `wgpuComputePassEncoderInsertDebugMarker` { #wgpuComputePassEncoderInsertDebugMarker }

**Prototype:** `void wgpuComputePassEncoderInsertDebugMarker(WGPUComputePassEncoder computePassEncoder, char const * markerLabel)`


TODO


**Arguments:**

 - TODO




#### `wgpuComputePassEncoderPopDebugGroup` { #wgpuComputePassEncoderPopDebugGroup }

**Prototype:** `void wgpuComputePassEncoderPopDebugGroup(WGPUComputePassEncoder computePassEncoder)`


TODO


**Arguments:**

 - TODO




#### `wgpuComputePassEncoderPushDebugGroup` { #wgpuComputePassEncoderPushDebugGroup }

**Prototype:** `void wgpuComputePassEncoderPushDebugGroup(WGPUComputePassEncoder computePassEncoder, char const * groupLabel)`


TODO


**Arguments:**

 - TODO




#### `wgpuComputePassEncoderSetBindGroup` { #wgpuComputePassEncoderSetBindGroup }

**Prototype:** `void wgpuComputePassEncoderSetBindGroup(WGPUComputePassEncoder computePassEncoder, uint32_t groupIndex, WGPU_NULLABLE WGPUBindGroup group, size_t dynamicOffsetCount, uint32_t const * dynamicOffsets)`


TODO


**Arguments:**

 - TODO




#### `wgpuComputePassEncoderSetLabel` { #wgpuComputePassEncoderSetLabel }

**Prototype:** `void wgpuComputePassEncoderSetLabel(WGPUComputePassEncoder computePassEncoder, char const * label)`


TODO


**Arguments:**

 - TODO




#### `wgpuComputePassEncoderSetPipeline` { #wgpuComputePassEncoderSetPipeline }

**Prototype:** `void wgpuComputePassEncoderSetPipeline(WGPUComputePassEncoder computePassEncoder, WGPUComputePipeline pipeline)`


TODO


**Arguments:**

 - TODO




 - `void wgpuComputePassEncoderReference(WGPUComputePassEncoder computePassEncoder)`
 - `void wgpuComputePassEncoderRelease(WGPUComputePassEncoder computePassEncoder)`




### ComputePipeline { #WGPUComputePipeline }


TODO




#### `wgpuComputePipelineGetBindGroupLayout` { #wgpuComputePipelineGetBindGroupLayout }

**Prototype:** `WGPUBindGroupLayout wgpuComputePipelineGetBindGroupLayout(WGPUComputePipeline computePipeline, uint32_t groupIndex)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPUBindGroupLayout 
TODO





#### `wgpuComputePipelineSetLabel` { #wgpuComputePipelineSetLabel }

**Prototype:** `void wgpuComputePipelineSetLabel(WGPUComputePipeline computePipeline, char const * label)`


TODO


**Arguments:**

 - TODO




 - `void wgpuComputePipelineReference(WGPUComputePipeline computePipeline)`
 - `void wgpuComputePipelineRelease(WGPUComputePipeline computePipeline)`




### Device { #WGPUDevice }


TODO




#### `wgpuDeviceCreateBindGroup` { #wgpuDeviceCreateBindGroup }

**Prototype:** `WGPUBindGroup wgpuDeviceCreateBindGroup(WGPUDevice device, WGPUBindGroupDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPUBindGroup 
TODO





#### `wgpuDeviceCreateBindGroupLayout` { #wgpuDeviceCreateBindGroupLayout }

**Prototype:** `WGPUBindGroupLayout wgpuDeviceCreateBindGroupLayout(WGPUDevice device, WGPUBindGroupLayoutDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPUBindGroupLayout 
TODO





#### `wgpuDeviceCreateBuffer` { #wgpuDeviceCreateBuffer }

**Prototype:** `WGPUBuffer wgpuDeviceCreateBuffer(WGPUDevice device, WGPUBufferDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPUBuffer 
TODO





#### `wgpuDeviceCreateCommandEncoder` { #wgpuDeviceCreateCommandEncoder }

**Prototype:** `WGPUCommandEncoder wgpuDeviceCreateCommandEncoder(WGPUDevice device, WGPU_NULLABLE WGPUCommandEncoderDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPUCommandEncoder 
TODO





#### `wgpuDeviceCreateComputePipeline` { #wgpuDeviceCreateComputePipeline }

**Prototype:** `WGPUComputePipeline wgpuDeviceCreateComputePipeline(WGPUDevice device, WGPUComputePipelineDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPUComputePipeline 
TODO





#### `wgpuDeviceCreateComputePipelineAsync` { #wgpuDeviceCreateComputePipelineAsync }

**Prototype:** `void wgpuDeviceCreateComputePipelineAsync(WGPUDevice device, WGPUComputePipelineDescriptor const * descriptor, WGPUDeviceCreateComputePipelineAsyncCallback callback, WGPU_NULLABLE void * userdata)`


TODO


**Arguments:**

 - TODO




#### `wgpuDeviceCreatePipelineLayout` { #wgpuDeviceCreatePipelineLayout }

**Prototype:** `WGPUPipelineLayout wgpuDeviceCreatePipelineLayout(WGPUDevice device, WGPUPipelineLayoutDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPUPipelineLayout 
TODO





#### `wgpuDeviceCreateQuerySet` { #wgpuDeviceCreateQuerySet }

**Prototype:** `WGPUQuerySet wgpuDeviceCreateQuerySet(WGPUDevice device, WGPUQuerySetDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPUQuerySet 
TODO





#### `wgpuDeviceCreateRenderBundleEncoder` { #wgpuDeviceCreateRenderBundleEncoder }

**Prototype:** `WGPURenderBundleEncoder wgpuDeviceCreateRenderBundleEncoder(WGPUDevice device, WGPURenderBundleEncoderDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPURenderBundleEncoder 
TODO





#### `wgpuDeviceCreateRenderPipeline` { #wgpuDeviceCreateRenderPipeline }

**Prototype:** `WGPURenderPipeline wgpuDeviceCreateRenderPipeline(WGPUDevice device, WGPURenderPipelineDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPURenderPipeline 
TODO





#### `wgpuDeviceCreateRenderPipelineAsync` { #wgpuDeviceCreateRenderPipelineAsync }

**Prototype:** `void wgpuDeviceCreateRenderPipelineAsync(WGPUDevice device, WGPURenderPipelineDescriptor const * descriptor, WGPUDeviceCreateRenderPipelineAsyncCallback callback, WGPU_NULLABLE void * userdata)`


TODO


**Arguments:**

 - TODO




#### `wgpuDeviceCreateSampler` { #wgpuDeviceCreateSampler }

**Prototype:** `WGPUSampler wgpuDeviceCreateSampler(WGPUDevice device, WGPU_NULLABLE WGPUSamplerDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPUSampler 
TODO





#### `wgpuDeviceCreateShaderModule` { #wgpuDeviceCreateShaderModule }

**Prototype:** `WGPUShaderModule wgpuDeviceCreateShaderModule(WGPUDevice device, WGPUShaderModuleDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPUShaderModule 
TODO





#### `wgpuDeviceCreateTexture` { #wgpuDeviceCreateTexture }

**Prototype:** `WGPUTexture wgpuDeviceCreateTexture(WGPUDevice device, WGPUTextureDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPUTexture 
TODO





#### `wgpuDeviceDestroy` { #wgpuDeviceDestroy }

**Prototype:** `void wgpuDeviceDestroy(WGPUDevice device)`


TODO


**Arguments:**

 - TODO




#### `wgpuDeviceEnumerateFeatures` { #wgpuDeviceEnumerateFeatures }

**Prototype:** `size_t wgpuDeviceEnumerateFeatures(WGPUDevice device, WGPUFeatureName * features)`


TODO


**Arguments:**

 - TODO



**Returns:** size_t 
TODO





#### `wgpuDeviceGetLimits` { #wgpuDeviceGetLimits }

**Prototype:** `WGPUBool wgpuDeviceGetLimits(WGPUDevice device, WGPUSupportedLimits * limits)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPUBool 
TODO





#### `wgpuDeviceGetQueue` { #wgpuDeviceGetQueue }

**Prototype:** `WGPUQueue wgpuDeviceGetQueue(WGPUDevice device)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPUQueue 
TODO





#### `wgpuDeviceHasFeature` { #wgpuDeviceHasFeature }

**Prototype:** `WGPUBool wgpuDeviceHasFeature(WGPUDevice device, WGPUFeatureName feature)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPUBool 
TODO





#### `wgpuDevicePopErrorScope` { #wgpuDevicePopErrorScope }

**Prototype:** `void wgpuDevicePopErrorScope(WGPUDevice device, WGPUErrorCallback callback, void * userdata)`


TODO


**Arguments:**

 - TODO




#### `wgpuDevicePushErrorScope` { #wgpuDevicePushErrorScope }

**Prototype:** `void wgpuDevicePushErrorScope(WGPUDevice device, WGPUErrorFilter filter)`


TODO


**Arguments:**

 - TODO




#### `wgpuDeviceSetLabel` { #wgpuDeviceSetLabel }

**Prototype:** `void wgpuDeviceSetLabel(WGPUDevice device, char const * label)`


TODO


**Arguments:**

 - TODO




#### `wgpuDeviceSetUncapturedErrorCallback` { #wgpuDeviceSetUncapturedErrorCallback }

**Prototype:** `void wgpuDeviceSetUncapturedErrorCallback(WGPUDevice device, WGPUErrorCallback callback, void * userdata)`


TODO


**Arguments:**

 - TODO




 - `void wgpuDeviceReference(WGPUDevice device)`
 - `void wgpuDeviceRelease(WGPUDevice device)`




### Instance { #WGPUInstance }


TODO




#### `wgpuInstanceCreateSurface` { #wgpuInstanceCreateSurface }

**Prototype:** `WGPUSurface wgpuInstanceCreateSurface(WGPUInstance instance, WGPUSurfaceDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPUSurface 
TODO





#### `wgpuInstanceHasWGSLLanguageFeature` { #wgpuInstanceHasWGSLLanguageFeature }

**Prototype:** `WGPUBool wgpuInstanceHasWGSLLanguageFeature(WGPUInstance instance, WGPUWGSLFeatureName feature)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPUBool 
TODO





#### `wgpuInstanceProcessEvents` { #wgpuInstanceProcessEvents }

**Prototype:** `void wgpuInstanceProcessEvents(WGPUInstance instance)`


TODO


**Arguments:**

 - TODO




#### `wgpuInstanceRequestAdapter` { #wgpuInstanceRequestAdapter }

**Prototype:** `void wgpuInstanceRequestAdapter(WGPUInstance instance, WGPU_NULLABLE WGPURequestAdapterOptions const * options, WGPUInstanceRequestAdapterCallback callback, WGPU_NULLABLE void * userdata)`


TODO


**Arguments:**

 - TODO




 - `void wgpuInstanceReference(WGPUInstance instance)`
 - `void wgpuInstanceRelease(WGPUInstance instance)`




### PipelineLayout { #WGPUPipelineLayout }


TODO




#### `wgpuPipelineLayoutSetLabel` { #wgpuPipelineLayoutSetLabel }

**Prototype:** `void wgpuPipelineLayoutSetLabel(WGPUPipelineLayout pipelineLayout, char const * label)`


TODO


**Arguments:**

 - TODO




 - `void wgpuPipelineLayoutReference(WGPUPipelineLayout pipelineLayout)`
 - `void wgpuPipelineLayoutRelease(WGPUPipelineLayout pipelineLayout)`




### QuerySet { #WGPUQuerySet }


TODO




#### `wgpuQuerySetDestroy` { #wgpuQuerySetDestroy }

**Prototype:** `void wgpuQuerySetDestroy(WGPUQuerySet querySet)`


TODO



**Arguments:**

 - TODO




#### `wgpuQuerySetGetCount` { #wgpuQuerySetGetCount }

**Prototype:** `uint32_t wgpuQuerySetGetCount(WGPUQuerySet querySet)`


TODO


**Arguments:**

 - TODO



**Returns:** uint32_t 
TODO





#### `wgpuQuerySetGetType` { #wgpuQuerySetGetType }

**Prototype:** `WGPUQueryType wgpuQuerySetGetType(WGPUQuerySet querySet)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPUQueryType 
TODO





#### `wgpuQuerySetSetLabel` { #wgpuQuerySetSetLabel }

**Prototype:** `void wgpuQuerySetSetLabel(WGPUQuerySet querySet, char const * label)`


TODO


**Arguments:**

 - TODO




 - `void wgpuQuerySetReference(WGPUQuerySet querySet)`
 - `void wgpuQuerySetRelease(WGPUQuerySet querySet)`




### Queue { #WGPUQueue }


TODO




#### `wgpuQueueOnSubmittedWorkDone` { #wgpuQueueOnSubmittedWorkDone }

**Prototype:** `void wgpuQueueOnSubmittedWorkDone(WGPUQueue queue, WGPUQueueOnSubmittedWorkDoneCallback callback, WGPU_NULLABLE void * userdata)`


TODO


**Arguments:**

 - TODO




#### `wgpuQueueSetLabel` { #wgpuQueueSetLabel }

**Prototype:** `void wgpuQueueSetLabel(WGPUQueue queue, char const * label)`


TODO


**Arguments:**

 - TODO




#### `wgpuQueueSubmit` { #wgpuQueueSubmit }

**Prototype:** `void wgpuQueueSubmit(WGPUQueue queue, size_t commandCount, WGPUCommandBuffer const * commands)`


TODO


**Arguments:**

 - TODO




#### `wgpuQueueWriteBuffer` { #wgpuQueueWriteBuffer }

**Prototype:** `void wgpuQueueWriteBuffer(WGPUQueue queue, WGPUBuffer buffer, uint64_t bufferOffset, void const * data, size_t size)`


TODO


**Arguments:**

 - TODO




#### `wgpuQueueWriteTexture` { #wgpuQueueWriteTexture }

**Prototype:** `void wgpuQueueWriteTexture(WGPUQueue queue, WGPUImageCopyTexture const * destination, void const * data, size_t dataSize, WGPUTextureDataLayout const * dataLayout, WGPUExtent3D const * writeSize)`


TODO


**Arguments:**

 - TODO




 - `void wgpuQueueReference(WGPUQueue queue)`
 - `void wgpuQueueRelease(WGPUQueue queue)`




### RenderBundle { #WGPURenderBundle }


TODO




#### `wgpuRenderBundleSetLabel` { #wgpuRenderBundleSetLabel }

**Prototype:** `void wgpuRenderBundleSetLabel(WGPURenderBundle renderBundle, char const * label)`


TODO


**Arguments:**

 - TODO




 - `void wgpuRenderBundleReference(WGPURenderBundle renderBundle)`
 - `void wgpuRenderBundleRelease(WGPURenderBundle renderBundle)`




### RenderBundleEncoder { #WGPURenderBundleEncoder }


TODO




#### `wgpuRenderBundleEncoderDraw` { #wgpuRenderBundleEncoderDraw }

**Prototype:** `void wgpuRenderBundleEncoderDraw(WGPURenderBundleEncoder renderBundleEncoder, uint32_t vertexCount, uint32_t instanceCount, uint32_t firstVertex, uint32_t firstInstance)`


TODO


**Arguments:**

 - TODO




#### `wgpuRenderBundleEncoderDrawIndexed` { #wgpuRenderBundleEncoderDrawIndexed }

**Prototype:** `void wgpuRenderBundleEncoderDrawIndexed(WGPURenderBundleEncoder renderBundleEncoder, uint32_t indexCount, uint32_t instanceCount, uint32_t firstIndex, int32_t baseVertex, uint32_t firstInstance)`


TODO


**Arguments:**

 - TODO




#### `wgpuRenderBundleEncoderDrawIndexedIndirect` { #wgpuRenderBundleEncoderDrawIndexedIndirect }

**Prototype:** `void wgpuRenderBundleEncoderDrawIndexedIndirect(WGPURenderBundleEncoder renderBundleEncoder, WGPUBuffer indirectBuffer, uint64_t indirectOffset)`


TODO


**Arguments:**

 - TODO




#### `wgpuRenderBundleEncoderDrawIndirect` { #wgpuRenderBundleEncoderDrawIndirect }

**Prototype:** `void wgpuRenderBundleEncoderDrawIndirect(WGPURenderBundleEncoder renderBundleEncoder, WGPUBuffer indirectBuffer, uint64_t indirectOffset)`


TODO


**Arguments:**

 - TODO




#### `wgpuRenderBundleEncoderFinish` { #wgpuRenderBundleEncoderFinish }

**Prototype:** `WGPURenderBundle wgpuRenderBundleEncoderFinish(WGPURenderBundleEncoder renderBundleEncoder, WGPU_NULLABLE WGPURenderBundleDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPURenderBundle 
TODO





#### `wgpuRenderBundleEncoderInsertDebugMarker` { #wgpuRenderBundleEncoderInsertDebugMarker }

**Prototype:** `void wgpuRenderBundleEncoderInsertDebugMarker(WGPURenderBundleEncoder renderBundleEncoder, char const * markerLabel)`


TODO


**Arguments:**

 - TODO




#### `wgpuRenderBundleEncoderPopDebugGroup` { #wgpuRenderBundleEncoderPopDebugGroup }

**Prototype:** `void wgpuRenderBundleEncoderPopDebugGroup(WGPURenderBundleEncoder renderBundleEncoder)`


TODO


**Arguments:**

 - TODO




#### `wgpuRenderBundleEncoderPushDebugGroup` { #wgpuRenderBundleEncoderPushDebugGroup }

**Prototype:** `void wgpuRenderBundleEncoderPushDebugGroup(WGPURenderBundleEncoder renderBundleEncoder, char const * groupLabel)`


TODO


**Arguments:**

 - TODO




#### `wgpuRenderBundleEncoderSetBindGroup` { #wgpuRenderBundleEncoderSetBindGroup }

**Prototype:** `void wgpuRenderBundleEncoderSetBindGroup(WGPURenderBundleEncoder renderBundleEncoder, uint32_t groupIndex, WGPU_NULLABLE WGPUBindGroup group, size_t dynamicOffsetCount, uint32_t const * dynamicOffsets)`


TODO


**Arguments:**

 - TODO




#### `wgpuRenderBundleEncoderSetIndexBuffer` { #wgpuRenderBundleEncoderSetIndexBuffer }

**Prototype:** `void wgpuRenderBundleEncoderSetIndexBuffer(WGPURenderBundleEncoder renderBundleEncoder, WGPUBuffer buffer, WGPUIndexFormat format, uint64_t offset, uint64_t size)`


TODO


**Arguments:**

 - TODO




#### `wgpuRenderBundleEncoderSetLabel` { #wgpuRenderBundleEncoderSetLabel }

**Prototype:** `void wgpuRenderBundleEncoderSetLabel(WGPURenderBundleEncoder renderBundleEncoder, char const * label)`


TODO


**Arguments:**

 - TODO




#### `wgpuRenderBundleEncoderSetPipeline` { #wgpuRenderBundleEncoderSetPipeline }

**Prototype:** `void wgpuRenderBundleEncoderSetPipeline(WGPURenderBundleEncoder renderBundleEncoder, WGPURenderPipeline pipeline)`


TODO


**Arguments:**

 - TODO




#### `wgpuRenderBundleEncoderSetVertexBuffer` { #wgpuRenderBundleEncoderSetVertexBuffer }

**Prototype:** `void wgpuRenderBundleEncoderSetVertexBuffer(WGPURenderBundleEncoder renderBundleEncoder, uint32_t slot, WGPU_NULLABLE WGPUBuffer buffer, uint64_t offset, uint64_t size)`


TODO


**Arguments:**

 - TODO




 - `void wgpuRenderBundleEncoderReference(WGPURenderBundleEncoder renderBundleEncoder)`
 - `void wgpuRenderBundleEncoderRelease(WGPURenderBundleEncoder renderBundleEncoder)`




### RenderPassEncoder { #WGPURenderPassEncoder }


TODO




#### `wgpuRenderPassEncoderBeginOcclusionQuery` { #wgpuRenderPassEncoderBeginOcclusionQuery }

**Prototype:** `void wgpuRenderPassEncoderBeginOcclusionQuery(WGPURenderPassEncoder renderPassEncoder, uint32_t queryIndex)`


TODO


**Arguments:**

 - TODO




#### `wgpuRenderPassEncoderDraw` { #wgpuRenderPassEncoderDraw }

**Prototype:** `void wgpuRenderPassEncoderDraw(WGPURenderPassEncoder renderPassEncoder, uint32_t vertexCount, uint32_t instanceCount, uint32_t firstVertex, uint32_t firstInstance)`


TODO


**Arguments:**

 - TODO




#### `wgpuRenderPassEncoderDrawIndexed` { #wgpuRenderPassEncoderDrawIndexed }

**Prototype:** `void wgpuRenderPassEncoderDrawIndexed(WGPURenderPassEncoder renderPassEncoder, uint32_t indexCount, uint32_t instanceCount, uint32_t firstIndex, int32_t baseVertex, uint32_t firstInstance)`


TODO


**Arguments:**

 - TODO




#### `wgpuRenderPassEncoderDrawIndexedIndirect` { #wgpuRenderPassEncoderDrawIndexedIndirect }

**Prototype:** `void wgpuRenderPassEncoderDrawIndexedIndirect(WGPURenderPassEncoder renderPassEncoder, WGPUBuffer indirectBuffer, uint64_t indirectOffset)`


TODO


**Arguments:**

 - TODO




#### `wgpuRenderPassEncoderDrawIndirect` { #wgpuRenderPassEncoderDrawIndirect }

**Prototype:** `void wgpuRenderPassEncoderDrawIndirect(WGPURenderPassEncoder renderPassEncoder, WGPUBuffer indirectBuffer, uint64_t indirectOffset)`


TODO


**Arguments:**

 - TODO




#### `wgpuRenderPassEncoderEnd` { #wgpuRenderPassEncoderEnd }

**Prototype:** `void wgpuRenderPassEncoderEnd(WGPURenderPassEncoder renderPassEncoder)`


TODO


**Arguments:**

 - TODO




#### `wgpuRenderPassEncoderEndOcclusionQuery` { #wgpuRenderPassEncoderEndOcclusionQuery }

**Prototype:** `void wgpuRenderPassEncoderEndOcclusionQuery(WGPURenderPassEncoder renderPassEncoder)`


TODO


**Arguments:**

 - TODO




#### `wgpuRenderPassEncoderExecuteBundles` { #wgpuRenderPassEncoderExecuteBundles }

**Prototype:** `void wgpuRenderPassEncoderExecuteBundles(WGPURenderPassEncoder renderPassEncoder, size_t bundleCount, WGPURenderBundle const * bundles)`


TODO


**Arguments:**

 - TODO




#### `wgpuRenderPassEncoderInsertDebugMarker` { #wgpuRenderPassEncoderInsertDebugMarker }

**Prototype:** `void wgpuRenderPassEncoderInsertDebugMarker(WGPURenderPassEncoder renderPassEncoder, char const * markerLabel)`


TODO


**Arguments:**

 - TODO




#### `wgpuRenderPassEncoderPopDebugGroup` { #wgpuRenderPassEncoderPopDebugGroup }

**Prototype:** `void wgpuRenderPassEncoderPopDebugGroup(WGPURenderPassEncoder renderPassEncoder)`


TODO


**Arguments:**

 - TODO




#### `wgpuRenderPassEncoderPushDebugGroup` { #wgpuRenderPassEncoderPushDebugGroup }

**Prototype:** `void wgpuRenderPassEncoderPushDebugGroup(WGPURenderPassEncoder renderPassEncoder, char const * groupLabel)`


TODO


**Arguments:**

 - TODO




#### `wgpuRenderPassEncoderSetBindGroup` { #wgpuRenderPassEncoderSetBindGroup }

**Prototype:** `void wgpuRenderPassEncoderSetBindGroup(WGPURenderPassEncoder renderPassEncoder, uint32_t groupIndex, WGPU_NULLABLE WGPUBindGroup group, size_t dynamicOffsetCount, uint32_t const * dynamicOffsets)`


TODO


**Arguments:**

 - TODO




#### `wgpuRenderPassEncoderSetBlendConstant` { #wgpuRenderPassEncoderSetBlendConstant }

**Prototype:** `void wgpuRenderPassEncoderSetBlendConstant(WGPURenderPassEncoder renderPassEncoder, WGPUColor const * color)`


TODO


**Arguments:**

 - TODO




#### `wgpuRenderPassEncoderSetIndexBuffer` { #wgpuRenderPassEncoderSetIndexBuffer }

**Prototype:** `void wgpuRenderPassEncoderSetIndexBuffer(WGPURenderPassEncoder renderPassEncoder, WGPUBuffer buffer, WGPUIndexFormat format, uint64_t offset, uint64_t size)`


TODO


**Arguments:**

 - TODO




#### `wgpuRenderPassEncoderSetLabel` { #wgpuRenderPassEncoderSetLabel }

**Prototype:** `void wgpuRenderPassEncoderSetLabel(WGPURenderPassEncoder renderPassEncoder, char const * label)`


TODO


**Arguments:**

 - TODO




#### `wgpuRenderPassEncoderSetPipeline` { #wgpuRenderPassEncoderSetPipeline }

**Prototype:** `void wgpuRenderPassEncoderSetPipeline(WGPURenderPassEncoder renderPassEncoder, WGPURenderPipeline pipeline)`


TODO


**Arguments:**

 - TODO




#### `wgpuRenderPassEncoderSetScissorRect` { #wgpuRenderPassEncoderSetScissorRect }

**Prototype:** `void wgpuRenderPassEncoderSetScissorRect(WGPURenderPassEncoder renderPassEncoder, uint32_t x, uint32_t y, uint32_t width, uint32_t height)`


TODO


**Arguments:**

 - TODO




#### `wgpuRenderPassEncoderSetStencilReference` { #wgpuRenderPassEncoderSetStencilReference }

**Prototype:** `void wgpuRenderPassEncoderSetStencilReference(WGPURenderPassEncoder renderPassEncoder, uint32_t reference)`


TODO


**Arguments:**

 - TODO




#### `wgpuRenderPassEncoderSetVertexBuffer` { #wgpuRenderPassEncoderSetVertexBuffer }

**Prototype:** `void wgpuRenderPassEncoderSetVertexBuffer(WGPURenderPassEncoder renderPassEncoder, uint32_t slot, WGPU_NULLABLE WGPUBuffer buffer, uint64_t offset, uint64_t size)`


TODO


**Arguments:**

 - TODO




#### `wgpuRenderPassEncoderSetViewport` { #wgpuRenderPassEncoderSetViewport }

**Prototype:** `void wgpuRenderPassEncoderSetViewport(WGPURenderPassEncoder renderPassEncoder, float x, float y, float width, float height, float minDepth, float maxDepth)`


TODO


**Arguments:**

 - TODO




 - `void wgpuRenderPassEncoderReference(WGPURenderPassEncoder renderPassEncoder)`
 - `void wgpuRenderPassEncoderRelease(WGPURenderPassEncoder renderPassEncoder)`




### RenderPipeline { #WGPURenderPipeline }


TODO




#### `wgpuRenderPipelineGetBindGroupLayout` { #wgpuRenderPipelineGetBindGroupLayout }

**Prototype:** `WGPUBindGroupLayout wgpuRenderPipelineGetBindGroupLayout(WGPURenderPipeline renderPipeline, uint32_t groupIndex)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPUBindGroupLayout 
TODO





#### `wgpuRenderPipelineSetLabel` { #wgpuRenderPipelineSetLabel }

**Prototype:** `void wgpuRenderPipelineSetLabel(WGPURenderPipeline renderPipeline, char const * label)`


TODO


**Arguments:**

 - TODO




 - `void wgpuRenderPipelineReference(WGPURenderPipeline renderPipeline)`
 - `void wgpuRenderPipelineRelease(WGPURenderPipeline renderPipeline)`




### Sampler { #WGPUSampler }


TODO




#### `wgpuSamplerSetLabel` { #wgpuSamplerSetLabel }

**Prototype:** `void wgpuSamplerSetLabel(WGPUSampler sampler, char const * label)`


TODO


**Arguments:**

 - TODO




 - `void wgpuSamplerReference(WGPUSampler sampler)`
 - `void wgpuSamplerRelease(WGPUSampler sampler)`




### ShaderModule { #WGPUShaderModule }


TODO




#### `wgpuShaderModuleGetCompilationInfo` { #wgpuShaderModuleGetCompilationInfo }

**Prototype:** `void wgpuShaderModuleGetCompilationInfo(WGPUShaderModule shaderModule, WGPUShaderModuleGetCompilationInfoCallback callback, WGPU_NULLABLE void * userdata)`


TODO


**Arguments:**

 - TODO




#### `wgpuShaderModuleSetLabel` { #wgpuShaderModuleSetLabel }

**Prototype:** `void wgpuShaderModuleSetLabel(WGPUShaderModule shaderModule, char const * label)`


TODO


**Arguments:**

 - TODO




 - `void wgpuShaderModuleReference(WGPUShaderModule shaderModule)`
 - `void wgpuShaderModuleRelease(WGPUShaderModule shaderModule)`




### Surface { #WGPUSurface }


TODO




#### `wgpuSurfaceConfigure` { #wgpuSurfaceConfigure }

**Prototype:** `void wgpuSurfaceConfigure(WGPUSurface surface, WGPUSurfaceConfiguration const * config)`


TODO


**Arguments:**

 - TODO




#### `wgpuSurfaceGetCapabilities` { #wgpuSurfaceGetCapabilities }

**Prototype:** `void wgpuSurfaceGetCapabilities(WGPUSurface surface, WGPUAdapter adapter, WGPUSurfaceCapabilities * capabilities)`


TODO


**Arguments:**

 - TODO




#### `wgpuSurfaceGetCurrentTexture` { #wgpuSurfaceGetCurrentTexture }

**Prototype:** `void wgpuSurfaceGetCurrentTexture(WGPUSurface surface, WGPUSurfaceTexture * surfaceTexture)`


TODO


**Arguments:**

 - TODO




#### `wgpuSurfaceGetPreferredFormat` { #wgpuSurfaceGetPreferredFormat }

**Prototype:** `WGPUTextureFormat wgpuSurfaceGetPreferredFormat(WGPUSurface surface, WGPUAdapter adapter)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPUTextureFormat 
TODO





#### `wgpuSurfacePresent` { #wgpuSurfacePresent }

**Prototype:** `void wgpuSurfacePresent(WGPUSurface surface)`


TODO


**Arguments:**

 - TODO




#### `wgpuSurfaceSetLabel` { #wgpuSurfaceSetLabel }

**Prototype:** `void wgpuSurfaceSetLabel(WGPUSurface surface, char const * label)`


TODO


**Arguments:**

 - TODO




#### `wgpuSurfaceUnconfigure` { #wgpuSurfaceUnconfigure }

**Prototype:** `void wgpuSurfaceUnconfigure(WGPUSurface surface)`


TODO


**Arguments:**

 - TODO




 - `void wgpuSurfaceReference(WGPUSurface surface)`
 - `void wgpuSurfaceRelease(WGPUSurface surface)`





### Texture { #WGPUTexture }


TODO




#### `wgpuTextureCreateView` { #wgpuTextureCreateView }

**Prototype:** `WGPUTextureView wgpuTextureCreateView(WGPUTexture texture, WGPU_NULLABLE WGPUTextureViewDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPUTextureView 
TODO





#### `wgpuTextureDestroy` { #wgpuTextureDestroy }

**Prototype:** `void wgpuTextureDestroy(WGPUTexture texture)`


TODO



**Arguments:**

 - TODO




#### `wgpuTextureGetDepthOrArrayLayers` { #wgpuTextureGetDepthOrArrayLayers }

**Prototype:** `uint32_t wgpuTextureGetDepthOrArrayLayers(WGPUTexture texture)`


TODO


**Arguments:**

 - TODO



**Returns:** uint32_t 
TODO





#### `wgpuTextureGetDimension` { #wgpuTextureGetDimension }

**Prototype:** `WGPUTextureDimension wgpuTextureGetDimension(WGPUTexture texture)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPUTextureDimension 
TODO





#### `wgpuTextureGetFormat` { #wgpuTextureGetFormat }

**Prototype:** `WGPUTextureFormat wgpuTextureGetFormat(WGPUTexture texture)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPUTextureFormat 
TODO





#### `wgpuTextureGetHeight` { #wgpuTextureGetHeight }

**Prototype:** `uint32_t wgpuTextureGetHeight(WGPUTexture texture)`


TODO


**Arguments:**

 - TODO



**Returns:** uint32_t 
TODO





#### `wgpuTextureGetMipLevelCount` { #wgpuTextureGetMipLevelCount }

**Prototype:** `uint32_t wgpuTextureGetMipLevelCount(WGPUTexture texture)`


TODO


**Arguments:**

 - TODO



**Returns:** uint32_t 
TODO





#### `wgpuTextureGetSampleCount` { #wgpuTextureGetSampleCount }

**Prototype:** `uint32_t wgpuTextureGetSampleCount(WGPUTexture texture)`


TODO


**Arguments:**

 - TODO



**Returns:** uint32_t 
TODO





#### `wgpuTextureGetUsage` { #wgpuTextureGetUsage }

**Prototype:** `WGPUTextureUsageFlags wgpuTextureGetUsage(WGPUTexture texture)`


TODO


**Arguments:**

 - TODO



**Returns:** WGPUTextureUsageFlags 
TODO





#### `wgpuTextureGetWidth` { #wgpuTextureGetWidth }

**Prototype:** `uint32_t wgpuTextureGetWidth(WGPUTexture texture)`


TODO


**Arguments:**

 - TODO



**Returns:** uint32_t 
TODO





#### `wgpuTextureSetLabel` { #wgpuTextureSetLabel }

**Prototype:** `void wgpuTextureSetLabel(WGPUTexture texture, char const * label)`


TODO


**Arguments:**

 - TODO




 - `void wgpuTextureReference(WGPUTexture texture)`
 - `void wgpuTextureRelease(WGPUTexture texture)`




### TextureView { #WGPUTextureView }


TODO




#### `wgpuTextureViewSetLabel` { #wgpuTextureViewSetLabel }

**Prototype:** `void wgpuTextureViewSetLabel(WGPUTextureView textureView, char const * label)`


TODO


**Arguments:**

 - TODO




 - `void wgpuTextureViewReference(WGPUTextureView textureView)`
 - `void wgpuTextureViewRelease(WGPUTextureView textureView)`





## Structures


 - `WGPUAdapterInfo`
 - `WGPUAdapterProperties`
 - `WGPUBindGroupEntry`
 - `WGPUBlendComponent`
 - `WGPUBufferBindingLayout`
 - `WGPUBufferDescriptor`
 - `WGPUColor`
 - `WGPUCommandBufferDescriptor`
 - `WGPUCommandEncoderDescriptor`
 - `WGPUCompilationMessage`
 - `WGPUComputePassTimestampWrites`
 - `WGPUConstantEntry`
 - `WGPUExtent3D`
 - `WGPUInstanceDescriptor`
 - `WGPULimits`
 - `WGPUMultisampleState`
 - `WGPUOrigin3D`
 - `WGPUPipelineLayoutDescriptor`
 - `WGPUPrimitiveDepthClipControl`
 - `WGPUPrimitiveState`
 - `WGPUQuerySetDescriptor`
 - `WGPUQueueDescriptor`
 - `WGPURenderBundleDescriptor`
 - `WGPURenderBundleEncoderDescriptor`
 - `WGPURenderPassDepthStencilAttachment`
 - `WGPURenderPassDescriptorMaxDrawCount`
 - `WGPURenderPassTimestampWrites`
 - `WGPURequestAdapterOptions`
 - `WGPUSamplerBindingLayout`
 - `WGPUSamplerDescriptor`
 - `WGPUShaderModuleCompilationHint`
 - `WGPUShaderModuleSPIRVDescriptor`
 - `WGPUShaderModuleWGSLDescriptor`
 - `WGPUStencilFaceState`
 - `WGPUStorageTextureBindingLayout`
 - `WGPUSurfaceCapabilities`
 - `WGPUSurfaceConfiguration`
 - `WGPUSurfaceDescriptor`
 - `WGPUSurfaceDescriptorFromAndroidNativeWindow`
 - `WGPUSurfaceDescriptorFromCanvasHTMLSelector`
 - `WGPUSurfaceDescriptorFromMetalLayer`
 - `WGPUSurfaceDescriptorFromWaylandSurface`
 - `WGPUSurfaceDescriptorFromWindowsHWND`
 - `WGPUSurfaceDescriptorFromXcbWindow`
 - `WGPUSurfaceDescriptorFromXlibWindow`
 - `WGPUSurfaceTexture`
 - `WGPUTextureBindingLayout`
 - `WGPUTextureDataLayout`
 - `WGPUTextureViewDescriptor`
 - `WGPUVertexAttribute`
 - `WGPUBindGroupDescriptor`
 - `WGPUBindGroupLayoutEntry`
 - `WGPUBlendState`
 - `WGPUCompilationInfo`
 - `WGPUComputePassDescriptor`
 - `WGPUDepthStencilState`
 - `WGPUImageCopyBuffer`
 - `WGPUImageCopyTexture`
 - `WGPUProgrammableStageDescriptor`
 - `WGPURenderPassColorAttachment`
 - `WGPURequiredLimits`
 - `WGPUShaderModuleDescriptor`
 - `WGPUSupportedLimits`
 - `WGPUTextureDescriptor`
 - `WGPUVertexBufferLayout`
 - `WGPUBindGroupLayoutDescriptor`
 - `WGPUColorTargetState`
 - `WGPUComputePipelineDescriptor`
 - `WGPUDeviceDescriptor`
 - `WGPURenderPassDescriptor`
 - `WGPUVertexState`
 - `WGPUFragmentState`
 - `WGPURenderPipelineDescriptor`


## Enumerations

**NB** All enumerations also have an extra values ending with `_Force32` that is here to ensure that all compilers use the same underlying representation for the enum values. This value must never be used.


```C
enum WGPUAdapterType {
    WGPUAdapterType_DiscreteGPU = 0x00000000,
    WGPUAdapterType_IntegratedGPU = 0x00000001,
    WGPUAdapterType_CPU = 0x00000002,
    WGPUAdapterType_Unknown = 0x00000003,
}
```

```C
enum WGPUAddressMode {
    WGPUAddressMode_Repeat = 0x00000000,
    WGPUAddressMode_MirrorRepeat = 0x00000001,
    WGPUAddressMode_ClampToEdge = 0x00000002,
}
```

```C
enum WGPUBackendType {
    WGPUBackendType_Undefined = 0x00000000,
    WGPUBackendType_Null = 0x00000001,
    WGPUBackendType_WebGPU = 0x00000002,
    WGPUBackendType_D3D11 = 0x00000003,
    WGPUBackendType_D3D12 = 0x00000004,
    WGPUBackendType_Metal = 0x00000005,
    WGPUBackendType_Vulkan = 0x00000006,
    WGPUBackendType_OpenGL = 0x00000007,
    WGPUBackendType_OpenGLES = 0x00000008,
}
```

```C
enum WGPUBlendFactor {
    WGPUBlendFactor_Zero = 0x00000000,
    WGPUBlendFactor_One = 0x00000001,
    WGPUBlendFactor_Src = 0x00000002,
    WGPUBlendFactor_OneMinusSrc = 0x00000003,
    WGPUBlendFactor_SrcAlpha = 0x00000004,
    WGPUBlendFactor_OneMinusSrcAlpha = 0x00000005,
    WGPUBlendFactor_Dst = 0x00000006,
    WGPUBlendFactor_OneMinusDst = 0x00000007,
    WGPUBlendFactor_DstAlpha = 0x00000008,
    WGPUBlendFactor_OneMinusDstAlpha = 0x00000009,
    WGPUBlendFactor_SrcAlphaSaturated = 0x0000000A,
    WGPUBlendFactor_Constant = 0x0000000B,
    WGPUBlendFactor_OneMinusConstant = 0x0000000C,
}
```

```C
enum WGPUBlendOperation {
    WGPUBlendOperation_Add = 0x00000000,
    WGPUBlendOperation_Subtract = 0x00000001,
    WGPUBlendOperation_ReverseSubtract = 0x00000002,
    WGPUBlendOperation_Min = 0x00000003,
    WGPUBlendOperation_Max = 0x00000004,
}
```

```C
enum WGPUBufferBindingType {
    WGPUBufferBindingType_Undefined = 0x00000000,
    WGPUBufferBindingType_Uniform = 0x00000001,
    WGPUBufferBindingType_Storage = 0x00000002,
    WGPUBufferBindingType_ReadOnlyStorage = 0x00000003,
}
```

```C
enum WGPUBufferMapAsyncStatus {
    WGPUBufferMapAsyncStatus_Success = 0x00000000,
    WGPUBufferMapAsyncStatus_ValidationError = 0x00000001,
    WGPUBufferMapAsyncStatus_Unknown = 0x00000002,
    WGPUBufferMapAsyncStatus_DeviceLost = 0x00000003,
    WGPUBufferMapAsyncStatus_DestroyedBeforeCallback = 0x00000004,
    WGPUBufferMapAsyncStatus_UnmappedBeforeCallback = 0x00000005,
    WGPUBufferMapAsyncStatus_MappingAlreadyPending = 0x00000006,
    WGPUBufferMapAsyncStatus_OffsetOutOfRange = 0x00000007,
    WGPUBufferMapAsyncStatus_SizeOutOfRange = 0x00000008,
}
```

```C
enum WGPUBufferMapState {
    WGPUBufferMapState_Unmapped = 0x00000000,
    WGPUBufferMapState_Pending = 0x00000001,
    WGPUBufferMapState_Mapped = 0x00000002,
}
```

```C
enum WGPUCompareFunction {
    WGPUCompareFunction_Undefined = 0x00000000,
    WGPUCompareFunction_Never = 0x00000001,
    WGPUCompareFunction_Less = 0x00000002,
    WGPUCompareFunction_LessEqual = 0x00000003,
    WGPUCompareFunction_Greater = 0x00000004,
    WGPUCompareFunction_GreaterEqual = 0x00000005,
    WGPUCompareFunction_Equal = 0x00000006,
    WGPUCompareFunction_NotEqual = 0x00000007,
    WGPUCompareFunction_Always = 0x00000008,
}
```

```C
enum WGPUCompilationInfoRequestStatus {
    WGPUCompilationInfoRequestStatus_Success = 0x00000000,
    WGPUCompilationInfoRequestStatus_Error = 0x00000001,
    WGPUCompilationInfoRequestStatus_DeviceLost = 0x00000002,
    WGPUCompilationInfoRequestStatus_Unknown = 0x00000003,
}
```

```C
enum WGPUCompilationMessageType {
    WGPUCompilationMessageType_Error = 0x00000000,
    WGPUCompilationMessageType_Warning = 0x00000001,
    WGPUCompilationMessageType_Info = 0x00000002,
}
```

```C
enum WGPUCompositeAlphaMode {
    WGPUCompositeAlphaMode_Auto = 0x00000000,
    WGPUCompositeAlphaMode_Opaque = 0x00000001,
    WGPUCompositeAlphaMode_Premultiplied = 0x00000002,
    WGPUCompositeAlphaMode_Unpremultiplied = 0x00000003,
    WGPUCompositeAlphaMode_Inherit = 0x00000004,
}
```

```C
enum WGPUCreatePipelineAsyncStatus {
    WGPUCreatePipelineAsyncStatus_Success = 0x00000000,
    WGPUCreatePipelineAsyncStatus_ValidationError = 0x00000001,
    WGPUCreatePipelineAsyncStatus_InternalError = 0x00000002,
    WGPUCreatePipelineAsyncStatus_DeviceLost = 0x00000003,
    WGPUCreatePipelineAsyncStatus_DeviceDestroyed = 0x00000004,
    WGPUCreatePipelineAsyncStatus_Unknown = 0x00000005,
}
```

```C
enum WGPUCullMode {
    WGPUCullMode_None = 0x00000000,
    WGPUCullMode_Front = 0x00000001,
    WGPUCullMode_Back = 0x00000002,
}
```

```C
enum WGPUDeviceLostReason {
    WGPUDeviceLostReason_Undefined = 0x00000000,
    WGPUDeviceLostReason_Destroyed = 0x00000001,
}
```

```C
enum WGPUErrorFilter {
    WGPUErrorFilter_Validation = 0x00000000,
    WGPUErrorFilter_OutOfMemory = 0x00000001,
    WGPUErrorFilter_Internal = 0x00000002,
}
```

```C
enum WGPUErrorType {
    WGPUErrorType_NoError = 0x00000000,
    WGPUErrorType_Validation = 0x00000001,
    WGPUErrorType_OutOfMemory = 0x00000002,
    WGPUErrorType_Internal = 0x00000003,
    WGPUErrorType_Unknown = 0x00000004,
    WGPUErrorType_DeviceLost = 0x00000005,
}
```

```C
enum WGPUFeatureName {
    WGPUFeatureName_Undefined = 0x00000000,
    WGPUFeatureName_DepthClipControl = 0x00000001,
    WGPUFeatureName_Depth32FloatStencil8 = 0x00000002,
    WGPUFeatureName_TimestampQuery = 0x00000003,
    WGPUFeatureName_TextureCompressionBC = 0x00000004,
    WGPUFeatureName_TextureCompressionETC2 = 0x00000005,
    WGPUFeatureName_TextureCompressionASTC = 0x00000006,
    WGPUFeatureName_IndirectFirstInstance = 0x00000007,
    WGPUFeatureName_ShaderF16 = 0x00000008,
    WGPUFeatureName_RG11B10UfloatRenderable = 0x00000009,
    WGPUFeatureName_BGRA8UnormStorage = 0x0000000A,
    WGPUFeatureName_Float32Filterable = 0x0000000B,
}
```

```C
enum WGPUFilterMode {
    WGPUFilterMode_Nearest = 0x00000000,
    WGPUFilterMode_Linear = 0x00000001,
}
```

```C
enum WGPUFrontFace {
    WGPUFrontFace_CCW = 0x00000000,
    WGPUFrontFace_CW = 0x00000001,
}
```

```C
enum WGPUIndexFormat {
    WGPUIndexFormat_Undefined = 0x00000000,
    WGPUIndexFormat_Uint16 = 0x00000001,
    WGPUIndexFormat_Uint32 = 0x00000002,
}
```

```C
enum WGPULoadOp {
    WGPULoadOp_Undefined = 0x00000000,
    WGPULoadOp_Clear = 0x00000001,
    WGPULoadOp_Load = 0x00000002,
}
```

```C
enum WGPUMipmapFilterMode {
    WGPUMipmapFilterMode_Nearest = 0x00000000,
    WGPUMipmapFilterMode_Linear = 0x00000001,
}
```

```C
enum WGPUPowerPreference {
    WGPUPowerPreference_Undefined = 0x00000000,
    WGPUPowerPreference_LowPower = 0x00000001,
    WGPUPowerPreference_HighPerformance = 0x00000002,
}
```

```C
enum WGPUPresentMode {
    WGPUPresentMode_Fifo = 0x00000000,
    WGPUPresentMode_FifoRelaxed = 0x00000001,
    WGPUPresentMode_Immediate = 0x00000002,
    WGPUPresentMode_Mailbox = 0x00000003,
}
```

```C
enum WGPUPrimitiveTopology {
    WGPUPrimitiveTopology_PointList = 0x00000000,
    WGPUPrimitiveTopology_LineList = 0x00000001,
    WGPUPrimitiveTopology_LineStrip = 0x00000002,
    WGPUPrimitiveTopology_TriangleList = 0x00000003,
    WGPUPrimitiveTopology_TriangleStrip = 0x00000004,
}
```

```C
enum WGPUQueryType {
    WGPUQueryType_Occlusion = 0x00000000,
    WGPUQueryType_Timestamp = 0x00000001,
}
```

```C
enum WGPUQueueWorkDoneStatus {
    WGPUQueueWorkDoneStatus_Success = 0x00000000,
    WGPUQueueWorkDoneStatus_Error = 0x00000001,
    WGPUQueueWorkDoneStatus_Unknown = 0x00000002,
    WGPUQueueWorkDoneStatus_DeviceLost = 0x00000003,
}
```

```C
enum WGPURequestAdapterStatus {
    WGPURequestAdapterStatus_Success = 0x00000000,
    WGPURequestAdapterStatus_Unavailable = 0x00000001,
    WGPURequestAdapterStatus_Error = 0x00000002,
    WGPURequestAdapterStatus_Unknown = 0x00000003,
}
```

```C
enum WGPURequestDeviceStatus {
    WGPURequestDeviceStatus_Success = 0x00000000,
    WGPURequestDeviceStatus_Error = 0x00000001,
    WGPURequestDeviceStatus_Unknown = 0x00000002,
}
```

```C
enum WGPUSType {
    WGPUSType_Invalid = 0x00000000,
    WGPUSType_SurfaceDescriptorFromMetalLayer = 0x00000001,
    WGPUSType_SurfaceDescriptorFromWindowsHWND = 0x00000002,
    WGPUSType_SurfaceDescriptorFromXlibWindow = 0x00000003,
    WGPUSType_SurfaceDescriptorFromCanvasHTMLSelector = 0x00000004,
    WGPUSType_ShaderModuleSPIRVDescriptor = 0x00000005,
    WGPUSType_ShaderModuleWGSLDescriptor = 0x00000006,
    WGPUSType_PrimitiveDepthClipControl = 0x00000007,
    WGPUSType_SurfaceDescriptorFromWaylandSurface = 0x00000008,
    WGPUSType_SurfaceDescriptorFromAndroidNativeWindow = 0x00000009,
    WGPUSType_SurfaceDescriptorFromXcbWindow = 0x0000000A,
    WGPUSType_RenderPassDescriptorMaxDrawCount = 0x0000000F,
}
```

```C
enum WGPUSamplerBindingType {
    WGPUSamplerBindingType_Undefined = 0x00000000,
    WGPUSamplerBindingType_Filtering = 0x00000001,
    WGPUSamplerBindingType_NonFiltering = 0x00000002,
    WGPUSamplerBindingType_Comparison = 0x00000003,
}
```

```C
enum WGPUStencilOperation {
    WGPUStencilOperation_Keep = 0x00000000,
    WGPUStencilOperation_Zero = 0x00000001,
    WGPUStencilOperation_Replace = 0x00000002,
    WGPUStencilOperation_Invert = 0x00000003,
    WGPUStencilOperation_IncrementClamp = 0x00000004,
    WGPUStencilOperation_DecrementClamp = 0x00000005,
    WGPUStencilOperation_IncrementWrap = 0x00000006,
    WGPUStencilOperation_DecrementWrap = 0x00000007,
}
```

```C
enum WGPUStorageTextureAccess {
    WGPUStorageTextureAccess_Undefined = 0x00000000,
    WGPUStorageTextureAccess_WriteOnly = 0x00000001,
    WGPUStorageTextureAccess_ReadOnly = 0x00000002,
    WGPUStorageTextureAccess_ReadWrite = 0x00000003,
}
```

```C
enum WGPUStoreOp {
    WGPUStoreOp_Undefined = 0x00000000,
    WGPUStoreOp_Store = 0x00000001,
    WGPUStoreOp_Discard = 0x00000002,
}
```

```C
enum WGPUSurfaceGetCurrentTextureStatus {
    WGPUSurfaceGetCurrentTextureStatus_Success = 0x00000000,
    WGPUSurfaceGetCurrentTextureStatus_Timeout = 0x00000001,
    WGPUSurfaceGetCurrentTextureStatus_Outdated = 0x00000002,
    WGPUSurfaceGetCurrentTextureStatus_Lost = 0x00000003,
    WGPUSurfaceGetCurrentTextureStatus_OutOfMemory = 0x00000004,
    WGPUSurfaceGetCurrentTextureStatus_DeviceLost = 0x00000005,
}
```

```C
enum WGPUTextureAspect {
    WGPUTextureAspect_All = 0x00000000,
    WGPUTextureAspect_StencilOnly = 0x00000001,
    WGPUTextureAspect_DepthOnly = 0x00000002,
}
```

```C
enum WGPUTextureDimension {
    WGPUTextureDimension_1D = 0x00000000,
    WGPUTextureDimension_2D = 0x00000001,
    WGPUTextureDimension_3D = 0x00000002,
}
```

```C
enum WGPUTextureFormat {
    WGPUTextureFormat_Undefined = 0x00000000,
    WGPUTextureFormat_R8Unorm = 0x00000001,
    WGPUTextureFormat_R8Snorm = 0x00000002,
    WGPUTextureFormat_R8Uint = 0x00000003,
    WGPUTextureFormat_R8Sint = 0x00000004,
    WGPUTextureFormat_R16Uint = 0x00000005,
    WGPUTextureFormat_R16Sint = 0x00000006,
    WGPUTextureFormat_R16Float = 0x00000007,
    WGPUTextureFormat_RG8Unorm = 0x00000008,
    WGPUTextureFormat_RG8Snorm = 0x00000009,
    WGPUTextureFormat_RG8Uint = 0x0000000A,
    WGPUTextureFormat_RG8Sint = 0x0000000B,
    WGPUTextureFormat_R32Float = 0x0000000C,
    WGPUTextureFormat_R32Uint = 0x0000000D,
    WGPUTextureFormat_R32Sint = 0x0000000E,
    WGPUTextureFormat_RG16Uint = 0x0000000F,
    WGPUTextureFormat_RG16Sint = 0x00000010,
    WGPUTextureFormat_RG16Float = 0x00000011,
    WGPUTextureFormat_RGBA8Unorm = 0x00000012,
    WGPUTextureFormat_RGBA8UnormSrgb = 0x00000013,
    WGPUTextureFormat_RGBA8Snorm = 0x00000014,
    WGPUTextureFormat_RGBA8Uint = 0x00000015,
    WGPUTextureFormat_RGBA8Sint = 0x00000016,
    WGPUTextureFormat_BGRA8Unorm = 0x00000017,
    WGPUTextureFormat_BGRA8UnormSrgb = 0x00000018,
    WGPUTextureFormat_RGB10A2Uint = 0x00000019,
    WGPUTextureFormat_RGB10A2Unorm = 0x0000001A,
    WGPUTextureFormat_RG11B10Ufloat = 0x0000001B,
    WGPUTextureFormat_RGB9E5Ufloat = 0x0000001C,
    WGPUTextureFormat_RG32Float = 0x0000001D,
    WGPUTextureFormat_RG32Uint = 0x0000001E,
    WGPUTextureFormat_RG32Sint = 0x0000001F,
    WGPUTextureFormat_RGBA16Uint = 0x00000020,
    WGPUTextureFormat_RGBA16Sint = 0x00000021,
    WGPUTextureFormat_RGBA16Float = 0x00000022,
    WGPUTextureFormat_RGBA32Float = 0x00000023,
    WGPUTextureFormat_RGBA32Uint = 0x00000024,
    WGPUTextureFormat_RGBA32Sint = 0x00000025,
    WGPUTextureFormat_Stencil8 = 0x00000026,
    WGPUTextureFormat_Depth16Unorm = 0x00000027,
    WGPUTextureFormat_Depth24Plus = 0x00000028,
    WGPUTextureFormat_Depth24PlusStencil8 = 0x00000029,
    WGPUTextureFormat_Depth32Float = 0x0000002A,
    WGPUTextureFormat_Depth32FloatStencil8 = 0x0000002B,
    WGPUTextureFormat_BC1RGBAUnorm = 0x0000002C,
    WGPUTextureFormat_BC1RGBAUnormSrgb = 0x0000002D,
    WGPUTextureFormat_BC2RGBAUnorm = 0x0000002E,
    WGPUTextureFormat_BC2RGBAUnormSrgb = 0x0000002F,
    WGPUTextureFormat_BC3RGBAUnorm = 0x00000030,
    WGPUTextureFormat_BC3RGBAUnormSrgb = 0x00000031,
    WGPUTextureFormat_BC4RUnorm = 0x00000032,
    WGPUTextureFormat_BC4RSnorm = 0x00000033,
    WGPUTextureFormat_BC5RGUnorm = 0x00000034,
    WGPUTextureFormat_BC5RGSnorm = 0x00000035,
    WGPUTextureFormat_BC6HRGBUfloat = 0x00000036,
    WGPUTextureFormat_BC6HRGBFloat = 0x00000037,
    WGPUTextureFormat_BC7RGBAUnorm = 0x00000038,
    WGPUTextureFormat_BC7RGBAUnormSrgb = 0x00000039,
    WGPUTextureFormat_ETC2RGB8Unorm = 0x0000003A,
    WGPUTextureFormat_ETC2RGB8UnormSrgb = 0x0000003B,
    WGPUTextureFormat_ETC2RGB8A1Unorm = 0x0000003C,
    WGPUTextureFormat_ETC2RGB8A1UnormSrgb = 0x0000003D,
    WGPUTextureFormat_ETC2RGBA8Unorm = 0x0000003E,
    WGPUTextureFormat_ETC2RGBA8UnormSrgb = 0x0000003F,
    WGPUTextureFormat_EACR11Unorm = 0x00000040,
    WGPUTextureFormat_EACR11Snorm = 0x00000041,
    WGPUTextureFormat_EACRG11Unorm = 0x00000042,
    WGPUTextureFormat_EACRG11Snorm = 0x00000043,
    WGPUTextureFormat_ASTC4x4Unorm = 0x00000044,
    WGPUTextureFormat_ASTC4x4UnormSrgb = 0x00000045,
    WGPUTextureFormat_ASTC5x4Unorm = 0x00000046,
    WGPUTextureFormat_ASTC5x4UnormSrgb = 0x00000047,
    WGPUTextureFormat_ASTC5x5Unorm = 0x00000048,
    WGPUTextureFormat_ASTC5x5UnormSrgb = 0x00000049,
    WGPUTextureFormat_ASTC6x5Unorm = 0x0000004A,
    WGPUTextureFormat_ASTC6x5UnormSrgb = 0x0000004B,
    WGPUTextureFormat_ASTC6x6Unorm = 0x0000004C,
    WGPUTextureFormat_ASTC6x6UnormSrgb = 0x0000004D,
    WGPUTextureFormat_ASTC8x5Unorm = 0x0000004E,
    WGPUTextureFormat_ASTC8x5UnormSrgb = 0x0000004F,
    WGPUTextureFormat_ASTC8x6Unorm = 0x00000050,
    WGPUTextureFormat_ASTC8x6UnormSrgb = 0x00000051,
    WGPUTextureFormat_ASTC8x8Unorm = 0x00000052,
    WGPUTextureFormat_ASTC8x8UnormSrgb = 0x00000053,
    WGPUTextureFormat_ASTC10x5Unorm = 0x00000054,
    WGPUTextureFormat_ASTC10x5UnormSrgb = 0x00000055,
    WGPUTextureFormat_ASTC10x6Unorm = 0x00000056,
    WGPUTextureFormat_ASTC10x6UnormSrgb = 0x00000057,
    WGPUTextureFormat_ASTC10x8Unorm = 0x00000058,
    WGPUTextureFormat_ASTC10x8UnormSrgb = 0x00000059,
    WGPUTextureFormat_ASTC10x10Unorm = 0x0000005A,
    WGPUTextureFormat_ASTC10x10UnormSrgb = 0x0000005B,
    WGPUTextureFormat_ASTC12x10Unorm = 0x0000005C,
    WGPUTextureFormat_ASTC12x10UnormSrgb = 0x0000005D,
    WGPUTextureFormat_ASTC12x12Unorm = 0x0000005E,
    WGPUTextureFormat_ASTC12x12UnormSrgb = 0x0000005F,
}
```

```C
enum WGPUTextureSampleType {
    WGPUTextureSampleType_Undefined = 0x00000000,
    WGPUTextureSampleType_Float = 0x00000001,
    WGPUTextureSampleType_UnfilterableFloat = 0x00000002,
    WGPUTextureSampleType_Depth = 0x00000003,
    WGPUTextureSampleType_Sint = 0x00000004,
    WGPUTextureSampleType_Uint = 0x00000005,
}
```

```C
enum WGPUTextureViewDimension {
    WGPUTextureViewDimension_Undefined = 0x00000000,
    WGPUTextureViewDimension_1D = 0x00000001,
    WGPUTextureViewDimension_2D = 0x00000002,
    WGPUTextureViewDimension_2DArray = 0x00000003,
    WGPUTextureViewDimension_Cube = 0x00000004,
    WGPUTextureViewDimension_CubeArray = 0x00000005,
    WGPUTextureViewDimension_3D = 0x00000006,
}
```

```C
enum WGPUVertexFormat {
    WGPUVertexFormat_Undefined = 0x00000000,
    WGPUVertexFormat_Uint8x2 = 0x00000001,
    WGPUVertexFormat_Uint8x4 = 0x00000002,
    WGPUVertexFormat_Sint8x2 = 0x00000003,
    WGPUVertexFormat_Sint8x4 = 0x00000004,
    WGPUVertexFormat_Unorm8x2 = 0x00000005,
    WGPUVertexFormat_Unorm8x4 = 0x00000006,
    WGPUVertexFormat_Snorm8x2 = 0x00000007,
    WGPUVertexFormat_Snorm8x4 = 0x00000008,
    WGPUVertexFormat_Uint16x2 = 0x00000009,
    WGPUVertexFormat_Uint16x4 = 0x0000000A,
    WGPUVertexFormat_Sint16x2 = 0x0000000B,
    WGPUVertexFormat_Sint16x4 = 0x0000000C,
    WGPUVertexFormat_Unorm16x2 = 0x0000000D,
    WGPUVertexFormat_Unorm16x4 = 0x0000000E,
    WGPUVertexFormat_Snorm16x2 = 0x0000000F,
    WGPUVertexFormat_Snorm16x4 = 0x00000010,
    WGPUVertexFormat_Float16x2 = 0x00000011,
    WGPUVertexFormat_Float16x4 = 0x00000012,
    WGPUVertexFormat_Float32 = 0x00000013,
    WGPUVertexFormat_Float32x2 = 0x00000014,
    WGPUVertexFormat_Float32x3 = 0x00000015,
    WGPUVertexFormat_Float32x4 = 0x00000016,
    WGPUVertexFormat_Uint32 = 0x00000017,
    WGPUVertexFormat_Uint32x2 = 0x00000018,
    WGPUVertexFormat_Uint32x3 = 0x00000019,
    WGPUVertexFormat_Uint32x4 = 0x0000001A,
    WGPUVertexFormat_Sint32 = 0x0000001B,
    WGPUVertexFormat_Sint32x2 = 0x0000001C,
    WGPUVertexFormat_Sint32x3 = 0x0000001D,
    WGPUVertexFormat_Sint32x4 = 0x0000001E,
}
```

```C
enum WGPUVertexStepMode {
    WGPUVertexStepMode_Vertex = 0x00000000,
    WGPUVertexStepMode_Instance = 0x00000001,
    WGPUVertexStepMode_VertexBufferNotUsed = 0x00000002,
}
```

```C
enum WGPUWGSLFeatureName {
    WGPUWGSLFeatureName_Undefined = 0x00000000,
    WGPUWGSLFeatureName_ReadonlyAndReadwriteStorageTextures = 0x00000001,
    WGPUWGSLFeatureName_Packed4x8IntegerDotProduct = 0x00000002,
    WGPUWGSLFeatureName_UnrestrictedPointerParameters = 0x00000003,
    WGPUWGSLFeatureName_PointerCompositeAccess = 0x00000004,
}
```



## Constants
 - `WGPU_ARRAY_LAYER_COUNT_UNDEFINED` = `0xffffffffUL`
 - `WGPU_COPY_STRIDE_UNDEFINED` = `0xffffffffUL`
 - `WGPU_DEPTH_SLICE_UNDEFINED` = `0xffffffffUL`
 - `WGPU_LIMIT_U32_UNDEFINED` = `0xffffffffUL`
 - `WGPU_LIMIT_U64_UNDEFINED` = `0xffffffffffffffffULL`
 - `WGPU_MIP_LEVEL_COUNT_UNDEFINED` = `0xffffffffUL`
 - `WGPU_QUERY_SET_INDEX_UNDEFINED` = `0xffffffffUL`
 - `WGPU_WHOLE_MAP_SIZE` = `SIZE_MAX`
 - `WGPU_WHOLE_SIZE` = `0xffffffffffffffffULL`



## Typedefs

 - `typedef uint32_t WGPUFlags`
 - `typedef uint32_t WGPUBool`



## Callback types


 - `typedef void (*WGPUAdapterRequestAdapterInfoCallback)(struct WGPUAdapterInfo adapterInfo, WGPU_NULLABLE void * userdata)`
 - `typedef void (*WGPUAdapterRequestDeviceCallback)(WGPURequestDeviceStatus status, WGPUDevice device, char const * message, WGPU_NULLABLE void * userdata)`
 - `typedef void (*WGPUBufferMapAsyncCallback)(WGPUBufferMapAsyncStatus status, WGPU_NULLABLE void * userdata)`
 - `typedef void (*WGPUDeviceCreateComputePipelineAsyncCallback)(WGPUCreatePipelineAsyncStatus status, WGPUComputePipeline pipeline, char const * message, WGPU_NULLABLE void * userdata)`
 - `typedef void (*WGPUDeviceCreateRenderPipelineAsyncCallback)(WGPUCreatePipelineAsyncStatus status, WGPURenderPipeline pipeline, char const * message, WGPU_NULLABLE void * userdata)`
 - `typedef void (*WGPUInstanceRequestAdapterCallback)(WGPURequestAdapterStatus status, WGPUAdapter adapter, char const * message, WGPU_NULLABLE void * userdata)`
 - `typedef void (*WGPUQueueOnSubmittedWorkDoneCallback)(WGPUQueueWorkDoneStatus status, WGPU_NULLABLE void * userdata)`
 - `typedef void (*WGPUShaderModuleGetCompilationInfoCallback)(WGPUCompilationInfoRequestStatus status, struct WGPUCompilationInfo const * compilationInfo, WGPU_NULLABLE void * userdata)`

