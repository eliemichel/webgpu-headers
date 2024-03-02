

# Buffer { #WGPUBuffer }


TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## Destroy { #wgpuBufferDestroy }

```C
void wgpuBufferDestroy(WGPUBuffer buffer)
```


TODO



**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## GetConstMappedRange { #wgpuBufferGetConstMappedRange }

```C
void const * wgpuBufferGetConstMappedRange(WGPUBuffer buffer, size_t offset, size_t size)
```


TODO


**Arguments:**

 - TODO



**Returns:** `void const *` 
TODO





<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## GetMapState { #wgpuBufferGetMapState }

```C
WGPUBufferMapState wgpuBufferGetMapState(WGPUBuffer buffer)
```


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUBufferMapState` 
TODO





<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## GetMappedRange { #wgpuBufferGetMappedRange }

```C
void * wgpuBufferGetMappedRange(WGPUBuffer buffer, size_t offset, size_t size)
```


TODO


**Arguments:**

 - TODO



**Returns:** `void *` 
TODO





<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## GetSize { #wgpuBufferGetSize }

```C
uint64_t wgpuBufferGetSize(WGPUBuffer buffer)
```


TODO


**Arguments:**

 - TODO



**Returns:** `uint64_t` 
TODO





<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## GetUsage { #wgpuBufferGetUsage }

```C
WGPUBufferUsageFlags wgpuBufferGetUsage(WGPUBuffer buffer)
```


TODO


**Arguments:**

 - TODO



**Returns:** `WGPUBufferUsageFlags` 
TODO





<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## MapAsync { #wgpuBufferMapAsync }

```C
void wgpuBufferMapAsync(WGPUBuffer buffer, WGPUMapModeFlags mode, size_t offset, size_t size, WGPUBufferMapAsyncCallback callback, WGPU_NULLABLE void * userdata)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## SetLabel { #wgpuBufferSetLabel }

```C
void wgpuBufferSetLabel(WGPUBuffer buffer, char const * label)
```


TODO


**Arguments:**

 - TODO




<br/><!-- poor man's styling, just for the demo before we use a non default theme -->
***

## Unmap { #wgpuBufferUnmap }

```C
void wgpuBufferUnmap(WGPUBuffer buffer)
```


TODO


**Arguments:**

 - TODO



