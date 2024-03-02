

## Device { #WGPUDevice }


TODO




### `wgpuDeviceCreateBindGroup` { #wgpuDeviceCreateBindGroup }

**Prototype:** `WGPUBindGroup wgpuDeviceCreateBindGroup(WGPUDevice device, WGPUBindGroupDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUBindGroup` 
TODO





### `wgpuDeviceCreateBindGroupLayout` { #wgpuDeviceCreateBindGroupLayout }

**Prototype:** `WGPUBindGroupLayout wgpuDeviceCreateBindGroupLayout(WGPUDevice device, WGPUBindGroupLayoutDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUBindGroupLayout` 
TODO





### `wgpuDeviceCreateBuffer` { #wgpuDeviceCreateBuffer }

**Prototype:** `WGPUBuffer wgpuDeviceCreateBuffer(WGPUDevice device, WGPUBufferDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUBuffer` 
TODO





### `wgpuDeviceCreateCommandEncoder` { #wgpuDeviceCreateCommandEncoder }

**Prototype:** `WGPUCommandEncoder wgpuDeviceCreateCommandEncoder(WGPUDevice device, WGPU_NULLABLE WGPUCommandEncoderDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUCommandEncoder` 
TODO





### `wgpuDeviceCreateComputePipeline` { #wgpuDeviceCreateComputePipeline }

**Prototype:** `WGPUComputePipeline wgpuDeviceCreateComputePipeline(WGPUDevice device, WGPUComputePipelineDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUComputePipeline` 
TODO





### `wgpuDeviceCreateComputePipelineAsync` { #wgpuDeviceCreateComputePipelineAsync }

**Prototype:** `void wgpuDeviceCreateComputePipelineAsync(WGPUDevice device, WGPUComputePipelineDescriptor const * descriptor, WGPUDeviceCreateComputePipelineAsyncCallback callback, WGPU_NULLABLE void * userdata)`


TODO


**Arguments:**

 - TODO




### `wgpuDeviceCreatePipelineLayout` { #wgpuDeviceCreatePipelineLayout }

**Prototype:** `WGPUPipelineLayout wgpuDeviceCreatePipelineLayout(WGPUDevice device, WGPUPipelineLayoutDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUPipelineLayout` 
TODO





### `wgpuDeviceCreateQuerySet` { #wgpuDeviceCreateQuerySet }

**Prototype:** `WGPUQuerySet wgpuDeviceCreateQuerySet(WGPUDevice device, WGPUQuerySetDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUQuerySet` 
TODO





### `wgpuDeviceCreateRenderBundleEncoder` { #wgpuDeviceCreateRenderBundleEncoder }

**Prototype:** `WGPURenderBundleEncoder wgpuDeviceCreateRenderBundleEncoder(WGPUDevice device, WGPURenderBundleEncoderDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPURenderBundleEncoder` 
TODO





### `wgpuDeviceCreateRenderPipeline` { #wgpuDeviceCreateRenderPipeline }

**Prototype:** `WGPURenderPipeline wgpuDeviceCreateRenderPipeline(WGPUDevice device, WGPURenderPipelineDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPURenderPipeline` 
TODO





### `wgpuDeviceCreateRenderPipelineAsync` { #wgpuDeviceCreateRenderPipelineAsync }

**Prototype:** `void wgpuDeviceCreateRenderPipelineAsync(WGPUDevice device, WGPURenderPipelineDescriptor const * descriptor, WGPUDeviceCreateRenderPipelineAsyncCallback callback, WGPU_NULLABLE void * userdata)`


TODO


**Arguments:**

 - TODO




### `wgpuDeviceCreateSampler` { #wgpuDeviceCreateSampler }

**Prototype:** `WGPUSampler wgpuDeviceCreateSampler(WGPUDevice device, WGPU_NULLABLE WGPUSamplerDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUSampler` 
TODO





### `wgpuDeviceCreateShaderModule` { #wgpuDeviceCreateShaderModule }

**Prototype:** `WGPUShaderModule wgpuDeviceCreateShaderModule(WGPUDevice device, WGPUShaderModuleDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUShaderModule` 
TODO





### `wgpuDeviceCreateTexture` { #wgpuDeviceCreateTexture }

**Prototype:** `WGPUTexture wgpuDeviceCreateTexture(WGPUDevice device, WGPUTextureDescriptor const * descriptor)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUTexture` 
TODO





### `wgpuDeviceDestroy` { #wgpuDeviceDestroy }

**Prototype:** `void wgpuDeviceDestroy(WGPUDevice device)`


TODO


**Arguments:**

 - TODO




### `wgpuDeviceEnumerateFeatures` { #wgpuDeviceEnumerateFeatures }

**Prototype:** `size_t wgpuDeviceEnumerateFeatures(WGPUDevice device, WGPUFeatureName * features)`


TODO


**Arguments:**

 - TODO



**Returns:** `size_t` 
TODO





### `wgpuDeviceGetLimits` { #wgpuDeviceGetLimits }

**Prototype:** `WGPUBool wgpuDeviceGetLimits(WGPUDevice device, WGPUSupportedLimits * limits)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUBool` 
TODO





### `wgpuDeviceGetQueue` { #wgpuDeviceGetQueue }

**Prototype:** `WGPUQueue wgpuDeviceGetQueue(WGPUDevice device)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUQueue` 
TODO





### `wgpuDeviceHasFeature` { #wgpuDeviceHasFeature }

**Prototype:** `WGPUBool wgpuDeviceHasFeature(WGPUDevice device, WGPUFeatureName feature)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUBool` 
TODO





### `wgpuDevicePopErrorScope` { #wgpuDevicePopErrorScope }

**Prototype:** `void wgpuDevicePopErrorScope(WGPUDevice device, WGPUErrorCallback callback, void * userdata)`


TODO


**Arguments:**

 - TODO




### `wgpuDevicePushErrorScope` { #wgpuDevicePushErrorScope }

**Prototype:** `void wgpuDevicePushErrorScope(WGPUDevice device, WGPUErrorFilter filter)`


TODO


**Arguments:**

 - TODO




### `wgpuDeviceSetLabel` { #wgpuDeviceSetLabel }

**Prototype:** `void wgpuDeviceSetLabel(WGPUDevice device, char const * label)`


TODO


**Arguments:**

 - TODO




### `wgpuDeviceSetUncapturedErrorCallback` { #wgpuDeviceSetUncapturedErrorCallback }

**Prototype:** `void wgpuDeviceSetUncapturedErrorCallback(WGPUDevice device, WGPUErrorCallback callback, void * userdata)`


TODO


**Arguments:**

 - TODO




