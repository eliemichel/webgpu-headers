

## Adapter { #WGPUAdapter }


This is a test.




### `wgpuAdapterEnumerateFeatures` { #wgpuAdapterEnumerateFeatures }

**Prototype:** `size_t wgpuAdapterEnumerateFeatures(WGPUAdapter adapter, WGPUFeatureName * features)`


TODO


**Arguments:**

 - TODO



**Returns:** `size_t` 
TODO





### `wgpuAdapterGetLimits` { #wgpuAdapterGetLimits }

**Prototype:** `WGPUBool wgpuAdapterGetLimits(WGPUAdapter adapter, WGPUSupportedLimits * limits)`


Get the limits that the adapter supports. This is used to drive the
choice of required limits in {device_descriptor.required_limits}.


**Arguments:**

 - TODO



**Returns:** `WGPUBool` 
True iff limits were successfully retrieved.





### `wgpuAdapterGetProperties` { #wgpuAdapterGetProperties }

**Prototype:** `void wgpuAdapterGetProperties(WGPUAdapter adapter, WGPUAdapterProperties * properties)`


TODO


**Arguments:**

 - TODO




### `wgpuAdapterHasFeature` { #wgpuAdapterHasFeature }

**Prototype:** `WGPUBool wgpuAdapterHasFeature(WGPUAdapter adapter, WGPUFeatureName feature)`


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUBool` 
TODO





### `wgpuAdapterRequestAdapterInfo` { #wgpuAdapterRequestAdapterInfo }

**Prototype:** `void wgpuAdapterRequestAdapterInfo(WGPUAdapter adapter, WGPUAdapterRequestAdapterInfoCallback callback, WGPU_NULLABLE void * userdata)`


TODO


**Arguments:**

 - TODO




### `wgpuAdapterRequestDevice` { #wgpuAdapterRequestDevice }

**Prototype:** `void wgpuAdapterRequestDevice(WGPUAdapter adapter, WGPU_NULLABLE WGPUDeviceDescriptor const * descriptor, WGPUAdapterRequestDeviceCallback callback, WGPU_NULLABLE void * userdata)`


TODO


**Arguments:**

 - TODO




